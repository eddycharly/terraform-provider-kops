package structures

import (
	"k8s.io/kops/pkg/apis/kops"
)

func ExpandIAMProfileSpec(in map[string]interface{}) kops.IAMProfileSpec {
	if in == nil {
		panic("expand IAMProfileSpec failure, in is nil")
	}
	return kops.IAMProfileSpec{
		Profile: func(in interface{}) *string {
			return func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in string) *string {
					return &in
				}(string(ExpandString(in)))
			}(in)
		}(in["profile"]),
	}
}

func FlattenIAMProfileSpec(in kops.IAMProfileSpec) map[string]interface{} {
	return map[string]interface{}{
		"profile": func(in *string) interface{} {
			return func(in *string) interface{} {
				if in == nil {
					return nil
				}
				return func(in string) interface{} {
					return FlattenString(string(in))
				}(*in)
			}(in)
		}(in.Profile),
	}
}
