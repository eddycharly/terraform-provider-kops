package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceCloudConfiguration() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"manage_storage_classes":         Computed(Nullable(Bool())),
			"multizone":                      Computed(Nullable(Bool())),
			"node_tags":                      Computed(Nullable(String())),
			"node_instance_prefix":           Computed(Nullable(String())),
			"gce_service_account":            Computed(String()),
			"disable_security_group_ingress": Computed(Nullable(Bool())),
			"elb_security_group":             Computed(Nullable(String())),
			"v_sphere_username":              Computed(Nullable(String())),
			"v_sphere_password":              Computed(Nullable(String())),
			"v_sphere_server":                Computed(Nullable(String())),
			"v_sphere_datacenter":            Computed(Nullable(String())),
			"v_sphere_resource_pool":         Computed(Nullable(String())),
			"v_sphere_datastore":             Computed(Nullable(String())),
			"v_sphere_core_dns_server":       Computed(Nullable(String())),
			"spotinst_product":               Computed(Nullable(String())),
			"spotinst_orientation":           Computed(Nullable(String())),
			"openstack":                      Computed(Struct(DataSourceOpenstackConfiguration())),
			"azure":                          Computed(Struct(DataSourceAzureConfiguration())),
			"aws_ebs_csi_driver":             Computed(Struct(DataSourceAWSEBSCSIDriver())),
		},
	}
}

func ExpandDataSourceCloudConfiguration(in map[string]interface{}) kops.CloudConfiguration {
	if in == nil {
		panic("expand CloudConfiguration failure, in is nil")
	}
	out := kops.CloudConfiguration{}
	if in, ok := in["manage_storage_classes"]; ok && in != nil {
		out.ManageStorageClasses = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["multizone"]; ok && in != nil {
		out.Multizone = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["node_tags"]; ok && in != nil {
		out.NodeTags = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["node_instance_prefix"]; ok && in != nil {
		out.NodeInstancePrefix = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["gce_service_account"]; ok && in != nil {
		out.GCEServiceAccount = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["disable_security_group_ingress"]; ok && in != nil {
		out.DisableSecurityGroupIngress = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["elb_security_group"]; ok && in != nil {
		out.ElbSecurityGroup = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["v_sphere_username"]; ok && in != nil {
		out.VSphereUsername = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["v_sphere_password"]; ok && in != nil {
		out.VSpherePassword = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["v_sphere_server"]; ok && in != nil {
		out.VSphereServer = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["v_sphere_datacenter"]; ok && in != nil {
		out.VSphereDatacenter = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["v_sphere_resource_pool"]; ok && in != nil {
		out.VSphereResourcePool = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["v_sphere_datastore"]; ok && in != nil {
		out.VSphereDatastore = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["v_sphere_core_dns_server"]; ok && in != nil {
		out.VSphereCoreDNSServer = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["spotinst_product"]; ok && in != nil {
		out.SpotinstProduct = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["spotinst_orientation"]; ok && in != nil {
		out.SpotinstOrientation = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["openstack"]; ok && in != nil {
		out.Openstack = func(in interface{}) *kops.OpenstackConfiguration {
			if in == nil {
				return nil
			}
			return func(in kops.OpenstackConfiguration) *kops.OpenstackConfiguration { return &in }(func(in interface{}) kops.OpenstackConfiguration {
				if in == nil {
					return kops.OpenstackConfiguration{}
				}
				return ExpandDataSourceOpenstackConfiguration(in.(map[string]interface{}))
			}(in))
		}(in)
	}
	if in, ok := in["azure"]; ok && in != nil {
		out.Azure = func(in interface{}) *kops.AzureConfiguration {
			if in == nil {
				return nil
			}
			return func(in kops.AzureConfiguration) *kops.AzureConfiguration { return &in }(func(in interface{}) kops.AzureConfiguration {
				if in == nil {
					return kops.AzureConfiguration{}
				}
				return ExpandDataSourceAzureConfiguration(in.(map[string]interface{}))
			}(in))
		}(in)
	}
	if in, ok := in["aws_ebs_csi_driver"]; ok && in != nil {
		out.AWSEBSCSIDriver = func(in interface{}) *kops.AWSEBSCSIDriver {
			if in == nil {
				return nil
			}
			return func(in kops.AWSEBSCSIDriver) *kops.AWSEBSCSIDriver { return &in }(func(in interface{}) kops.AWSEBSCSIDriver {
				if in == nil {
					return kops.AWSEBSCSIDriver{}
				}
				return ExpandDataSourceAWSEBSCSIDriver(in.(map[string]interface{}))
			}(in))
		}(in)
	}
	return out
}

func FlattenDataSourceCloudConfigurationInto(in kops.CloudConfiguration, out map[string]interface{}) {
	out["manage_storage_classes"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in bool) interface{} { return in }(*in)}
	}(in.ManageStorageClasses)
	out["multizone"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in bool) interface{} { return in }(*in)}
	}(in.Multizone)
	out["node_tags"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.NodeTags)
	out["node_instance_prefix"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.NodeInstancePrefix)
	out["gce_service_account"] = func(in string) interface{} { return string(in) }(in.GCEServiceAccount)
	out["disable_security_group_ingress"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in bool) interface{} { return in }(*in)}
	}(in.DisableSecurityGroupIngress)
	out["elb_security_group"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.ElbSecurityGroup)
	out["v_sphere_username"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.VSphereUsername)
	out["v_sphere_password"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.VSpherePassword)
	out["v_sphere_server"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.VSphereServer)
	out["v_sphere_datacenter"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.VSphereDatacenter)
	out["v_sphere_resource_pool"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.VSphereResourcePool)
	out["v_sphere_datastore"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.VSphereDatastore)
	out["v_sphere_core_dns_server"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.VSphereCoreDNSServer)
	out["spotinst_product"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.SpotinstProduct)
	out["spotinst_orientation"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.SpotinstOrientation)
	out["openstack"] = func(in *kops.OpenstackConfiguration) interface{} {
		if in == nil {
			return nil
		}
		return func(in kops.OpenstackConfiguration) interface{} { return FlattenDataSourceOpenstackConfiguration(in) }(*in)
	}(in.Openstack)
	out["azure"] = func(in *kops.AzureConfiguration) interface{} {
		if in == nil {
			return nil
		}
		return func(in kops.AzureConfiguration) interface{} { return FlattenDataSourceAzureConfiguration(in) }(*in)
	}(in.Azure)
	out["aws_ebs_csi_driver"] = func(in *kops.AWSEBSCSIDriver) interface{} {
		if in == nil {
			return nil
		}
		return func(in kops.AWSEBSCSIDriver) interface{} { return FlattenDataSourceAWSEBSCSIDriver(in) }(*in)
	}(in.AWSEBSCSIDriver)
}

func FlattenDataSourceCloudConfiguration(in kops.CloudConfiguration) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceCloudConfigurationInto(in, out)
	return out
}
