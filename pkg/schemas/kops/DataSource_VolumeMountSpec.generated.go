package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceVolumeMountSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"device":         Computed(String()),
			"filesystem":     Computed(String()),
			"format_options": Computed(List(String())),
			"mount_options":  Computed(List(String())),
			"path":           Computed(String()),
		},
	}
}

func ExpandDataSourceVolumeMountSpec(in map[string]interface{}) kops.VolumeMountSpec {
	if in == nil {
		panic("expand VolumeMountSpec failure, in is nil")
	}
	out := kops.VolumeMountSpec{}
	if in, ok := in["device"]; ok && in != nil {
		out.Device = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["filesystem"]; ok && in != nil {
		out.Filesystem = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["format_options"]; ok && in != nil {
		out.FormatOptions = func(in interface{}) []string {
			var out []string
			for _, in := range in.([]interface{}) {
				out = append(out, func(in interface{}) string { return string(in.(string)) }(in))
			}
			return out
		}(in)
	}
	if in, ok := in["mount_options"]; ok && in != nil {
		out.MountOptions = func(in interface{}) []string {
			var out []string
			for _, in := range in.([]interface{}) {
				out = append(out, func(in interface{}) string { return string(in.(string)) }(in))
			}
			return out
		}(in)
	}
	if in, ok := in["path"]; ok && in != nil {
		out.Path = func(in interface{}) string { return string(in.(string)) }(in)
	}
	return out
}

func FlattenDataSourceVolumeMountSpecInto(in kops.VolumeMountSpec, out map[string]interface{}) {
	out["device"] = func(in string) interface{} { return string(in) }(in.Device)
	out["filesystem"] = func(in string) interface{} { return string(in) }(in.Filesystem)
	out["format_options"] = func(in []string) interface{} {
		var out []interface{}
		for _, in := range in {
			out = append(out, func(in string) interface{} { return string(in) }(in))
		}
		return out
	}(in.FormatOptions)
	out["mount_options"] = func(in []string) interface{} {
		var out []interface{}
		for _, in := range in {
			out = append(out, func(in string) interface{} { return string(in) }(in))
		}
		return out
	}(in.MountOptions)
	out["path"] = func(in string) interface{} { return string(in) }(in.Path)
}

func FlattenDataSourceVolumeMountSpec(in kops.VolumeMountSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceVolumeMountSpecInto(in, out)
	return out
}
