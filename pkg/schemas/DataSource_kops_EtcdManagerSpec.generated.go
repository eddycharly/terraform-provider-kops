package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

func DataSourceKopsEtcdManagerSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"image": ComputedString(),
			"env":   ComputedList(DataSourceKopsEnvVar()),
		},
	}
}

func ExpandDataSourceKopsEtcdManagerSpec(in map[string]interface{}) kops.EtcdManagerSpec {
	if in == nil {
		panic("expand EtcdManagerSpec failure, in is nil")
	}
	return kops.EtcdManagerSpec{
		Image: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["image"]),
		Env: func(in interface{}) []kops.EnvVar {
			return func(in interface{}) []kops.EnvVar {
				var out []kops.EnvVar
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) kops.EnvVar {
						if in == nil {
							return kops.EnvVar{}
						}
						return (ExpandDataSourceKopsEnvVar(in.(map[string]interface{})))
					}(in))
				}
				return out
			}(in)
		}(in["env"]),
	}
}

func FlattenDataSourceKopsEtcdManagerSpecInto(in kops.EtcdManagerSpec, out map[string]interface{}) {
	out["image"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Image)
	out["env"] = func(in []kops.EnvVar) interface{} {
		return func(in []kops.EnvVar) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, func(in kops.EnvVar) interface{} {
					return FlattenDataSourceKopsEnvVar(in)
				}(in))
			}
			return out
		}(in)
	}(in.Env)
}

func FlattenDataSourceKopsEtcdManagerSpec(in kops.EtcdManagerSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceKopsEtcdManagerSpecInto(in, out)
	return out
}
