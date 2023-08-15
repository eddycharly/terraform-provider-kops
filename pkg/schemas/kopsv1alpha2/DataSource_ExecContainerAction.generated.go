package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

var _ = Schema

func DataSourceExecContainerAction() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"image":       ComputedString(),
			"command":     ComputedList(String()),
			"environment": ComputedMap(String()),
		},
	}

	return res
}

func ExpandDataSourceExecContainerAction(in map[string]interface{}) kopsv1alpha2.ExecContainerAction {
	if in == nil {
		panic("expand ExecContainerAction failure, in is nil")
	}
	return kopsv1alpha2.ExecContainerAction{
		Image: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["image"]),
		Command: func(in interface{}) []string {
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
		}(in["command"]),
		Environment: func(in interface{}) map[string]string {
			return func(in interface{}) map[string]string {
				if in == nil {
					return nil
				}
				if in, ok := in.(map[string]interface{}); ok {
					if len(in) > 0 {
						out := map[string]string{}
						for key, in := range in {
							out[key] = string(ExpandString(in))
						}
						return out
					}
				}
				return nil
			}(in)
		}(in["environment"]),
	}
}

func FlattenDataSourceExecContainerActionInto(in kopsv1alpha2.ExecContainerAction, out map[string]interface{}) {
	out["image"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Image)
	out["command"] = func(in []string) interface{} {
		return func(in []string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.Command)
	out["environment"] = func(in map[string]string) interface{} {
		return func(in map[string]string) map[string]interface{} {
			if in == nil {
				return nil
			}
			out := map[string]interface{}{}
			for key, in := range in {
				out[key] = FlattenString(string(in))
			}
			return out
		}(in)
	}(in.Environment)
}

func FlattenDataSourceExecContainerAction(in kopsv1alpha2.ExecContainerAction) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceExecContainerActionInto(in, out)
	return out
}
