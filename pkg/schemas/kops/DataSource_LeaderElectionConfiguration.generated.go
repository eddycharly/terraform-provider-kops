package schemas

import (
	"reflect"

	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceLeaderElectionConfiguration() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"leader_elect":                         ComputedBool(),
			"leader_elect_lease_duration":          ComputedDuration(),
			"leader_elect_renew_deadline_duration": ComputedDuration(),
			"leader_elect_resource_lock":           ComputedString(),
			"leader_elect_resource_name":           ComputedString(),
			"leader_elect_resource_namespace":      ComputedString(),
			"leader_elect_retry_period":            ComputedDuration(),
		},
	}

	return res
}

func ExpandDataSourceLeaderElectionConfiguration(in map[string]interface{}) kops.LeaderElectionConfiguration {
	if in == nil {
		panic("expand LeaderElectionConfiguration failure, in is nil")
	}
	return kops.LeaderElectionConfiguration{
		LeaderElect: func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in bool) *bool {
					return &in
				}(bool(ExpandBool(in)))
			}(in)
		}(in["leader_elect"]),
		LeaderElectLeaseDuration: func(in interface{}) *meta.Duration {
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
		}(in["leader_elect_lease_duration"]),
		LeaderElectRenewDeadlineDuration: func(in interface{}) *meta.Duration {
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
		}(in["leader_elect_renew_deadline_duration"]),
		LeaderElectResourceLock: func(in interface{}) *string {
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
		}(in["leader_elect_resource_lock"]),
		LeaderElectResourceName: func(in interface{}) *string {
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
		}(in["leader_elect_resource_name"]),
		LeaderElectResourceNamespace: func(in interface{}) *string {
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
		}(in["leader_elect_resource_namespace"]),
		LeaderElectRetryPeriod: func(in interface{}) *meta.Duration {
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
		}(in["leader_elect_retry_period"]),
	}
}

func FlattenDataSourceLeaderElectionConfigurationInto(in kops.LeaderElectionConfiguration, out map[string]interface{}) {
	out["leader_elect"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.LeaderElect)
	out["leader_elect_lease_duration"] = func(in *meta.Duration) interface{} {
		return func(in *meta.Duration) interface{} {
			if in == nil {
				return nil
			}
			return func(in meta.Duration) interface{} {
				return FlattenDuration(in)
			}(*in)
		}(in)
	}(in.LeaderElectLeaseDuration)
	out["leader_elect_renew_deadline_duration"] = func(in *meta.Duration) interface{} {
		return func(in *meta.Duration) interface{} {
			if in == nil {
				return nil
			}
			return func(in meta.Duration) interface{} {
				return FlattenDuration(in)
			}(*in)
		}(in)
	}(in.LeaderElectRenewDeadlineDuration)
	out["leader_elect_resource_lock"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.LeaderElectResourceLock)
	out["leader_elect_resource_name"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.LeaderElectResourceName)
	out["leader_elect_resource_namespace"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.LeaderElectResourceNamespace)
	out["leader_elect_retry_period"] = func(in *meta.Duration) interface{} {
		return func(in *meta.Duration) interface{} {
			if in == nil {
				return nil
			}
			return func(in meta.Duration) interface{} {
				return FlattenDuration(in)
			}(*in)
		}(in)
	}(in.LeaderElectRetryPeriod)
}

func FlattenDataSourceLeaderElectionConfiguration(in kops.LeaderElectionConfiguration) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceLeaderElectionConfigurationInto(in, out)
	return out
}
