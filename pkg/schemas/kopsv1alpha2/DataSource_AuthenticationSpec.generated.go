package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

var _ = Schema

func DataSourceAuthenticationSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"kopeio": ComputedStruct(DataSourceKopeioAuthenticationSpec()),
			"aws":    ComputedStruct(DataSourceAWSAuthenticationSpec()),
		},
	}

	return res
}

func ExpandDataSourceAuthenticationSpec(in map[string]interface{}) kopsv1alpha2.AuthenticationSpec {
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
						return ExpandDataSourceKopeioAuthenticationSpec(in[0].(map[string]interface{}))
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
						return ExpandDataSourceAWSAuthenticationSpec(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.AWSAuthenticationSpec{}
				}(in))
			}(in)
		}(in["aws"]),
	}
}

func FlattenDataSourceAuthenticationSpecInto(in kopsv1alpha2.AuthenticationSpec, out map[string]interface{}) {
	out["kopeio"] = func(in *kopsv1alpha2.KopeioAuthenticationSpec) interface{} {
		return func(in *kopsv1alpha2.KopeioAuthenticationSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.KopeioAuthenticationSpec) interface{} {
				return func(in kopsv1alpha2.KopeioAuthenticationSpec) []interface{} {
					return []interface{}{FlattenDataSourceKopeioAuthenticationSpec(in)}
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
					return []interface{}{FlattenDataSourceAWSAuthenticationSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.AWS)
}

func FlattenDataSourceAuthenticationSpec(in kopsv1alpha2.AuthenticationSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceAuthenticationSpecInto(in, out)
	return out
}
