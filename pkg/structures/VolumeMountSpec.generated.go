package structures

import (
	"log"

	"k8s.io/kops/pkg/apis/kops"
)

func ExpandVolumeMountSpec(in map[string]interface{}) kops.VolumeMountSpec {
	if in == nil {
		panic("expand VolumeMountSpec failure, in is nil")
	}
	return kops.VolumeMountSpec{
		Device: func(in interface{}) string {
			value := string(ExpandString(in))
			log.Printf("%s - %#v", "device", value)
			return value
		}(in["device"]),
		Filesystem: func(in interface{}) string {
			value := string(ExpandString(in))
			log.Printf("%s - %#v", "filesystem", value)
			return value
		}(in["filesystem"]),
		FormatOptions: func(in interface{}) []string {
			value := func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
			log.Printf("%s - %#v", "format_options", value)
			return value
		}(in["format_options"]),
		MountOptions: func(in interface{}) []string {
			value := func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
			log.Printf("%s - %#v", "mount_options", value)
			return value
		}(in["mount_options"]),
		Path: func(in interface{}) string {
			value := string(ExpandString(in))
			log.Printf("%s - %#v", "path", value)
			return value
		}(in["path"]),
	}
}

func FlattenVolumeMountSpec(in kops.VolumeMountSpec) map[string]interface{} {
	return map[string]interface{}{
		"device": func(in string) interface{} {
			value := FlattenString(string(in))
			log.Printf("%s - %v", "device", value)
			return value
		}(in.Device),
		"filesystem": func(in string) interface{} {
			value := FlattenString(string(in))
			log.Printf("%s - %v", "filesystem", value)
			return value
		}(in.Filesystem),
		"format_options": func(in []string) interface{} {
			value := func(in []string) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, FlattenString(string(in)))
				}
				return out
			}(in)
			log.Printf("%s - %v", "format_options", value)
			return value
		}(in.FormatOptions),
		"mount_options": func(in []string) interface{} {
			value := func(in []string) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, FlattenString(string(in)))
				}
				return out
			}(in)
			log.Printf("%s - %v", "mount_options", value)
			return value
		}(in.MountOptions),
		"path": func(in string) interface{} {
			value := FlattenString(string(in))
			log.Printf("%s - %v", "path", value)
			return value
		}(in.Path),
	}
}
