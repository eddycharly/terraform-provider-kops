package schemas

import (
	"github.com/eddycharly/terraform-provider-kops/pkg/api/config"
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var _ = Schema

func ConfigProvider() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"state_store": Required(String()),
			"aws":         Optional(Ptr(Struct(ConfigAws()))),
			"openstack":   Optional(Ptr(Struct(ConfigOpenstack()))),
			"klog":        Optional(Ptr(Struct(ConfigKlog()))),
		},
	}
}

func ExpandConfigProvider(in map[string]interface{}) config.Provider {
	if in == nil {
		panic("expand Provider failure, in is nil")
	}
	out := config.Provider{}
	if in, ok := in["state_store"]; ok && in != nil {
		out.StateStore = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["aws"]; ok && in != nil {
		out.Aws = func(in interface{}) *config.Aws {
			if in == nil {
				return nil
			}
			return func(in config.Aws) *config.Aws { return &in }(func(in interface{}) config.Aws {
				if in == nil {
					return config.Aws{}
				}
				return ExpandConfigAws(in.(map[string]interface{}))
			}(in))
		}(in)
	}
	if in, ok := in["openstack"]; ok && in != nil {
		out.Openstack = func(in interface{}) *config.Openstack {
			if in == nil {
				return nil
			}
			return func(in config.Openstack) *config.Openstack { return &in }(func(in interface{}) config.Openstack {
				if in == nil {
					return config.Openstack{}
				}
				return ExpandConfigOpenstack(in.(map[string]interface{}))
			}(in))
		}(in)
	}
	if in, ok := in["klog"]; ok && in != nil {
		out.Klog = func(in interface{}) *config.Klog {
			if in == nil {
				return nil
			}
			return func(in config.Klog) *config.Klog { return &in }(func(in interface{}) config.Klog {
				if in == nil {
					return config.Klog{}
				}
				return ExpandConfigKlog(in.(map[string]interface{}))
			}(in))
		}(in)
	}
	return out
}
