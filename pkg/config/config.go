package config

import (
	"fmt"
	"os"

	"github.com/eddycharly/terraform-provider-kops/pkg/structures"
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

func ConfigureProvider(d *schema.ResourceData) (interface{}, error) {
	providerConfig := structures.ExpandProviderConfig(d.Get("").(map[string]interface{}))
	if providerConfig.Aws != nil {
		if providerConfig.Aws.Profile != "" {
			os.Setenv("AWS_SDK_LOAD_CONFIG", "1")
			os.Setenv("AWS_PROFILE", providerConfig.Aws.Profile)
		}
	}
	basePath, err := vfs.Context.BuildVfsPath(providerConfig.StateStore)
	if err != nil {
		return nil, fmt.Errorf("error building path for %q: %v", providerConfig.StateStore, err)
	}
	if !vfs.IsClusterReadable(basePath) {
		return nil, field.Invalid(field.NewPath("State Store"), providerConfig.StateStore, invalidStateError)
	}
	return vfsclientset.NewVFSClientset(basePath), nil
}

func Clientset(in interface{}) simple.Clientset {
	return in.(simple.Clientset)
}
