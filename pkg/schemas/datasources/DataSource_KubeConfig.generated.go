package schemas

import (
	"time"

	"github.com/eddycharly/terraform-provider-kops/pkg/api/datasources"
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	kubeschemas "github.com/eddycharly/terraform-provider-kops/pkg/schemas/kube"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var _ = Schema

func DataSourceKubeConfig() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cluster_name":  Required(String()),
			"admin":         Optional(Computed(Nullable(Int()))),
			"internal":      Optional(Computed(Bool())),
			"server":        Computed(String()),
			"context":       Computed(String()),
			"namespace":     Computed(String()),
			"kube_user":     Computed(String()),
			"kube_password": Sensitive(Computed(String())),
			"ca_cert":       Sensitive(Computed(String())),
			"client_cert":   Sensitive(Computed(String())),
			"client_key":    Sensitive(Computed(String())),
		},
	}
}

func ExpandDataSourceKubeConfig(in map[string]interface{}) datasources.KubeConfig {
	if in == nil {
		panic("expand KubeConfig failure, in is nil")
	}
	out := datasources.KubeConfig{}
	if in, ok := in["cluster_name"]; ok && in != nil {
		out.ClusterName = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["admin"]; ok && in != nil {
		out.Admin = func(in interface{}) *time.Duration {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in time.Duration) *time.Duration { return &in }(func(in interface{}) time.Duration { return time.Duration(in.(int)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["internal"]; ok && in != nil {
		out.Internal = func(in interface{}) bool { return in.(bool) }(in)
	}
	out.Config = kubeschemas.ExpandDataSourceConfig(in)
	return out
}

func FlattenDataSourceKubeConfigInto(in datasources.KubeConfig, out map[string]interface{}) {
	out["cluster_name"] = func(in string) interface{} { return string(in) }(in.ClusterName)
	out["admin"] = func(in *time.Duration) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in time.Duration) interface{} { return int(in) }(*in)}
	}(in.Admin)
	out["internal"] = func(in bool) interface{} { return in }(in.Internal)
	kubeschemas.FlattenDataSourceConfigInto(in.Config, out)
}

func FlattenDataSourceKubeConfig(in datasources.KubeConfig) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceKubeConfigInto(in, out)
	return out
}
