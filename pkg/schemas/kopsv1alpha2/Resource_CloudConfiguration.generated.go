package schemas

import (
	"reflect"

	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

var _ = Schema

func ResourceCloudConfiguration() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"manage_storage_classes":         Nullable(OptionalBool()),
			"multizone":                      OptionalBool(),
			"node_tags":                      OptionalString(),
			"node_instance_prefix":           OptionalString(),
			"node_ip_families":               OptionalList(String()),
			"gce_service_account":            OptionalString(),
			"disable_security_group_ingress": OptionalBool(),
			"elb_security_group":             OptionalString(),
			"v_sphere_username":              OptionalString(),
			"v_sphere_password":              OptionalString(),
			"v_sphere_server":                OptionalString(),
			"v_sphere_datacenter":            OptionalString(),
			"v_sphere_resource_pool":         OptionalString(),
			"v_sphere_datastore":             OptionalString(),
			"v_sphere_core_dns_server":       OptionalString(),
			"spotinst_product":               OptionalString(),
			"spotinst_orientation":           OptionalString(),
			"openstack":                      OptionalStruct(ResourceOpenstackSpec()),
			"azure":                          OptionalStruct(ResourceAzureSpec()),
			"aws_ebs_csi_driver":             OptionalStruct(ResourceEBSCSIDriverSpec()),
			"gcp_pd_csi_driver":              OptionalStruct(ResourcePDCSIDriver()),
		},
	}

	return res
}

func ExpandResourceCloudConfiguration(in map[string]interface{}) kopsv1alpha2.CloudConfiguration {
	if in == nil {
		panic("expand CloudConfiguration failure, in is nil")
	}
	return kopsv1alpha2.CloudConfiguration{
		ManageStorageClasses: func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && len(in) == 1 {
				return func(in interface{}) *bool {
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
				}(in[0].(map[string]interface{})["value"])
			}
			return nil
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
		VSphereUsername: func(in interface{}) *string {
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
		}(in["v_sphere_username"]),
		VSpherePassword: func(in interface{}) *string {
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
		}(in["v_sphere_password"]),
		VSphereServer: func(in interface{}) *string {
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
		}(in["v_sphere_server"]),
		VSphereDatacenter: func(in interface{}) *string {
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
		}(in["v_sphere_datacenter"]),
		VSphereResourcePool: func(in interface{}) *string {
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
		}(in["v_sphere_resource_pool"]),
		VSphereDatastore: func(in interface{}) *string {
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
		}(in["v_sphere_datastore"]),
		VSphereCoreDNSServer: func(in interface{}) *string {
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
		}(in["v_sphere_core_dns_server"]),
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
		Openstack: func(in interface{}) *kopsv1alpha2.OpenstackSpec {
			return func(in interface{}) *kopsv1alpha2.OpenstackSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.OpenstackSpec) *kopsv1alpha2.OpenstackSpec {
					return &in
				}(func(in interface{}) kopsv1alpha2.OpenstackSpec {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandResourceOpenstackSpec(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.OpenstackSpec{}
				}(in))
			}(in)
		}(in["openstack"]),
		Azure: func(in interface{}) *kopsv1alpha2.AzureSpec {
			return func(in interface{}) *kopsv1alpha2.AzureSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.AzureSpec) *kopsv1alpha2.AzureSpec {
					return &in
				}(func(in interface{}) kopsv1alpha2.AzureSpec {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandResourceAzureSpec(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.AzureSpec{}
				}(in))
			}(in)
		}(in["azure"]),
		AWSEBSCSIDriver: func(in interface{}) *kopsv1alpha2.EBSCSIDriverSpec {
			return func(in interface{}) *kopsv1alpha2.EBSCSIDriverSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.EBSCSIDriverSpec) *kopsv1alpha2.EBSCSIDriverSpec {
					return &in
				}(func(in interface{}) kopsv1alpha2.EBSCSIDriverSpec {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandResourceEBSCSIDriverSpec(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.EBSCSIDriverSpec{}
				}(in))
			}(in)
		}(in["aws_ebs_csi_driver"]),
		GCPPDCSIDriver: func(in interface{}) *kopsv1alpha2.PDCSIDriver {
			return func(in interface{}) *kopsv1alpha2.PDCSIDriver {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.PDCSIDriver) *kopsv1alpha2.PDCSIDriver {
					return &in
				}(func(in interface{}) kopsv1alpha2.PDCSIDriver {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandResourcePDCSIDriver(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.PDCSIDriver{}
				}(in))
			}(in)
		}(in["gcp_pd_csi_driver"]),
	}
}

func FlattenResourceCloudConfigurationInto(in kopsv1alpha2.CloudConfiguration, out map[string]interface{}) {
	out["manage_storage_classes"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return []interface{}{map[string]interface{}{"value": func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)}}
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
	out["openstack"] = func(in *kopsv1alpha2.OpenstackSpec) interface{} {
		return func(in *kopsv1alpha2.OpenstackSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.OpenstackSpec) interface{} {
				return func(in kopsv1alpha2.OpenstackSpec) []interface{} {
					return []interface{}{FlattenResourceOpenstackSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Openstack)
	out["azure"] = func(in *kopsv1alpha2.AzureSpec) interface{} {
		return func(in *kopsv1alpha2.AzureSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.AzureSpec) interface{} {
				return func(in kopsv1alpha2.AzureSpec) []interface{} {
					return []interface{}{FlattenResourceAzureSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Azure)
	out["aws_ebs_csi_driver"] = func(in *kopsv1alpha2.EBSCSIDriverSpec) interface{} {
		return func(in *kopsv1alpha2.EBSCSIDriverSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.EBSCSIDriverSpec) interface{} {
				return func(in kopsv1alpha2.EBSCSIDriverSpec) []interface{} {
					return []interface{}{FlattenResourceEBSCSIDriverSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.AWSEBSCSIDriver)
	out["gcp_pd_csi_driver"] = func(in *kopsv1alpha2.PDCSIDriver) interface{} {
		return func(in *kopsv1alpha2.PDCSIDriver) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.PDCSIDriver) interface{} {
				return func(in kopsv1alpha2.PDCSIDriver) []interface{} {
					return []interface{}{FlattenResourcePDCSIDriver(in)}
				}(in)
			}(*in)
		}(in)
	}(in.GCPPDCSIDriver)
}

func FlattenResourceCloudConfiguration(in kopsv1alpha2.CloudConfiguration) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceCloudConfigurationInto(in, out)
	return out
}
