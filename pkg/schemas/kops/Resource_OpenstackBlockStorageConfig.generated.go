package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceOpenstackBlockStorageConfig() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"version":              Optional(Nullable(String())),
			"ignore_az":            Optional(Nullable(Bool())),
			"override_az":          Optional(Nullable(String())),
			"create_storage_class": Optional(Nullable(Bool())),
			"csi_plugin_image":     Optional(String()),
			"csi_topology_support": Optional(Nullable(Bool())),
		},
	}
}

func ExpandResourceOpenstackBlockStorageConfig(in map[string]interface{}) kops.OpenstackBlockStorageConfig {
	if in == nil {
		panic("expand OpenstackBlockStorageConfig failure, in is nil")
	}
	out := kops.OpenstackBlockStorageConfig{}
	if in, ok := in["version"]; ok && in != nil {
		out.Version = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["ignore_az"]; ok && in != nil {
		out.IgnoreAZ = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["override_az"]; ok && in != nil {
		out.OverrideAZ = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["create_storage_class"]; ok && in != nil {
		out.CreateStorageClass = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["csi_plugin_image"]; ok && in != nil {
		out.CSIPluginImage = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["csi_topology_support"]; ok && in != nil {
		out.CSITopologySupport = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	return out
}

func FlattenResourceOpenstackBlockStorageConfigInto(in kops.OpenstackBlockStorageConfig, out map[string]interface{}) {
	out["version"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.Version)
	out["ignore_az"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in bool) interface{} { return in }(*in)}
	}(in.IgnoreAZ)
	out["override_az"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.OverrideAZ)
	out["create_storage_class"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in bool) interface{} { return in }(*in)}
	}(in.CreateStorageClass)
	out["csi_plugin_image"] = func(in string) interface{} { return string(in) }(in.CSIPluginImage)
	out["csi_topology_support"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in bool) interface{} { return in }(*in)}
	}(in.CSITopologySupport)
}

func FlattenResourceOpenstackBlockStorageConfig(in kops.OpenstackBlockStorageConfig) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceOpenstackBlockStorageConfigInto(in, out)
	return out
}
