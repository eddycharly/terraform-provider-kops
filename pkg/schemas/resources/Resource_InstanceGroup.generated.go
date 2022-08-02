package schemas

import (
	"context"

	"github.com/eddycharly/terraform-provider-kops/pkg/api/resources"
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	kopsschemas "github.com/eddycharly/terraform-provider-kops/pkg/schemas/kops"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceInstanceGroup() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"manager":                        OptionalString(),
			"role":                           RequiredString(),
			"image":                          OptionalComputedString(),
			"min_size":                       RequiredInt(),
			"max_size":                       RequiredInt(),
			"autoscale":                      OptionalBool(),
			"machine_type":                   RequiredString(),
			"root_volume_size":               OptionalInt(),
			"root_volume_type":               OptionalString(),
			"root_volume_iops":               OptionalInt(),
			"root_volume_throughput":         OptionalInt(),
			"root_volume_optimization":       OptionalBool(),
			"root_volume_encryption":         OptionalBool(),
			"root_volume_encryption_key":     OptionalString(),
			"volumes":                        OptionalList(kopsschemas.ResourceVolumeSpec()),
			"volume_mounts":                  OptionalList(kopsschemas.ResourceVolumeMountSpec()),
			"subnets":                        RequiredList(String()),
			"zones":                          OptionalList(String()),
			"hooks":                          OptionalList(kopsschemas.ResourceHookSpec()),
			"max_price":                      OptionalString(),
			"spot_duration_in_minutes":       OptionalInt(),
			"cpu_credits":                    OptionalString(),
			"associate_public_ip":            OptionalBool(),
			"additional_security_groups":     OptionalList(String()),
			"cloud_labels":                   OptionalMap(String()),
			"node_labels":                    OptionalMap(String()),
			"file_assets":                    OptionalList(kopsschemas.ResourceFileAssetSpec()),
			"tenancy":                        OptionalString(),
			"kubelet":                        OptionalStruct(kopsschemas.ResourceKubeletConfigSpec()),
			"taints":                         OptionalList(String()),
			"mixed_instances_policy":         OptionalStruct(kopsschemas.ResourceMixedInstancesPolicySpec()),
			"additional_user_data":           OptionalList(kopsschemas.ResourceUserData()),
			"suspend_processes":              OptionalList(String()),
			"external_load_balancers":        OptionalList(kopsschemas.ResourceLoadBalancer()),
			"detailed_instance_monitoring":   OptionalBool(),
			"iam":                            OptionalStruct(kopsschemas.ResourceIAMProfileSpec()),
			"security_group_override":        OptionalString(),
			"instance_protection":            OptionalBool(),
			"sysctl_parameters":              OptionalList(String()),
			"rolling_update":                 OptionalStruct(kopsschemas.ResourceRollingUpdate()),
			"instance_interruption_behavior": OptionalString(),
			"compress_user_data":             OptionalBool(),
			"instance_metadata":              OptionalStruct(kopsschemas.ResourceInstanceMetadataOptions()),
			"update_policy":                  OptionalString(),
			"warm_pool":                      OptionalStruct(kopsschemas.ResourceWarmPoolSpec()),
			"containerd":                     OptionalStruct(kopsschemas.ResourceContainerdConfig()),
			"packages":                       OptionalList(String()),
			"guest_accelerators":             OptionalList(kopsschemas.ResourceAcceleratorConfig()),
			"max_instance_lifetime":          OptionalDuration(),
			"labels":                         OptionalMap(String()),
			"annotations":                    OptionalMap(String()),
			"cluster_name":                   ForceNew(RequiredString()),
			"name":                           ForceNew(RequiredString()),
			"revision":                       ComputedInt(),
		},
	}
	res.SchemaVersion = 2
	res.StateUpgraders = []schema.StateUpgrader{
		{
			Type: res.CoreConfigSchema().ImpliedType(),
			Upgrade: func(ctx context.Context, rawState map[string]interface{}, meta interface{}) (map[string]interface{}, error) {
				ret := FlattenResourceInstanceGroup(ExpandResourceInstanceGroup(rawState))
				ret["id"] = rawState["id"]
				return ret, nil
			},
			Version: 0,
		}, {
			Type: res.CoreConfigSchema().ImpliedType(),
			Upgrade: func(ctx context.Context, rawState map[string]interface{}, meta interface{}) (map[string]interface{}, error) {
				ret := FlattenResourceInstanceGroup(ExpandResourceInstanceGroup(rawState))
				ret["id"] = rawState["id"]
				return ret, nil
			},
			Version: 1,
		},
	}
	return res
}

func ExpandResourceInstanceGroup(in map[string]interface{}) resources.InstanceGroup {
	if in == nil {
		panic("expand InstanceGroup failure, in is nil")
	}
	return resources.InstanceGroup{
		InstanceGroupSpec: func(in interface{}) kops.InstanceGroupSpec {
			return kopsschemas.ExpandResourceInstanceGroupSpec(in.(map[string]interface{}))
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
		Revision: func(in interface{}) int {
			return int(ExpandInt(in))
		}(in["revision"]),
	}
}

func FlattenResourceInstanceGroupInto(in resources.InstanceGroup, out map[string]interface{}) {
	kopsschemas.FlattenResourceInstanceGroupSpecInto(in.InstanceGroupSpec, out)
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
	out["revision"] = func(in int) interface{} {
		return FlattenInt(int(in))
	}(in.Revision)
}

func FlattenResourceInstanceGroup(in resources.InstanceGroup) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceInstanceGroupInto(in, out)
	return out
}
