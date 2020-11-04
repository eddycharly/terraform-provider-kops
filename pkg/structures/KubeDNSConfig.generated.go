package structures

import (
	"reflect"

	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/kops/pkg/apis/kops"
)

func ExpandKubeDNSConfig(in map[string]interface{}) kops.KubeDNSConfig {
	if in == nil {
		panic("expand KubeDNSConfig failure, in is nil")
	}
	return kops.KubeDNSConfig{
		CacheMaxSize: func(in interface{}) int {
			return int(ExpandInt(in))
		}(in["cache_max_size"]),
		CacheMaxConcurrent: func(in interface{}) int {
			return int(ExpandInt(in))
		}(in["cache_max_concurrent"]),
		CoreDNSImage: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["core_dns_image"]),
		Domain: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["domain"]),
		ExternalCoreFile: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["external_core_file"]),
		Image: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["image"]),
		Replicas: func(in interface{}) int {
			return int(ExpandInt(in))
		}(in["replicas"]),
		Provider: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["provider"]),
		ServerIP: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["server_ip"]),
		StubDomains: func(in interface{}) map[string][]string {
			return func(in interface{}) map[string][]string {
				if in == nil {
					return nil
				}
				out := map[string][]string{}
				for key, in := range in.(map[string]interface{}) {
					out[key] = func(in interface{}) []string {
						var out []string
						for _, in := range in.([]interface{}) {
							out = append(out, string(ExpandString(in)))
						}
						return out
					}(in)
				}
				return out
			}(in)
		}(in["stub_domains"]),
		UpstreamNameservers: func(in interface{}) []string {
			return func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["upstream_nameservers"]),
		MemoryRequest: func(in interface{}) *resource.Quantity {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *resource.Quantity {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in resource.Quantity) *resource.Quantity {
					return &in
				}(ExpandQuantity(in))
			}(in)
		}(in["memory_request"]),
		CPURequest: func(in interface{}) *resource.Quantity {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *resource.Quantity {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in resource.Quantity) *resource.Quantity {
					return &in
				}(ExpandQuantity(in))
			}(in)
		}(in["cpu_request"]),
		MemoryLimit: func(in interface{}) *resource.Quantity {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *resource.Quantity {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in resource.Quantity) *resource.Quantity {
					return &in
				}(ExpandQuantity(in))
			}(in)
		}(in["memory_limit"]),
		NodeLocalDNS: func(in interface{}) *kops.NodeLocalDNSConfig {
			return func(in interface{}) *kops.NodeLocalDNSConfig {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.NodeLocalDNSConfig) *kops.NodeLocalDNSConfig {
					return &in
				}(func(in interface{}) kops.NodeLocalDNSConfig {
					if in.([]interface{})[0] == nil {
						return kops.NodeLocalDNSConfig{}
					}
					return (ExpandNodeLocalDNSConfig(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["node_local_dns"]),
	}
}

func FlattenKubeDNSConfig(in kops.KubeDNSConfig) map[string]interface{} {
	return map[string]interface{}{
		"cache_max_size": func(in int) interface{} {
			return FlattenInt(int(in))
		}(in.CacheMaxSize),
		"cache_max_concurrent": func(in int) interface{} {
			return FlattenInt(int(in))
		}(in.CacheMaxConcurrent),
		"core_dns_image": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.CoreDNSImage),
		"domain": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.Domain),
		"external_core_file": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.ExternalCoreFile),
		"image": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.Image),
		"replicas": func(in int) interface{} {
			return FlattenInt(int(in))
		}(in.Replicas),
		"provider": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.Provider),
		"server_ip": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.ServerIP),
		"stub_domains": func(in map[string][]string) interface{} {
			return func(in map[string][]string) map[string]interface{} {
				if in == nil {
					return nil
				}
				out := map[string]interface{}{}
				for key, in := range in {
					out[key] = func(in []string) []interface{} {
						var out []interface{}
						for _, in := range in {
							out = append(out, FlattenString(string(in)))
						}
						return out
					}(in)
				}
				return out
			}(in)
		}(in.StubDomains),
		"upstream_nameservers": func(in []string) interface{} {
			return func(in []string) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, FlattenString(string(in)))
				}
				return out
			}(in)
		}(in.UpstreamNameservers),
		"memory_request": func(in *resource.Quantity) interface{} {
			return func(in *resource.Quantity) interface{} {
				if in == nil {
					return nil
				}
				return func(in resource.Quantity) interface{} {
					return FlattenQuantity(in)
				}(*in)
			}(in)
		}(in.MemoryRequest),
		"cpu_request": func(in *resource.Quantity) interface{} {
			return func(in *resource.Quantity) interface{} {
				if in == nil {
					return nil
				}
				return func(in resource.Quantity) interface{} {
					return FlattenQuantity(in)
				}(*in)
			}(in)
		}(in.CPURequest),
		"memory_limit": func(in *resource.Quantity) interface{} {
			return func(in *resource.Quantity) interface{} {
				if in == nil {
					return nil
				}
				return func(in resource.Quantity) interface{} {
					return FlattenQuantity(in)
				}(*in)
			}(in)
		}(in.MemoryLimit),
		"node_local_dns": func(in *kops.NodeLocalDNSConfig) interface{} {
			return func(in *kops.NodeLocalDNSConfig) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.NodeLocalDNSConfig) interface{} {
					return func(in kops.NodeLocalDNSConfig) []map[string]interface{} {
						return []map[string]interface{}{FlattenNodeLocalDNSConfig(in)}
					}(in)
				}(*in)
			}(in)
		}(in.NodeLocalDNS),
	}
}
