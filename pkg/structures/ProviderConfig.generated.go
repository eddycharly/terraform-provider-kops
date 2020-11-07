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
	}
}

func FlattenProviderConfig(in api.ProviderConfig) map[string]interface{} {
	return map[string]interface{}{
		"state_store": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.StateStore),
		"aws": func(in *api.AwsConfig) interface{} {
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
		}(in.Aws),
	}
}
