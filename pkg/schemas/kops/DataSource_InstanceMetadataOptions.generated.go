package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceInstanceMetadataOptions() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"http_put_response_hop_limit": Computed(Ptr(Int())),
			"http_tokens":                 Computed(Ptr(String())),
		},
	}
}

func ExpandDataSourceInstanceMetadataOptions(in map[string]interface{}) kops.InstanceMetadataOptions {
	if in == nil {
		panic("expand InstanceMetadataOptions failure, in is nil")
	}
	out := kops.InstanceMetadataOptions{}
	if in, ok := in["http_put_response_hop_limit"]; ok && in != nil {
		out.HTTPPutResponseHopLimit = func(in interface{}) *int64 {
			if in == nil {
				return nil
			}
			return func(in int64) *int64 { return &in }(func(in interface{}) int64 { return int64(in.(int)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["http_tokens"]; ok && in != nil {
		out.HTTPTokens = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	return out
}
