package schemas

import (
	"reflect"

	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceCloudControllerManagerConfig() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"master":                          OptionalString(),
			"log_level":                       OptionalInt(),
			"image":                           OptionalString(),
			"cloud_provider":                  OptionalString(),
			"cluster_name":                    OptionalString(),
			"cluster_cidr":                    OptionalString(),
			"allocate_node_cidrs":             OptionalBool(),
			"configure_cloud_routes":          OptionalBool(),
			"cidr_allocator_type":             OptionalString(),
			"leader_election":                 OptionalStruct(ResourceLeaderElectionConfiguration()),
			"use_service_account_credentials": OptionalBool(),
		},
	}

	return res
}

func ExpandResourceCloudControllerManagerConfig(in map[string]interface{}) kops.CloudControllerManagerConfig {
	if in == nil {
		panic("expand CloudControllerManagerConfig failure, in is nil")
	}
	return kops.CloudControllerManagerConfig{
		Master: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["master"]),
		LogLevel: func(in interface{}) int32 {
			return int32(ExpandInt(in))
		}(in["log_level"]),
		Image: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["image"]),
		CloudProvider: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["cloud_provider"]),
		ClusterName: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["cluster_name"]),
		ClusterCIDR: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["cluster_cidr"]),
		AllocateNodeCIDRs: func(in interface{}) *bool {
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
		}(in["allocate_node_cidrs"]),
		ConfigureCloudRoutes: func(in interface{}) *bool {
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
		}(in["configure_cloud_routes"]),
		CIDRAllocatorType: func(in interface{}) *string {
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
		}(in["cidr_allocator_type"]),
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
					return (ExpandResourceLeaderElectionConfiguration(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["leader_election"]),
		UseServiceAccountCredentials: func(in interface{}) *bool {
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
		}(in["use_service_account_credentials"]),
	}
}

func FlattenResourceCloudControllerManagerConfigInto(in kops.CloudControllerManagerConfig, out map[string]interface{}) {
	out["master"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Master)
	out["log_level"] = func(in int32) interface{} {
		return FlattenInt(int(in))
	}(in.LogLevel)
	out["image"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Image)
	out["cloud_provider"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.CloudProvider)
	out["cluster_name"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.ClusterName)
	out["cluster_cidr"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.ClusterCIDR)
	out["allocate_node_cidrs"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.AllocateNodeCIDRs)
	out["configure_cloud_routes"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.ConfigureCloudRoutes)
	out["cidr_allocator_type"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.CIDRAllocatorType)
	out["leader_election"] = func(in *kops.LeaderElectionConfiguration) interface{} {
		return func(in *kops.LeaderElectionConfiguration) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.LeaderElectionConfiguration) interface{} {
				return func(in kops.LeaderElectionConfiguration) []interface{} {
					return []interface{}{FlattenResourceLeaderElectionConfiguration(in)}
				}(in)
			}(*in)
		}(in)
	}(in.LeaderElection)
	out["use_service_account_credentials"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.UseServiceAccountCredentials)
}

func FlattenResourceCloudControllerManagerConfig(in kops.CloudControllerManagerConfig) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceCloudControllerManagerConfigInto(in, out)
	return out
}
