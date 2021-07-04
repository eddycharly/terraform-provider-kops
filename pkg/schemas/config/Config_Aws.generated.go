package schemas

import (
	"github.com/eddycharly/terraform-provider-kops/pkg/api/config"
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var _ = Schema

func ConfigAws() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"profile":           Optional(String()),
			"region":            Optional(String()),
			"access_key":        Optional(String()),
			"secret_key":        Optional(String()),
			"assume_role":       Optional(Struct(ConfigAwsAssumeRole())),
			"s3_endpoint":       Optional(String()),
			"s3_region":         Optional(String()),
			"s3_access_key":     Optional(String()),
			"s3_secret_key":     Optional(String()),
			"skip_region_check": Optional(Bool()),
		},
	}
}

func ExpandConfigAws(in map[string]interface{}) config.Aws {
	if in == nil {
		panic("expand Aws failure, in is nil")
	}
	out := config.Aws{}
	if in, ok := in["profile"]; ok && in != nil {
		out.Profile = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["region"]; ok && in != nil {
		out.Region = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["access_key"]; ok && in != nil {
		out.AccessKey = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["secret_key"]; ok && in != nil {
		out.SecretKey = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["assume_role"]; ok && in != nil {
		out.AssumeRole = func(in interface{}) *config.AwsAssumeRole {
			if in == nil {
				return nil
			}
			return func(in config.AwsAssumeRole) *config.AwsAssumeRole { return &in }(func(in interface{}) config.AwsAssumeRole {
				if in == nil {
					return config.AwsAssumeRole{}
				}
				return ExpandConfigAwsAssumeRole(in.(map[string]interface{}))
			}(in))
		}(in)
	}
	if in, ok := in["s3_endpoint"]; ok && in != nil {
		out.S3Endpoint = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["s3_region"]; ok && in != nil {
		out.S3Region = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["s3_access_key"]; ok && in != nil {
		out.S3AccessKey = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["s3_secret_key"]; ok && in != nil {
		out.S3SecretKey = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["skip_region_check"]; ok && in != nil {
		out.SkipRegionCheck = func(in interface{}) bool { return in.(bool) }(in)
	}
	return out
}

func FlattenConfigAwsInto(in config.Aws, out map[string]interface{}) {
	out["profile"] = func(in string) interface{} { return string(in) }(in.Profile)
	out["region"] = func(in string) interface{} { return string(in) }(in.Region)
	out["access_key"] = func(in string) interface{} { return string(in) }(in.AccessKey)
	out["secret_key"] = func(in string) interface{} { return string(in) }(in.SecretKey)
	out["assume_role"] = func(in *config.AwsAssumeRole) interface{} {
		if in == nil {
			return nil
		}
		return func(in config.AwsAssumeRole) interface{} { return FlattenConfigAwsAssumeRole(in) }(*in)
	}(in.AssumeRole)
	out["s3_endpoint"] = func(in string) interface{} { return string(in) }(in.S3Endpoint)
	out["s3_region"] = func(in string) interface{} { return string(in) }(in.S3Region)
	out["s3_access_key"] = func(in string) interface{} { return string(in) }(in.S3AccessKey)
	out["s3_secret_key"] = func(in string) interface{} { return string(in) }(in.S3SecretKey)
	out["skip_region_check"] = func(in bool) interface{} { return in }(in.SkipRegionCheck)
}

func FlattenConfigAws(in config.Aws) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenConfigAwsInto(in, out)
	return out
}
