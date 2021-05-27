package schemas

import (
	"reflect"
	"time"

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
			"cluster_name":  RequiredString(),
			"admin":         OptionalComputedInt(),
			"internal":      OptionalComputedBool(),
			"server":        ComputedString(),
			"context":       ComputedString(),
			"namespace":     ComputedString(),
			"kube_user":     ComputedString(),
			"kube_password": Sensitive(ComputedString()),
			"ca_cert":       Sensitive(ComputedString()),
			"client_cert":   Sensitive(ComputedString()),
			"client_key":    Sensitive(ComputedString()),
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
		Admin: func(in interface{}) *time.Duration {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *time.Duration {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in time.Duration) *time.Duration {
					return &in
				}(time.Duration(ExpandInt(in)))
			}(in)
		}(in["admin"]),
		Internal: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["internal"]),
		Config: func(in interface{}) kube.Config {
			return kubeschemas.ExpandDataSourceConfig(in.(map[string]interface{}))
		}(in),
	}
}

func FlattenDataSourceKubeConfigInto(in datasources.KubeConfig, out map[string]interface{}) {
	out["cluster_name"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.ClusterName)
	out["admin"] = func(in *time.Duration) interface{} {
		return func(in *time.Duration) interface{} {
			if in == nil {
				return nil
			}
			return func(in time.Duration) interface{} {
				return FlattenInt(int(in))
			}(*in)
		}(in)
	}(in.Admin)
	out["internal"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.Internal)
	kubeschemas.FlattenDataSourceConfigInto(in.Config, out)
}

func FlattenDataSourceKubeConfig(in datasources.KubeConfig) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceKubeConfigInto(in, out)
	return out
}
