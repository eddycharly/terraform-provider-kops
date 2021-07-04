package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceExecContainerAction() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"image":       Required(String()),
			"command":     Optional(List(String())),
			"environment": Optional(Map(String())),
		},
	}
}

func ExpandResourceExecContainerAction(in map[string]interface{}) kops.ExecContainerAction {
	if in == nil {
		panic("expand ExecContainerAction failure, in is nil")
	}
	out := kops.ExecContainerAction{}
	if in, ok := in["image"]; ok && in != nil {
		out.Image = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["command"]; ok && in != nil {
		out.Command = func(in interface{}) []string {
			var out []string
			for _, in := range in.([]interface{}) {
				out = append(out, func(in interface{}) string { return string(in.(string)) }(in))
			}
			return out
		}(in)
	}
	if in, ok := in["environment"]; ok && in != nil {
		out.Environment = func(in interface{}) map[string]string {
			if in == nil {
				return nil
			}
			out := map[string]string{}
			for key, in := range in.(map[string]interface{}) {
				out[key] = func(in interface{}) string { return string(in.(string)) }(in)
			}
			return out
		}(in)
	}
	return out
}

func FlattenResourceExecContainerActionInto(in kops.ExecContainerAction, out map[string]interface{}) {
	out["image"] = func(in string) interface{} { return string(in) }(in.Image)
	out["command"] = func(in []string) interface{} {
		var out []interface{}
		for _, in := range in {
			out = append(out, func(in string) interface{} { return string(in) }(in))
		}
		return out
	}(in.Command)
	out["environment"] = func(in map[string]string) interface{} {
		if in == nil {
			return nil
		}
		out := map[string]interface{}{}
		for key, in := range in {
			out[key] = func(in string) interface{} { return string(in) }(in)
		}
		return out
	}(in.Environment)
}

func FlattenResourceExecContainerAction(in kops.ExecContainerAction) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceExecContainerActionInto(in, out)
	return out
}
