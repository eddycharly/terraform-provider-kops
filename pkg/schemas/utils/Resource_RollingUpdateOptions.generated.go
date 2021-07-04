package schemas

import (
	"github.com/eddycharly/terraform-provider-kops/pkg/api/utils"
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ = Schema

func ExpandResourceRollingUpdateOptions(in map[string]interface{}) utils.RollingUpdateOptions {
	if in == nil {
		panic("expand RollingUpdateOptions failure, in is nil")
	}
	out := utils.RollingUpdateOptions{}
	if in, ok := in["master_interval"]; ok && in != nil {
		out.MasterInterval = func(in interface{}) *v1.Duration {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in v1.Duration) *v1.Duration { return &in }(func(in interface{}) v1.Duration { return ExpandDuration(in.(string)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["node_interval"]; ok && in != nil {
		out.NodeInterval = func(in interface{}) *v1.Duration {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in v1.Duration) *v1.Duration { return &in }(func(in interface{}) v1.Duration { return ExpandDuration(in.(string)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["bastion_interval"]; ok && in != nil {
		out.BastionInterval = func(in interface{}) *v1.Duration {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in v1.Duration) *v1.Duration { return &in }(func(in interface{}) v1.Duration { return ExpandDuration(in.(string)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["fail_on_drain_error"]; ok && in != nil {
		out.FailOnDrainError = func(in interface{}) bool { return in.(bool) }(in)
	}
	if in, ok := in["fail_on_validate"]; ok && in != nil {
		out.FailOnValidate = func(in interface{}) bool { return in.(bool) }(in)
	}
	if in, ok := in["post_drain_delay"]; ok && in != nil {
		out.PostDrainDelay = func(in interface{}) *v1.Duration {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in v1.Duration) *v1.Duration { return &in }(func(in interface{}) v1.Duration { return ExpandDuration(in.(string)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["validation_timeout"]; ok && in != nil {
		out.ValidationTimeout = func(in interface{}) *v1.Duration {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in v1.Duration) *v1.Duration { return &in }(func(in interface{}) v1.Duration { return ExpandDuration(in.(string)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["validate_count"]; ok && in != nil {
		out.ValidateCount = func(in interface{}) *int {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in int) *int { return &in }(func(in interface{}) int { return int(in.(int)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["cloud_only"]; ok && in != nil {
		out.CloudOnly = func(in interface{}) bool { return in.(bool) }(in)
	}
	return out
}

func FlattenResourceRollingUpdateOptionsInto(in utils.RollingUpdateOptions, out map[string]interface{}) {
	out["master_interval"] = func(in *v1.Duration) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in v1.Duration) interface{} { return FlattenDuration(in) }(*in)}
	}(in.MasterInterval)
	out["node_interval"] = func(in *v1.Duration) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in v1.Duration) interface{} { return FlattenDuration(in) }(*in)}
	}(in.NodeInterval)
	out["bastion_interval"] = func(in *v1.Duration) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in v1.Duration) interface{} { return FlattenDuration(in) }(*in)}
	}(in.BastionInterval)
	out["fail_on_drain_error"] = func(in bool) interface{} { return in }(in.FailOnDrainError)
	out["fail_on_validate"] = func(in bool) interface{} { return in }(in.FailOnValidate)
	out["post_drain_delay"] = func(in *v1.Duration) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in v1.Duration) interface{} { return FlattenDuration(in) }(*in)}
	}(in.PostDrainDelay)
	out["validation_timeout"] = func(in *v1.Duration) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in v1.Duration) interface{} { return FlattenDuration(in) }(*in)}
	}(in.ValidationTimeout)
	out["validate_count"] = func(in *int) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in int) interface{} { return int(in) }(*in)}
	}(in.ValidateCount)
	out["cloud_only"] = func(in bool) interface{} { return in }(in.CloudOnly)
}

func FlattenResourceRollingUpdateOptions(in utils.RollingUpdateOptions) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceRollingUpdateOptionsInto(in, out)
	return out
}
