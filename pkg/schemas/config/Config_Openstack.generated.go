package schemas

import (
	"github.com/eddycharly/terraform-provider-kops/pkg/api/config"
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var _ = Schema

func ConfigOpenstack() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"tenant_id":                     Optional(String()),
			"tenant_name":                   Optional(String()),
			"project_id":                    Optional(String()),
			"project_name":                  Optional(String()),
			"project_domain_id":             Optional(String()),
			"project_domain_name":           Optional(String()),
			"domain_id":                     Optional(String()),
			"domain_name":                   Optional(String()),
			"username":                      Optional(String()),
			"password":                      Optional(String()),
			"auth_url":                      Optional(String()),
			"region_name":                   Optional(String()),
			"application_credential_id":     Optional(String()),
			"application_credential_secret": Optional(String()),
		},
	}
}

func ExpandConfigOpenstack(in map[string]interface{}) config.Openstack {
	if in == nil {
		panic("expand Openstack failure, in is nil")
	}
	out := config.Openstack{}
	if in, ok := in["tenant_id"]; ok && in != nil {
		out.TenantId = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["tenant_name"]; ok && in != nil {
		out.TenantName = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["project_id"]; ok && in != nil {
		out.ProjectId = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["project_name"]; ok && in != nil {
		out.ProjectName = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["project_domain_id"]; ok && in != nil {
		out.ProjectDomainId = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["project_domain_name"]; ok && in != nil {
		out.ProjectDomainName = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["domain_id"]; ok && in != nil {
		out.DomainId = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["domain_name"]; ok && in != nil {
		out.DomainName = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["username"]; ok && in != nil {
		out.Username = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["password"]; ok && in != nil {
		out.Password = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["auth_url"]; ok && in != nil {
		out.AuthUrl = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["region_name"]; ok && in != nil {
		out.RegionName = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["application_credential_id"]; ok && in != nil {
		out.ApplicationCredentialId = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["application_credential_secret"]; ok && in != nil {
		out.ApplicationCredentialSecret = func(in interface{}) string { return string(in.(string)) }(in)
	}
	return out
}

func FlattenConfigOpenstackInto(in config.Openstack, out map[string]interface{}) {
	out["tenant_id"] = func(in string) interface{} { return string(in) }(in.TenantId)
	out["tenant_name"] = func(in string) interface{} { return string(in) }(in.TenantName)
	out["project_id"] = func(in string) interface{} { return string(in) }(in.ProjectId)
	out["project_name"] = func(in string) interface{} { return string(in) }(in.ProjectName)
	out["project_domain_id"] = func(in string) interface{} { return string(in) }(in.ProjectDomainId)
	out["project_domain_name"] = func(in string) interface{} { return string(in) }(in.ProjectDomainName)
	out["domain_id"] = func(in string) interface{} { return string(in) }(in.DomainId)
	out["domain_name"] = func(in string) interface{} { return string(in) }(in.DomainName)
	out["username"] = func(in string) interface{} { return string(in) }(in.Username)
	out["password"] = func(in string) interface{} { return string(in) }(in.Password)
	out["auth_url"] = func(in string) interface{} { return string(in) }(in.AuthUrl)
	out["region_name"] = func(in string) interface{} { return string(in) }(in.RegionName)
	out["application_credential_id"] = func(in string) interface{} { return string(in) }(in.ApplicationCredentialId)
	out["application_credential_secret"] = func(in string) interface{} { return string(in) }(in.ApplicationCredentialSecret)
}

func FlattenConfigOpenstack(in config.Openstack) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenConfigOpenstackInto(in, out)
	return out
}
