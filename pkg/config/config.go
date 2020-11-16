package config

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/eddycharly/terraform-provider-kops/pkg/api/config"
	"github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/kops/pkg/client/simple"
	"k8s.io/kops/pkg/client/simple/vfsclientset"
	"k8s.io/kops/util/pkg/vfs"
)

const (
	invalidStateError = `Unable to read state store s3 bucket.
Please use a valid s3 bucket uri on state_store attribute or KOPS_STATE_STORE env var.
A valid value follows the format s3://<bucket>.
Trailing slash will be trimmed.`
)

type options struct {
	clientset            simple.Clientset
	rollingUpdateOptions config.RollingUpdate
	validateOptions      config.Validate
}

func ConfigureProvider(d *schema.ResourceData) (interface{}, error) {
	providerConfig := schemas.ExpandConfigProvider(d.Get("").(map[string]interface{}))
	err := initCredentials(providerConfig.Aws)
	if err != nil {
		return nil, err
	}
	basePath, err := vfs.Context.BuildVfsPath(providerConfig.StateStore)
	if err != nil {
		return nil, fmt.Errorf("error building path for %q: %v", providerConfig.StateStore, err)
	}
	if !vfs.IsClusterReadable(basePath) {
		return nil, field.Invalid(field.NewPath("State Store"), providerConfig.StateStore, invalidStateError)
	}
	return &options{
		clientset:            vfsclientset.NewVFSClientset(basePath),
		rollingUpdateOptions: providerConfig.RollingUpdate,
		validateOptions:      providerConfig.Validate,
	}, nil
}

func Clientset(in interface{}) simple.Clientset {
	return in.(*options).clientset
}

func RollingUpdateOptions(in interface{}) config.RollingUpdate {
	return in.(*options).rollingUpdateOptions
}

func ValidateOptions(in interface{}) config.Validate {
	return in.(*options).validateOptions
}

func initCredentials(config *config.Aws) error {
	if config == nil {
		return nil
	}
	if config.Region != "" {
		os.Setenv("AWS_DEFAULT_REGION", config.Region)
	}
	if config.Profile != "" {
		os.Setenv("AWS_SDK_LOAD_CONFIG", "1")
		os.Setenv("AWS_PROFILE", config.Profile)
	}
	if config.AssumeRole != nil {
		svc := sts.New(session.New())
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
