package schemas

import (
	"github.com/eddycharly/terraform-provider-kops/pkg/api/resources"
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var _ = Schema

func DataSourceClusterSecrets() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"docker_config":   Sensitive(ComputedString()),
			"cluster_ca_cert": Sensitive(ComputedString()),
			"cluster_ca_key":  Sensitive(ComputedString()),
		},
	}

	return res
}

func ExpandDataSourceClusterSecrets(in map[string]interface{}) resources.ClusterSecrets {
	if in == nil {
		panic("expand ClusterSecrets failure, in is nil")
	}
	return resources.ClusterSecrets{
		DockerConfig: func(in interface{}) string /**/ {
			return string(ExpandString(in))
		}(in["docker_config"]),
		ClusterCaCert: func(in interface{}) string /**/ {
			return string(ExpandString(in))
		}(in["cluster_ca_cert"]),
		ClusterCaKey: func(in interface{}) string /**/ {
			return string(ExpandString(in))
		}(in["cluster_ca_key"]),
	}
}

func FlattenDataSourceClusterSecretsInto(in resources.ClusterSecrets, out map[string]interface{}) {
	out["docker_config"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.DockerConfig)
	out["cluster_ca_cert"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.ClusterCaCert)
	out["cluster_ca_key"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.ClusterCaKey)
}

func FlattenDataSourceClusterSecrets(in resources.ClusterSecrets) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceClusterSecretsInto(in, out)
	return out
}
