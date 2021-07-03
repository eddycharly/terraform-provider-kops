package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceAWSEBSCSIDriver() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enabled":             Computed(Ptr(Bool())),
			"version":             Computed(Ptr(String())),
			"volume_attach_limit": Computed(Ptr(Int())),
		},
	}
}

func ExpandDataSourceAWSEBSCSIDriver(in map[string]interface{}) kops.AWSEBSCSIDriver {
	if in == nil {
		panic("expand AWSEBSCSIDriver failure, in is nil")
	}
	out := kops.AWSEBSCSIDriver{}
	if in, ok := in["enabled"]; ok && in != nil {
		out.Enabled = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["version"]; ok && in != nil {
		out.Version = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["volume_attach_limit"]; ok && in != nil {
		out.VolumeAttachLimit = func(in interface{}) *int {
			if in == nil {
				return nil
			}
			return func(in int) *int { return &in }(func(in interface{}) int { return int(in.(int)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	return out
}
