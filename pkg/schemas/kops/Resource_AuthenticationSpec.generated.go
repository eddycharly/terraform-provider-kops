package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceAuthenticationSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"kopeio": Optional(Struct(ResourceKopeioAuthenticationSpec())),
			"aws":    Optional(Struct(ResourceAwsAuthenticationSpec())),
		},
	}
}

func ExpandResourceAuthenticationSpec(in map[string]interface{}) kops.AuthenticationSpec {
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
					return ExpandResourceKopeioAuthenticationSpec(in.(map[string]interface{}))
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
					return ExpandResourceAwsAuthenticationSpec(in.(map[string]interface{}))
				}(in[0]))
			}
			return nil
		}(in)
	}
	return out
}

func FlattenResourceAuthenticationSpecInto(in kops.AuthenticationSpec, out map[string]interface{}) {
	out["kopeio"] = func(in *kops.KopeioAuthenticationSpec) interface{} {
		if in == nil {
			return nil
		}
		return func(in kops.KopeioAuthenticationSpec) interface{} { return FlattenResourceKopeioAuthenticationSpec(in) }(*in)
	}(in.Kopeio)
	out["aws"] = func(in *kops.AwsAuthenticationSpec) interface{} {
		if in == nil {
			return nil
		}
		return func(in kops.AwsAuthenticationSpec) interface{} { return FlattenResourceAwsAuthenticationSpec(in) }(*in)
	}(in.Aws)
}

func FlattenResourceAuthenticationSpec(in kops.AuthenticationSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceAuthenticationSpecInto(in, out)
	return out
}
