package structures

import (
	"k8s.io/kops/pkg/apis/kops"
)

func ExpandEtcdMemberSpec(in map[string]interface{}) kops.EtcdMemberSpec {
	if in == nil {
		panic("expand EtcdMemberSpec failure, in is nil")
	}
	return kops.EtcdMemberSpec{
		Name: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["name"]),
		InstanceGroup: func(in interface{}) *string {
			value := func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in string) *string {
					// if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
					// 	return nil
					// }
					return &in
				}(string(ExpandString(in)))
				return tmp
			}(in)
			return value
		}(in["instance_group"]),
		VolumeType: func(in interface{}) *string {
			value := func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in string) *string {
					// if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
					// 	return nil
					// }
					return &in
				}(string(ExpandString(in)))
				return tmp
			}(in)
			return value
		}(in["volume_type"]),
		VolumeIops: func(in interface{}) *int32 {
			value := func(in interface{}) *int32 {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in int32) *int32 {
					// if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
					// 	return nil
					// }
					return &in
				}(int32(ExpandInt(in)))
				return tmp
			}(in)
			return value
		}(in["volume_iops"]),
		VolumeSize: func(in interface{}) *int32 {
			value := func(in interface{}) *int32 {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in int32) *int32 {
					// if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
					// 	return nil
					// }
					return &in
				}(int32(ExpandInt(in)))
				return tmp
			}(in)
			return value
		}(in["volume_size"]),
		KmsKeyId: func(in interface{}) *string {
			value := func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in string) *string {
					// if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
					// 	return nil
					// }
					return &in
				}(string(ExpandString(in)))
				return tmp
			}(in)
			return value
		}(in["kms_key_id"]),
		EncryptedVolume: func(in interface{}) *bool {
			value := func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in bool) *bool {
					// if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
					// 	return nil
					// }
					return &in
				}(bool(ExpandBool(in)))
				return tmp
			}(in)
			return value
		}(in["encrypted_volume"]),
	}
}

func FlattenEtcdMemberSpec(in kops.EtcdMemberSpec) map[string]interface{} {
	return map[string]interface{}{
		"name": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.Name),
		"instance_group": func(in *string) interface{} {
			value := func(in *string) interface{} {
				if in == nil {
					return nil
				}
				return func(in string) interface{} {
					return FlattenString(string(in))
				}(*in)
			}(in)
			return value
		}(in.InstanceGroup),
		"volume_type": func(in *string) interface{} {
			value := func(in *string) interface{} {
				if in == nil {
					return nil
				}
				return func(in string) interface{} {
					return FlattenString(string(in))
				}(*in)
			}(in)
			return value
		}(in.VolumeType),
		"volume_iops": func(in *int32) interface{} {
			value := func(in *int32) interface{} {
				if in == nil {
					return nil
				}
				return func(in int32) interface{} {
					return FlattenInt(int(in))
				}(*in)
			}(in)
			return value
		}(in.VolumeIops),
		"volume_size": func(in *int32) interface{} {
			value := func(in *int32) interface{} {
				if in == nil {
					return nil
				}
				return func(in int32) interface{} {
					return FlattenInt(int(in))
				}(*in)
			}(in)
			return value
		}(in.VolumeSize),
		"kms_key_id": func(in *string) interface{} {
			value := func(in *string) interface{} {
				if in == nil {
					return nil
				}
				return func(in string) interface{} {
					return FlattenString(string(in))
				}(*in)
			}(in)
			return value
		}(in.KmsKeyId),
		"encrypted_volume": func(in *bool) interface{} {
			value := func(in *bool) interface{} {
				if in == nil {
					return nil
				}
				return func(in bool) interface{} {
					return FlattenBool(bool(in))
				}(*in)
			}(in)
			return value
		}(in.EncryptedVolume),
	}
}
