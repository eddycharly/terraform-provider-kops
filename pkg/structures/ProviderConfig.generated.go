package structures

import (
	"github.com/eddycharly/terraform-provider-kops/pkg/api"
)

func ExpandProviderConfig(in map[string]interface{}) api.ProviderConfig {
	if in == nil {
		panic("expand ProviderConfig failure, in is nil")
	}
	return api.ProviderConfig{
		StateStore: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["state_store"]),
		Aws: func(in interface{}) *api.AwsConfig {
			return func(in interface{}) *api.AwsConfig {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in api.AwsConfig) *api.AwsConfig {
					return &in
				}(func(in interface{}) api.AwsConfig {
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return api.AwsConfig{}
					}
					return (ExpandAwsConfig(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["aws"]),
		RollingUpdateOptions: func(in interface{}) api.RollingUpdateOptions {
			return func(in interface{}) api.RollingUpdateOptions {
				if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
					return api.RollingUpdateOptions{}
				}
				return (ExpandRollingUpdateOptions(in.([]interface{})[0].(map[string]interface{})))
			}(in)
		}(in["rolling_update_options"]),
		ValidateOptions: func(in interface{}) api.ValidateOptions {
			return func(in interface{}) api.ValidateOptions {
				if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
					return api.ValidateOptions{}
				}
				return (ExpandValidateOptions(in.([]interface{})[0].(map[string]interface{})))
			}(in)
		}(in["validate_options"]),
	}
}

func FlattenProviderConfigInto(in api.ProviderConfig, out map[string]interface{}) {
	out["state_store"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.StateStore)
	out["aws"] = func(in *api.AwsConfig) interface{} {
		return func(in *api.AwsConfig) interface{} {
			if in == nil {
				return nil
			}
			return func(in api.AwsConfig) interface{} {
				return func(in api.AwsConfig) []map[string]interface{} {
					return []map[string]interface{}{FlattenAwsConfig(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Aws)
	out["rolling_update_options"] = func(in api.RollingUpdateOptions) interface{} {
		return func(in api.RollingUpdateOptions) []map[string]interface{} {
			return []map[string]interface{}{FlattenRollingUpdateOptions(in)}
		}(in)
	}(in.RollingUpdateOptions)
	out["validate_options"] = func(in api.ValidateOptions) interface{} {
		return func(in api.ValidateOptions) []map[string]interface{} {
			return []map[string]interface{}{FlattenValidateOptions(in)}
		}(in)
	}(in.ValidateOptions)
}

func FlattenProviderConfig(in api.ProviderConfig) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenProviderConfigInto(in, out)
	return out
}
