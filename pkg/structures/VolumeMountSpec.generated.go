package structures

import (
	"k8s.io/kops/pkg/apis/kops"
)

func ExpandVolumeMountSpec(in map[string]interface{}) kops.VolumeMountSpec {
	if in == nil {
		panic("expand VolumeMountSpec failure, in is nil")
	}
	return kops.VolumeMountSpec{
		Device: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["device"]),
		Filesystem: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["filesystem"]),
		FormatOptions: func(in interface{}) []string {
			return func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["format_options"]),
		MountOptions: func(in interface{}) []string {
			return func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["mount_options"]),
		Path: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["path"]),
	}
}

func FlattenVolumeMountSpecInto(in kops.VolumeMountSpec, out map[string]interface{}) {
	out["device"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Device)
	out["filesystem"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Filesystem)
	out["format_options"] = func(in []string) interface{} {
		return func(in []string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.FormatOptions)
	out["mount_options"] = func(in []string) interface{} {
		return func(in []string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.MountOptions)
	out["path"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Path)
}

func FlattenVolumeMountSpec(in kops.VolumeMountSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenVolumeMountSpecInto(in, out)
	return out
}
