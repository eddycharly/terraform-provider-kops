package schemas

import (
	"reflect"

	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceInstanceMetadataOptions() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"http_put_response_hop_limit": ComputedInt(),
			"http_tokens":                 ComputedString(),
		},
	}

	return res
}

func ExpandDataSourceInstanceMetadataOptions(in map[string]interface{}) kops.InstanceMetadataOptions {
	if in == nil {
		panic("expand InstanceMetadataOptions failure, in is nil")
	}
	return kops.InstanceMetadataOptions{
		HTTPPutResponseHopLimit: func(in interface{}) *int64 {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *int64 {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in int64) *int64 {
					return &in
				}(int64(ExpandInt(in)))
			}(in)
		}(in["http_put_response_hop_limit"]),
		HTTPTokens: func(in interface{}) *string {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in string) *string {
					return &in
				}(string(ExpandString(in)))
			}(in)
		}(in["http_tokens"]),
	}
}

func FlattenDataSourceInstanceMetadataOptionsInto(in kops.InstanceMetadataOptions, out map[string]interface{}) {
	out["http_put_response_hop_limit"] = func(in *int64) interface{} {
		return func(in *int64) interface{} {
			if in == nil {
				return nil
			}
			return func(in int64) interface{} {
				return FlattenInt(int(in))
			}(*in)
		}(in)
	}(in.HTTPPutResponseHopLimit)
	out["http_tokens"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.HTTPTokens)
}

func FlattenDataSourceInstanceMetadataOptions(in kops.InstanceMetadataOptions) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceInstanceMetadataOptionsInto(in, out)
	return out
}
