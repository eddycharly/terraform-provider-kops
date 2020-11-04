package structures

import (
	"k8s.io/kops/pkg/apis/kops"
)

func ExpandExecContainerAction(in map[string]interface{}) kops.ExecContainerAction {
	if in == nil {
		panic("expand ExecContainerAction failure, in is nil")
	}
	return kops.ExecContainerAction{
		Image: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["image"]),
		Command: func(in interface{}) []string {
			return func(in interface{}) []string {
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
				out := map[string]string{}
				for key, in := range in.(map[string]interface{}) {
					out[key] = string(ExpandString(in))
				}
				return out
			}(in)
		}(in["environment"]),
	}
}

func FlattenExecContainerAction(in kops.ExecContainerAction) map[string]interface{} {
	return map[string]interface{}{
		"image": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.Image),
		"command": func(in []string) interface{} {
			return func(in []string) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, FlattenString(string(in)))
				}
				return out
			}(in)
		}(in.Command),
		"environment": func(in map[string]string) interface{} {
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
		}(in.Environment),
	}
}
