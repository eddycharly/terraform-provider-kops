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
			"hash_amd64": Computed(Ptr(String())),
			"hash_arm64": Computed(Ptr(String())),
			"url_amd64":  Computed(Ptr(String())),
			"url_arm64":  Computed(Ptr(String())),
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
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["hash_arm64"]; ok && in != nil {
		out.HashArm64 = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["url_amd64"]; ok && in != nil {
		out.UrlAmd64 = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["url_arm64"]; ok && in != nil {
		out.UrlArm64 = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	return out
}
