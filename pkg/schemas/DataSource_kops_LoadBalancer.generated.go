package schemas

import (
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

func DataSourceKopsLoadBalancer() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"load_balancer_name": ComputedString(),
			"target_group_arn":   ComputedString(),
		},
	}
}

func ExpandDataSourceKopsLoadBalancer(in map[string]interface{}) kops.LoadBalancer {
	if in == nil {
		panic("expand LoadBalancer failure, in is nil")
	}
	return kops.LoadBalancer{
		LoadBalancerName: func(in interface{}) *string {
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

func FlattenDataSourceKopsLoadBalancerInto(in kops.LoadBalancer, out map[string]interface{}) {
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

func FlattenDataSourceKopsLoadBalancer(in kops.LoadBalancer) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceKopsLoadBalancerInto(in, out)
	return out
}
