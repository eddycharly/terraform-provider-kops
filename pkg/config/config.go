package config

import (
	"fmt"

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

type Config struct {
	stateStore string
	Clientset  simple.Clientset
}

func ConfigureProvider(data *schema.ResourceData) (interface{}, error) {
	registryPath := data.Get("state_store").(string)

	basePath, err := vfs.Context.BuildVfsPath(registryPath)
	if err != nil {
		return nil, fmt.Errorf("error building path for %q: %v", registryPath, err)
	}

	if !vfs.IsClusterReadable(basePath) {
		return nil, field.Invalid(field.NewPath("State Store"), registryPath, invalidStateError)
	}

	clientset := vfsclientset.NewVFSClientset(basePath)

	return &Config{
		Clientset:  clientset,
		stateStore: registryPath,
	}, nil
}
