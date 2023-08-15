package schemas

import (
	"reflect"

	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/apimachinery/pkg/api/resource"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

var _ = Schema

func ResourceAWSAuthenticationSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"image":             OptionalString(),
			"backend_mode":      OptionalString(),
			"cluster_id":        OptionalString(),
			"memory_request":    OptionalQuantity(),
			"cpu_request":       OptionalQuantity(),
			"memory_limit":      OptionalQuantity(),
			"cpu_limit":         OptionalQuantity(),
			"identity_mappings": OptionalList(ResourceAWSAuthenticationIdentityMappingSpec()),
		},
	}

	return res
}

func ExpandResourceAWSAuthenticationSpec(in map[string]interface{}) kopsv1alpha2.AWSAuthenticationSpec {
	if in == nil {
		panic("expand AWSAuthenticationSpec failure, in is nil")
	}
	return kopsv1alpha2.AWSAuthenticationSpec{
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
		IdentityMappings: func(in interface{}) []kopsv1alpha2.AWSAuthenticationIdentityMappingSpec {
			return func(in interface{}) []kopsv1alpha2.AWSAuthenticationIdentityMappingSpec {
				if in == nil {
					return nil
				}
				var out []kopsv1alpha2.AWSAuthenticationIdentityMappingSpec
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) kopsv1alpha2.AWSAuthenticationIdentityMappingSpec {
						if in == nil {
							return kopsv1alpha2.AWSAuthenticationIdentityMappingSpec{}
						}
						return ExpandResourceAWSAuthenticationIdentityMappingSpec(in.(map[string]interface{}))
					}(in))
				}
				return out
			}(in)
		}(in["identity_mappings"]),
	}
}

func FlattenResourceAWSAuthenticationSpecInto(in kopsv1alpha2.AWSAuthenticationSpec, out map[string]interface{}) {
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
	out["identity_mappings"] = func(in []kopsv1alpha2.AWSAuthenticationIdentityMappingSpec) interface{} {
		return func(in []kopsv1alpha2.AWSAuthenticationIdentityMappingSpec) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, func(in kopsv1alpha2.AWSAuthenticationIdentityMappingSpec) interface{} {
					return FlattenResourceAWSAuthenticationIdentityMappingSpec(in)
				}(in))
			}
			return out
		}(in)
	}(in.IdentityMappings)
}

func FlattenResourceAWSAuthenticationSpec(in kopsv1alpha2.AWSAuthenticationSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceAWSAuthenticationSpecInto(in, out)
	return out
}
