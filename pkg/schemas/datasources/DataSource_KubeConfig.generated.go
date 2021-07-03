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
			"admin":         Optional(Computed(Ptr(Int()))),
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
			return func(in time.Duration) *time.Duration { return &in }(func(in interface{}) time.Duration { return time.Duration(in.(int)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["internal"]; ok && in != nil {
		out.Internal = func(in interface{}) bool { return in.(bool) }(in)
	}
	out.Config = kubeschemas.ExpandDataSourceConfig(in)
	return out
}
