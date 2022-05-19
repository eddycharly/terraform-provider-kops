package config

import (
	"context"
	"flag"
	"os"
	"strconv"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/eddycharly/terraform-provider-kops/pkg/api/config"
	configschemas "github.com/eddycharly/terraform-provider-kops/pkg/schemas/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/klog"
	"k8s.io/kops/cloudmock/aws/mockec2"
	"k8s.io/kops/pkg/client/simple"
	"k8s.io/kops/pkg/client/simple/vfsclientset"
	"k8s.io/kops/pkg/featureflag"
	"k8s.io/kops/pkg/testutils"
	"k8s.io/kops/upup/pkg/fi/cloudup/awsup"
	"k8s.io/kops/util/pkg/vfs"
)

const (
	invalidStateError = `Unable to read state store s3 bucket.
Please use a valid s3 bucket uri on state_store attribute or KOPS_STATE_STORE env var.
A valid value follows the format s3://<bucket>.
Trailing slash will be trimmed.`
)

type options struct {
	clientset simple.Clientset
}

func ConfigureProvider(_ context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	providerConfig := configschemas.ExpandConfigProvider(d.Get("").(map[string]interface{}))
	if err := initKlog(providerConfig.Klog); err != nil {
		return nil, diag.FromErr(err)
	}
	if err := initFeatureFlags(providerConfig.FeatureFlags); err != nil {
		return nil, diag.FromErr(err)
	}
	if err := initAwsCredentials(providerConfig.Aws); err != nil {
		return nil, diag.FromErr(err)
	}
	if err := initOpenstackCredentials(providerConfig.Openstack); err != nil {
		return nil, diag.FromErr(err)
	}
	if providerConfig.Mock {
		initMock()
	}
	basePath, err := vfs.Context.BuildVfsPath(providerConfig.StateStore)
	if err != nil {
		return nil, diag.Errorf("error building path for %q: %v", providerConfig.StateStore, err)
	}
	if !vfs.IsClusterReadable(basePath) {
		return nil, diag.FromErr(field.Invalid(field.NewPath("State Store"), providerConfig.StateStore, invalidStateError))
	}
	return &options{
		clientset: vfsclientset.NewVFSClientset(basePath),
	}, nil
}

func Clientset(in interface{}) simple.Clientset {
	return in.(*options).clientset
}

func setEnvVarSimple(name, value string) {
	if value != "" {
		os.Setenv(name, value)
	}
}

func initAwsCredentials(config *config.Aws) error {
	if config == nil {
		return nil
	}
	setEnvVarSimple("AWS_DEFAULT_REGION", config.Region)
	setEnvVarSimple("AWS_ACCESS_KEY_ID", config.AccessKey)
	setEnvVarSimple("AWS_SECRET_ACCESS_KEY", config.SecretKey)
	setEnvVarSimple("S3_ENDPOINT", config.S3Endpoint)
	setEnvVarSimple("S3_REGION", config.S3Region)
	setEnvVarSimple("S3_ACCESS_KEY_ID", config.S3AccessKey)
	setEnvVarSimple("S3_SECRET_ACCESS_KEY", config.S3SecretKey)
	if config.SkipRegionCheck {
		os.Setenv("SKIP_REGION_CHECK", "1")
	}
	if config.Profile != "" {
		os.Setenv("AWS_SDK_LOAD_CONFIG", "1")
		os.Setenv("AWS_PROFILE", config.Profile)
	}
	if config.AssumeRole != nil {
		ses, err := session.NewSession()
		if err != nil {
			return err
		}
		svc := sts.New(ses)
		input := &sts.AssumeRoleInput{
			RoleArn:         aws.String(config.AssumeRole.RoleArn),
			RoleSessionName: aws.String("TF-PROVIDER-KOPS"),
		}
		result, err := svc.AssumeRole(input)
		if err != nil {
			return err
		}
		os.Setenv("AWS_ACCESS_KEY_ID", *result.Credentials.AccessKeyId)
		os.Setenv("AWS_SECRET_ACCESS_KEY", *result.Credentials.SecretAccessKey)
		os.Setenv("AWS_SESSION_TOKEN", *result.Credentials.SessionToken)
	}
	return nil
}

func initOpenstackCredentials(config *config.Openstack) error {
	if config == nil {
		return nil
	}
	setEnvVarSimple("OS_TENANT_ID", config.TenantId)
	setEnvVarSimple("OS_TENANT_NAME", config.TenantName)
	setEnvVarSimple("OS_PROJECT_ID", config.ProjectId)
	setEnvVarSimple("OS_PROJECT_NAME", config.ProjectName)
	setEnvVarSimple("OS_PROJECT_DOMAIN_NAME", config.ProjectDomainName)
	setEnvVarSimple("OS_PROJECT_DOMAIN_ID", config.ProjectDomainId)
	setEnvVarSimple("OS_DOMAIN_NAME", config.DomainName)
	setEnvVarSimple("OS_DOMAIN_ID", config.DomainId)
	setEnvVarSimple("OS_USERNAME", config.Username)
	setEnvVarSimple("OS_PASSWORD", config.Password)
	setEnvVarSimple("OS_AUTH_URL", config.AuthUrl)
	setEnvVarSimple("OS_REGION_NAME", config.RegionName)
	setEnvVarSimple("OS_APPLICATION_CREDENTIAL_ID", config.ApplicationCredentialId)
	setEnvVarSimple("OS_APPLICATION_CREDENTIAL_SECRET", config.ApplicationCredentialSecret)
	return nil
}

func initKlog(config *config.Klog) error {
	if config == nil {
		return nil
	}
	flags := flag.NewFlagSet("klog", flag.ExitOnError)
	if config.Verbosity != nil {
		if err := flags.Set("v", strconv.Itoa(*config.Verbosity)); err != nil {
			return err
		}
	}
	klog.InitFlags(flags)
	return nil
}

func initFeatureFlags(config []string) error {
	if len(config) == 0 {
		return nil
	}
	featureflag.ParseFlags(strings.Join(config, ","))
	return nil
}

// nolint
func initMock() {
	h := &testutils.IntegrationTestHarness{}

	h.SetupMockAWS()
	h.SetupMockGCE()

	cloudTags := map[string]string{}
	awsCloud, _ := awsup.NewAWSCloud("us-test-1", cloudTags)
	(awsCloud.EC2().(*mockec2.MockEC2)).CreateVpcWithId(&ec2.CreateVpcInput{
		CidrBlock: aws.String("10.0.0.0/12"),
	}, "vpc-12345678")

	awsCloud.EC2().CreateSubnet(&ec2.CreateSubnetInput{
		AvailabilityZone: aws.String("us-test-1a"),
		VpcId:            aws.String("vpc-12345678"),
		CidrBlock:        aws.String("10.10.0.0/24"),
	})

	awsCloud.EC2().CreateSubnet(&ec2.CreateSubnetInput{
		AvailabilityZone: aws.String("us-test-1a"),
		VpcId:            aws.String("vpc-12345678"),
		CidrBlock:        aws.String("10.11.0.0/24"),
	})
}
