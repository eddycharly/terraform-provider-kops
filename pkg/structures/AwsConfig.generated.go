package structures

import (
	"github.com/eddycharly/terraform-provider-kops/pkg/api"
)

func ExpandAwsConfig(in map[string]interface{}) api.AwsConfig {
	if in == nil {
		panic("expand AwsConfig failure, in is nil")
	}
	return api.AwsConfig{
		Profile: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["profile"]),
		Region: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["region"]),
		AssumeRole: func(in interface{}) *api.AwsAssumeRole {
			return func(in interface{}) *api.AwsAssumeRole {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in api.AwsAssumeRole) *api.AwsAssumeRole {
					return &in
				}(func(in interface{}) api.AwsAssumeRole {
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return api.AwsAssumeRole{}
					}
					return (ExpandAwsAssumeRole(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["assume_role"]),
	}
}

func FlattenAwsConfigInto(in api.AwsConfig, out map[string]interface{}) {
	out["profile"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Profile)
	out["region"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Region)
	out["assume_role"] = func(in *api.AwsAssumeRole) interface{} {
		return func(in *api.AwsAssumeRole) interface{} {
			if in == nil {
				return nil
			}
			return func(in api.AwsAssumeRole) interface{} {
				return func(in api.AwsAssumeRole) []map[string]interface{} {
					return []map[string]interface{}{FlattenAwsAssumeRole(in)}
				}(in)
			}(*in)
		}(in)
	}(in.AssumeRole)
}

func FlattenAwsConfig(in api.AwsConfig) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenAwsConfigInto(in, out)
	return out
}
