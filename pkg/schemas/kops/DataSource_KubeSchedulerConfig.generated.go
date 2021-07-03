package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceKubeSchedulerConfig() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"master":                           Computed(String()),
			"log_level":                        Computed(Int()),
			"image":                            Computed(String()),
			"leader_election":                  Computed(Ptr(Struct(DataSourceLeaderElectionConfiguration()))),
			"use_policy_config_map":            Computed(Ptr(Bool())),
			"feature_gates":                    Computed(Map(String())),
			"max_persistent_volumes":           Computed(Ptr(Int())),
			"qps":                              Computed(Ptr(Quantity())),
			"burst":                            Computed(Int()),
			"authentication_kubeconfig":        Computed(String()),
			"authorization_kubeconfig":         Computed(String()),
			"authorization_always_allow_paths": Computed(List(String())),
			"enable_profiling":                 Computed(Ptr(Bool())),
		},
	}
}

func ExpandDataSourceKubeSchedulerConfig(in map[string]interface{}) kops.KubeSchedulerConfig {
	if in == nil {
		panic("expand KubeSchedulerConfig failure, in is nil")
	}
	out := kops.KubeSchedulerConfig{}
	if in, ok := in["master"]; ok && in != nil {
		out.Master = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["log_level"]; ok && in != nil {
		out.LogLevel = func(in interface{}) int32 { return int32(in.(int)) }(in)
	}
	if in, ok := in["image"]; ok && in != nil {
		out.Image = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["leader_election"]; ok && in != nil {
		out.LeaderElection = func(in interface{}) *kops.LeaderElectionConfiguration {
			if in == nil {
				return nil
			}
			return func(in kops.LeaderElectionConfiguration) *kops.LeaderElectionConfiguration { return &in }(func(in interface{}) kops.LeaderElectionConfiguration {
				if in == nil {
					return kops.LeaderElectionConfiguration{}
				}
				return ExpandDataSourceLeaderElectionConfiguration(in.(map[string]interface{}))
			}(in))
		}(in)
	}
	if in, ok := in["use_policy_config_map"]; ok && in != nil {
		out.UsePolicyConfigMap = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["feature_gates"]; ok && in != nil {
		out.FeatureGates = func(in interface{}) map[string]string {
			if in == nil {
				return nil
			}
			out := map[string]string{}
			for key, in := range in.(map[string]interface{}) {
				out[key] = func(in interface{}) string { return string(in.(string)) }(in)
			}
			return out
		}(in)
	}
	if in, ok := in["max_persistent_volumes"]; ok && in != nil {
		out.MaxPersistentVolumes = func(in interface{}) *int32 {
			if in == nil {
				return nil
			}
			return func(in int32) *int32 { return &in }(func(in interface{}) int32 { return int32(in.(int)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["qps"]; ok && in != nil {
		out.Qps = func(in interface{}) *resource.Quantity {
			if in == nil {
				return nil
			}
			return func(in resource.Quantity) *resource.Quantity { return &in }(func(in interface{}) resource.Quantity { return ExpandQuantity(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["burst"]; ok && in != nil {
		out.Burst = func(in interface{}) int32 { return int32(in.(int)) }(in)
	}
	if in, ok := in["authentication_kubeconfig"]; ok && in != nil {
		out.AuthenticationKubeconfig = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["authorization_kubeconfig"]; ok && in != nil {
		out.AuthorizationKubeconfig = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["authorization_always_allow_paths"]; ok && in != nil {
		out.AuthorizationAlwaysAllowPaths = func(in interface{}) []string {
			var out []string
			for _, in := range in.([]interface{}) {
				out = append(out, func(in interface{}) string { return string(in.(string)) }(in))
			}
			return out
		}(in)
	}
	if in, ok := in["enable_profiling"]; ok && in != nil {
		out.EnableProfiling = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	return out
}
