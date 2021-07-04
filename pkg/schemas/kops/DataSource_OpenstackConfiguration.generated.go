package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceOpenstackConfiguration() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"loadbalancer":         Computed(Struct(DataSourceOpenstackLoadbalancerConfig())),
			"monitor":              Computed(Struct(DataSourceOpenstackMonitor())),
			"router":               Computed(Struct(DataSourceOpenstackRouter())),
			"block_storage":        Computed(Struct(DataSourceOpenstackBlockStorageConfig())),
			"insecure_skip_verify": Computed(Nullable(Bool())),
			"network":              Computed(Struct(DataSourceOpenstackNetwork())),
		},
	}
}

func ExpandDataSourceOpenstackConfiguration(in map[string]interface{}) kops.OpenstackConfiguration {
	if in == nil {
		panic("expand OpenstackConfiguration failure, in is nil")
	}
	out := kops.OpenstackConfiguration{}
	if in, ok := in["loadbalancer"]; ok && in != nil {
		out.Loadbalancer = func(in interface{}) *kops.OpenstackLoadbalancerConfig {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in kops.OpenstackLoadbalancerConfig) *kops.OpenstackLoadbalancerConfig { return &in }(func(in interface{}) kops.OpenstackLoadbalancerConfig {
					return ExpandDataSourceOpenstackLoadbalancerConfig(in.(map[string]interface{}))
				}(in[0]))
			}
			return nil
		}(in)
	}
	if in, ok := in["monitor"]; ok && in != nil {
		out.Monitor = func(in interface{}) *kops.OpenstackMonitor {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in kops.OpenstackMonitor) *kops.OpenstackMonitor { return &in }(func(in interface{}) kops.OpenstackMonitor {
					return ExpandDataSourceOpenstackMonitor(in.(map[string]interface{}))
				}(in[0]))
			}
			return nil
		}(in)
	}
	if in, ok := in["router"]; ok && in != nil {
		out.Router = func(in interface{}) *kops.OpenstackRouter {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in kops.OpenstackRouter) *kops.OpenstackRouter { return &in }(func(in interface{}) kops.OpenstackRouter {
					return ExpandDataSourceOpenstackRouter(in.(map[string]interface{}))
				}(in[0]))
			}
			return nil
		}(in)
	}
	if in, ok := in["block_storage"]; ok && in != nil {
		out.BlockStorage = func(in interface{}) *kops.OpenstackBlockStorageConfig {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in kops.OpenstackBlockStorageConfig) *kops.OpenstackBlockStorageConfig { return &in }(func(in interface{}) kops.OpenstackBlockStorageConfig {
					return ExpandDataSourceOpenstackBlockStorageConfig(in.(map[string]interface{}))
				}(in[0]))
			}
			return nil
		}(in)
	}
	if in, ok := in["insecure_skip_verify"]; ok && in != nil {
		out.InsecureSkipVerify = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["network"]; ok && in != nil {
		out.Network = func(in interface{}) *kops.OpenstackNetwork {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in kops.OpenstackNetwork) *kops.OpenstackNetwork { return &in }(func(in interface{}) kops.OpenstackNetwork {
					return ExpandDataSourceOpenstackNetwork(in.(map[string]interface{}))
				}(in[0]))
			}
			return nil
		}(in)
	}
	return out
}

func FlattenDataSourceOpenstackConfigurationInto(in kops.OpenstackConfiguration, out map[string]interface{}) {
	out["loadbalancer"] = func(in *kops.OpenstackLoadbalancerConfig) interface{} {
		if in == nil {
			return nil
		}
		return func(in kops.OpenstackLoadbalancerConfig) interface{} {
			return FlattenDataSourceOpenstackLoadbalancerConfig(in)
		}(*in)
	}(in.Loadbalancer)
	out["monitor"] = func(in *kops.OpenstackMonitor) interface{} {
		if in == nil {
			return nil
		}
		return func(in kops.OpenstackMonitor) interface{} { return FlattenDataSourceOpenstackMonitor(in) }(*in)
	}(in.Monitor)
	out["router"] = func(in *kops.OpenstackRouter) interface{} {
		if in == nil {
			return nil
		}
		return func(in kops.OpenstackRouter) interface{} { return FlattenDataSourceOpenstackRouter(in) }(*in)
	}(in.Router)
	out["block_storage"] = func(in *kops.OpenstackBlockStorageConfig) interface{} {
		if in == nil {
			return nil
		}
		return func(in kops.OpenstackBlockStorageConfig) interface{} {
			return FlattenDataSourceOpenstackBlockStorageConfig(in)
		}(*in)
	}(in.BlockStorage)
	out["insecure_skip_verify"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in bool) interface{} { return in }(*in)}
	}(in.InsecureSkipVerify)
	out["network"] = func(in *kops.OpenstackNetwork) interface{} {
		if in == nil {
			return nil
		}
		return func(in kops.OpenstackNetwork) interface{} { return FlattenDataSourceOpenstackNetwork(in) }(*in)
	}(in.Network)
}

func FlattenDataSourceOpenstackConfiguration(in kops.OpenstackConfiguration) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceOpenstackConfigurationInto(in, out)
	return out
}
