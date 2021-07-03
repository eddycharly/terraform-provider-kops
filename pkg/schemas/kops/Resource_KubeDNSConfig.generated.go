package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceKubeDNSConfig() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cache_max_size":       Optional(Int()),
			"cache_max_concurrent": Optional(Int()),
			"core_dns_image":       Optional(String()),
			"cpa_image":            Optional(String()),
			"domain":               Optional(String()),
			"external_core_file":   Optional(String()),
			"image":                Optional(String()),
			"replicas":             Optional(Int()),
			"provider":             Optional(String()),
			"server_ip":            Optional(String()),
			"stub_domains":         Optional(Map(List(String()))),
			"upstream_nameservers": Optional(List(String())),
			"memory_request":       Optional(Ptr(Quantity())),
			"cpu_request":          Optional(Ptr(Quantity())),
			"memory_limit":         Optional(Ptr(Quantity())),
			"node_local_dns":       Optional(Ptr(Struct(ResourceNodeLocalDNSConfig()))),
		},
	}
}

func ExpandResourceKubeDNSConfig(in map[string]interface{}) kops.KubeDNSConfig {
	if in == nil {
		panic("expand KubeDNSConfig failure, in is nil")
	}
	out := kops.KubeDNSConfig{}
	if in, ok := in["cache_max_size"]; ok && in != nil {
		out.CacheMaxSize = func(in interface{}) int { return int(in.(int)) }(in)
	}
	if in, ok := in["cache_max_concurrent"]; ok && in != nil {
		out.CacheMaxConcurrent = func(in interface{}) int { return int(in.(int)) }(in)
	}
	if in, ok := in["core_dns_image"]; ok && in != nil {
		out.CoreDNSImage = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["cpa_image"]; ok && in != nil {
		out.CPAImage = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["domain"]; ok && in != nil {
		out.Domain = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["external_core_file"]; ok && in != nil {
		out.ExternalCoreFile = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["image"]; ok && in != nil {
		out.Image = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["replicas"]; ok && in != nil {
		out.Replicas = func(in interface{}) int { return int(in.(int)) }(in)
	}
	if in, ok := in["provider"]; ok && in != nil {
		out.Provider = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["server_ip"]; ok && in != nil {
		out.ServerIP = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["stub_domains"]; ok && in != nil {
		out.StubDomains = func(in interface{}) map[string][]string {
			if in == nil {
				return nil
			}
			out := map[string][]string{}
			for key, in := range in.(map[string]interface{}) {
				out[key] = func(in interface{}) []string {
					var out []string
					for _, in := range in.([]interface{}) {
						out = append(out, func(in interface{}) string { return string(in.(string)) }(in))
					}
					return out
				}(in)
			}
			return out
		}(in)
	}
	if in, ok := in["upstream_nameservers"]; ok && in != nil {
		out.UpstreamNameservers = func(in interface{}) []string {
			var out []string
			for _, in := range in.([]interface{}) {
				out = append(out, func(in interface{}) string { return string(in.(string)) }(in))
			}
			return out
		}(in)
	}
	if in, ok := in["memory_request"]; ok && in != nil {
		out.MemoryRequest = func(in interface{}) *resource.Quantity {
			if in == nil {
				return nil
			}
			return func(in resource.Quantity) *resource.Quantity { return &in }(func(in interface{}) resource.Quantity { return ExpandQuantity(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["cpu_request"]; ok && in != nil {
		out.CPURequest = func(in interface{}) *resource.Quantity {
			if in == nil {
				return nil
			}
			return func(in resource.Quantity) *resource.Quantity { return &in }(func(in interface{}) resource.Quantity { return ExpandQuantity(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["memory_limit"]; ok && in != nil {
		out.MemoryLimit = func(in interface{}) *resource.Quantity {
			if in == nil {
				return nil
			}
			return func(in resource.Quantity) *resource.Quantity { return &in }(func(in interface{}) resource.Quantity { return ExpandQuantity(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["node_local_dns"]; ok && in != nil {
		out.NodeLocalDNS = func(in interface{}) *kops.NodeLocalDNSConfig {
			if in == nil {
				return nil
			}
			return func(in kops.NodeLocalDNSConfig) *kops.NodeLocalDNSConfig { return &in }(func(in interface{}) kops.NodeLocalDNSConfig {
				if in == nil {
					return kops.NodeLocalDNSConfig{}
				}
				return ExpandResourceNodeLocalDNSConfig(in.(map[string]interface{}))
			}(in))
		}(in)
	}
	return out
}
