package schemas

import (
	"reflect"

	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

var _ = Schema

func ResourceLoadBalancerControllerSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enabled":       OptionalBool(),
			"version":       OptionalString(),
			"enable_waf":    OptionalBool(),
			"enable_wa_fv2": OptionalBool(),
			"enable_shield": OptionalBool(),
		},
	}

	return res
}

func ExpandResourceLoadBalancerControllerSpec(in map[string]interface{}) kopsv1alpha2.LoadBalancerControllerSpec {
	if in == nil {
		panic("expand LoadBalancerControllerSpec failure, in is nil")
	}
	return kopsv1alpha2.LoadBalancerControllerSpec{
		Enabled: func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in bool) *bool {
					return &in
				}(bool(ExpandBool(in)))
			}(in)
		}(in["enabled"]),
		Version: func(in interface{}) *string {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in string) *string {
					return &in
				}(string(ExpandString(in)))
			}(in)
		}(in["version"]),
		EnableWAF: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["enable_waf"]),
		EnableWAFv2: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["enable_wa_fv2"]),
		EnableShield: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["enable_shield"]),
	}
}

func FlattenResourceLoadBalancerControllerSpecInto(in kopsv1alpha2.LoadBalancerControllerSpec, out map[string]interface{}) {
	out["enabled"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.Enabled)
	out["version"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.Version)
	out["enable_waf"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.EnableWAF)
	out["enable_wa_fv2"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.EnableWAFv2)
	out["enable_shield"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.EnableShield)
}

func FlattenResourceLoadBalancerControllerSpec(in kopsv1alpha2.LoadBalancerControllerSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceLoadBalancerControllerSpecInto(in, out)
	return out
}
