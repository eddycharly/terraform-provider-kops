package schemas

import (
	"reflect"

	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

var _ = Schema

func DataSourceOpenstackSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"loadbalancer":         ComputedStruct(DataSourceOpenstackLoadbalancerConfig()),
			"monitor":              ComputedStruct(DataSourceOpenstackMonitor()),
			"router":               ComputedStruct(DataSourceOpenstackRouter()),
			"block_storage":        ComputedStruct(DataSourceOpenstackBlockStorageConfig()),
			"insecure_skip_verify": ComputedBool(),
			"network":              ComputedStruct(DataSourceOpenstackNetwork()),
			"metadata":             ComputedStruct(DataSourceOpenstackMetadata()),
		},
	}

	return res
}

func ExpandDataSourceOpenstackSpec(in map[string]interface{}) kopsv1alpha2.OpenstackSpec {
	if in == nil {
		panic("expand OpenstackSpec failure, in is nil")
	}
	return kopsv1alpha2.OpenstackSpec{
		Loadbalancer: func(in interface{}) *kopsv1alpha2.OpenstackLoadbalancerConfig {
			return func(in interface{}) *kopsv1alpha2.OpenstackLoadbalancerConfig {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.OpenstackLoadbalancerConfig) *kopsv1alpha2.OpenstackLoadbalancerConfig {
					return &in
				}(func(in interface{}) kopsv1alpha2.OpenstackLoadbalancerConfig {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandDataSourceOpenstackLoadbalancerConfig(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.OpenstackLoadbalancerConfig{}
				}(in))
			}(in)
		}(in["loadbalancer"]),
		Monitor: func(in interface{}) *kopsv1alpha2.OpenstackMonitor {
			return func(in interface{}) *kopsv1alpha2.OpenstackMonitor {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.OpenstackMonitor) *kopsv1alpha2.OpenstackMonitor {
					return &in
				}(func(in interface{}) kopsv1alpha2.OpenstackMonitor {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandDataSourceOpenstackMonitor(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.OpenstackMonitor{}
				}(in))
			}(in)
		}(in["monitor"]),
		Router: func(in interface{}) *kopsv1alpha2.OpenstackRouter {
			return func(in interface{}) *kopsv1alpha2.OpenstackRouter {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.OpenstackRouter) *kopsv1alpha2.OpenstackRouter {
					return &in
				}(func(in interface{}) kopsv1alpha2.OpenstackRouter {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandDataSourceOpenstackRouter(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.OpenstackRouter{}
				}(in))
			}(in)
		}(in["router"]),
		BlockStorage: func(in interface{}) *kopsv1alpha2.OpenstackBlockStorageConfig {
			return func(in interface{}) *kopsv1alpha2.OpenstackBlockStorageConfig {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.OpenstackBlockStorageConfig) *kopsv1alpha2.OpenstackBlockStorageConfig {
					return &in
				}(func(in interface{}) kopsv1alpha2.OpenstackBlockStorageConfig {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandDataSourceOpenstackBlockStorageConfig(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.OpenstackBlockStorageConfig{}
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
		Network: func(in interface{}) *kopsv1alpha2.OpenstackNetwork {
			return func(in interface{}) *kopsv1alpha2.OpenstackNetwork {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.OpenstackNetwork) *kopsv1alpha2.OpenstackNetwork {
					return &in
				}(func(in interface{}) kopsv1alpha2.OpenstackNetwork {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandDataSourceOpenstackNetwork(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.OpenstackNetwork{}
				}(in))
			}(in)
		}(in["network"]),
		Metadata: func(in interface{}) *kopsv1alpha2.OpenstackMetadata {
			return func(in interface{}) *kopsv1alpha2.OpenstackMetadata {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.OpenstackMetadata) *kopsv1alpha2.OpenstackMetadata {
					return &in
				}(func(in interface{}) kopsv1alpha2.OpenstackMetadata {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandDataSourceOpenstackMetadata(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.OpenstackMetadata{}
				}(in))
			}(in)
		}(in["metadata"]),
	}
}

func FlattenDataSourceOpenstackSpecInto(in kopsv1alpha2.OpenstackSpec, out map[string]interface{}) {
	out["loadbalancer"] = func(in *kopsv1alpha2.OpenstackLoadbalancerConfig) interface{} {
		return func(in *kopsv1alpha2.OpenstackLoadbalancerConfig) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.OpenstackLoadbalancerConfig) interface{} {
				return func(in kopsv1alpha2.OpenstackLoadbalancerConfig) []interface{} {
					return []interface{}{FlattenDataSourceOpenstackLoadbalancerConfig(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Loadbalancer)
	out["monitor"] = func(in *kopsv1alpha2.OpenstackMonitor) interface{} {
		return func(in *kopsv1alpha2.OpenstackMonitor) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.OpenstackMonitor) interface{} {
				return func(in kopsv1alpha2.OpenstackMonitor) []interface{} {
					return []interface{}{FlattenDataSourceOpenstackMonitor(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Monitor)
	out["router"] = func(in *kopsv1alpha2.OpenstackRouter) interface{} {
		return func(in *kopsv1alpha2.OpenstackRouter) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.OpenstackRouter) interface{} {
				return func(in kopsv1alpha2.OpenstackRouter) []interface{} {
					return []interface{}{FlattenDataSourceOpenstackRouter(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Router)
	out["block_storage"] = func(in *kopsv1alpha2.OpenstackBlockStorageConfig) interface{} {
		return func(in *kopsv1alpha2.OpenstackBlockStorageConfig) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.OpenstackBlockStorageConfig) interface{} {
				return func(in kopsv1alpha2.OpenstackBlockStorageConfig) []interface{} {
					return []interface{}{FlattenDataSourceOpenstackBlockStorageConfig(in)}
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
	out["network"] = func(in *kopsv1alpha2.OpenstackNetwork) interface{} {
		return func(in *kopsv1alpha2.OpenstackNetwork) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.OpenstackNetwork) interface{} {
				return func(in kopsv1alpha2.OpenstackNetwork) []interface{} {
					return []interface{}{FlattenDataSourceOpenstackNetwork(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Network)
	out["metadata"] = func(in *kopsv1alpha2.OpenstackMetadata) interface{} {
		return func(in *kopsv1alpha2.OpenstackMetadata) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.OpenstackMetadata) interface{} {
				return func(in kopsv1alpha2.OpenstackMetadata) []interface{} {
					return []interface{}{FlattenDataSourceOpenstackMetadata(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Metadata)
}

func FlattenDataSourceOpenstackSpec(in kopsv1alpha2.OpenstackSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceOpenstackSpecInto(in, out)
	return out
}
