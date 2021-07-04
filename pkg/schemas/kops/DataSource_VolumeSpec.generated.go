package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceVolumeSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"delete_on_termination": Computed(Nullable(Bool())),
			"device":                Computed(String()),
			"encrypted":             Computed(Nullable(Bool())),
			"iops":                  Computed(Nullable(Int())),
			"throughput":            Computed(Nullable(Int())),
			"key":                   Computed(Nullable(String())),
			"size":                  Computed(Int()),
			"type":                  Computed(String()),
		},
	}
}

func ExpandDataSourceVolumeSpec(in map[string]interface{}) kops.VolumeSpec {
	if in == nil {
		panic("expand VolumeSpec failure, in is nil")
	}
	out := kops.VolumeSpec{}
	if in, ok := in["delete_on_termination"]; ok && in != nil {
		out.DeleteOnTermination = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["device"]; ok && in != nil {
		out.Device = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["encrypted"]; ok && in != nil {
		out.Encrypted = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["iops"]; ok && in != nil {
		out.Iops = func(in interface{}) *int64 {
			if in == nil {
				return nil
			}
			return func(in int64) *int64 { return &in }(func(in interface{}) int64 { return int64(in.(int)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["throughput"]; ok && in != nil {
		out.Throughput = func(in interface{}) *int64 {
			if in == nil {
				return nil
			}
			return func(in int64) *int64 { return &in }(func(in interface{}) int64 { return int64(in.(int)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["key"]; ok && in != nil {
		out.Key = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["size"]; ok && in != nil {
		out.Size = func(in interface{}) int64 { return int64(in.(int)) }(in)
	}
	if in, ok := in["type"]; ok && in != nil {
		out.Type = func(in interface{}) string { return string(in.(string)) }(in)
	}
	return out
}

func FlattenDataSourceVolumeSpecInto(in kops.VolumeSpec, out map[string]interface{}) {
	out["delete_on_termination"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in bool) interface{} { return in }(*in)}
	}(in.DeleteOnTermination)
	out["device"] = func(in string) interface{} { return string(in) }(in.Device)
	out["encrypted"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in bool) interface{} { return in }(*in)}
	}(in.Encrypted)
	out["iops"] = func(in *int64) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in int64) interface{} { return int(in) }(*in)}
	}(in.Iops)
	out["throughput"] = func(in *int64) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in int64) interface{} { return int(in) }(*in)}
	}(in.Throughput)
	out["key"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.Key)
	out["size"] = func(in int64) interface{} { return int(in) }(in.Size)
	out["type"] = func(in string) interface{} { return string(in) }(in.Type)
}

func FlattenDataSourceVolumeSpec(in kops.VolumeSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceVolumeSpecInto(in, out)
	return out
}
