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
			"manage_storage_classes":         Computed(Ptr(Bool())),
			"multizone":                      Computed(Ptr(Bool())),
			"node_tags":                      Computed(Ptr(String())),
			"node_instance_prefix":           Computed(Ptr(String())),
			"gce_service_account":            Computed(String()),
			"disable_security_group_ingress": Computed(Ptr(Bool())),
			"elb_security_group":             Computed(Ptr(String())),
			"v_sphere_username":              Computed(Ptr(String())),
			"v_sphere_password":              Computed(Ptr(String())),
			"v_sphere_server":                Computed(Ptr(String())),
			"v_sphere_datacenter":            Computed(Ptr(String())),
			"v_sphere_resource_pool":         Computed(Ptr(String())),
			"v_sphere_datastore":             Computed(Ptr(String())),
			"v_sphere_core_dns_server":       Computed(Ptr(String())),
			"spotinst_product":               Computed(Ptr(String())),
			"spotinst_orientation":           Computed(Ptr(String())),
			"openstack":                      Computed(Ptr(Struct(DataSourceOpenstackConfiguration()))),
			"azure":                          Computed(Ptr(Struct(DataSourceAzureConfiguration()))),
			"aws_ebs_csi_driver":             Computed(Ptr(Struct(DataSourceAWSEBSCSIDriver()))),
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
