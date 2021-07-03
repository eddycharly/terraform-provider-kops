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
			return func(in v1.Duration) *v1.Duration { return &in }(func(in interface{}) v1.Duration { return ExpandDuration(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["poll_interval"]; ok && in != nil {
		out.PollInterval = func(in interface{}) *v1.Duration {
			if in == nil {
				return nil
			}
			return func(in v1.Duration) *v1.Duration { return &in }(func(in interface{}) v1.Duration { return ExpandDuration(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	return out
}
