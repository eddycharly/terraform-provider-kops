package structures

import (
	"reflect"

	"k8s.io/kops/pkg/apis/kops"
)

func ExpandLoadBalancer(in map[string]interface{}) kops.LoadBalancer {
	if in == nil {
		panic("expand LoadBalancer failure, in is nil")
	}
	return kops.LoadBalancer{
		LoadBalancerName: func(in interface{}) *string {
			value := func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in string) *string {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(string(ExpandString(in)))
			}(in)
			return value
		}(in["load_balancer_name"]),
		TargetGroupARN: func(in interface{}) *string {
			value := func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in string) *string {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(string(ExpandString(in)))
			}(in)
			return value
		}(in["target_group_arn"]),
	}
}

func FlattenLoadBalancer(in kops.LoadBalancer) map[string]interface{} {
	return map[string]interface{}{
		"load_balancer_name": func(in *string) interface{} {
			value := func(in *string) interface{} {
				if in == nil {
					return nil
				}
				return func(in string) interface{} {
					return FlattenString(string(in))
				}(*in)
			}(in)
			return value
		}(in.LoadBalancerName),
		"target_group_arn": func(in *string) interface{} {
			value := func(in *string) interface{} {
				if in == nil {
					return nil
				}
				return func(in string) interface{} {
					return FlattenString(string(in))
				}(*in)
			}(in)
			return value
		}(in.TargetGroupARN),
	}
}
