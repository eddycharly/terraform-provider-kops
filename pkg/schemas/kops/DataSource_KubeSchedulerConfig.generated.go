package schemas

import (
	"reflect"

	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceKubeSchedulerConfig() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"master":                           ComputedString(),
			"log_level":                        ComputedInt(),
			"image":                            ComputedString(),
			"leader_election":                  ComputedStruct(DataSourceLeaderElectionConfiguration()),
			"use_policy_config_map":            ComputedBool(),
			"feature_gates":                    ComputedMap(String()),
			"max_persistent_volumes":           ComputedInt(),
			"qps":                              ComputedQuantity(),
			"burst":                            ComputedInt(),
			"authentication_kubeconfig":        ComputedString(),
			"authorization_kubeconfig":         ComputedString(),
			"authorization_always_allow_paths": ComputedList(String()),
			"enable_profiling":                 ComputedBool(),
		},
	}

	return res
}

func ExpandDataSourceKubeSchedulerConfig(in map[string]interface{}) kops.KubeSchedulerConfig {
	if in == nil {
		panic("expand KubeSchedulerConfig failure, in is nil")
	}
	return kops.KubeSchedulerConfig{
		Master: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["master"]),
		LogLevel: func(in interface{}) int32 {
			return int32(ExpandInt(in))
		}(in["log_level"]),
		Image: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["image"]),
		LeaderElection: func(in interface{}) *kops.LeaderElectionConfiguration {
			return func(in interface{}) *kops.LeaderElectionConfiguration {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.LeaderElectionConfiguration) *kops.LeaderElectionConfiguration {
					return &in
				}(func(in interface{}) kops.LeaderElectionConfiguration {
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return kops.LeaderElectionConfiguration{}
					}
					return (ExpandDataSourceLeaderElectionConfiguration(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["leader_election"]),
		UsePolicyConfigMap: func(in interface{}) *bool {
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
		}(in["use_policy_config_map"]),
		FeatureGates: func(in interface{}) map[string]string {
			return func(in interface{}) map[string]string {
				if in == nil {
					return nil
				}
				out := map[string]string{}
				for key, in := range in.(map[string]interface{}) {
					out[key] = string(ExpandString(in))
				}
				return out
			}(in)
		}(in["feature_gates"]),
		MaxPersistentVolumes: func(in interface{}) *int32 {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *int32 {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in int32) *int32 {
					return &in
				}(int32(ExpandInt(in)))
			}(in)
		}(in["max_persistent_volumes"]),
		Qps: func(in interface{}) *resource.Quantity {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *resource.Quantity {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in resource.Quantity) *resource.Quantity {
					return &in
				}(ExpandQuantity(in))
			}(in)
		}(in["qps"]),
		Burst: func(in interface{}) int32 {
			return int32(ExpandInt(in))
		}(in["burst"]),
		AuthenticationKubeconfig: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["authentication_kubeconfig"]),
		AuthorizationKubeconfig: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["authorization_kubeconfig"]),
		AuthorizationAlwaysAllowPaths: func(in interface{}) []string {
			return func(in interface{}) []string {
				if in == nil {
					return nil
				}
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["authorization_always_allow_paths"]),
		EnableProfiling: func(in interface{}) *bool {
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
		}(in["enable_profiling"]),
	}
}

func FlattenDataSourceKubeSchedulerConfigInto(in kops.KubeSchedulerConfig, out map[string]interface{}) {
	out["master"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Master)
	out["log_level"] = func(in int32) interface{} {
		return FlattenInt(int(in))
	}(in.LogLevel)
	out["image"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Image)
	out["leader_election"] = func(in *kops.LeaderElectionConfiguration) interface{} {
		return func(in *kops.LeaderElectionConfiguration) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.LeaderElectionConfiguration) interface{} {
				return func(in kops.LeaderElectionConfiguration) []map[string]interface{} {
					return []map[string]interface{}{FlattenDataSourceLeaderElectionConfiguration(in)}
				}(in)
			}(*in)
		}(in)
	}(in.LeaderElection)
	out["use_policy_config_map"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.UsePolicyConfigMap)
	out["feature_gates"] = func(in map[string]string) interface{} {
		return func(in map[string]string) map[string]interface{} {
			if in == nil {
				return nil
			}
			out := map[string]interface{}{}
			for key, in := range in {
				out[key] = FlattenString(string(in))
			}
			return out
		}(in)
	}(in.FeatureGates)
	out["max_persistent_volumes"] = func(in *int32) interface{} {
		return func(in *int32) interface{} {
			if in == nil {
				return nil
			}
			return func(in int32) interface{} {
				return FlattenInt(int(in))
			}(*in)
		}(in)
	}(in.MaxPersistentVolumes)
	out["qps"] = func(in *resource.Quantity) interface{} {
		return func(in *resource.Quantity) interface{} {
			if in == nil {
				return nil
			}
			return func(in resource.Quantity) interface{} {
				return FlattenQuantity(in)
			}(*in)
		}(in)
	}(in.Qps)
	out["burst"] = func(in int32) interface{} {
		return FlattenInt(int(in))
	}(in.Burst)
	out["authentication_kubeconfig"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.AuthenticationKubeconfig)
	out["authorization_kubeconfig"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.AuthorizationKubeconfig)
	out["authorization_always_allow_paths"] = func(in []string) interface{} {
		return func(in []string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.AuthorizationAlwaysAllowPaths)
	out["enable_profiling"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.EnableProfiling)
}

func FlattenDataSourceKubeSchedulerConfig(in kops.KubeSchedulerConfig) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceKubeSchedulerConfigInto(in, out)
	return out
}
