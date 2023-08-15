package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

var _ = Schema

func ResourceAWSAuthenticationIdentityMappingSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"arn":      OptionalString(),
			"username": OptionalString(),
			"groups":   OptionalList(String()),
		},
	}

	return res
}

func ExpandResourceAWSAuthenticationIdentityMappingSpec(in map[string]interface{}) kopsv1alpha2.AWSAuthenticationIdentityMappingSpec {
	if in == nil {
		panic("expand AWSAuthenticationIdentityMappingSpec failure, in is nil")
	}
	return kopsv1alpha2.AWSAuthenticationIdentityMappingSpec{
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

func FlattenResourceAWSAuthenticationIdentityMappingSpecInto(in kopsv1alpha2.AWSAuthenticationIdentityMappingSpec, out map[string]interface{}) {
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

func FlattenResourceAWSAuthenticationIdentityMappingSpec(in kopsv1alpha2.AWSAuthenticationIdentityMappingSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceAWSAuthenticationIdentityMappingSpecInto(in, out)
	return out
}
