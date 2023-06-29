package schemas

import (
	"reflect"

	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

var _ = Schema

func ResourceAccessLogSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"interval":      OptionalInt(),
			"bucket":        OptionalString(),
			"bucket_prefix": OptionalString(),
		},
	}

	return res
}

func ExpandResourceAccessLogSpec(in map[string]interface{}) kopsv1alpha2.AccessLogSpec {
	if in == nil {
		panic("expand AccessLogSpec failure, in is nil")
	}
	return kopsv1alpha2.AccessLogSpec{
		Interval: func(in interface{}) int {
			return int(ExpandInt(in))
		}(in["interval"]),
		Bucket: func(in interface{}) *string {
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
		}(in["bucket"]),
		BucketPrefix: func(in interface{}) *string {
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
		}(in["bucket_prefix"]),
	}
}

func FlattenResourceAccessLogSpecInto(in kopsv1alpha2.AccessLogSpec, out map[string]interface{}) {
	out["interval"] = func(in int) interface{} {
		return FlattenInt(int(in))
	}(in.Interval)
	out["bucket"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.Bucket)
	out["bucket_prefix"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.BucketPrefix)
}

func FlattenResourceAccessLogSpec(in kopsv1alpha2.AccessLogSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceAccessLogSpecInto(in, out)
	return out
}
