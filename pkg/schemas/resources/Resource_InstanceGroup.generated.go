package schemas

import (
	"github.com/eddycharly/terraform-provider-kops/pkg/api/resources"
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	kopsschemas "github.com/eddycharly/terraform-provider-kops/pkg/schemas/kops"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var _ = Schema

func ResourceInstanceGroup() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"revision":                          Computed(Int()),
			"cluster_name":                      ForceNew(Required(String())),
			"name":                              ForceNew(Required(String())),
			"role":                              Required(String()),
			"image":                             Optional(Computed(String())),
			"min_size":                          Required(Ptr(Int())),
			"max_size":                          Required(Ptr(Int())),
			"autoscale":                         Optional(Ptr(Bool())),
			"machine_type":                      Required(String()),
			"root_volume_size":                  Optional(Ptr(Int())),
			"root_volume_type":                  Optional(Ptr(String())),
			"root_volume_iops":                  Optional(Ptr(Int())),
			"root_volume_throughput":            Optional(Ptr(Int())),
			"root_volume_optimization":          Optional(Ptr(Bool())),
			"root_volume_delete_on_termination": Optional(Ptr(Bool())),
			"root_volume_encryption":            Optional(Ptr(Bool())),
			"root_volume_encryption_key":        Optional(Ptr(String())),
			"volumes":                           Optional(List(Struct(kopsschemas.ResourceVolumeSpec()))),
			"volume_mounts":                     Optional(List(Struct(kopsschemas.ResourceVolumeMountSpec()))),
			"subnets":                           Required(List(String())),
			"zones":                             Optional(List(String())),
			"hooks":                             Optional(List(Struct(kopsschemas.ResourceHookSpec()))),
			"max_price":                         Optional(Ptr(String())),
			"spot_duration_in_minutes":          Optional(Ptr(Int())),
			"cpu_credits":                       Optional(Ptr(String())),
			"associate_public_ip":               Optional(Ptr(Bool())),
			"additional_security_groups":        Optional(List(String())),
			"cloud_labels":                      Optional(Map(String())),
			"node_labels":                       Optional(Map(String())),
			"file_assets":                       Optional(List(Struct(kopsschemas.ResourceFileAssetSpec()))),
			"tenancy":                           Optional(String()),
			"kubelet":                           Optional(Ptr(Struct(kopsschemas.ResourceKubeletConfigSpec()))),
			"taints":                            Optional(List(String())),
			"mixed_instances_policy":            Optional(Ptr(Struct(kopsschemas.ResourceMixedInstancesPolicySpec()))),
			"additional_user_data":              Optional(List(Struct(kopsschemas.ResourceUserData()))),
			"suspend_processes":                 Optional(List(String())),
			"external_load_balancers":           Optional(List(Struct(kopsschemas.ResourceLoadBalancer()))),
			"detailed_instance_monitoring":      Optional(Ptr(Bool())),
			"iam":                               Optional(Ptr(Struct(kopsschemas.ResourceIAMProfileSpec()))),
			"security_group_override":           Optional(Ptr(String())),
			"instance_protection":               Optional(Ptr(Bool())),
			"sysctl_parameters":                 Optional(List(String())),
			"rolling_update":                    Optional(Ptr(Struct(kopsschemas.ResourceRollingUpdate()))),
			"instance_interruption_behavior":    Optional(Ptr(String())),
			"compress_user_data":                Optional(Ptr(Bool())),
			"instance_metadata":                 Optional(Ptr(Struct(kopsschemas.ResourceInstanceMetadataOptions()))),
			"update_policy":                     Optional(Ptr(String())),
		},
	}
}

func ExpandResourceInstanceGroup(in map[string]interface{}) resources.InstanceGroup {
	if in == nil {
		panic("expand InstanceGroup failure, in is nil")
	}
	out := resources.InstanceGroup{}
	if in, ok := in["revision"]; ok && in != nil {
		out.Revision = func(in interface{}) int { return int(in.(int)) }(in)
	}
	if in, ok := in["cluster_name"]; ok && in != nil {
		out.ClusterName = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["name"]; ok && in != nil {
		out.Name = func(in interface{}) string { return string(in.(string)) }(in)
	}
	out.InstanceGroupSpec = kopsschemas.ExpandResourceInstanceGroupSpec(in)
	return out
}
