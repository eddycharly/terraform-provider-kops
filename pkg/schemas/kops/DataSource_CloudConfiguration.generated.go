package schemas

import (
	"reflect"

	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceCloudConfiguration() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"manage_storage_classes":         ComputedBool(),
			"multizone":                      ComputedBool(),
			"node_tags":                      ComputedString(),
			"node_instance_prefix":           ComputedString(),
			"node_ip_families":               ComputedList(String()),
			"gce_service_account":            ComputedString(),
			"disable_security_group_ingress": ComputedBool(),
			"elb_security_group":             ComputedString(),
			"spotinst_product":               ComputedString(),
			"spotinst_orientation":           ComputedString(),
			"aws_ebs_csi_driver":             ComputedStruct(DataSourceAWSEBSCSIDriver()),
			"gcp_pd_csi_driver":              ComputedStruct(DataSourceGCPPDCSIDriver()),
		},
	}

	return res
}

func ExpandDataSourceCloudConfiguration(in map[string]interface{}) kops.CloudConfiguration {
	if in == nil {
		panic("expand CloudConfiguration failure, in is nil")
	}
	return kops.CloudConfiguration{
		ManageStorageClasses: func(in interface{}) *bool {
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
		}(in["manage_storage_classes"]),
		Multizone: func(in interface{}) *bool {
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
		}(in["multizone"]),
		NodeTags: func(in interface{}) *string {
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
		}(in["node_tags"]),
		NodeInstancePrefix: func(in interface{}) *string {
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
		}(in["node_instance_prefix"]),
		NodeIPFamilies: func(in interface{}) []string {
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
		}(in["node_ip_families"]),
		GCEServiceAccount: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["gce_service_account"]),
		DisableSecurityGroupIngress: func(in interface{}) *bool {
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
		}(in["disable_security_group_ingress"]),
		ElbSecurityGroup: func(in interface{}) *string {
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
		}(in["elb_security_group"]),
		SpotinstProduct: func(in interface{}) *string {
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
		}(in["spotinst_product"]),
		SpotinstOrientation: func(in interface{}) *string {
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
		}(in["spotinst_orientation"]),
		AWSEBSCSIDriver: func(in interface{}) *kops.AWSEBSCSIDriver {
			return func(in interface{}) *kops.AWSEBSCSIDriver {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.AWSEBSCSIDriver) *kops.AWSEBSCSIDriver {
					return &in
				}(func(in interface{}) kops.AWSEBSCSIDriver {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandDataSourceAWSEBSCSIDriver(in[0].(map[string]interface{}))
					}
					return kops.AWSEBSCSIDriver{}
				}(in))
			}(in)
		}(in["aws_ebs_csi_driver"]),
		GCPPDCSIDriver: func(in interface{}) *kops.GCPPDCSIDriver {
			return func(in interface{}) *kops.GCPPDCSIDriver {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.GCPPDCSIDriver) *kops.GCPPDCSIDriver {
					return &in
				}(func(in interface{}) kops.GCPPDCSIDriver {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandDataSourceGCPPDCSIDriver(in[0].(map[string]interface{}))
					}
					return kops.GCPPDCSIDriver{}
				}(in))
			}(in)
		}(in["gcp_pd_csi_driver"]),
	}
}

func FlattenDataSourceCloudConfigurationInto(in kops.CloudConfiguration, out map[string]interface{}) {
	out["manage_storage_classes"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.ManageStorageClasses)
	out["multizone"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.Multizone)
	out["node_tags"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.NodeTags)
	out["node_instance_prefix"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.NodeInstancePrefix)
	out["node_ip_families"] = func(in []string) interface{} {
		return func(in []string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.NodeIPFamilies)
	out["gce_service_account"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.GCEServiceAccount)
	out["disable_security_group_ingress"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.DisableSecurityGroupIngress)
	out["elb_security_group"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.ElbSecurityGroup)
	out["spotinst_product"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.SpotinstProduct)
	out["spotinst_orientation"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.SpotinstOrientation)
	out["aws_ebs_csi_driver"] = func(in *kops.AWSEBSCSIDriver) interface{} {
		return func(in *kops.AWSEBSCSIDriver) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.AWSEBSCSIDriver) interface{} {
				return func(in kops.AWSEBSCSIDriver) []interface{} {
					return []interface{}{FlattenDataSourceAWSEBSCSIDriver(in)}
				}(in)
			}(*in)
		}(in)
	}(in.AWSEBSCSIDriver)
	out["gcp_pd_csi_driver"] = func(in *kops.GCPPDCSIDriver) interface{} {
		return func(in *kops.GCPPDCSIDriver) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.GCPPDCSIDriver) interface{} {
				return func(in kops.GCPPDCSIDriver) []interface{} {
					return []interface{}{FlattenDataSourceGCPPDCSIDriver(in)}
				}(in)
			}(*in)
		}(in)
	}(in.GCPPDCSIDriver)
}

func FlattenDataSourceCloudConfiguration(in kops.CloudConfiguration) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceCloudConfigurationInto(in, out)
	return out
}
