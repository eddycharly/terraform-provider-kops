package config

import (
	"fmt"
	"os"

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

type config struct {
	stateStore string
	clientset  simple.Clientset
}

func ConfigureProvider(data *schema.ResourceData) (interface{}, error) {
	profile := data.Get("aws_profile").(string)

	if profile != "" {
		os.Setenv("AWS_SDK_LOAD_CONFIG", "1")
		os.Setenv("AWS_PROFILE", profile)
	}

	registryPath := data.Get("state_store").(string)

	basePath, err := vfs.Context.BuildVfsPath(registryPath)
	if err != nil {
		return nil, fmt.Errorf("error building path for %q: %v", registryPath, err)
	}

	if !vfs.IsClusterReadable(basePath) {
		return nil, field.Invalid(field.NewPath("State Store"), registryPath, invalidStateError)
	}

	clientset := vfsclientset.NewVFSClientset(basePath)

	return &config{
		clientset:  clientset,
		stateStore: registryPath,
	}, nil
}

func Clientset(in interface{}) simple.Clientset {
	return in.(*config).clientset
}
