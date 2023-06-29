package schemas

import (
	"reflect"

	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

var _ = Schema

func ResourceInstanceMetadataOptions() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"http_put_response_hop_limit": OptionalInt(),
			"http_tokens":                 OptionalString(),
		},
	}

	return res
}

func ExpandResourceInstanceMetadataOptions(in map[string]interface{}) kopsv1alpha2.InstanceMetadataOptions {
	if in == nil {
		panic("expand InstanceMetadataOptions failure, in is nil")
	}
	return kopsv1alpha2.InstanceMetadataOptions{
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

func FlattenResourceInstanceMetadataOptionsInto(in kopsv1alpha2.InstanceMetadataOptions, out map[string]interface{}) {
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

func FlattenResourceInstanceMetadataOptions(in kopsv1alpha2.InstanceMetadataOptions) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceInstanceMetadataOptionsInto(in, out)
	return out
}
