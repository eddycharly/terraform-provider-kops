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
			"kopeio": Optional(Ptr(Struct(ResourceKopeioAuthenticationSpec()))),
			"aws":    Optional(Ptr(Struct(ResourceAwsAuthenticationSpec()))),
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
			return func(in kops.KopeioAuthenticationSpec) *kops.KopeioAuthenticationSpec { return &in }(func(in interface{}) kops.KopeioAuthenticationSpec {
				if in == nil {
					return kops.KopeioAuthenticationSpec{}
				}
				return ExpandResourceKopeioAuthenticationSpec(in.(map[string]interface{}))
			}(in))
		}(in)
	}
	if in, ok := in["aws"]; ok && in != nil {
		out.Aws = func(in interface{}) *kops.AwsAuthenticationSpec {
			if in == nil {
				return nil
			}
			return func(in kops.AwsAuthenticationSpec) *kops.AwsAuthenticationSpec { return &in }(func(in interface{}) kops.AwsAuthenticationSpec {
				if in == nil {
					return kops.AwsAuthenticationSpec{}
				}
				return ExpandResourceAwsAuthenticationSpec(in.(map[string]interface{}))
			}(in))
		}(in)
	}
	return out
}
