package structures

import (
	"reflect"

	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/kops/pkg/apis/kops"
)

func ExpandKubeSchedulerConfig(in map[string]interface{}) kops.KubeSchedulerConfig {
	if in == nil {
		panic("expand KubeSchedulerConfig failure, in is nil")
	}
	return kops.KubeSchedulerConfig{
		Master: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["master"]),
		LogLevel: func(in interface{}) int32 {
			value := int32(ExpandInt(in))
			return value
		}(in["log_level"]),
		Image: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["image"]),
		LeaderElection: func(in interface{}) *kops.LeaderElectionConfiguration {
			value := func(in interface{}) *kops.LeaderElectionConfiguration {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in kops.LeaderElectionConfiguration) *kops.LeaderElectionConfiguration {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(func(in interface{}) kops.LeaderElectionConfiguration {
					if in.([]interface{})[0] == nil {
						return kops.LeaderElectionConfiguration{}
					}
					return (ExpandLeaderElectionConfiguration(in.([]interface{})[0].(map[string]interface{})))
				}(in))
				return tmp
			}(in)
			return value
		}(in["leader_election"]),
		UsePolicyConfigMap: func(in interface{}) *bool {
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
		}(in["use_policy_config_map"]),
		FeatureGates: func(in interface{}) map[string]string {
			value := func(in interface{}) map[string]string {
				if in == nil {
					return nil
				}
				out := map[string]string{}
				for key, in := range in.(map[string]interface{}) {
					out[key] = string(ExpandString(in))
				}
				return out
			}(in)
			return value
		}(in["feature_gates"]),
		MaxPersistentVolumes: func(in interface{}) *int32 {
			value := func(in interface{}) *int32 {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in int32) *int32 {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(int32(ExpandInt(in)))
				return tmp
			}(in)
			return value
		}(in["max_persistent_volumes"]),
		Qps: func(in interface{}) *resource.Quantity {
			value := func(in interface{}) *resource.Quantity {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in resource.Quantity) *resource.Quantity {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(ExpandQuantity(in))
				return tmp
			}(in)
			return value
		}(in["qps"]),
		Burst: func(in interface{}) int32 {
			value := int32(ExpandInt(in))
			return value
		}(in["burst"]),
		EnableProfiling: func(in interface{}) *bool {
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
		}(in["enable_profiling"]),
	}
}

func FlattenKubeSchedulerConfig(in kops.KubeSchedulerConfig) map[string]interface{} {
	return map[string]interface{}{
		"master": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.Master),
		"log_level": func(in int32) interface{} {
			value := FlattenInt(int(in))
			return value
		}(in.LogLevel),
		"image": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.Image),
		"leader_election": func(in *kops.LeaderElectionConfiguration) interface{} {
			value := func(in *kops.LeaderElectionConfiguration) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.LeaderElectionConfiguration) interface{} {
					return func(in kops.LeaderElectionConfiguration) []map[string]interface{} {
						return []map[string]interface{}{FlattenLeaderElectionConfiguration(in)}
					}(in)
				}(*in)
			}(in)
			return value
		}(in.LeaderElection),
		"use_policy_config_map": func(in *bool) interface{} {
			value := func(in *bool) interface{} {
				if in == nil {
					return nil
				}
				return func(in bool) interface{} {
					return FlattenBool(bool(in))
				}(*in)
			}(in)
			return value
		}(in.UsePolicyConfigMap),
		"feature_gates": func(in map[string]string) interface{} {
			value := func(in map[string]string) map[string]interface{} {
				if in == nil {
					return nil
				}
				// TODO
				return nil
			}(in)
			return value
		}(in.FeatureGates),
		"max_persistent_volumes": func(in *int32) interface{} {
			value := func(in *int32) interface{} {
				if in == nil {
					return nil
				}
				return func(in int32) interface{} {
					return FlattenInt(int(in))
				}(*in)
			}(in)
			return value
		}(in.MaxPersistentVolumes),
		"qps": func(in *resource.Quantity) interface{} {
			value := func(in *resource.Quantity) interface{} {
				if in == nil {
					return nil
				}
				return func(in resource.Quantity) interface{} {
					return FlattenQuantity(in)
				}(*in)
			}(in)
			return value
		}(in.Qps),
		"burst": func(in int32) interface{} {
			value := FlattenInt(int(in))
			return value
		}(in.Burst),
		"enable_profiling": func(in *bool) interface{} {
			value := func(in *bool) interface{} {
				if in == nil {
					return nil
				}
				return func(in bool) interface{} {
					return FlattenBool(bool(in))
				}(*in)
			}(in)
			return value
		}(in.EnableProfiling),
	}
}
