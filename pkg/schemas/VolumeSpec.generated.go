package schemas

import (
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

func VolumeSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"delete_on_termination": OptionalBool(),
			"device":                RequiredString(),
			"encrypted":             OptionalBool(),
			"iops":                  OptionalInt(),
			"size":                  OptionalInt(),
			"type":                  OptionalString(),
		},
	}
}

func ExpandVolumeSpec(in map[string]interface{}) kops.VolumeSpec {
	if in == nil {
		panic("expand VolumeSpec failure, in is nil")
	}
	return kops.VolumeSpec{
		DeleteOnTermination: func(in interface{}) *bool {
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
		}(in["delete_on_termination"]),
		Device: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["device"]),
		Encrypted: func(in interface{}) *bool {
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
		}(in["encrypted"]),
		Iops: func(in interface{}) *int64 {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *int64 {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in int64) *int64 {
					return &in
				}(int64(ExpandInt(in)))
			}(in)
		}(in["iops"]),
		Size: func(in interface{}) int64 {
			return int64(ExpandInt(in))
		}(in["size"]),
		Type: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["type"]),
	}
}

func FlattenVolumeSpecInto(in kops.VolumeSpec, out map[string]interface{}) {
	out["delete_on_termination"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.DeleteOnTermination)
	out["device"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Device)
	out["encrypted"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.Encrypted)
	out["iops"] = func(in *int64) interface{} {
		return func(in *int64) interface{} {
			if in == nil {
				return nil
			}
			return func(in int64) interface{} {
				return FlattenInt(int(in))
			}(*in)
		}(in)
	}(in.Iops)
	out["size"] = func(in int64) interface{} {
		return FlattenInt(int(in))
	}(in.Size)
	out["type"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Type)
}

func FlattenVolumeSpec(in kops.VolumeSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenVolumeSpecInto(in, out)
	return out
}
