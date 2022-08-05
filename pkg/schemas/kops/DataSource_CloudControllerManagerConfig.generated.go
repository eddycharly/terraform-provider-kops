package schemas

import (
	"reflect"

	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceCloudControllerManagerConfig() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"master":                          ComputedString(),
			"log_level":                       ComputedInt(),
			"image":                           ComputedString(),
			"cloud_provider":                  ComputedString(),
			"cluster_name":                    ComputedString(),
			"cluster_cidr":                    ComputedString(),
			"allocate_node_cidrs":             ComputedBool(),
			"configure_cloud_routes":          ComputedBool(),
			"controllers":                     ComputedList(String()),
			"cidr_allocator_type":             ComputedString(),
			"leader_election":                 ComputedStruct(DataSourceLeaderElectionConfiguration()),
			"use_service_account_credentials": ComputedBool(),
			"enable_leader_migration":         ComputedBool(),
			"cpu_request":                     ComputedQuantity(),
		},
	}

	return res
}

func ExpandDataSourceCloudControllerManagerConfig(in map[string]interface{}) kops.CloudControllerManagerConfig {
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
		Controllers: func(in interface{}) []string {
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
		}(in["controllers"]),
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
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandDataSourceLeaderElectionConfiguration(in[0].(map[string]interface{}))
					}
					return kops.LeaderElectionConfiguration{}
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
		EnableLeaderMigration: func(in interface{}) *bool {
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
		}(in["enable_leader_migration"]),
		CPURequest: func(in interface{}) *resource.Quantity {
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
		}(in["cpu_request"]),
	}
}

func FlattenDataSourceCloudControllerManagerConfigInto(in kops.CloudControllerManagerConfig, out map[string]interface{}) {
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
	out["controllers"] = func(in []string) interface{} {
		return func(in []string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.Controllers)
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
					return []interface{}{FlattenDataSourceLeaderElectionConfiguration(in)}
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
	out["enable_leader_migration"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.EnableLeaderMigration)
	out["cpu_request"] = func(in *resource.Quantity) interface{} {
		return func(in *resource.Quantity) interface{} {
			if in == nil {
				return nil
			}
			return func(in resource.Quantity) interface{} {
				return FlattenQuantity(in)
			}(*in)
		}(in)
	}(in.CPURequest)
}

func FlattenDataSourceCloudControllerManagerConfig(in kops.CloudControllerManagerConfig) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceCloudControllerManagerConfigInto(in, out)
	return out
}
