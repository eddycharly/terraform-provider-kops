package schemas

import (
	"reflect"

	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	v1 "k8s.io/api/core/v1"
)

var _ = Schema

func DataSourceToleration() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"key":                ComputedString(),
			"operator":           ComputedString(),
			"value":              ComputedString(),
			"effect":             ComputedString(),
			"toleration_seconds": ComputedInt(),
		},
	}

	return res
}

func ExpandDataSourceToleration(in map[string]interface{}) v1.Toleration {
	if in == nil {
		panic("expand Toleration failure, in is nil")
	}
	return v1.Toleration{
		Key: func(in interface{}) string /**/ {
			return string(ExpandString(in))
		}(in["key"]),
		Operator: func(in interface{}) v1.TolerationOperator /*k8s.io/api/core/v1*/ {
			return v1.TolerationOperator(ExpandString(in))
		}(in["operator"]),
		Value: func(in interface{}) string /**/ {
			return string(ExpandString(in))
		}(in["value"]),
		Effect: func(in interface{}) v1.TaintEffect /*k8s.io/api/core/v1*/ {
			return v1.TaintEffect(ExpandString(in))
		}(in["effect"]),
		TolerationSeconds: func(in interface{}) *int64 /**/ {
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
		}(in["toleration_seconds"]),
	}
}

func FlattenDataSourceTolerationInto(in v1.Toleration, out map[string]interface{}) {
	out["key"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Key)
	out["operator"] = func(in v1.TolerationOperator) interface{} {
		return FlattenString(string(in))
	}(in.Operator)
	out["value"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Value)
	out["effect"] = func(in v1.TaintEffect) interface{} {
		return FlattenString(string(in))
	}(in.Effect)
	out["toleration_seconds"] = func(in *int64) interface{} {
		return func(in *int64) interface{} {
			if in == nil {
				return nil
			}
			return func(in int64) interface{} {
				return FlattenInt(int(in))
			}(*in)
		}(in)
	}(in.TolerationSeconds)
}

func FlattenDataSourceToleration(in v1.Toleration) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceTolerationInto(in, out)
	return out
}
