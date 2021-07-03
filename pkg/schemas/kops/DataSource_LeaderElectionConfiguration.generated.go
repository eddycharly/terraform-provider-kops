package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceLeaderElectionConfiguration() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"leader_elect":                         Computed(Ptr(Bool())),
			"leader_elect_lease_duration":          Computed(Ptr(Duration())),
			"leader_elect_renew_deadline_duration": Computed(Ptr(Duration())),
			"leader_elect_resource_lock":           Computed(Ptr(String())),
			"leader_elect_resource_name":           Computed(Ptr(String())),
			"leader_elect_resource_namespace":      Computed(Ptr(String())),
			"leader_elect_retry_period":            Computed(Ptr(Duration())),
		},
	}
}

func ExpandDataSourceLeaderElectionConfiguration(in map[string]interface{}) kops.LeaderElectionConfiguration {
	if in == nil {
		panic("expand LeaderElectionConfiguration failure, in is nil")
	}
	out := kops.LeaderElectionConfiguration{}
	if in, ok := in["leader_elect"]; ok && in != nil {
		out.LeaderElect = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["leader_elect_lease_duration"]; ok && in != nil {
		out.LeaderElectLeaseDuration = func(in interface{}) *v1.Duration {
			if in == nil {
				return nil
			}
			return func(in v1.Duration) *v1.Duration { return &in }(func(in interface{}) v1.Duration { return ExpandDuration(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["leader_elect_renew_deadline_duration"]; ok && in != nil {
		out.LeaderElectRenewDeadlineDuration = func(in interface{}) *v1.Duration {
			if in == nil {
				return nil
			}
			return func(in v1.Duration) *v1.Duration { return &in }(func(in interface{}) v1.Duration { return ExpandDuration(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["leader_elect_resource_lock"]; ok && in != nil {
		out.LeaderElectResourceLock = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["leader_elect_resource_name"]; ok && in != nil {
		out.LeaderElectResourceName = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["leader_elect_resource_namespace"]; ok && in != nil {
		out.LeaderElectResourceNamespace = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["leader_elect_retry_period"]; ok && in != nil {
		out.LeaderElectRetryPeriod = func(in interface{}) *v1.Duration {
			if in == nil {
				return nil
			}
			return func(in v1.Duration) *v1.Duration { return &in }(func(in interface{}) v1.Duration { return ExpandDuration(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	return out
}
