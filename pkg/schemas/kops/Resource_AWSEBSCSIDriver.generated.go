package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceAWSEBSCSIDriver() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enabled":             Optional(Nullable(Bool())),
			"version":             Optional(Nullable(String())),
			"volume_attach_limit": Optional(Nullable(Int())),
		},
	}
}

func ExpandResourceAWSEBSCSIDriver(in map[string]interface{}) kops.AWSEBSCSIDriver {
	if in == nil {
		panic("expand AWSEBSCSIDriver failure, in is nil")
	}
	out := kops.AWSEBSCSIDriver{}
	if in, ok := in["enabled"]; ok && in != nil {
		out.Enabled = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["version"]; ok && in != nil {
		out.Version = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["volume_attach_limit"]; ok && in != nil {
		out.VolumeAttachLimit = func(in interface{}) *int {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in int) *int { return &in }(func(in interface{}) int { return int(in.(int)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	return out
}

func FlattenResourceAWSEBSCSIDriverInto(in kops.AWSEBSCSIDriver, out map[string]interface{}) {
	out["enabled"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in bool) interface{} { return in }(*in)}
	}(in.Enabled)
	out["version"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.Version)
	out["volume_attach_limit"] = func(in *int) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in int) interface{} { return int(in) }(*in)}
	}(in.VolumeAttachLimit)
}

func FlattenResourceAWSEBSCSIDriver(in kops.AWSEBSCSIDriver) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceAWSEBSCSIDriverInto(in, out)
	return out
}
