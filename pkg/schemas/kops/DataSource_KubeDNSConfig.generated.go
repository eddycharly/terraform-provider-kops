package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceKubeDNSConfig() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cache_max_size":       Computed(Int()),
			"cache_max_concurrent": Computed(Int()),
			"core_dns_image":       Computed(String()),
			"cpa_image":            Computed(String()),
			"domain":               Computed(String()),
			"external_core_file":   Computed(String()),
			"image":                Computed(String()),
			"replicas":             Computed(Int()),
			"provider":             Computed(String()),
			"server_ip":            Computed(String()),
			"stub_domains":         Computed(Map(List(String()))),
			"upstream_nameservers": Computed(List(String())),
			"memory_request":       Computed(Nullable(Quantity())),
			"cpu_request":          Computed(Nullable(Quantity())),
			"memory_limit":         Computed(Nullable(Quantity())),
			"node_local_dns":       Computed(Struct(DataSourceNodeLocalDNSConfig())),
		},
	}
}

func ExpandDataSourceKubeDNSConfig(in map[string]interface{}) kops.KubeDNSConfig {
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
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in resource.Quantity) *resource.Quantity { return &in }(func(in interface{}) resource.Quantity { return ExpandQuantity(in.(string)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["cpu_request"]; ok && in != nil {
		out.CPURequest = func(in interface{}) *resource.Quantity {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in resource.Quantity) *resource.Quantity { return &in }(func(in interface{}) resource.Quantity { return ExpandQuantity(in.(string)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["memory_limit"]; ok && in != nil {
		out.MemoryLimit = func(in interface{}) *resource.Quantity {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in resource.Quantity) *resource.Quantity { return &in }(func(in interface{}) resource.Quantity { return ExpandQuantity(in.(string)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["node_local_dns"]; ok && in != nil {
		out.NodeLocalDNS = func(in interface{}) *kops.NodeLocalDNSConfig {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in kops.NodeLocalDNSConfig) *kops.NodeLocalDNSConfig { return &in }(func(in interface{}) kops.NodeLocalDNSConfig {
					return ExpandDataSourceNodeLocalDNSConfig(in.(map[string]interface{}))
				}(in[0]))
			}
			return nil
		}(in)
	}
	return out
}

func FlattenDataSourceKubeDNSConfigInto(in kops.KubeDNSConfig, out map[string]interface{}) {
	out["cache_max_size"] = func(in int) interface{} { return int(in) }(in.CacheMaxSize)
	out["cache_max_concurrent"] = func(in int) interface{} { return int(in) }(in.CacheMaxConcurrent)
	out["core_dns_image"] = func(in string) interface{} { return string(in) }(in.CoreDNSImage)
	out["cpa_image"] = func(in string) interface{} { return string(in) }(in.CPAImage)
	out["domain"] = func(in string) interface{} { return string(in) }(in.Domain)
	out["external_core_file"] = func(in string) interface{} { return string(in) }(in.ExternalCoreFile)
	out["image"] = func(in string) interface{} { return string(in) }(in.Image)
	out["replicas"] = func(in int) interface{} { return int(in) }(in.Replicas)
	out["provider"] = func(in string) interface{} { return string(in) }(in.Provider)
	out["server_ip"] = func(in string) interface{} { return string(in) }(in.ServerIP)
	out["stub_domains"] = func(in map[string][]string) interface{} {
		if in == nil {
			return nil
		}
		out := map[string]interface{}{}
		for key, in := range in {
			out[key] = func(in []string) interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, func(in string) interface{} { return string(in) }(in))
				}
				return out
			}(in)
		}
		return out
	}(in.StubDomains)
	out["upstream_nameservers"] = func(in []string) interface{} {
		var out []interface{}
		for _, in := range in {
			out = append(out, func(in string) interface{} { return string(in) }(in))
		}
		return out
	}(in.UpstreamNameservers)
	out["memory_request"] = func(in *resource.Quantity) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in resource.Quantity) interface{} { return FlattenQuantity(in) }(*in)}
	}(in.MemoryRequest)
	out["cpu_request"] = func(in *resource.Quantity) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in resource.Quantity) interface{} { return FlattenQuantity(in) }(*in)}
	}(in.CPURequest)
	out["memory_limit"] = func(in *resource.Quantity) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in resource.Quantity) interface{} { return FlattenQuantity(in) }(*in)}
	}(in.MemoryLimit)
	out["node_local_dns"] = func(in *kops.NodeLocalDNSConfig) interface{} {
		if in == nil {
			return nil
		}
		return func(in kops.NodeLocalDNSConfig) interface{} { return FlattenDataSourceNodeLocalDNSConfig(in) }(*in)
	}(in.NodeLocalDNS)
}

func FlattenDataSourceKubeDNSConfig(in kops.KubeDNSConfig) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceKubeDNSConfigInto(in, out)
	return out
}
