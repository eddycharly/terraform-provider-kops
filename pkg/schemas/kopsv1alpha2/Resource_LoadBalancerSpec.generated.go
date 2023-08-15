package schemas

import (
	"reflect"

	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

var _ = Schema

func ResourceLoadBalancerSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"load_balancer_name": OptionalString(),
			"target_group_arn":   OptionalString(),
		},
	}

	return res
}

func ExpandResourceLoadBalancerSpec(in map[string]interface{}) kopsv1alpha2.LoadBalancerSpec {
	if in == nil {
		panic("expand LoadBalancerSpec failure, in is nil")
	}
	return kopsv1alpha2.LoadBalancerSpec{
		LoadBalancerName: func(in interface{}) *string {
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
		}(in["load_balancer_name"]),
		TargetGroupARN: func(in interface{}) *string {
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
		}(in["target_group_arn"]),
	}
}

func FlattenResourceLoadBalancerSpecInto(in kopsv1alpha2.LoadBalancerSpec, out map[string]interface{}) {
	out["load_balancer_name"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.LoadBalancerName)
	out["target_group_arn"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.TargetGroupARN)
}

func FlattenResourceLoadBalancerSpec(in kopsv1alpha2.LoadBalancerSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceLoadBalancerSpecInto(in, out)
	return out
}
