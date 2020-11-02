package structures

import (
	"log"
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
				tmp := func(in bool) *bool {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(bool(ExpandBool(in)))
				return tmp
			}(in)
			log.Printf("%s - %#v", "delete_on_termination", value)
			return value
		}(in["delete_on_termination"]),
		Device: func(in interface{}) string {
			value := string(ExpandString(in))
			log.Printf("%s - %#v", "device", value)
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
				tmp := func(in bool) *bool {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(bool(ExpandBool(in)))
				return tmp
			}(in)
			log.Printf("%s - %#v", "encrypted", value)
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
				tmp := func(in int64) *int64 {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(int64(ExpandInt(in)))
				return tmp
			}(in)
			log.Printf("%s - %#v", "iops", value)
			return value
		}(in["iops"]),
		Size: func(in interface{}) int64 {
			value := int64(ExpandInt(in))
			log.Printf("%s - %#v", "size", value)
			return value
		}(in["size"]),
		Type: func(in interface{}) string {
			value := string(ExpandString(in))
			log.Printf("%s - %#v", "type", value)
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
			log.Printf("%s - %v", "delete_on_termination", value)
			return value
		}(in.DeleteOnTermination),
		"device": func(in string) interface{} {
			value := FlattenString(string(in))
			log.Printf("%s - %v", "device", value)
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
			log.Printf("%s - %v", "encrypted", value)
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
			log.Printf("%s - %v", "iops", value)
			return value
		}(in.Iops),
		"size": func(in int64) interface{} {
			value := FlattenInt(int(in))
			log.Printf("%s - %v", "size", value)
			return value
		}(in.Size),
		"type": func(in string) interface{} {
			value := FlattenString(string(in))
			log.Printf("%s - %v", "type", value)
			return value
		}(in.Type),
	}
}
