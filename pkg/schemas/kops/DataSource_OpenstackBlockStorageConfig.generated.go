package schemas

import (
	"reflect"

	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceOpenstackBlockStorageConfig() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"version":     ComputedString(),
			"ignore_az":   ComputedBool(),
			"override_az": ComputedString(),
		},
	}
}

func ExpandDataSourceOpenstackBlockStorageConfig(in map[string]interface{}) kops.OpenstackBlockStorageConfig {
	if in == nil {
		panic("expand OpenstackBlockStorageConfig failure, in is nil")
	}
	return kops.OpenstackBlockStorageConfig{
		Version: func(in interface{}) *string {
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
	}
}

func FlattenDataSourceOpenstackBlockStorageConfigInto(in kops.OpenstackBlockStorageConfig, out map[string]interface{}) {
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
}

func FlattenDataSourceOpenstackBlockStorageConfig(in kops.OpenstackBlockStorageConfig) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceOpenstackBlockStorageConfigInto(in, out)
	return out
}
