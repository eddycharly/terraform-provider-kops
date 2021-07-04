package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourcePackagesConfig() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"hash_amd64": Computed(Nullable(String())),
			"hash_arm64": Computed(Nullable(String())),
			"url_amd64":  Computed(Nullable(String())),
			"url_arm64":  Computed(Nullable(String())),
		},
	}
}

func ExpandDataSourcePackagesConfig(in map[string]interface{}) kops.PackagesConfig {
	if in == nil {
		panic("expand PackagesConfig failure, in is nil")
	}
	out := kops.PackagesConfig{}
	if in, ok := in["hash_amd64"]; ok && in != nil {
		out.HashAmd64 = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["hash_arm64"]; ok && in != nil {
		out.HashArm64 = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["url_amd64"]; ok && in != nil {
		out.UrlAmd64 = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["url_arm64"]; ok && in != nil {
		out.UrlArm64 = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	return out
}

func FlattenDataSourcePackagesConfigInto(in kops.PackagesConfig, out map[string]interface{}) {
	out["hash_amd64"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.HashAmd64)
	out["hash_arm64"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.HashArm64)
	out["url_amd64"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.UrlAmd64)
	out["url_arm64"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.UrlArm64)
}

func FlattenDataSourcePackagesConfig(in kops.PackagesConfig) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourcePackagesConfigInto(in, out)
	return out
}
