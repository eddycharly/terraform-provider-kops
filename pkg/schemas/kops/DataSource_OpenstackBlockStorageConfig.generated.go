package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceOpenstackBlockStorageConfig() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"version":              Computed(Ptr(String())),
			"ignore_az":            Computed(Ptr(Bool())),
			"override_az":          Computed(Ptr(String())),
			"create_storage_class": Computed(Ptr(Bool())),
			"csi_plugin_image":     Computed(String()),
			"csi_topology_support": Computed(Ptr(Bool())),
		},
	}
}

func ExpandDataSourceOpenstackBlockStorageConfig(in map[string]interface{}) kops.OpenstackBlockStorageConfig {
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
