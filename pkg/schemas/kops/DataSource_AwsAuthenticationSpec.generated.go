package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceAwsAuthenticationSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"image":          Computed(String()),
			"backend_mode":   Computed(String()),
			"cluster_id":     Computed(String()),
			"memory_request": Computed(Ptr(Quantity())),
			"cpu_request":    Computed(Ptr(Quantity())),
			"memory_limit":   Computed(Ptr(Quantity())),
			"cpu_limit":      Computed(Ptr(Quantity())),
		},
	}
}

func ExpandDataSourceAwsAuthenticationSpec(in map[string]interface{}) kops.AwsAuthenticationSpec {
	if in == nil {
		panic("expand AwsAuthenticationSpec failure, in is nil")
	}
	out := kops.AwsAuthenticationSpec{}
	if in, ok := in["image"]; ok && in != nil {
		out.Image = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["backend_mode"]; ok && in != nil {
		out.BackendMode = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["cluster_id"]; ok && in != nil {
		out.ClusterID = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["memory_request"]; ok && in != nil {
		out.MemoryRequest = func(in interface{}) *resource.Quantity {
			if in == nil {
				return nil
			}
			return func(in resource.Quantity) *resource.Quantity { return &in }(func(in interface{}) resource.Quantity { return ExpandQuantity(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["cpu_request"]; ok && in != nil {
		out.CPURequest = func(in interface{}) *resource.Quantity {
			if in == nil {
				return nil
			}
			return func(in resource.Quantity) *resource.Quantity { return &in }(func(in interface{}) resource.Quantity { return ExpandQuantity(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["memory_limit"]; ok && in != nil {
		out.MemoryLimit = func(in interface{}) *resource.Quantity {
			if in == nil {
				return nil
			}
			return func(in resource.Quantity) *resource.Quantity { return &in }(func(in interface{}) resource.Quantity { return ExpandQuantity(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["cpu_limit"]; ok && in != nil {
		out.CPULimit = func(in interface{}) *resource.Quantity {
			if in == nil {
				return nil
			}
			return func(in resource.Quantity) *resource.Quantity { return &in }(func(in interface{}) resource.Quantity { return ExpandQuantity(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	return out
}
