package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceInstanceMetadataOptions() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"http_put_response_hop_limit": Optional(Nullable(Int())),
			"http_tokens":                 Optional(Nullable(String())),
		},
	}
}

func ExpandResourceInstanceMetadataOptions(in map[string]interface{}) kops.InstanceMetadataOptions {
	if in == nil {
		panic("expand InstanceMetadataOptions failure, in is nil")
	}
	out := kops.InstanceMetadataOptions{}
	if in, ok := in["http_put_response_hop_limit"]; ok && in != nil {
		out.HTTPPutResponseHopLimit = func(in interface{}) *int64 {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in int64) *int64 { return &in }(func(in interface{}) int64 { return int64(in.(int)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["http_tokens"]; ok && in != nil {
		out.HTTPTokens = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	return out
}

func FlattenResourceInstanceMetadataOptionsInto(in kops.InstanceMetadataOptions, out map[string]interface{}) {
	out["http_put_response_hop_limit"] = func(in *int64) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in int64) interface{} { return int(in) }(*in)}
	}(in.HTTPPutResponseHopLimit)
	out["http_tokens"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.HTTPTokens)
}

func FlattenResourceInstanceMetadataOptions(in kops.InstanceMetadataOptions) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceInstanceMetadataOptionsInto(in, out)
	return out
}
