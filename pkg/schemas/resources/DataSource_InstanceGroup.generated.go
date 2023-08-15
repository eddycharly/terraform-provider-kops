package schemas

import (
	"context"

	"github.com/eddycharly/terraform-provider-kops/pkg/api/resources"
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	kopsv1alpha2schemas "github.com/eddycharly/terraform-provider-kops/pkg/schemas/kopsv1alpha2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

var _ = Schema

func DataSourceInstanceGroup() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"manager":                           ComputedString(),
			"role":                              ComputedString(),
			"image":                             ComputedString(),
			"min_size":                          ComputedInt(),
			"max_size":                          ComputedInt(),
			"autoscale":                         ComputedBool(),
			"autoscale_priority":                ComputedInt(),
			"machine_type":                      ComputedString(),
			"root_volume_size":                  ComputedInt(),
			"root_volume_type":                  ComputedString(),
			"root_volume_iops":                  ComputedInt(),
			"root_volume_throughput":            ComputedInt(),
			"root_volume_optimization":          ComputedBool(),
			"root_volume_delete_on_termination": ComputedBool(),
			"root_volume_encryption":            ComputedBool(),
			"root_volume_encryption_key":        ComputedString(),
			"volumes":                           ComputedList(kopsv1alpha2schemas.DataSourceVolumeSpec()),
			"volume_mounts":                     ComputedList(kopsv1alpha2schemas.DataSourceVolumeMountSpec()),
			"subnets":                           ComputedList(String()),
			"zones":                             ComputedList(String()),
			"hooks":                             ComputedList(kopsv1alpha2schemas.DataSourceHookSpec()),
			"max_price":                         ComputedString(),
			"spot_duration_in_minutes":          ComputedInt(),
			"cpu_credits":                       ComputedString(),
			"associate_public_ip":               ComputedBool(),
			"additional_security_groups":        ComputedList(String()),
			"cloud_labels":                      ComputedMap(String()),
			"node_labels":                       ComputedMap(String()),
			"file_assets":                       ComputedList(kopsv1alpha2schemas.DataSourceFileAssetSpec()),
			"tenancy":                           ComputedString(),
			"kubelet":                           ComputedStruct(kopsv1alpha2schemas.DataSourceKubeletConfigSpec()),
			"taints":                            ComputedList(String()),
			"mixed_instances_policy":            ComputedStruct(kopsv1alpha2schemas.DataSourceMixedInstancesPolicySpec()),
			"capacity_rebalance":                ComputedBool(),
			"additional_user_data":              ComputedList(kopsv1alpha2schemas.DataSourceUserData()),
			"suspend_processes":                 ComputedList(String()),
			"external_load_balancers":           ComputedList(kopsv1alpha2schemas.DataSourceLoadBalancerSpec()),
			"detailed_instance_monitoring":      ComputedBool(),
			"iam":                               ComputedStruct(kopsv1alpha2schemas.DataSourceIAMProfileSpec()),
			"security_group_override":           ComputedString(),
			"instance_protection":               ComputedBool(),
			"sysctl_parameters":                 ComputedList(String()),
			"rolling_update":                    ComputedStruct(kopsv1alpha2schemas.DataSourceRollingUpdate()),
			"instance_interruption_behavior":    ComputedString(),
			"compress_user_data":                ComputedBool(),
			"instance_metadata":                 ComputedStruct(kopsv1alpha2schemas.DataSourceInstanceMetadataOptions()),
			"update_policy":                     ComputedString(),
			"warm_pool":                         ComputedStruct(kopsv1alpha2schemas.DataSourceWarmPoolSpec()),
			"containerd":                        ComputedStruct(kopsv1alpha2schemas.DataSourceContainerdConfig()),
			"packages":                          ComputedList(String()),
			"guest_accelerators":                ComputedList(kopsv1alpha2schemas.DataSourceAcceleratorConfig()),
			"max_instance_lifetime":             ComputedDuration(),
			"gcp_provisioning_model":            ComputedString(),
			"labels":                            ComputedMap(String()),
			"annotations":                       ComputedMap(String()),
			"cluster_name":                      RequiredString(),
			"name":                              RequiredString(),
		},
	}
	res.SchemaVersion = 2
	res.StateUpgraders = []schema.StateUpgrader{
		{
			Type: res.CoreConfigSchema().ImpliedType(),
			Upgrade: func(ctx context.Context, rawState map[string]interface{}, meta interface{}) (map[string]interface{}, error) {
				ret := FlattenDataSourceInstanceGroup(ExpandDataSourceInstanceGroup(rawState))
				ret["id"] = rawState["id"]
				return ret, nil
			},
			Version: 0,
		}, {
			Type: res.CoreConfigSchema().ImpliedType(),
			Upgrade: func(ctx context.Context, rawState map[string]interface{}, meta interface{}) (map[string]interface{}, error) {
				ret := FlattenDataSourceInstanceGroup(ExpandDataSourceInstanceGroup(rawState))
				ret["id"] = rawState["id"]
				return ret, nil
			},
			Version: 1,
		},
	}
	return res
}

func ExpandDataSourceInstanceGroup(in map[string]interface{}) resources.InstanceGroup {
	if in == nil {
		panic("expand InstanceGroup failure, in is nil")
	}
	return resources.InstanceGroup{
		InstanceGroupSpec: func(in interface{}) kopsv1alpha2.InstanceGroupSpec {
			return kopsv1alpha2schemas.ExpandDataSourceInstanceGroupSpec(in.(map[string]interface{}))
		}(in),
		Labels: func(in interface{}) map[string]string {
			return func(in interface{}) map[string]string {
				if in == nil {
					return nil
				}
				if in, ok := in.(map[string]interface{}); ok {
					if len(in) > 0 {
						out := map[string]string{}
						for key, in := range in {
							out[key] = string(ExpandString(in))
						}
						return out
					}
				}
				return nil
			}(in)
		}(in["labels"]),
		Annotations: func(in interface{}) map[string]string {
			return func(in interface{}) map[string]string {
				if in == nil {
					return nil
				}
				if in, ok := in.(map[string]interface{}); ok {
					if len(in) > 0 {
						out := map[string]string{}
						for key, in := range in {
							out[key] = string(ExpandString(in))
						}
						return out
					}
				}
				return nil
			}(in)
		}(in["annotations"]),
		ClusterName: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["cluster_name"]),
		Name: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["name"]),
	}
}

func FlattenDataSourceInstanceGroupInto(in resources.InstanceGroup, out map[string]interface{}) {
	kopsv1alpha2schemas.FlattenDataSourceInstanceGroupSpecInto(in.InstanceGroupSpec, out)
	out["labels"] = func(in map[string]string) interface{} {
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
	}(in.Labels)
	out["annotations"] = func(in map[string]string) interface{} {
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
	}(in.Annotations)
	out["cluster_name"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.ClusterName)
	out["name"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Name)
}

func FlattenDataSourceInstanceGroup(in resources.InstanceGroup) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceInstanceGroupInto(in, out)
	return out
}
