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
			"delete_on_termination": Computed(Ptr(Bool())),
			"device":                Computed(String()),
			"encrypted":             Computed(Ptr(Bool())),
			"iops":                  Computed(Ptr(Int())),
			"throughput":            Computed(Ptr(Int())),
			"key":                   Computed(Ptr(String())),
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
