package schemas

import (
	"github.com/eddycharly/terraform-provider-kops/pkg/api/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ConfigConfigProvider() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"state_store":    RequiredString(),
			"aws":            OptionalStruct(ConfigConfigAws()),
			"rolling_update": OptionalStruct(ConfigConfigRollingUpdate()),
			"validate":       OptionalStruct(ConfigConfigValidate()),
		},
	}
}

func ExpandConfigConfigProvider(in map[string]interface{}) config.Provider {
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
					return (ExpandConfigConfigAws(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["aws"]),
		RollingUpdate: func(in interface{}) config.RollingUpdate {
			return func(in interface{}) config.RollingUpdate {
				if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
					return config.RollingUpdate{}
				}
				return (ExpandConfigConfigRollingUpdate(in.([]interface{})[0].(map[string]interface{})))
			}(in)
		}(in["rolling_update"]),
		Validate: func(in interface{}) config.Validate {
			return func(in interface{}) config.Validate {
				if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
					return config.Validate{}
				}
				return (ExpandConfigConfigValidate(in.([]interface{})[0].(map[string]interface{})))
			}(in)
		}(in["validate"]),
	}
}

func FlattenConfigConfigProviderInto(in config.Provider, out map[string]interface{}) {
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
					return []map[string]interface{}{FlattenConfigConfigAws(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Aws)
	out["rolling_update"] = func(in config.RollingUpdate) interface{} {
		return func(in config.RollingUpdate) []map[string]interface{} {
			return []map[string]interface{}{FlattenConfigConfigRollingUpdate(in)}
		}(in)
	}(in.RollingUpdate)
	out["validate"] = func(in config.Validate) interface{} {
		return func(in config.Validate) []map[string]interface{} {
			return []map[string]interface{}{FlattenConfigConfigValidate(in)}
		}(in)
	}(in.Validate)
}

func FlattenConfigConfigProvider(in config.Provider) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenConfigConfigProviderInto(in, out)
	return out
}
