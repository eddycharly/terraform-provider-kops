package schemas

import (
	"github.com/eddycharly/terraform-provider-kops/pkg/api/resources"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

func ResourceInstanceGroup() *schema.Resource {
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
			"volumes":                           OptionalList(ResourceVolumeSpec()),
			"volume_mounts":                     OptionalList(ResourceVolumeMountSpec()),
			"subnets":                           RequiredList(String()),
			"zones":                             OptionalList(String()),
			"hooks":                             OptionalList(ResourceHookSpec()),
			"max_price":                         OptionalString(),
			"spot_duration_in_minutes":          OptionalInt(),
			"associate_public_ip":               OptionalBool(),
			"additional_security_groups":        OptionalList(String()),
			"cloud_labels":                      OptionalMap(String()),
			"node_labels":                       OptionalMap(String()),
			"file_assets":                       OptionalList(ResourceFileAssetSpec()),
			"tenancy":                           OptionalString(),
			"kubelet":                           OptionalStruct(ResourceKubeletConfigSpec()),
			"taints":                            OptionalList(String()),
			"mixed_instances_policy":            OptionalStruct(ResourceMixedInstancesPolicySpec()),
			"additional_user_data":              OptionalList(ResourceUserData()),
			"suspend_processes":                 OptionalList(String()),
			"external_load_balancers":           OptionalList(ResourceLoadBalancer()),
			"detailed_instance_monitoring":      OptionalBool(),
			"iam":                               OptionalStruct(ResourceIAMProfileSpec()),
			"security_group_override":           OptionalString(),
			"instance_protection":               OptionalBool(),
			"sysctl_parameters":                 OptionalList(String()),
			"rolling_update":                    OptionalStruct(ResourceRollingUpdate()),
			"instance_interruption_behavior":    OptionalString(),
		},
	}
}

func ExpandResourceInstanceGroup(in map[string]interface{}) resources.InstanceGroup {
	if in == nil {
		panic("expand InstanceGroup failure, in is nil")
	}
	return resources.InstanceGroup{
		Name: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["name"]),
		InstanceGroupSpec: func(in interface{}) kops.InstanceGroupSpec {
			return ExpandResourceInstanceGroupSpec(in.(map[string]interface{}))
		}(in),
	}
}

func FlattenResourceInstanceGroupInto(in resources.InstanceGroup, out map[string]interface{}) {
	out["name"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Name)
	FlattenResourceInstanceGroupSpecInto(in.InstanceGroupSpec, out)
}

func FlattenResourceInstanceGroup(in resources.InstanceGroup) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceInstanceGroupInto(in, out)
	return out
}
