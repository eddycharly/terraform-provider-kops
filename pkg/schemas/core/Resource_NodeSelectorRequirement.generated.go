package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	v1 "k8s.io/api/core/v1"
)

var _ = Schema

func ResourceNodeSelectorRequirement() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"key":      OptionalString(),
			"operator": OptionalString(),
			"values":   OptionalList(String()),
		},
	}

	return res
}

func ExpandResourceNodeSelectorRequirement(in map[string]interface{}) v1.NodeSelectorRequirement {
	if in == nil {
		panic("expand NodeSelectorRequirement failure, in is nil")
	}
	return v1.NodeSelectorRequirement{
		Key: func(in interface{}) string /**/ {
			return string(ExpandString(in))
		}(in["key"]),
		Operator: func(in interface{}) v1.NodeSelectorOperator /*k8s.io/api/core/v1*/ {
			return v1.NodeSelectorOperator(ExpandString(in))
		}(in["operator"]),
		Values: func(in interface{}) []string /**/ {
			return func(in interface{}) []string {
				if in == nil {
					return nil
				}
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["values"]),
	}
}

func FlattenResourceNodeSelectorRequirementInto(in v1.NodeSelectorRequirement, out map[string]interface{}) {
	out["key"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Key)
	out["operator"] = func(in v1.NodeSelectorOperator) interface{} {
		return FlattenString(string(in))
	}(in.Operator)
	out["values"] = func(in []string) interface{} {
		return func(in []string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.Values)
}

func FlattenResourceNodeSelectorRequirement(in v1.NodeSelectorRequirement) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceNodeSelectorRequirementInto(in, out)
	return out
}
