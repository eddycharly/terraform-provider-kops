package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceAWSAuthenticationIdentityMappingSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"arn":      ComputedString(),
			"username": ComputedString(),
			"groups":   ComputedList(String()),
		},
	}

	return res
}

func ExpandDataSourceAWSAuthenticationIdentityMappingSpec(in map[string]interface{}) kops.AWSAuthenticationIdentityMappingSpec {
	if in == nil {
		panic("expand AWSAuthenticationIdentityMappingSpec failure, in is nil")
	}
	return kops.AWSAuthenticationIdentityMappingSpec{
		ARN: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["arn"]),
		Username: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["username"]),
		Groups: func(in interface{}) []string {
			return func(in interface{}) []string {
				if in == nil {
					return nil
				}
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["groups"]),
	}
}

func FlattenDataSourceAWSAuthenticationIdentityMappingSpecInto(in kops.AWSAuthenticationIdentityMappingSpec, out map[string]interface{}) {
	out["arn"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.ARN)
	out["username"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Username)
	out["groups"] = func(in []string) interface{} {
		return func(in []string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.Groups)
}

func FlattenDataSourceAWSAuthenticationIdentityMappingSpec(in kops.AWSAuthenticationIdentityMappingSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceAWSAuthenticationIdentityMappingSpecInto(in, out)
	return out
}
