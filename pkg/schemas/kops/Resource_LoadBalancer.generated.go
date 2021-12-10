package schemas

import (
	"reflect"

	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceLoadBalancer() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"load_balancer_name": OptionalString(),
			"target_group_arn":   OptionalString(),
		},
	}

	return res
}

func ExpandResourceLoadBalancer(in map[string]interface{}) kops.LoadBalancer {
	if in == nil {
		panic("expand LoadBalancer failure, in is nil")
	}
	return kops.LoadBalancer{
		LoadBalancerName: func(in interface{}) *string /**/ {
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
		TargetGroupARN: func(in interface{}) *string /**/ {
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

func FlattenResourceLoadBalancerInto(in kops.LoadBalancer, out map[string]interface{}) {
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

func FlattenResourceLoadBalancer(in kops.LoadBalancer) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceLoadBalancerInto(in, out)
	return out
}
