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
			return func(in v1.Duration) *v1.Duration { return &in }(func(in interface{}) v1.Duration { return ExpandDuration(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["node_interval"]; ok && in != nil {
		out.NodeInterval = func(in interface{}) *v1.Duration {
			if in == nil {
				return nil
			}
			return func(in v1.Duration) *v1.Duration { return &in }(func(in interface{}) v1.Duration { return ExpandDuration(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["bastion_interval"]; ok && in != nil {
		out.BastionInterval = func(in interface{}) *v1.Duration {
			if in == nil {
				return nil
			}
			return func(in v1.Duration) *v1.Duration { return &in }(func(in interface{}) v1.Duration { return ExpandDuration(in.(string)) }(in.(map[string]interface{})["value"]))
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
			return func(in v1.Duration) *v1.Duration { return &in }(func(in interface{}) v1.Duration { return ExpandDuration(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["validation_timeout"]; ok && in != nil {
		out.ValidationTimeout = func(in interface{}) *v1.Duration {
			if in == nil {
				return nil
			}
			return func(in v1.Duration) *v1.Duration { return &in }(func(in interface{}) v1.Duration { return ExpandDuration(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["validate_count"]; ok && in != nil {
		out.ValidateCount = func(in interface{}) *int {
			if in == nil {
				return nil
			}
			return func(in int) *int { return &in }(func(in interface{}) int { return int(in.(int)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["cloud_only"]; ok && in != nil {
		out.CloudOnly = func(in interface{}) bool { return in.(bool) }(in)
	}
	return out
}
