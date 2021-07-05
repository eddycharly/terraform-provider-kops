package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceAWSPermission() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"policy_ar_ns":  ComputedList(String()),
			"inline_policy": ComputedString(),
		},
	}

	return res
}

func ExpandDataSourceAWSPermission(in map[string]interface{}) kops.AWSPermission {
	if in == nil {
		panic("expand AWSPermission failure, in is nil")
	}
	return kops.AWSPermission{
		PolicyARNs: func(in interface{}) []string {
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
		}(in["policy_ar_ns"]),
		InlinePolicy: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["inline_policy"]),
	}
}

func FlattenDataSourceAWSPermissionInto(in kops.AWSPermission, out map[string]interface{}) {
	out["policy_ar_ns"] = func(in []string) interface{} {
		return func(in []string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.PolicyARNs)
	out["inline_policy"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.InlinePolicy)
}

func FlattenDataSourceAWSPermission(in kops.AWSPermission) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceAWSPermissionInto(in, out)
	return out
}
