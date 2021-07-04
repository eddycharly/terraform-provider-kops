package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceLoadBalancer() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"load_balancer_name": Computed(Nullable(String())),
			"target_group_arn":   Computed(Nullable(String())),
		},
	}
}

func ExpandDataSourceLoadBalancer(in map[string]interface{}) kops.LoadBalancer {
	if in == nil {
		panic("expand LoadBalancer failure, in is nil")
	}
	out := kops.LoadBalancer{}
	if in, ok := in["load_balancer_name"]; ok && in != nil {
		out.LoadBalancerName = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["target_group_arn"]; ok && in != nil {
		out.TargetGroupARN = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	return out
}

func FlattenDataSourceLoadBalancerInto(in kops.LoadBalancer, out map[string]interface{}) {
	out["load_balancer_name"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.LoadBalancerName)
	out["target_group_arn"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.TargetGroupARN)
}

func FlattenDataSourceLoadBalancer(in kops.LoadBalancer) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceLoadBalancerInto(in, out)
	return out
}
