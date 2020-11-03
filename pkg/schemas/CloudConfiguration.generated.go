package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func CloudConfiguration() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"multizone":                      OptionalBool(),
			"node_tags":                      OptionalString(),
			"node_instance_prefix":           OptionalString(),
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
			"openstack":                      OptionalStruct(OpenstackConfiguration()),
		},
	}
}
