package schemas

import (
	"reflect"

	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceCloudConfiguration() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"manage_storage_classes":         ComputedBool(),
			"multizone":                      ComputedBool(),
			"node_tags":                      ComputedString(),
			"node_instance_prefix":           ComputedString(),
			"gce_service_account":            ComputedString(),
			"disable_security_group_ingress": ComputedBool(),
			"elb_security_group":             ComputedString(),
			"v_sphere_username":              ComputedString(),
			"v_sphere_password":              ComputedString(),
			"v_sphere_server":                ComputedString(),
			"v_sphere_datacenter":            ComputedString(),
			"v_sphere_resource_pool":         ComputedString(),
			"v_sphere_datastore":             ComputedString(),
			"v_sphere_core_dns_server":       ComputedString(),
			"spotinst_product":               ComputedString(),
			"spotinst_orientation":           ComputedString(),
			"openstack":                      ComputedStruct(DataSourceOpenstackConfiguration()),
			"azure":                          ComputedStruct(DataSourceAzureConfiguration()),
			"awsebscsi_driver":               ComputedStruct(DataSourceAWSEBSCSIDriver()),
		},
	}
}

func ExpandDataSourceCloudConfiguration(in map[string]interface{}) kops.CloudConfiguration {
	if in == nil {
		panic("expand CloudConfiguration failure, in is nil")
	}
	return kops.CloudConfiguration{
		ManageStorageClasses: func(in interface{}) *bool {
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
		GCEServiceAccount: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["gce_service_account"]),
		DisableSecurityGroupIngress: func(in interface{}) *bool {
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
		VSphereUsername: func(in interface{}) *string {
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
		}(in["v_sphere_username"]),
		VSpherePassword: func(in interface{}) *string {
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
		}(in["v_sphere_password"]),
		VSphereServer: func(in interface{}) *string {
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
		}(in["v_sphere_server"]),
		VSphereDatacenter: func(in interface{}) *string {
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
		}(in["v_sphere_datacenter"]),
		VSphereResourcePool: func(in interface{}) *string {
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
		}(in["v_sphere_resource_pool"]),
		VSphereDatastore: func(in interface{}) *string {
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
		}(in["v_sphere_datastore"]),
		VSphereCoreDNSServer: func(in interface{}) *string {
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
		}(in["v_sphere_core_dns_server"]),
		SpotinstProduct: func(in interface{}) *string {
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
		Openstack: func(in interface{}) *kops.OpenstackConfiguration {
			return func(in interface{}) *kops.OpenstackConfiguration {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.OpenstackConfiguration) *kops.OpenstackConfiguration {
					return &in
				}(func(in interface{}) kops.OpenstackConfiguration {
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return kops.OpenstackConfiguration{}
					}
					return (ExpandDataSourceOpenstackConfiguration(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["openstack"]),
		Azure: func(in interface{}) *kops.AzureConfiguration {
			return func(in interface{}) *kops.AzureConfiguration {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.AzureConfiguration) *kops.AzureConfiguration {
					return &in
				}(func(in interface{}) kops.AzureConfiguration {
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return kops.AzureConfiguration{}
					}
					return (ExpandDataSourceAzureConfiguration(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["azure"]),
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
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return kops.AWSEBSCSIDriver{}
					}
					return (ExpandDataSourceAWSEBSCSIDriver(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["awsebscsi_driver"]),
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
	out["v_sphere_username"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.VSphereUsername)
	out["v_sphere_password"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.VSpherePassword)
	out["v_sphere_server"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.VSphereServer)
	out["v_sphere_datacenter"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.VSphereDatacenter)
	out["v_sphere_resource_pool"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.VSphereResourcePool)
	out["v_sphere_datastore"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.VSphereDatastore)
	out["v_sphere_core_dns_server"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.VSphereCoreDNSServer)
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
	out["openstack"] = func(in *kops.OpenstackConfiguration) interface{} {
		return func(in *kops.OpenstackConfiguration) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.OpenstackConfiguration) interface{} {
				return func(in kops.OpenstackConfiguration) []map[string]interface{} {
					return []map[string]interface{}{FlattenDataSourceOpenstackConfiguration(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Openstack)
	out["azure"] = func(in *kops.AzureConfiguration) interface{} {
		return func(in *kops.AzureConfiguration) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.AzureConfiguration) interface{} {
				return func(in kops.AzureConfiguration) []map[string]interface{} {
					return []map[string]interface{}{FlattenDataSourceAzureConfiguration(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Azure)
	out["awsebscsi_driver"] = func(in *kops.AWSEBSCSIDriver) interface{} {
		return func(in *kops.AWSEBSCSIDriver) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.AWSEBSCSIDriver) interface{} {
				return func(in kops.AWSEBSCSIDriver) []map[string]interface{} {
					return []map[string]interface{}{FlattenDataSourceAWSEBSCSIDriver(in)}
				}(in)
			}(*in)
		}(in)
	}(in.AWSEBSCSIDriver)
}

func FlattenDataSourceCloudConfiguration(in kops.CloudConfiguration) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceCloudConfigurationInto(in, out)
	return out
}
