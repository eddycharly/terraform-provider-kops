package schemas

import (
	"github.com/eddycharly/terraform-provider-kops/pkg/api/config"
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var _ = Schema

func ConfigProvider() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"state_store": RequiredString(),
			"aws":         OptionalStruct(ConfigAws()),
		},
	}
}

func ExpandConfigProvider(in map[string]interface{}) config.Provider {
	if in == nil {
		panic("expand Provider failure, in is nil")
	}
	return config.Provider{
		StateStore: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["state_store"]),
		Aws: func(in interface{}) *config.Aws {
			return func(in interface{}) *config.Aws {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in config.Aws) *config.Aws {
					return &in
				}(func(in interface{}) config.Aws {
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return config.Aws{}
					}
					return (ExpandConfigAws(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["aws"]),
	}
}

func FlattenConfigProviderInto(in config.Provider, out map[string]interface{}) {
	out["state_store"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.StateStore)
	out["aws"] = func(in *config.Aws) interface{} {
		return func(in *config.Aws) interface{} {
			if in == nil {
				return nil
			}
			return func(in config.Aws) interface{} {
				return func(in config.Aws) []map[string]interface{} {
					return []map[string]interface{}{FlattenConfigAws(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Aws)
}

func FlattenConfigProvider(in config.Provider) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenConfigProviderInto(in, out)
	return out
}
