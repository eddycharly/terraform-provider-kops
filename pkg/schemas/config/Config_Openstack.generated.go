package schemas

import (
	"github.com/eddycharly/terraform-provider-kops/pkg/api/config"
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var _ = Schema

func ConfigOpenstack() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"tenant_id":                     OptionalString(),
			"tenant_name":                   OptionalString(),
			"project_id":                    OptionalString(),
			"project_name":                  OptionalString(),
			"project_domain_id":             OptionalString(),
			"project_domain_name":           OptionalString(),
			"domain_id":                     OptionalString(),
			"domain_name":                   OptionalString(),
			"username":                      OptionalString(),
			"password":                      OptionalString(),
			"auth_url":                      OptionalString(),
			"region_name":                   OptionalString(),
			"application_credential_id":     OptionalString(),
			"application_credential_secret": OptionalString(),
		},
	}

	return res
}

func ExpandConfigOpenstack(in map[string]interface{}) config.Openstack {
	if in == nil {
		panic("expand Openstack failure, in is nil")
	}
	return config.Openstack{
		TenantId: func(in interface{}) string /**/ {
			return string(ExpandString(in))
		}(in["tenant_id"]),
		TenantName: func(in interface{}) string /**/ {
			return string(ExpandString(in))
		}(in["tenant_name"]),
		ProjectId: func(in interface{}) string /**/ {
			return string(ExpandString(in))
		}(in["project_id"]),
		ProjectName: func(in interface{}) string /**/ {
			return string(ExpandString(in))
		}(in["project_name"]),
		ProjectDomainId: func(in interface{}) string /**/ {
			return string(ExpandString(in))
		}(in["project_domain_id"]),
		ProjectDomainName: func(in interface{}) string /**/ {
			return string(ExpandString(in))
		}(in["project_domain_name"]),
		DomainId: func(in interface{}) string /**/ {
			return string(ExpandString(in))
		}(in["domain_id"]),
		DomainName: func(in interface{}) string /**/ {
			return string(ExpandString(in))
		}(in["domain_name"]),
		Username: func(in interface{}) string /**/ {
			return string(ExpandString(in))
		}(in["username"]),
		Password: func(in interface{}) string /**/ {
			return string(ExpandString(in))
		}(in["password"]),
		AuthUrl: func(in interface{}) string /**/ {
			return string(ExpandString(in))
		}(in["auth_url"]),
		RegionName: func(in interface{}) string /**/ {
			return string(ExpandString(in))
		}(in["region_name"]),
		ApplicationCredentialId: func(in interface{}) string /**/ {
			return string(ExpandString(in))
		}(in["application_credential_id"]),
		ApplicationCredentialSecret: func(in interface{}) string /**/ {
			return string(ExpandString(in))
		}(in["application_credential_secret"]),
	}
}

func FlattenConfigOpenstackInto(in config.Openstack, out map[string]interface{}) {
	out["tenant_id"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.TenantId)
	out["tenant_name"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.TenantName)
	out["project_id"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.ProjectId)
	out["project_name"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.ProjectName)
	out["project_domain_id"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.ProjectDomainId)
	out["project_domain_name"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.ProjectDomainName)
	out["domain_id"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.DomainId)
	out["domain_name"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.DomainName)
	out["username"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Username)
	out["password"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Password)
	out["auth_url"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.AuthUrl)
	out["region_name"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.RegionName)
	out["application_credential_id"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.ApplicationCredentialId)
	out["application_credential_secret"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.ApplicationCredentialSecret)
}

func FlattenConfigOpenstack(in config.Openstack) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenConfigOpenstackInto(in, out)
	return out
}
