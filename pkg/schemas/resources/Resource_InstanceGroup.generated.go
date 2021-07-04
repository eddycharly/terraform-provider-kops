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
			"min_size":                          Required(Nullable(Int())),
			"max_size":                          Required(Nullable(Int())),
			"autoscale":                         Optional(Nullable(Bool())),
			"machine_type":                      Required(String()),
			"root_volume_size":                  Optional(Nullable(Int())),
			"root_volume_type":                  Optional(Nullable(String())),
			"root_volume_iops":                  Optional(Nullable(Int())),
			"root_volume_throughput":            Optional(Nullable(Int())),
			"root_volume_optimization":          Optional(Nullable(Bool())),
			"root_volume_delete_on_termination": Optional(Nullable(Bool())),
			"root_volume_encryption":            Optional(Nullable(Bool())),
			"root_volume_encryption_key":        Optional(Nullable(String())),
			"volumes":                           Optional(List(kopsschemas.ResourceVolumeSpec())),
			"volume_mounts":                     Optional(List(kopsschemas.ResourceVolumeMountSpec())),
			"subnets":                           Required(List(String())),
			"zones":                             Optional(List(String())),
			"hooks":                             Optional(List(kopsschemas.ResourceHookSpec())),
			"max_price":                         Optional(Nullable(String())),
			"spot_duration_in_minutes":          Optional(Nullable(Int())),
			"cpu_credits":                       Optional(Nullable(String())),
			"associate_public_ip":               Optional(Nullable(Bool())),
			"additional_security_groups":        Optional(List(String())),
			"cloud_labels":                      Optional(Map(String())),
			"node_labels":                       Optional(Map(String())),
			"file_assets":                       Optional(List(kopsschemas.ResourceFileAssetSpec())),
			"tenancy":                           Optional(String()),
			"kubelet":                           Optional(Struct(kopsschemas.ResourceKubeletConfigSpec())),
			"taints":                            Optional(List(String())),
			"mixed_instances_policy":            Optional(Struct(kopsschemas.ResourceMixedInstancesPolicySpec())),
			"additional_user_data":              Optional(List(kopsschemas.ResourceUserData())),
			"suspend_processes":                 Optional(List(String())),
			"external_load_balancers":           Optional(List(kopsschemas.ResourceLoadBalancer())),
			"detailed_instance_monitoring":      Optional(Nullable(Bool())),
			"iam":                               Optional(Struct(kopsschemas.ResourceIAMProfileSpec())),
			"security_group_override":           Optional(Nullable(String())),
			"instance_protection":               Optional(Nullable(Bool())),
			"sysctl_parameters":                 Optional(List(String())),
			"rolling_update":                    Optional(Struct(kopsschemas.ResourceRollingUpdate())),
			"instance_interruption_behavior":    Optional(Nullable(String())),
			"compress_user_data":                Optional(Nullable(Bool())),
			"instance_metadata":                 Optional(Struct(kopsschemas.ResourceInstanceMetadataOptions())),
			"update_policy":                     Optional(Nullable(String())),
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

func FlattenResourceInstanceGroupInto(in resources.InstanceGroup, out map[string]interface{}) {
	out["revision"] = func(in int) interface{} { return int(in) }(in.Revision)
	out["cluster_name"] = func(in string) interface{} { return string(in) }(in.ClusterName)
	out["name"] = func(in string) interface{} { return string(in) }(in.Name)
	kopsschemas.FlattenResourceInstanceGroupSpecInto(in.InstanceGroupSpec, out)
}

func FlattenResourceInstanceGroup(in resources.InstanceGroup) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceInstanceGroupInto(in, out)
	return out
}
