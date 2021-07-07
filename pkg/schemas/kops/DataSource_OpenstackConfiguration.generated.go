package schemas

import (
	"reflect"

	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceOpenstackConfiguration() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"loadbalancer":         ComputedStruct(DataSourceOpenstackLoadbalancerConfig()),
			"monitor":              ComputedStruct(DataSourceOpenstackMonitor()),
			"router":               ComputedStruct(DataSourceOpenstackRouter()),
			"block_storage":        ComputedStruct(DataSourceOpenstackBlockStorageConfig()),
			"insecure_skip_verify": ComputedBool(),
			"network":              ComputedStruct(DataSourceOpenstackNetwork()),
		},
	}

	return res
}

func ExpandDataSourceOpenstackConfiguration(in map[string]interface{}) kops.OpenstackConfiguration {
	if in == nil {
		panic("expand OpenstackConfiguration failure, in is nil")
	}
	return kops.OpenstackConfiguration{
		Loadbalancer: func(in interface{}) *kops.OpenstackLoadbalancerConfig {
			return func(in interface{}) *kops.OpenstackLoadbalancerConfig {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.OpenstackLoadbalancerConfig) *kops.OpenstackLoadbalancerConfig {
					return &in
				}(func(in interface{}) kops.OpenstackLoadbalancerConfig {
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return kops.OpenstackLoadbalancerConfig{}
					}
					return (ExpandDataSourceOpenstackLoadbalancerConfig(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["loadbalancer"]),
		Monitor: func(in interface{}) *kops.OpenstackMonitor {
			return func(in interface{}) *kops.OpenstackMonitor {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.OpenstackMonitor) *kops.OpenstackMonitor {
					return &in
				}(func(in interface{}) kops.OpenstackMonitor {
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return kops.OpenstackMonitor{}
					}
					return (ExpandDataSourceOpenstackMonitor(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["monitor"]),
		Router: func(in interface{}) *kops.OpenstackRouter {
			return func(in interface{}) *kops.OpenstackRouter {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.OpenstackRouter) *kops.OpenstackRouter {
					return &in
				}(func(in interface{}) kops.OpenstackRouter {
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return kops.OpenstackRouter{}
					}
					return (ExpandDataSourceOpenstackRouter(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["router"]),
		BlockStorage: func(in interface{}) *kops.OpenstackBlockStorageConfig {
			return func(in interface{}) *kops.OpenstackBlockStorageConfig {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.OpenstackBlockStorageConfig) *kops.OpenstackBlockStorageConfig {
					return &in
				}(func(in interface{}) kops.OpenstackBlockStorageConfig {
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return kops.OpenstackBlockStorageConfig{}
					}
					return (ExpandDataSourceOpenstackBlockStorageConfig(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["block_storage"]),
		InsecureSkipVerify: func(in interface{}) *bool {
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
		}(in["insecure_skip_verify"]),
		Network: func(in interface{}) *kops.OpenstackNetwork {
			return func(in interface{}) *kops.OpenstackNetwork {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.OpenstackNetwork) *kops.OpenstackNetwork {
					return &in
				}(func(in interface{}) kops.OpenstackNetwork {
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return kops.OpenstackNetwork{}
					}
					return (ExpandDataSourceOpenstackNetwork(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["network"]),
	}
}

func FlattenDataSourceOpenstackConfigurationInto(in kops.OpenstackConfiguration, out map[string]interface{}) {
	out["loadbalancer"] = func(in *kops.OpenstackLoadbalancerConfig) interface{} {
		return func(in *kops.OpenstackLoadbalancerConfig) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.OpenstackLoadbalancerConfig) interface{} {
				return func(in kops.OpenstackLoadbalancerConfig) []map[string]interface{} {
					return []map[string]interface{}{FlattenDataSourceOpenstackLoadbalancerConfig(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Loadbalancer)
	out["monitor"] = func(in *kops.OpenstackMonitor) interface{} {
		return func(in *kops.OpenstackMonitor) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.OpenstackMonitor) interface{} {
				return func(in kops.OpenstackMonitor) []map[string]interface{} {
					return []map[string]interface{}{FlattenDataSourceOpenstackMonitor(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Monitor)
	out["router"] = func(in *kops.OpenstackRouter) interface{} {
		return func(in *kops.OpenstackRouter) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.OpenstackRouter) interface{} {
				return func(in kops.OpenstackRouter) []map[string]interface{} {
					return []map[string]interface{}{FlattenDataSourceOpenstackRouter(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Router)
	out["block_storage"] = func(in *kops.OpenstackBlockStorageConfig) interface{} {
		return func(in *kops.OpenstackBlockStorageConfig) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.OpenstackBlockStorageConfig) interface{} {
				return func(in kops.OpenstackBlockStorageConfig) []map[string]interface{} {
					return []map[string]interface{}{FlattenDataSourceOpenstackBlockStorageConfig(in)}
				}(in)
			}(*in)
		}(in)
	}(in.BlockStorage)
	out["insecure_skip_verify"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.InsecureSkipVerify)
	out["network"] = func(in *kops.OpenstackNetwork) interface{} {
		return func(in *kops.OpenstackNetwork) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.OpenstackNetwork) interface{} {
				return func(in kops.OpenstackNetwork) []map[string]interface{} {
					return []map[string]interface{}{FlattenDataSourceOpenstackNetwork(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Network)
}

func FlattenDataSourceOpenstackConfiguration(in kops.OpenstackConfiguration) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceOpenstackConfigurationInto(in, out)
	return out
}
