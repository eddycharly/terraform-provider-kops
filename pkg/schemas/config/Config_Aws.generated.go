package schemas

import (
	"github.com/eddycharly/terraform-provider-kops/pkg/api/config"
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var _ = Schema

func ConfigAws() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"profile":           OptionalString(),
			"region":            OptionalString(),
			"access_key":        OptionalString(),
			"secret_key":        OptionalString(),
			"assume_role":       OptionalStruct(ConfigAwsAssumeRole()),
			"s3_endpoint":       OptionalString(),
			"s3_region":         OptionalString(),
			"s3_access_key":     OptionalString(),
			"s3_secret_key":     OptionalString(),
			"skip_region_check": OptionalBool(),
		},
	}

	return res
}

func ExpandConfigAws(in map[string]interface{}) config.Aws {
	if in == nil {
		panic("expand Aws failure, in is nil")
	}
	return config.Aws{
		Profile: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["profile"]),
		Region: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["region"]),
		AccessKey: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["access_key"]),
		SecretKey: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["secret_key"]),
		AssumeRole: func(in interface{}) *config.AwsAssumeRole {
			return func(in interface{}) *config.AwsAssumeRole {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in config.AwsAssumeRole) *config.AwsAssumeRole {
					return &in
				}(func(in interface{}) config.AwsAssumeRole {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandConfigAwsAssumeRole(in[0].(map[string]interface{}))
					}
					return config.AwsAssumeRole{}
				}(in))
			}(in)
		}(in["assume_role"]),
		S3Endpoint: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["s3_endpoint"]),
		S3Region: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["s3_region"]),
		S3AccessKey: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["s3_access_key"]),
		S3SecretKey: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["s3_secret_key"]),
		SkipRegionCheck: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["skip_region_check"]),
	}
}

func FlattenConfigAwsInto(in config.Aws, out map[string]interface{}) {
	out["profile"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Profile)
	out["region"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Region)
	out["access_key"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.AccessKey)
	out["secret_key"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.SecretKey)
	out["assume_role"] = func(in *config.AwsAssumeRole) interface{} {
		return func(in *config.AwsAssumeRole) interface{} {
			if in == nil {
				return nil
			}
			return func(in config.AwsAssumeRole) interface{} {
				return func(in config.AwsAssumeRole) []interface{} {
					return []interface{}{FlattenConfigAwsAssumeRole(in)}
				}(in)
			}(*in)
		}(in)
	}(in.AssumeRole)
	out["s3_endpoint"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.S3Endpoint)
	out["s3_region"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.S3Region)
	out["s3_access_key"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.S3AccessKey)
	out["s3_secret_key"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.S3SecretKey)
	out["skip_region_check"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.SkipRegionCheck)
}

func FlattenConfigAws(in config.Aws) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenConfigAwsInto(in, out)
	return out
}
