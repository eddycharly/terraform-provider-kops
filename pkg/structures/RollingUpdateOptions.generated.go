package structures

import (
	"reflect"

	"github.com/eddycharly/terraform-provider-kops/pkg/api"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func ExpandRollingUpdateOptions(in map[string]interface{}) api.RollingUpdateOptions {
	if in == nil {
		panic("expand RollingUpdateOptions failure, in is nil")
	}
	return api.RollingUpdateOptions{
		Skip: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["skip"]),
		MasterInterval: func(in interface{}) *v1.Duration {
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
		}(in["master_interval"]),
		NodeInterval: func(in interface{}) *v1.Duration {
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
		}(in["node_interval"]),
		BastionInterval: func(in interface{}) *v1.Duration {
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
		}(in["bastion_interval"]),
		FailOnDrainError: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["fail_on_drain_error"]),
		FailOnValidate: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["fail_on_validate"]),
		PostDrainDelay: func(in interface{}) *v1.Duration {
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
		}(in["post_drain_delay"]),
		ValidationTimeout: func(in interface{}) *v1.Duration {
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
		}(in["validation_timeout"]),
		ValidateCount: func(in interface{}) *int {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *int {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in int) *int {
					return &in
				}(int(ExpandInt(in)))
			}(in)
		}(in["validate_count"]),
	}
}

func FlattenRollingUpdateOptions(in api.RollingUpdateOptions) map[string]interface{} {
	return map[string]interface{}{
		"skip": func(in bool) interface{} {
			return FlattenBool(bool(in))
		}(in.Skip),
		"master_interval": func(in *v1.Duration) interface{} {
			return func(in *v1.Duration) interface{} {
				if in == nil {
					return nil
				}
				return func(in v1.Duration) interface{} {
					return FlattenDuration(in)
				}(*in)
			}(in)
		}(in.MasterInterval),
		"node_interval": func(in *v1.Duration) interface{} {
			return func(in *v1.Duration) interface{} {
				if in == nil {
					return nil
				}
				return func(in v1.Duration) interface{} {
					return FlattenDuration(in)
				}(*in)
			}(in)
		}(in.NodeInterval),
		"bastion_interval": func(in *v1.Duration) interface{} {
			return func(in *v1.Duration) interface{} {
				if in == nil {
					return nil
				}
				return func(in v1.Duration) interface{} {
					return FlattenDuration(in)
				}(*in)
			}(in)
		}(in.BastionInterval),
		"fail_on_drain_error": func(in bool) interface{} {
			return FlattenBool(bool(in))
		}(in.FailOnDrainError),
		"fail_on_validate": func(in bool) interface{} {
			return FlattenBool(bool(in))
		}(in.FailOnValidate),
		"post_drain_delay": func(in *v1.Duration) interface{} {
			return func(in *v1.Duration) interface{} {
				if in == nil {
					return nil
				}
				return func(in v1.Duration) interface{} {
					return FlattenDuration(in)
				}(*in)
			}(in)
		}(in.PostDrainDelay),
		"validation_timeout": func(in *v1.Duration) interface{} {
			return func(in *v1.Duration) interface{} {
				if in == nil {
					return nil
				}
				return func(in v1.Duration) interface{} {
					return FlattenDuration(in)
				}(*in)
			}(in)
		}(in.ValidationTimeout),
		"validate_count": func(in *int) interface{} {
			return func(in *int) interface{} {
				if in == nil {
					return nil
				}
				return func(in int) interface{} {
					return FlattenInt(int(in))
				}(*in)
			}(in)
		}(in.ValidateCount),
	}
}
