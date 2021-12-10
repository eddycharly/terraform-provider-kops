package schemas

import (
	"github.com/eddycharly/terraform-provider-kops/pkg/api/resources"
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var _ = Schema

func ResourceClusterSecrets() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"docker_config": Sensitive(OptionalString()),
		},
	}

	return res
}

func ExpandResourceClusterSecrets(in map[string]interface{}) resources.ClusterSecrets {
	if in == nil {
		panic("expand ClusterSecrets failure, in is nil")
	}
	return resources.ClusterSecrets{
		DockerConfig: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["docker_config"]),
	}
}

func FlattenResourceClusterSecretsInto(in resources.ClusterSecrets, out map[string]interface{}) {
	out["docker_config"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.DockerConfig)
}

func FlattenResourceClusterSecrets(in resources.ClusterSecrets) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceClusterSecretsInto(in, out)
	return out
}
