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
			"cluster_name":  RequiredString(),
			"docker_config": RequiredString(),
		},
	}
}

func ExpandResourceDockerConfig(in map[string]interface{}) resources.DockerConfig {
	if in == nil {
		panic("expand DockerConfig failure, in is nil")
	}
	return resources.DockerConfig{
		ClusterName: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["cluster_name"]),
		DockerConfig: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["docker_config"]),
	}
}

func FlattenResourceDockerConfigInto(in resources.DockerConfig, out map[string]interface{}) {
	out["cluster_name"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.ClusterName)
	out["docker_config"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.DockerConfig)
}

func FlattenResourceDockerConfig(in resources.DockerConfig) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceDockerConfigInto(in, out)
	return out
}
