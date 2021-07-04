package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceCloudControllerManagerConfig() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"master":                          Computed(String()),
			"log_level":                       Computed(Int()),
			"image":                           Computed(String()),
			"cloud_provider":                  Computed(String()),
			"cluster_name":                    Computed(String()),
			"cluster_cidr":                    Computed(String()),
			"allocate_node_cidrs":             Computed(Nullable(Bool())),
			"configure_cloud_routes":          Computed(Nullable(Bool())),
			"cidr_allocator_type":             Computed(Nullable(String())),
			"leader_election":                 Computed(Struct(DataSourceLeaderElectionConfiguration())),
			"use_service_account_credentials": Computed(Nullable(Bool())),
		},
	}
}

func ExpandDataSourceCloudControllerManagerConfig(in map[string]interface{}) kops.CloudControllerManagerConfig {
	if in == nil {
		panic("expand CloudControllerManagerConfig failure, in is nil")
	}
	out := kops.CloudControllerManagerConfig{}
	if in, ok := in["master"]; ok && in != nil {
		out.Master = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["log_level"]; ok && in != nil {
		out.LogLevel = func(in interface{}) int32 { return int32(in.(int)) }(in)
	}
	if in, ok := in["image"]; ok && in != nil {
		out.Image = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["cloud_provider"]; ok && in != nil {
		out.CloudProvider = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["cluster_name"]; ok && in != nil {
		out.ClusterName = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["cluster_cidr"]; ok && in != nil {
		out.ClusterCIDR = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["allocate_node_cidrs"]; ok && in != nil {
		out.AllocateNodeCIDRs = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["configure_cloud_routes"]; ok && in != nil {
		out.ConfigureCloudRoutes = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["cidr_allocator_type"]; ok && in != nil {
		out.CIDRAllocatorType = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["leader_election"]; ok && in != nil {
		out.LeaderElection = func(in interface{}) *kops.LeaderElectionConfiguration {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in kops.LeaderElectionConfiguration) *kops.LeaderElectionConfiguration { return &in }(func(in interface{}) kops.LeaderElectionConfiguration {
					return ExpandDataSourceLeaderElectionConfiguration(in.(map[string]interface{}))
				}(in[0]))
			}
			return nil
		}(in)
	}
	if in, ok := in["use_service_account_credentials"]; ok && in != nil {
		out.UseServiceAccountCredentials = func(in interface{}) *bool {
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

func FlattenDataSourceCloudControllerManagerConfigInto(in kops.CloudControllerManagerConfig, out map[string]interface{}) {
	out["master"] = func(in string) interface{} { return string(in) }(in.Master)
	out["log_level"] = func(in int32) interface{} { return int(in) }(in.LogLevel)
	out["image"] = func(in string) interface{} { return string(in) }(in.Image)
	out["cloud_provider"] = func(in string) interface{} { return string(in) }(in.CloudProvider)
	out["cluster_name"] = func(in string) interface{} { return string(in) }(in.ClusterName)
	out["cluster_cidr"] = func(in string) interface{} { return string(in) }(in.ClusterCIDR)
	out["allocate_node_cidrs"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in bool) interface{} { return in }(*in)}
	}(in.AllocateNodeCIDRs)
	out["configure_cloud_routes"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in bool) interface{} { return in }(*in)}
	}(in.ConfigureCloudRoutes)
	out["cidr_allocator_type"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.CIDRAllocatorType)
	out["leader_election"] = func(in *kops.LeaderElectionConfiguration) interface{} {
		if in == nil {
			return nil
		}
		return func(in kops.LeaderElectionConfiguration) interface{} {
			return FlattenDataSourceLeaderElectionConfiguration(in)
		}(*in)
	}(in.LeaderElection)
	out["use_service_account_credentials"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in bool) interface{} { return in }(*in)}
	}(in.UseServiceAccountCredentials)
}

func FlattenDataSourceCloudControllerManagerConfig(in kops.CloudControllerManagerConfig) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceCloudControllerManagerConfigInto(in, out)
	return out
}
