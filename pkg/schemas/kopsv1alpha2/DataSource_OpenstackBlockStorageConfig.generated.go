package schemas

import (
	"reflect"

	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

var _ = Schema

func DataSourceOpenstackBlockStorageConfig() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"version":                     ComputedString(),
			"ignore_az":                   ComputedBool(),
			"override_az":                 ComputedString(),
			"ignore_volume_micro_version": ComputedBool(),
			"create_storage_class":        ComputedBool(),
			"csi_plugin_image":            ComputedString(),
			"csi_topology_support":        ComputedBool(),
			"cluster_name":                ComputedString(),
		},
	}

	return res
}

func ExpandDataSourceOpenstackBlockStorageConfig(in map[string]interface{}) kopsv1alpha2.OpenstackBlockStorageConfig {
	if in == nil {
		panic("expand OpenstackBlockStorageConfig failure, in is nil")
	}
	return kopsv1alpha2.OpenstackBlockStorageConfig{
		Version: func(in interface{}) *string {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
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
		}(in["version"]),
		IgnoreAZ: func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in bool) *bool {
					return &in
				}(bool(ExpandBool(in)))
			}(in)
		}(in["ignore_az"]),
		OverrideAZ: func(in interface{}) *string {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
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
		}(in["override_az"]),
		IgnoreVolumeMicroVersion: func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in bool) *bool {
					return &in
				}(bool(ExpandBool(in)))
			}(in)
		}(in["ignore_volume_micro_version"]),
		CreateStorageClass: func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in bool) *bool {
					return &in
				}(bool(ExpandBool(in)))
			}(in)
		}(in["create_storage_class"]),
		CSIPluginImage: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["csi_plugin_image"]),
		CSITopologySupport: func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in bool) *bool {
					return &in
				}(bool(ExpandBool(in)))
			}(in)
		}(in["csi_topology_support"]),
		ClusterName: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["cluster_name"]),
	}
}

func FlattenDataSourceOpenstackBlockStorageConfigInto(in kopsv1alpha2.OpenstackBlockStorageConfig, out map[string]interface{}) {
	out["version"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.Version)
	out["ignore_az"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.IgnoreAZ)
	out["override_az"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.OverrideAZ)
	out["ignore_volume_micro_version"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.IgnoreVolumeMicroVersion)
	out["create_storage_class"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.CreateStorageClass)
	out["csi_plugin_image"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.CSIPluginImage)
	out["csi_topology_support"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.CSITopologySupport)
	out["cluster_name"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.ClusterName)
}

func FlattenDataSourceOpenstackBlockStorageConfig(in kopsv1alpha2.OpenstackBlockStorageConfig) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceOpenstackBlockStorageConfigInto(in, out)
	return out
}
