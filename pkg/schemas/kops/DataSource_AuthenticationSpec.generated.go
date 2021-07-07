package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceAuthenticationSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"kopeio": ComputedStruct(DataSourceKopeioAuthenticationSpec()),
			"aws":    ComputedStruct(DataSourceAwsAuthenticationSpec()),
		},
	}

	return res
}

func ExpandDataSourceAuthenticationSpec(in map[string]interface{}) kops.AuthenticationSpec {
	if in == nil {
		panic("expand AuthenticationSpec failure, in is nil")
	}
	return kops.AuthenticationSpec{
		Kopeio: func(in interface{}) *kops.KopeioAuthenticationSpec {
			return func(in interface{}) *kops.KopeioAuthenticationSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.KopeioAuthenticationSpec) *kops.KopeioAuthenticationSpec {
					return &in
				}(func(in interface{}) kops.KopeioAuthenticationSpec {
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return kops.KopeioAuthenticationSpec{}
					}
					return (ExpandDataSourceKopeioAuthenticationSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["kopeio"]),
		Aws: func(in interface{}) *kops.AwsAuthenticationSpec {
			return func(in interface{}) *kops.AwsAuthenticationSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.AwsAuthenticationSpec) *kops.AwsAuthenticationSpec {
					return &in
				}(func(in interface{}) kops.AwsAuthenticationSpec {
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return kops.AwsAuthenticationSpec{}
					}
					return (ExpandDataSourceAwsAuthenticationSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["aws"]),
	}
}

func FlattenDataSourceAuthenticationSpecInto(in kops.AuthenticationSpec, out map[string]interface{}) {
	out["kopeio"] = func(in *kops.KopeioAuthenticationSpec) interface{} {
		return func(in *kops.KopeioAuthenticationSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.KopeioAuthenticationSpec) interface{} {
				return func(in kops.KopeioAuthenticationSpec) []map[string]interface{} {
					return []map[string]interface{}{FlattenDataSourceKopeioAuthenticationSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Kopeio)
	out["aws"] = func(in *kops.AwsAuthenticationSpec) interface{} {
		return func(in *kops.AwsAuthenticationSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.AwsAuthenticationSpec) interface{} {
				return func(in kops.AwsAuthenticationSpec) []map[string]interface{} {
					return []map[string]interface{}{FlattenDataSourceAwsAuthenticationSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Aws)
}

func FlattenDataSourceAuthenticationSpec(in kops.AuthenticationSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceAuthenticationSpecInto(in, out)
	return out
}
