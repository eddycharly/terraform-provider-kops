package schemas

import (
	"reflect"

	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceAWSAuthenticationSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"image":             ComputedString(),
			"backend_mode":      ComputedString(),
			"cluster_id":        ComputedString(),
			"memory_request":    ComputedQuantity(),
			"cpu_request":       ComputedQuantity(),
			"memory_limit":      ComputedQuantity(),
			"cpu_limit":         ComputedQuantity(),
			"identity_mappings": ComputedList(DataSourceAWSAuthenticationIdentityMappingSpec()),
		},
	}

	return res
}

func ExpandDataSourceAWSAuthenticationSpec(in map[string]interface{}) kops.AWSAuthenticationSpec {
	if in == nil {
		panic("expand AWSAuthenticationSpec failure, in is nil")
	}
	return kops.AWSAuthenticationSpec{
		Image: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["image"]),
		BackendMode: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["backend_mode"]),
		ClusterID: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["cluster_id"]),
		MemoryRequest: func(in interface{}) *resource.Quantity {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *resource.Quantity {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in resource.Quantity) *resource.Quantity {
					return &in
				}(ExpandQuantity(in))
			}(in)
		}(in["memory_request"]),
		CPURequest: func(in interface{}) *resource.Quantity {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *resource.Quantity {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in resource.Quantity) *resource.Quantity {
					return &in
				}(ExpandQuantity(in))
			}(in)
		}(in["cpu_request"]),
		MemoryLimit: func(in interface{}) *resource.Quantity {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *resource.Quantity {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in resource.Quantity) *resource.Quantity {
					return &in
				}(ExpandQuantity(in))
			}(in)
		}(in["memory_limit"]),
		CPULimit: func(in interface{}) *resource.Quantity {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *resource.Quantity {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in resource.Quantity) *resource.Quantity {
					return &in
				}(ExpandQuantity(in))
			}(in)
		}(in["cpu_limit"]),
		IdentityMappings: func(in interface{}) []kops.AWSAuthenticationIdentityMappingSpec {
			return func(in interface{}) []kops.AWSAuthenticationIdentityMappingSpec {
				if in == nil {
					return nil
				}
				var out []kops.AWSAuthenticationIdentityMappingSpec
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) kops.AWSAuthenticationIdentityMappingSpec {
						if in == nil {
							return kops.AWSAuthenticationIdentityMappingSpec{}
						}
						return (ExpandDataSourceAWSAuthenticationIdentityMappingSpec(in.(map[string]interface{})))
					}(in))
				}
				return out
			}(in)
		}(in["identity_mappings"]),
	}
}

func FlattenDataSourceAWSAuthenticationSpecInto(in kops.AWSAuthenticationSpec, out map[string]interface{}) {
	out["image"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Image)
	out["backend_mode"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.BackendMode)
	out["cluster_id"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.ClusterID)
	out["memory_request"] = func(in *resource.Quantity) interface{} {
		return func(in *resource.Quantity) interface{} {
			if in == nil {
				return nil
			}
			return func(in resource.Quantity) interface{} {
				return FlattenQuantity(in)
			}(*in)
		}(in)
	}(in.MemoryRequest)
	out["cpu_request"] = func(in *resource.Quantity) interface{} {
		return func(in *resource.Quantity) interface{} {
			if in == nil {
				return nil
			}
			return func(in resource.Quantity) interface{} {
				return FlattenQuantity(in)
			}(*in)
		}(in)
	}(in.CPURequest)
	out["memory_limit"] = func(in *resource.Quantity) interface{} {
		return func(in *resource.Quantity) interface{} {
			if in == nil {
				return nil
			}
			return func(in resource.Quantity) interface{} {
				return FlattenQuantity(in)
			}(*in)
		}(in)
	}(in.MemoryLimit)
	out["cpu_limit"] = func(in *resource.Quantity) interface{} {
		return func(in *resource.Quantity) interface{} {
			if in == nil {
				return nil
			}
			return func(in resource.Quantity) interface{} {
				return FlattenQuantity(in)
			}(*in)
		}(in)
	}(in.CPULimit)
	out["identity_mappings"] = func(in []kops.AWSAuthenticationIdentityMappingSpec) interface{} {
		return func(in []kops.AWSAuthenticationIdentityMappingSpec) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, func(in kops.AWSAuthenticationIdentityMappingSpec) interface{} {
					return FlattenDataSourceAWSAuthenticationIdentityMappingSpec(in)
				}(in))
			}
			return out
		}(in)
	}(in.IdentityMappings)
}

func FlattenDataSourceAWSAuthenticationSpec(in kops.AWSAuthenticationSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceAWSAuthenticationSpecInto(in, out)
	return out
}
