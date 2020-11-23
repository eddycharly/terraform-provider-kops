package schemas

import (
	"github.com/eddycharly/terraform-provider-kops/pkg/api/datasources"
	"github.com/eddycharly/terraform-provider-kops/pkg/api/kube"
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	kubeschemas "github.com/eddycharly/terraform-provider-kops/pkg/schemas/kube"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var _ = Schema

func DataSourceKubeConfig() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cluster_name":      RequiredString(),
			"server":            ComputedString(),
			"context":           ComputedString(),
			"namespace":         ComputedString(),
			"kube_bearer_token": Sensitive(ComputedString()),
			"kube_user":         ComputedString(),
			"kube_password":     Sensitive(ComputedString()),
			"ca_cert":           Sensitive(ComputedString()),
			"client_cert":       Sensitive(ComputedString()),
			"client_key":        Sensitive(ComputedString()),
		},
	}
}

func ExpandDataSourceKubeConfig(in map[string]interface{}) datasources.KubeConfig {
	if in == nil {
		panic("expand KubeConfig failure, in is nil")
	}
	return datasources.KubeConfig{
		ClusterName: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["cluster_name"]),
		Config: func(in interface{}) kube.Config {
			return kubeschemas.ExpandDataSourceConfig(in.(map[string]interface{}))
		}(in),
	}
}

func FlattenDataSourceKubeConfigInto(in datasources.KubeConfig, out map[string]interface{}) {
	out["cluster_name"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.ClusterName)
	kubeschemas.FlattenDataSourceConfigInto(in.Config, out)
}

func FlattenDataSourceKubeConfig(in datasources.KubeConfig) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceKubeConfigInto(in, out)
	return out
}
