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
			"allocate_node_cidrs":             Computed(Ptr(Bool())),
			"configure_cloud_routes":          Computed(Ptr(Bool())),
			"cidr_allocator_type":             Computed(Ptr(String())),
			"leader_election":                 Computed(Ptr(Struct(DataSourceLeaderElectionConfiguration()))),
			"use_service_account_credentials": Computed(Ptr(Bool())),
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
			return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["configure_cloud_routes"]; ok && in != nil {
		out.ConfigureCloudRoutes = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["cidr_allocator_type"]; ok && in != nil {
		out.CIDRAllocatorType = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
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
	if in, ok := in["use_service_account_credentials"]; ok && in != nil {
		out.UseServiceAccountCredentials = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	return out
}
