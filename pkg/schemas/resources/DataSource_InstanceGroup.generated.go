package schemas

import (
	"github.com/eddycharly/terraform-provider-kops/pkg/api/resources"
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	kopsschemas "github.com/eddycharly/terraform-provider-kops/pkg/schemas/kops"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var _ = Schema

func DataSourceInstanceGroup() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cluster_name":                      Required(String()),
			"name":                              Required(String()),
			"role":                              Computed(String()),
			"image":                             Computed(String()),
			"min_size":                          Computed(Ptr(Int())),
			"max_size":                          Computed(Ptr(Int())),
			"autoscale":                         Computed(Ptr(Bool())),
			"machine_type":                      Computed(String()),
			"root_volume_size":                  Computed(Ptr(Int())),
			"root_volume_type":                  Computed(Ptr(String())),
			"root_volume_iops":                  Computed(Ptr(Int())),
			"root_volume_throughput":            Computed(Ptr(Int())),
			"root_volume_optimization":          Computed(Ptr(Bool())),
			"root_volume_delete_on_termination": Computed(Ptr(Bool())),
			"root_volume_encryption":            Computed(Ptr(Bool())),
			"root_volume_encryption_key":        Computed(Ptr(String())),
			"volumes":                           Computed(List(Struct(kopsschemas.DataSourceVolumeSpec()))),
			"volume_mounts":                     Computed(List(Struct(kopsschemas.DataSourceVolumeMountSpec()))),
			"subnets":                           Computed(List(String())),
			"zones":                             Computed(List(String())),
			"hooks":                             Computed(List(Struct(kopsschemas.DataSourceHookSpec()))),
			"max_price":                         Computed(Ptr(String())),
			"spot_duration_in_minutes":          Computed(Ptr(Int())),
			"cpu_credits":                       Computed(Ptr(String())),
			"associate_public_ip":               Computed(Ptr(Bool())),
			"additional_security_groups":        Computed(List(String())),
			"cloud_labels":                      Computed(Map(String())),
			"node_labels":                       Computed(Map(String())),
			"file_assets":                       Computed(List(Struct(kopsschemas.DataSourceFileAssetSpec()))),
			"tenancy":                           Computed(String()),
			"kubelet":                           Computed(Ptr(Struct(kopsschemas.DataSourceKubeletConfigSpec()))),
			"taints":                            Computed(List(String())),
			"mixed_instances_policy":            Computed(Ptr(Struct(kopsschemas.DataSourceMixedInstancesPolicySpec()))),
			"additional_user_data":              Computed(List(Struct(kopsschemas.DataSourceUserData()))),
			"suspend_processes":                 Computed(List(String())),
			"external_load_balancers":           Computed(List(Struct(kopsschemas.DataSourceLoadBalancer()))),
			"detailed_instance_monitoring":      Computed(Ptr(Bool())),
			"iam":                               Computed(Ptr(Struct(kopsschemas.DataSourceIAMProfileSpec()))),
			"security_group_override":           Computed(Ptr(String())),
			"instance_protection":               Computed(Ptr(Bool())),
			"sysctl_parameters":                 Computed(List(String())),
			"rolling_update":                    Computed(Ptr(Struct(kopsschemas.DataSourceRollingUpdate()))),
			"instance_interruption_behavior":    Computed(Ptr(String())),
			"compress_user_data":                Computed(Ptr(Bool())),
			"instance_metadata":                 Computed(Ptr(Struct(kopsschemas.DataSourceInstanceMetadataOptions()))),
			"update_policy":                     Computed(Ptr(String())),
		},
	}
}

func ExpandDataSourceInstanceGroup(in map[string]interface{}) resources.InstanceGroup {
	if in == nil {
		panic("expand InstanceGroup failure, in is nil")
	}
	out := resources.InstanceGroup{}
	if in, ok := in["cluster_name"]; ok && in != nil {
		out.ClusterName = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["name"]; ok && in != nil {
		out.Name = func(in interface{}) string { return string(in.(string)) }(in)
	}
	out.InstanceGroupSpec = kopsschemas.ExpandDataSourceInstanceGroupSpec(in)
	return out
}
