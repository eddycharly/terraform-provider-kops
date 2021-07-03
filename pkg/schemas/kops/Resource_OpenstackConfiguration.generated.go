package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceOpenstackConfiguration() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"loadbalancer":         Optional(Ptr(Struct(ResourceOpenstackLoadbalancerConfig()))),
			"monitor":              Optional(Ptr(Struct(ResourceOpenstackMonitor()))),
			"router":               Optional(Ptr(Struct(ResourceOpenstackRouter()))),
			"block_storage":        Optional(Ptr(Struct(ResourceOpenstackBlockStorageConfig()))),
			"insecure_skip_verify": Optional(Ptr(Bool())),
			"network":              Optional(Ptr(Struct(ResourceOpenstackNetwork()))),
		},
	}
}

func ExpandResourceOpenstackConfiguration(in map[string]interface{}) kops.OpenstackConfiguration {
	if in == nil {
		panic("expand OpenstackConfiguration failure, in is nil")
	}
	out := kops.OpenstackConfiguration{}
	if in, ok := in["loadbalancer"]; ok && in != nil {
		out.Loadbalancer = func(in interface{}) *kops.OpenstackLoadbalancerConfig {
			if in == nil {
				return nil
			}
			return func(in kops.OpenstackLoadbalancerConfig) *kops.OpenstackLoadbalancerConfig { return &in }(func(in interface{}) kops.OpenstackLoadbalancerConfig {
				if in == nil {
					return kops.OpenstackLoadbalancerConfig{}
				}
				return ExpandResourceOpenstackLoadbalancerConfig(in.(map[string]interface{}))
			}(in))
		}(in)
	}
	if in, ok := in["monitor"]; ok && in != nil {
		out.Monitor = func(in interface{}) *kops.OpenstackMonitor {
			if in == nil {
				return nil
			}
			return func(in kops.OpenstackMonitor) *kops.OpenstackMonitor { return &in }(func(in interface{}) kops.OpenstackMonitor {
				if in == nil {
					return kops.OpenstackMonitor{}
				}
				return ExpandResourceOpenstackMonitor(in.(map[string]interface{}))
			}(in))
		}(in)
	}
	if in, ok := in["router"]; ok && in != nil {
		out.Router = func(in interface{}) *kops.OpenstackRouter {
			if in == nil {
				return nil
			}
			return func(in kops.OpenstackRouter) *kops.OpenstackRouter { return &in }(func(in interface{}) kops.OpenstackRouter {
				if in == nil {
					return kops.OpenstackRouter{}
				}
				return ExpandResourceOpenstackRouter(in.(map[string]interface{}))
			}(in))
		}(in)
	}
	if in, ok := in["block_storage"]; ok && in != nil {
		out.BlockStorage = func(in interface{}) *kops.OpenstackBlockStorageConfig {
			if in == nil {
				return nil
			}
			return func(in kops.OpenstackBlockStorageConfig) *kops.OpenstackBlockStorageConfig { return &in }(func(in interface{}) kops.OpenstackBlockStorageConfig {
				if in == nil {
					return kops.OpenstackBlockStorageConfig{}
				}
				return ExpandResourceOpenstackBlockStorageConfig(in.(map[string]interface{}))
			}(in))
		}(in)
	}
	if in, ok := in["insecure_skip_verify"]; ok && in != nil {
		out.InsecureSkipVerify = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["network"]; ok && in != nil {
		out.Network = func(in interface{}) *kops.OpenstackNetwork {
			if in == nil {
				return nil
			}
			return func(in kops.OpenstackNetwork) *kops.OpenstackNetwork { return &in }(func(in interface{}) kops.OpenstackNetwork {
				if in == nil {
					return kops.OpenstackNetwork{}
				}
				return ExpandResourceOpenstackNetwork(in.(map[string]interface{}))
			}(in))
		}(in)
	}
	return out
}
