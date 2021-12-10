package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ = Schema

func ResourceLabelSelectorRequirement() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"key":      OptionalString(),
			"operator": OptionalString(),
			"values":   OptionalList(String()),
		},
	}

	return res
}

func ExpandResourceLabelSelectorRequirement(in map[string]interface{}) v1.LabelSelectorRequirement {
	if in == nil {
		panic("expand LabelSelectorRequirement failure, in is nil")
	}
	return v1.LabelSelectorRequirement{
		Key: func(in interface{}) string /**/ {
			return string(ExpandString(in))
		}(in["key"]),
		Operator: func(in interface{}) v1.LabelSelectorOperator /*k8s.io/apimachinery/pkg/apis/meta/v1*/ {
			return v1.LabelSelectorOperator(ExpandString(in))
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

func FlattenResourceLabelSelectorRequirementInto(in v1.LabelSelectorRequirement, out map[string]interface{}) {
	out["key"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Key)
	out["operator"] = func(in v1.LabelSelectorOperator) interface{} {
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

func FlattenResourceLabelSelectorRequirement(in v1.LabelSelectorRequirement) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceLabelSelectorRequirementInto(in, out)
	return out
}
