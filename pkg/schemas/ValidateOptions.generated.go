package schemas

import (
	"reflect"

	"github.com/eddycharly/terraform-provider-kops/pkg/api"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func ValidateOptions() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"skip":          OptionalBool(),
			"timeout":       OptionalDuration(),
			"poll_interval": OptionalDuration(),
		},
	}
}

func ExpandValidateOptions(in map[string]interface{}) api.ValidateOptions {
	if in == nil {
		panic("expand ValidateOptions failure, in is nil")
	}
	return api.ValidateOptions{
		Skip: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["skip"]),
		Timeout: func(in interface{}) *v1.Duration {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *v1.Duration {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in v1.Duration) *v1.Duration {
					return &in
				}(ExpandDuration(in))
			}(in)
		}(in["timeout"]),
		PollInterval: func(in interface{}) *v1.Duration {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *v1.Duration {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in v1.Duration) *v1.Duration {
					return &in
				}(ExpandDuration(in))
			}(in)
		}(in["poll_interval"]),
	}
}

func FlattenValidateOptionsInto(in api.ValidateOptions, out map[string]interface{}) {
	out["skip"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.Skip)
	out["timeout"] = func(in *v1.Duration) interface{} {
		return func(in *v1.Duration) interface{} {
			if in == nil {
				return nil
			}
			return func(in v1.Duration) interface{} {
				return FlattenDuration(in)
			}(*in)
		}(in)
	}(in.Timeout)
	out["poll_interval"] = func(in *v1.Duration) interface{} {
		return func(in *v1.Duration) interface{} {
			if in == nil {
				return nil
			}
			return func(in v1.Duration) interface{} {
				return FlattenDuration(in)
			}(*in)
		}(in)
	}(in.PollInterval)
}

func FlattenValidateOptions(in api.ValidateOptions) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenValidateOptionsInto(in, out)
	return out
}
