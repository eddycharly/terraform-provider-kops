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
			"min_size":                          Computed(Nullable(Int())),
			"max_size":                          Computed(Nullable(Int())),
			"autoscale":                         Computed(Nullable(Bool())),
			"machine_type":                      Computed(String()),
			"root_volume_size":                  Computed(Nullable(Int())),
			"root_volume_type":                  Computed(Nullable(String())),
			"root_volume_iops":                  Computed(Nullable(Int())),
			"root_volume_throughput":            Computed(Nullable(Int())),
			"root_volume_optimization":          Computed(Nullable(Bool())),
			"root_volume_delete_on_termination": Computed(Nullable(Bool())),
			"root_volume_encryption":            Computed(Nullable(Bool())),
			"root_volume_encryption_key":        Computed(Nullable(String())),
			"volumes":                           Computed(List(kopsschemas.DataSourceVolumeSpec())),
			"volume_mounts":                     Computed(List(kopsschemas.DataSourceVolumeMountSpec())),
			"subnets":                           Computed(List(String())),
			"zones":                             Computed(List(String())),
			"hooks":                             Computed(List(kopsschemas.DataSourceHookSpec())),
			"max_price":                         Computed(Nullable(String())),
			"spot_duration_in_minutes":          Computed(Nullable(Int())),
			"cpu_credits":                       Computed(Nullable(String())),
			"associate_public_ip":               Computed(Nullable(Bool())),
			"additional_security_groups":        Computed(List(String())),
			"cloud_labels":                      Computed(Map(String())),
			"node_labels":                       Computed(Map(String())),
			"file_assets":                       Computed(List(kopsschemas.DataSourceFileAssetSpec())),
			"tenancy":                           Computed(String()),
			"kubelet":                           Computed(Struct(kopsschemas.DataSourceKubeletConfigSpec())),
			"taints":                            Computed(List(String())),
			"mixed_instances_policy":            Computed(Struct(kopsschemas.DataSourceMixedInstancesPolicySpec())),
			"additional_user_data":              Computed(List(kopsschemas.DataSourceUserData())),
			"suspend_processes":                 Computed(List(String())),
			"external_load_balancers":           Computed(List(kopsschemas.DataSourceLoadBalancer())),
			"detailed_instance_monitoring":      Computed(Nullable(Bool())),
			"iam":                               Computed(Struct(kopsschemas.DataSourceIAMProfileSpec())),
			"security_group_override":           Computed(Nullable(String())),
			"instance_protection":               Computed(Nullable(Bool())),
			"sysctl_parameters":                 Computed(List(String())),
			"rolling_update":                    Computed(Struct(kopsschemas.DataSourceRollingUpdate())),
			"instance_interruption_behavior":    Computed(Nullable(String())),
			"compress_user_data":                Computed(Nullable(Bool())),
			"instance_metadata":                 Computed(Struct(kopsschemas.DataSourceInstanceMetadataOptions())),
			"update_policy":                     Computed(Nullable(String())),
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

func FlattenDataSourceInstanceGroupInto(in resources.InstanceGroup, out map[string]interface{}) {
	out["cluster_name"] = func(in string) interface{} { return string(in) }(in.ClusterName)
	out["name"] = func(in string) interface{} { return string(in) }(in.Name)
	kopsschemas.FlattenDataSourceInstanceGroupSpecInto(in.InstanceGroupSpec, out)
}

func FlattenDataSourceInstanceGroup(in resources.InstanceGroup) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceInstanceGroupInto(in, out)
	return out
}
