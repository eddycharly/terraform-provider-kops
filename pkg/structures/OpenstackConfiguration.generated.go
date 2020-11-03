package structures

import (
	"reflect"

	"k8s.io/kops/pkg/apis/kops"
)

func ExpandOpenstackConfiguration(in map[string]interface{}) kops.OpenstackConfiguration {
	if in == nil {
		panic("expand OpenstackConfiguration failure, in is nil")
	}
	return kops.OpenstackConfiguration{
		Loadbalancer: func(in interface{}) *kops.OpenstackLoadbalancerConfig {
			value := func(in interface{}) *kops.OpenstackLoadbalancerConfig {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in kops.OpenstackLoadbalancerConfig) *kops.OpenstackLoadbalancerConfig {
					return &in
				}(func(in interface{}) kops.OpenstackLoadbalancerConfig {
					if in.([]interface{})[0] == nil {
						return kops.OpenstackLoadbalancerConfig{}
					}
					return (ExpandOpenstackLoadbalancerConfig(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
			return value
		}(in["loadbalancer"]),
		Monitor: func(in interface{}) *kops.OpenstackMonitor {
			value := func(in interface{}) *kops.OpenstackMonitor {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in kops.OpenstackMonitor) *kops.OpenstackMonitor {
					return &in
				}(func(in interface{}) kops.OpenstackMonitor {
					if in.([]interface{})[0] == nil {
						return kops.OpenstackMonitor{}
					}
					return (ExpandOpenstackMonitor(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
			return value
		}(in["monitor"]),
		Router: func(in interface{}) *kops.OpenstackRouter {
			value := func(in interface{}) *kops.OpenstackRouter {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in kops.OpenstackRouter) *kops.OpenstackRouter {
					return &in
				}(func(in interface{}) kops.OpenstackRouter {
					if in.([]interface{})[0] == nil {
						return kops.OpenstackRouter{}
					}
					return (ExpandOpenstackRouter(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
			return value
		}(in["router"]),
		BlockStorage: func(in interface{}) *kops.OpenstackBlockStorageConfig {
			value := func(in interface{}) *kops.OpenstackBlockStorageConfig {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in kops.OpenstackBlockStorageConfig) *kops.OpenstackBlockStorageConfig {
					return &in
				}(func(in interface{}) kops.OpenstackBlockStorageConfig {
					if in.([]interface{})[0] == nil {
						return kops.OpenstackBlockStorageConfig{}
					}
					return (ExpandOpenstackBlockStorageConfig(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
			return value
		}(in["block_storage"]),
		InsecureSkipVerify: func(in interface{}) *bool {
			value := func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in bool) *bool {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(bool(ExpandBool(in)))
			}(in)
			return value
		}(in["insecure_skip_verify"]),
	}
}

func FlattenOpenstackConfiguration(in kops.OpenstackConfiguration) map[string]interface{} {
	return map[string]interface{}{
		"loadbalancer": func(in *kops.OpenstackLoadbalancerConfig) interface{} {
			value := func(in *kops.OpenstackLoadbalancerConfig) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.OpenstackLoadbalancerConfig) interface{} {
					return func(in kops.OpenstackLoadbalancerConfig) []map[string]interface{} {
						return []map[string]interface{}{FlattenOpenstackLoadbalancerConfig(in)}
					}(in)
				}(*in)
			}(in)
			return value
		}(in.Loadbalancer),
		"monitor": func(in *kops.OpenstackMonitor) interface{} {
			value := func(in *kops.OpenstackMonitor) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.OpenstackMonitor) interface{} {
					return func(in kops.OpenstackMonitor) []map[string]interface{} {
						return []map[string]interface{}{FlattenOpenstackMonitor(in)}
					}(in)
				}(*in)
			}(in)
			return value
		}(in.Monitor),
		"router": func(in *kops.OpenstackRouter) interface{} {
			value := func(in *kops.OpenstackRouter) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.OpenstackRouter) interface{} {
					return func(in kops.OpenstackRouter) []map[string]interface{} {
						return []map[string]interface{}{FlattenOpenstackRouter(in)}
					}(in)
				}(*in)
			}(in)
			return value
		}(in.Router),
		"block_storage": func(in *kops.OpenstackBlockStorageConfig) interface{} {
			value := func(in *kops.OpenstackBlockStorageConfig) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.OpenstackBlockStorageConfig) interface{} {
					return func(in kops.OpenstackBlockStorageConfig) []map[string]interface{} {
						return []map[string]interface{}{FlattenOpenstackBlockStorageConfig(in)}
					}(in)
				}(*in)
			}(in)
			return value
		}(in.BlockStorage),
		"insecure_skip_verify": func(in *bool) interface{} {
			value := func(in *bool) interface{} {
				if in == nil {
					return nil
				}
				return func(in bool) interface{} {
					return FlattenBool(bool(in))
				}(*in)
			}(in)
			return value
		}(in.InsecureSkipVerify),
	}
}
