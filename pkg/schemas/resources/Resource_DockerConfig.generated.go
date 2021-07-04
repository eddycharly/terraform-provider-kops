package schemas

import (
	"github.com/eddycharly/terraform-provider-kops/pkg/api/resources"
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var _ = Schema

func ResourceDockerConfig() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cluster_name":  Required(String()),
			"docker_config": Required(String()),
		},
	}
}

func ExpandResourceDockerConfig(in map[string]interface{}) resources.DockerConfig {
	if in == nil {
		panic("expand DockerConfig failure, in is nil")
	}
	out := resources.DockerConfig{}
	if in, ok := in["cluster_name"]; ok && in != nil {
		out.ClusterName = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["docker_config"]; ok && in != nil {
		out.DockerConfig = func(in interface{}) string { return string(in.(string)) }(in)
	}
	return out
}

func FlattenResourceDockerConfigInto(in resources.DockerConfig, out map[string]interface{}) {
	out["cluster_name"] = func(in string) interface{} { return string(in) }(in.ClusterName)
	out["docker_config"] = func(in string) interface{} { return string(in) }(in.DockerConfig)
}

func FlattenResourceDockerConfig(in resources.DockerConfig) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceDockerConfigInto(in, out)
	return out
}
