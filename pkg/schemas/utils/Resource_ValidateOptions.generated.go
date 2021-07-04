package schemas

import (
	"github.com/eddycharly/terraform-provider-kops/pkg/api/utils"
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ = Schema

func ExpandResourceValidateOptions(in map[string]interface{}) utils.ValidateOptions {
	if in == nil {
		panic("expand ValidateOptions failure, in is nil")
	}
	out := utils.ValidateOptions{}
	if in, ok := in["timeout"]; ok && in != nil {
		out.Timeout = func(in interface{}) *v1.Duration {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in v1.Duration) *v1.Duration { return &in }(func(in interface{}) v1.Duration { return ExpandDuration(in.(string)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["poll_interval"]; ok && in != nil {
		out.PollInterval = func(in interface{}) *v1.Duration {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in v1.Duration) *v1.Duration { return &in }(func(in interface{}) v1.Duration { return ExpandDuration(in.(string)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	return out
}

func FlattenResourceValidateOptionsInto(in utils.ValidateOptions, out map[string]interface{}) {
	out["timeout"] = func(in *v1.Duration) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in v1.Duration) interface{} { return FlattenDuration(in) }(*in)}
	}(in.Timeout)
	out["poll_interval"] = func(in *v1.Duration) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in v1.Duration) interface{} { return FlattenDuration(in) }(*in)}
	}(in.PollInterval)
}

func FlattenResourceValidateOptions(in utils.ValidateOptions) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceValidateOptionsInto(in, out)
	return out
}
