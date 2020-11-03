package structures

import (
	"reflect"

	"k8s.io/kops/pkg/apis/kops"
)

func ExpandVolumeSpec(in map[string]interface{}) kops.VolumeSpec {
	if in == nil {
		panic("expand VolumeSpec failure, in is nil")
	}
	return kops.VolumeSpec{
		DeleteOnTermination: func(in interface{}) *bool {
			value := func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in bool) *bool {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(bool(ExpandBool(in)))
			}(in)
			return value
		}(in["delete_on_termination"]),
		Device: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["device"]),
		Encrypted: func(in interface{}) *bool {
			value := func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in bool) *bool {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(bool(ExpandBool(in)))
			}(in)
			return value
		}(in["encrypted"]),
		Iops: func(in interface{}) *int64 {
			value := func(in interface{}) *int64 {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in int64) *int64 {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(int64(ExpandInt(in)))
			}(in)
			return value
		}(in["iops"]),
		Size: func(in interface{}) int64 {
			value := int64(ExpandInt(in))
			return value
		}(in["size"]),
		Type: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["type"]),
	}
}

func FlattenVolumeSpec(in kops.VolumeSpec) map[string]interface{} {
	return map[string]interface{}{
		"delete_on_termination": func(in *bool) interface{} {
			value := func(in *bool) interface{} {
				if in == nil {
					return nil
				}
				return func(in bool) interface{} {
					return FlattenBool(bool(in))
				}(*in)
			}(in)
			return value
		}(in.DeleteOnTermination),
		"device": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.Device),
		"encrypted": func(in *bool) interface{} {
			value := func(in *bool) interface{} {
				if in == nil {
					return nil
				}
				return func(in bool) interface{} {
					return FlattenBool(bool(in))
				}(*in)
			}(in)
			return value
		}(in.Encrypted),
		"iops": func(in *int64) interface{} {
			value := func(in *int64) interface{} {
				if in == nil {
					return nil
				}
				return func(in int64) interface{} {
					return FlattenInt(int(in))
				}(*in)
			}(in)
			return value
		}(in.Iops),
		"size": func(in int64) interface{} {
			value := FlattenInt(int(in))
			return value
		}(in.Size),
		"type": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.Type),
	}
}
