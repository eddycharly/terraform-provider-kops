package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceKubeSchedulerConfig() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"master":                           Optional(String()),
			"log_level":                        Optional(Int()),
			"image":                            Optional(String()),
			"leader_election":                  Optional(Struct(ResourceLeaderElectionConfiguration())),
			"use_policy_config_map":            Optional(Nullable(Bool())),
			"feature_gates":                    Optional(Map(String())),
			"max_persistent_volumes":           Optional(Nullable(Int())),
			"qps":                              Optional(Nullable(Quantity())),
			"burst":                            Optional(Int()),
			"authentication_kubeconfig":        Optional(String()),
			"authorization_kubeconfig":         Optional(String()),
			"authorization_always_allow_paths": Optional(List(String())),
			"enable_profiling":                 Optional(Nullable(Bool())),
		},
	}
}

func ExpandResourceKubeSchedulerConfig(in map[string]interface{}) kops.KubeSchedulerConfig {
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
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in kops.LeaderElectionConfiguration) *kops.LeaderElectionConfiguration { return &in }(func(in interface{}) kops.LeaderElectionConfiguration {
					return ExpandResourceLeaderElectionConfiguration(in.(map[string]interface{}))
				}(in[0]))
			}
			return nil
		}(in)
	}
	if in, ok := in["use_policy_config_map"]; ok && in != nil {
		out.UsePolicyConfigMap = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
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
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in int32) *int32 { return &in }(func(in interface{}) int32 { return int32(in.(int)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["qps"]; ok && in != nil {
		out.Qps = func(in interface{}) *resource.Quantity {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in resource.Quantity) *resource.Quantity { return &in }(func(in interface{}) resource.Quantity { return ExpandQuantity(in.(string)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
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
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	return out
}

func FlattenResourceKubeSchedulerConfigInto(in kops.KubeSchedulerConfig, out map[string]interface{}) {
	out["master"] = func(in string) interface{} { return string(in) }(in.Master)
	out["log_level"] = func(in int32) interface{} { return int(in) }(in.LogLevel)
	out["image"] = func(in string) interface{} { return string(in) }(in.Image)
	out["leader_election"] = func(in *kops.LeaderElectionConfiguration) interface{} {
		if in == nil {
			return nil
		}
		return func(in kops.LeaderElectionConfiguration) interface{} {
			return FlattenResourceLeaderElectionConfiguration(in)
		}(*in)
	}(in.LeaderElection)
	out["use_policy_config_map"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in bool) interface{} { return in }(*in)}
	}(in.UsePolicyConfigMap)
	out["feature_gates"] = func(in map[string]string) interface{} {
		if in == nil {
			return nil
		}
		out := map[string]interface{}{}
		for key, in := range in {
			out[key] = func(in string) interface{} { return string(in) }(in)
		}
		return out
	}(in.FeatureGates)
	out["max_persistent_volumes"] = func(in *int32) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in int32) interface{} { return int(in) }(*in)}
	}(in.MaxPersistentVolumes)
	out["qps"] = func(in *resource.Quantity) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in resource.Quantity) interface{} { return FlattenQuantity(in) }(*in)}
	}(in.Qps)
	out["burst"] = func(in int32) interface{} { return int(in) }(in.Burst)
	out["authentication_kubeconfig"] = func(in string) interface{} { return string(in) }(in.AuthenticationKubeconfig)
	out["authorization_kubeconfig"] = func(in string) interface{} { return string(in) }(in.AuthorizationKubeconfig)
	out["authorization_always_allow_paths"] = func(in []string) interface{} {
		var out []interface{}
		for _, in := range in {
			out = append(out, func(in string) interface{} { return string(in) }(in))
		}
		return out
	}(in.AuthorizationAlwaysAllowPaths)
	out["enable_profiling"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in bool) interface{} { return in }(*in)}
	}(in.EnableProfiling)
}

func FlattenResourceKubeSchedulerConfig(in kops.KubeSchedulerConfig) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceKubeSchedulerConfigInto(in, out)
	return out
}
