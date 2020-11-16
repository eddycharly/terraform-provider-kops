package schemas

import (
	"github.com/eddycharly/terraform-provider-kops/pkg/api/resources"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

func ResourceResourcesInstanceGroup() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name":                              RequiredString(),
			"role":                              RequiredString(),
			"image":                             OptionalComputedString(),
			"min_size":                          RequiredInt(),
			"max_size":                          RequiredInt(),
			"machine_type":                      RequiredString(),
			"root_volume_size":                  OptionalInt(),
			"root_volume_type":                  OptionalString(),
			"root_volume_iops":                  OptionalInt(),
			"root_volume_optimization":          OptionalBool(),
			"root_volume_delete_on_termination": OptionalBool(),
			"root_volume_encryption":            OptionalBool(),
			"volumes":                           OptionalList(ResourceKopsVolumeSpec()),
			"volume_mounts":                     OptionalList(ResourceKopsVolumeMountSpec()),
			"subnets":                           RequiredList(String()),
			"zones":                             OptionalList(String()),
			"hooks":                             OptionalList(ResourceKopsHookSpec()),
			"max_price":                         OptionalString(),
			"spot_duration_in_minutes":          OptionalInt(),
			"associate_public_ip":               OptionalBool(),
			"additional_security_groups":        OptionalList(String()),
			"cloud_labels":                      OptionalMap(String()),
			"node_labels":                       OptionalMap(String()),
			"file_assets":                       OptionalList(ResourceKopsFileAssetSpec()),
			"tenancy":                           OptionalString(),
			"kubelet":                           OptionalStruct(ResourceKopsKubeletConfigSpec()),
			"taints":                            OptionalList(String()),
			"mixed_instances_policy":            OptionalStruct(ResourceKopsMixedInstancesPolicySpec()),
			"additional_user_data":              OptionalList(ResourceKopsUserData()),
			"suspend_processes":                 OptionalList(String()),
			"external_load_balancers":           OptionalList(ResourceKopsLoadBalancer()),
			"detailed_instance_monitoring":      OptionalBool(),
			"iam":                               OptionalStruct(ResourceKopsIAMProfileSpec()),
			"security_group_override":           OptionalString(),
			"instance_protection":               OptionalBool(),
			"sysctl_parameters":                 OptionalList(String()),
			"rolling_update":                    OptionalStruct(ResourceKopsRollingUpdate()),
			"instance_interruption_behavior":    OptionalString(),
		},
	}
}

func ExpandResourceResourcesInstanceGroup(in map[string]interface{}) resources.InstanceGroup {
	if in == nil {
		panic("expand InstanceGroup failure, in is nil")
	}
	return resources.InstanceGroup{
		Name: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["name"]),
		InstanceGroupSpec: func(in interface{}) kops.InstanceGroupSpec {
			return ExpandResourceKopsInstanceGroupSpec(in.(map[string]interface{}))
		}(in),
	}
}

func FlattenResourceResourcesInstanceGroupInto(in resources.InstanceGroup, out map[string]interface{}) {
	out["name"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Name)
	FlattenResourceKopsInstanceGroupSpecInto(in.InstanceGroupSpec, out)
}

func FlattenResourceResourcesInstanceGroup(in resources.InstanceGroup) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceResourcesInstanceGroupInto(in, out)
	return out
}
