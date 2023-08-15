package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

var _ = Schema

func ResourceAuthenticationSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"kopeio": OptionalStruct(ResourceKopeioAuthenticationSpec()),
			"aws":    OptionalStruct(ResourceAWSAuthenticationSpec()),
		},
	}

	return res
}

func ExpandResourceAuthenticationSpec(in map[string]interface{}) kopsv1alpha2.AuthenticationSpec {
	if in == nil {
		panic("expand AuthenticationSpec failure, in is nil")
	}
	return kopsv1alpha2.AuthenticationSpec{
		Kopeio: func(in interface{}) *kopsv1alpha2.KopeioAuthenticationSpec {
			return func(in interface{}) *kopsv1alpha2.KopeioAuthenticationSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.KopeioAuthenticationSpec) *kopsv1alpha2.KopeioAuthenticationSpec {
					return &in
				}(func(in interface{}) kopsv1alpha2.KopeioAuthenticationSpec {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandResourceKopeioAuthenticationSpec(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.KopeioAuthenticationSpec{}
				}(in))
			}(in)
		}(in["kopeio"]),
		AWS: func(in interface{}) *kopsv1alpha2.AWSAuthenticationSpec {
			return func(in interface{}) *kopsv1alpha2.AWSAuthenticationSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.AWSAuthenticationSpec) *kopsv1alpha2.AWSAuthenticationSpec {
					return &in
				}(func(in interface{}) kopsv1alpha2.AWSAuthenticationSpec {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandResourceAWSAuthenticationSpec(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.AWSAuthenticationSpec{}
				}(in))
			}(in)
		}(in["aws"]),
	}
}

func FlattenResourceAuthenticationSpecInto(in kopsv1alpha2.AuthenticationSpec, out map[string]interface{}) {
	out["kopeio"] = func(in *kopsv1alpha2.KopeioAuthenticationSpec) interface{} {
		return func(in *kopsv1alpha2.KopeioAuthenticationSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.KopeioAuthenticationSpec) interface{} {
				return func(in kopsv1alpha2.KopeioAuthenticationSpec) []interface{} {
					return []interface{}{FlattenResourceKopeioAuthenticationSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Kopeio)
	out["aws"] = func(in *kopsv1alpha2.AWSAuthenticationSpec) interface{} {
		return func(in *kopsv1alpha2.AWSAuthenticationSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.AWSAuthenticationSpec) interface{} {
				return func(in kopsv1alpha2.AWSAuthenticationSpec) []interface{} {
					return []interface{}{FlattenResourceAWSAuthenticationSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.AWS)
}

func FlattenResourceAuthenticationSpec(in kopsv1alpha2.AuthenticationSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceAuthenticationSpecInto(in, out)
	return out
}
