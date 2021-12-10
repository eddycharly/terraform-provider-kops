package schemas

import (
	"reflect"

	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceWarmPoolSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"min_size":              ComputedInt(),
			"max_size":              ComputedInt(),
			"enable_lifecycle_hook": ComputedBool(),
		},
	}

	return res
}

func ExpandDataSourceWarmPoolSpec(in map[string]interface{}) kops.WarmPoolSpec {
	if in == nil {
		panic("expand WarmPoolSpec failure, in is nil")
	}
	return kops.WarmPoolSpec{
		MinSize: func(in interface{}) int64 /**/ {
			return int64(ExpandInt(in))
		}(in["min_size"]),
		MaxSize: func(in interface{}) *int64 /**/ {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *int64 {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in int64) *int64 {
					return &in
				}(int64(ExpandInt(in)))
			}(in)
		}(in["max_size"]),
		EnableLifecycleHook: func(in interface{}) bool /**/ {
			return bool(ExpandBool(in))
		}(in["enable_lifecycle_hook"]),
	}
}

func FlattenDataSourceWarmPoolSpecInto(in kops.WarmPoolSpec, out map[string]interface{}) {
	out["min_size"] = func(in int64) interface{} {
		return FlattenInt(int(in))
	}(in.MinSize)
	out["max_size"] = func(in *int64) interface{} {
		return func(in *int64) interface{} {
			if in == nil {
				return nil
			}
			return func(in int64) interface{} {
				return FlattenInt(int(in))
			}(*in)
		}(in)
	}(in.MaxSize)
	out["enable_lifecycle_hook"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.EnableLifecycleHook)
}

func FlattenDataSourceWarmPoolSpec(in kops.WarmPoolSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceWarmPoolSpecInto(in, out)
	return out
}
