package schemas

import (
	"reflect"

	"github.com/eddycharly/terraform-provider-kops/pkg/api/utils"
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ = Schema

func ExpandResourceValidateOptions(in map[string]interface{}) utils.ValidateOptions {
	if in == nil {
		panic("expand ValidateOptions failure, in is nil")
	}
	return utils.ValidateOptions{
		Timeout: func(in interface{}) *meta.Duration {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *meta.Duration {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in meta.Duration) *meta.Duration {
					return &in
				}(ExpandDuration(in))
			}(in)
		}(in["timeout"]),
		PollInterval: func(in interface{}) *meta.Duration {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *meta.Duration {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in meta.Duration) *meta.Duration {
					return &in
				}(ExpandDuration(in))
			}(in)
		}(in["poll_interval"]),
	}
}

func FlattenResourceValidateOptionsInto(in utils.ValidateOptions, out map[string]interface{}) {
	out["timeout"] = func(in *meta.Duration) interface{} {
		return func(in *meta.Duration) interface{} {
			if in == nil {
				return nil
			}
			return func(in meta.Duration) interface{} {
				return FlattenDuration(in)
			}(*in)
		}(in)
	}(in.Timeout)
	out["poll_interval"] = func(in *meta.Duration) interface{} {
		return func(in *meta.Duration) interface{} {
			if in == nil {
				return nil
			}
			return func(in meta.Duration) interface{} {
				return FlattenDuration(in)
			}(*in)
		}(in)
	}(in.PollInterval)
}

func FlattenResourceValidateOptions(in utils.ValidateOptions) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceValidateOptionsInto(in, out)
	return out
}
