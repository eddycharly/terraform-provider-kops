package schemas

import (
	"github.com/eddycharly/terraform-provider-kops/pkg/api/config"
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var _ = Schema

func ConfigProvider() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"state_store":   RequiredString(),
			"aws":           OptionalStruct(ConfigAws()),
			"openstack":     OptionalStruct(ConfigOpenstack()),
			"klog":          OptionalStruct(ConfigKlog()),
			"mock":          OptionalBool(),
			"feature_flags": OptionalList(String()),
		},
	}

	return res
}

func ExpandConfigProvider(in map[string]interface{}) config.Provider {
	if in == nil {
		panic("expand Provider failure, in is nil")
	}
	return config.Provider{
		StateStore: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["state_store"]),
		Aws: func(in interface{}) *config.Aws {
			return func(in interface{}) *config.Aws {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in config.Aws) *config.Aws {
					return &in
				}(func(in interface{}) config.Aws {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandConfigAws(in[0].(map[string]interface{}))
					}
					return config.Aws{}
				}(in))
			}(in)
		}(in["aws"]),
		Openstack: func(in interface{}) *config.Openstack {
			return func(in interface{}) *config.Openstack {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in config.Openstack) *config.Openstack {
					return &in
				}(func(in interface{}) config.Openstack {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandConfigOpenstack(in[0].(map[string]interface{}))
					}
					return config.Openstack{}
				}(in))
			}(in)
		}(in["openstack"]),
		Klog: func(in interface{}) *config.Klog {
			return func(in interface{}) *config.Klog {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in config.Klog) *config.Klog {
					return &in
				}(func(in interface{}) config.Klog {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandConfigKlog(in[0].(map[string]interface{}))
					}
					return config.Klog{}
				}(in))
			}(in)
		}(in["klog"]),
		Mock: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["mock"]),
		FeatureFlags: func(in interface{}) []string {
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
		}(in["feature_flags"]),
	}
}

func FlattenConfigProviderInto(in config.Provider, out map[string]interface{}) {
	out["state_store"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.StateStore)
	out["aws"] = func(in *config.Aws) interface{} {
		return func(in *config.Aws) interface{} {
			if in == nil {
				return nil
			}
			return func(in config.Aws) interface{} {
				return func(in config.Aws) []interface{} {
					return []interface{}{FlattenConfigAws(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Aws)
	out["openstack"] = func(in *config.Openstack) interface{} {
		return func(in *config.Openstack) interface{} {
			if in == nil {
				return nil
			}
			return func(in config.Openstack) interface{} {
				return func(in config.Openstack) []interface{} {
					return []interface{}{FlattenConfigOpenstack(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Openstack)
	out["klog"] = func(in *config.Klog) interface{} {
		return func(in *config.Klog) interface{} {
			if in == nil {
				return nil
			}
			return func(in config.Klog) interface{} {
				return func(in config.Klog) []interface{} {
					return []interface{}{FlattenConfigKlog(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Klog)
	out["mock"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.Mock)
	out["feature_flags"] = func(in []string) interface{} {
		return func(in []string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.FeatureFlags)
}

func FlattenConfigProvider(in config.Provider) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenConfigProviderInto(in, out)
	return out
}
