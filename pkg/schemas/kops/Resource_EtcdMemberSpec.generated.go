package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceEtcdMemberSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name":              Required(String()),
			"instance_group":    Required(Nullable(String())),
			"volume_type":       Optional(Nullable(String())),
			"volume_iops":       Optional(Nullable(Int())),
			"volume_throughput": Optional(Nullable(Int())),
			"volume_size":       Optional(Nullable(Int())),
			"kms_key_id":        Optional(Nullable(String())),
			"encrypted_volume":  Optional(Nullable(Bool())),
		},
	}
}

func ExpandResourceEtcdMemberSpec(in map[string]interface{}) kops.EtcdMemberSpec {
	if in == nil {
		panic("expand EtcdMemberSpec failure, in is nil")
	}
	out := kops.EtcdMemberSpec{}
	if in, ok := in["name"]; ok && in != nil {
		out.Name = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["instance_group"]; ok && in != nil {
		out.InstanceGroup = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["volume_type"]; ok && in != nil {
		out.VolumeType = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["volume_iops"]; ok && in != nil {
		out.VolumeIops = func(in interface{}) *int32 {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in int32) *int32 { return &in }(func(in interface{}) int32 { return int32(in.(int)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["volume_throughput"]; ok && in != nil {
		out.VolumeThroughput = func(in interface{}) *int32 {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in int32) *int32 { return &in }(func(in interface{}) int32 { return int32(in.(int)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["volume_size"]; ok && in != nil {
		out.VolumeSize = func(in interface{}) *int32 {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in int32) *int32 { return &in }(func(in interface{}) int32 { return int32(in.(int)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["kms_key_id"]; ok && in != nil {
		out.KmsKeyId = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["encrypted_volume"]; ok && in != nil {
		out.EncryptedVolume = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	return out
}

func FlattenResourceEtcdMemberSpecInto(in kops.EtcdMemberSpec, out map[string]interface{}) {
	out["name"] = func(in string) interface{} { return string(in) }(in.Name)
	out["instance_group"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.InstanceGroup)
	out["volume_type"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.VolumeType)
	out["volume_iops"] = func(in *int32) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in int32) interface{} { return int(in) }(*in)}
	}(in.VolumeIops)
	out["volume_throughput"] = func(in *int32) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in int32) interface{} { return int(in) }(*in)}
	}(in.VolumeThroughput)
	out["volume_size"] = func(in *int32) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in int32) interface{} { return int(in) }(*in)}
	}(in.VolumeSize)
	out["kms_key_id"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.KmsKeyId)
	out["encrypted_volume"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in bool) interface{} { return in }(*in)}
	}(in.EncryptedVolume)
}

func FlattenResourceEtcdMemberSpec(in kops.EtcdMemberSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceEtcdMemberSpecInto(in, out)
	return out
}
