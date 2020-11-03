package structures

import (
	"reflect"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/kops/pkg/apis/kops"
)

func ExpandLeaderElectionConfiguration(in map[string]interface{}) kops.LeaderElectionConfiguration {
	if in == nil {
		panic("expand LeaderElectionConfiguration failure, in is nil")
	}
	return kops.LeaderElectionConfiguration{
		LeaderElect: func(in interface{}) *bool {
			value := func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in bool) *bool {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(bool(ExpandBool(in)))
				return tmp
			}(in)
			return value
		}(in["leader_elect"]),
		LeaderElectLeaseDuration: func(in interface{}) *v1.Duration {
			value := func(in interface{}) *v1.Duration {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in v1.Duration) *v1.Duration {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(ExpandDuration(in))
				return tmp
			}(in)
			return value
		}(in["leader_elect_lease_duration"]),
		LeaderElectRenewDeadlineDuration: func(in interface{}) *v1.Duration {
			value := func(in interface{}) *v1.Duration {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in v1.Duration) *v1.Duration {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(ExpandDuration(in))
				return tmp
			}(in)
			return value
		}(in["leader_elect_renew_deadline_duration"]),
		LeaderElectResourceLock: func(in interface{}) *string {
			value := func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in string) *string {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(string(ExpandString(in)))
				return tmp
			}(in)
			return value
		}(in["leader_elect_resource_lock"]),
		LeaderElectResourceName: func(in interface{}) *string {
			value := func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in string) *string {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(string(ExpandString(in)))
				return tmp
			}(in)
			return value
		}(in["leader_elect_resource_name"]),
		LeaderElectResourceNamespace: func(in interface{}) *string {
			value := func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in string) *string {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(string(ExpandString(in)))
				return tmp
			}(in)
			return value
		}(in["leader_elect_resource_namespace"]),
		LeaderElectRetryPeriod: func(in interface{}) *v1.Duration {
			value := func(in interface{}) *v1.Duration {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in v1.Duration) *v1.Duration {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(ExpandDuration(in))
				return tmp
			}(in)
			return value
		}(in["leader_elect_retry_period"]),
	}
}

func FlattenLeaderElectionConfiguration(in kops.LeaderElectionConfiguration) map[string]interface{} {
	return map[string]interface{}{
		"leader_elect": func(in *bool) interface{} {
			value := func(in *bool) interface{} {
				if in == nil {
					return nil
				}
				return func(in bool) interface{} {
					return FlattenBool(bool(in))
				}(*in)
			}(in)
			return value
		}(in.LeaderElect),
		"leader_elect_lease_duration": func(in *v1.Duration) interface{} {
			value := func(in *v1.Duration) interface{} {
				if in == nil {
					return nil
				}
				return func(in v1.Duration) interface{} {
					return FlattenDuration(in)
				}(*in)
			}(in)
			return value
		}(in.LeaderElectLeaseDuration),
		"leader_elect_renew_deadline_duration": func(in *v1.Duration) interface{} {
			value := func(in *v1.Duration) interface{} {
				if in == nil {
					return nil
				}
				return func(in v1.Duration) interface{} {
					return FlattenDuration(in)
				}(*in)
			}(in)
			return value
		}(in.LeaderElectRenewDeadlineDuration),
		"leader_elect_resource_lock": func(in *string) interface{} {
			value := func(in *string) interface{} {
				if in == nil {
					return nil
				}
				return func(in string) interface{} {
					return FlattenString(string(in))
				}(*in)
			}(in)
			return value
		}(in.LeaderElectResourceLock),
		"leader_elect_resource_name": func(in *string) interface{} {
			value := func(in *string) interface{} {
				if in == nil {
					return nil
				}
				return func(in string) interface{} {
					return FlattenString(string(in))
				}(*in)
			}(in)
			return value
		}(in.LeaderElectResourceName),
		"leader_elect_resource_namespace": func(in *string) interface{} {
			value := func(in *string) interface{} {
				if in == nil {
					return nil
				}
				return func(in string) interface{} {
					return FlattenString(string(in))
				}(*in)
			}(in)
			return value
		}(in.LeaderElectResourceNamespace),
		"leader_elect_retry_period": func(in *v1.Duration) interface{} {
			value := func(in *v1.Duration) interface{} {
				if in == nil {
					return nil
				}
				return func(in v1.Duration) interface{} {
					return FlattenDuration(in)
				}(*in)
			}(in)
			return value
		}(in.LeaderElectRetryPeriod),
	}
}
