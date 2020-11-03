package structures

import (
	"reflect"

	"k8s.io/kops/pkg/apis/kops"
)

func ExpandCloudControllerManagerConfig(in map[string]interface{}) kops.CloudControllerManagerConfig {
	if in == nil {
		panic("expand CloudControllerManagerConfig failure, in is nil")
	}
	return kops.CloudControllerManagerConfig{
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
		CloudProvider: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["cloud_provider"]),
		ClusterName: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["cluster_name"]),
		ClusterCIDR: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["cluster_cidr"]),
		AllocateNodeCIDRs: func(in interface{}) *bool {
			value := func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in bool) *bool {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(bool(ExpandBool(in)))
			}(in)
			return value
		}(in["allocate_node_cidrs"]),
		ConfigureCloudRoutes: func(in interface{}) *bool {
			value := func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in bool) *bool {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(bool(ExpandBool(in)))
			}(in)
			return value
		}(in["configure_cloud_routes"]),
		CIDRAllocatorType: func(in interface{}) *string {
			value := func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in string) *string {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(string(ExpandString(in)))
			}(in)
			return value
		}(in["cidr_allocator_type"]),
		LeaderElection: func(in interface{}) *kops.LeaderElectionConfiguration {
			value := func(in interface{}) *kops.LeaderElectionConfiguration {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in kops.LeaderElectionConfiguration) *kops.LeaderElectionConfiguration {
					return &in
				}(func(in interface{}) kops.LeaderElectionConfiguration {
					if in.([]interface{})[0] == nil {
						return kops.LeaderElectionConfiguration{}
					}
					return (ExpandLeaderElectionConfiguration(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
			return value
		}(in["leader_election"]),
		UseServiceAccountCredentials: func(in interface{}) *bool {
			value := func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in bool) *bool {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(bool(ExpandBool(in)))
			}(in)
			return value
		}(in["use_service_account_credentials"]),
	}
}

func FlattenCloudControllerManagerConfig(in kops.CloudControllerManagerConfig) map[string]interface{} {
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
		"cloud_provider": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.CloudProvider),
		"cluster_name": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.ClusterName),
		"cluster_cidr": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.ClusterCIDR),
		"allocate_node_cidrs": func(in *bool) interface{} {
			value := func(in *bool) interface{} {
				if in == nil {
					return nil
				}
				return func(in bool) interface{} {
					return FlattenBool(bool(in))
				}(*in)
			}(in)
			return value
		}(in.AllocateNodeCIDRs),
		"configure_cloud_routes": func(in *bool) interface{} {
			value := func(in *bool) interface{} {
				if in == nil {
					return nil
				}
				return func(in bool) interface{} {
					return FlattenBool(bool(in))
				}(*in)
			}(in)
			return value
		}(in.ConfigureCloudRoutes),
		"cidr_allocator_type": func(in *string) interface{} {
			value := func(in *string) interface{} {
				if in == nil {
					return nil
				}
				return func(in string) interface{} {
					return FlattenString(string(in))
				}(*in)
			}(in)
			return value
		}(in.CIDRAllocatorType),
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
		"use_service_account_credentials": func(in *bool) interface{} {
			value := func(in *bool) interface{} {
				if in == nil {
					return nil
				}
				return func(in bool) interface{} {
					return FlattenBool(bool(in))
				}(*in)
			}(in)
			return value
		}(in.UseServiceAccountCredentials),
	}
}
