package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceAuthenticationSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"kopeio": Computed(Struct(DataSourceKopeioAuthenticationSpec())),
			"aws":    Computed(Struct(DataSourceAwsAuthenticationSpec())),
		},
	}
}

func ExpandDataSourceAuthenticationSpec(in map[string]interface{}) kops.AuthenticationSpec {
	if in == nil {
		panic("expand AuthenticationSpec failure, in is nil")
	}
	out := kops.AuthenticationSpec{}
	if in, ok := in["kopeio"]; ok && in != nil {
		out.Kopeio = func(in interface{}) *kops.KopeioAuthenticationSpec {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in kops.KopeioAuthenticationSpec) *kops.KopeioAuthenticationSpec { return &in }(func(in interface{}) kops.KopeioAuthenticationSpec {
					return ExpandDataSourceKopeioAuthenticationSpec(in.(map[string]interface{}))
				}(in[0]))
			}
			return nil
		}(in)
	}
	if in, ok := in["aws"]; ok && in != nil {
		out.Aws = func(in interface{}) *kops.AwsAuthenticationSpec {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in kops.AwsAuthenticationSpec) *kops.AwsAuthenticationSpec { return &in }(func(in interface{}) kops.AwsAuthenticationSpec {
					return ExpandDataSourceAwsAuthenticationSpec(in.(map[string]interface{}))
				}(in[0]))
			}
			return nil
		}(in)
	}
	return out
}

func FlattenDataSourceAuthenticationSpecInto(in kops.AuthenticationSpec, out map[string]interface{}) {
	out["kopeio"] = func(in *kops.KopeioAuthenticationSpec) interface{} {
		if in == nil {
			return nil
		}
		return func(in kops.KopeioAuthenticationSpec) interface{} {
			return FlattenDataSourceKopeioAuthenticationSpec(in)
		}(*in)
	}(in.Kopeio)
	out["aws"] = func(in *kops.AwsAuthenticationSpec) interface{} {
		if in == nil {
			return nil
		}
		return func(in kops.AwsAuthenticationSpec) interface{} { return FlattenDataSourceAwsAuthenticationSpec(in) }(*in)
	}(in.Aws)
}

func FlattenDataSourceAuthenticationSpec(in kops.AuthenticationSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceAuthenticationSpecInto(in, out)
	return out
}
