package schemas

import (
	"reflect"

	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceKubeDNSConfig() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cache_max_size":       ComputedInt(),
			"cache_max_concurrent": ComputedInt(),
			"core_dns_image":       ComputedString(),
			"cpa_image":            ComputedString(),
			"domain":               ComputedString(),
			"external_core_file":   ComputedString(),
			"image":                ComputedString(),
			"replicas":             ComputedInt(),
			"provider":             ComputedString(),
			"server_ip":            ComputedString(),
			"stub_domains":         ComputedMap(List(String())),
			"upstream_nameservers": ComputedList(String()),
			"memory_request":       ComputedQuantity(),
			"cpu_request":          ComputedQuantity(),
			"memory_limit":         ComputedQuantity(),
			"node_local_dns":       ComputedStruct(DataSourceNodeLocalDNSConfig()),
		},
	}
}

func ExpandDataSourceKubeDNSConfig(in map[string]interface{}) kops.KubeDNSConfig {
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
		CPAImage: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["cpa_image"]),
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
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return kops.NodeLocalDNSConfig{}
					}
					return (ExpandDataSourceNodeLocalDNSConfig(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["node_local_dns"]),
	}
}

func FlattenDataSourceKubeDNSConfigInto(in kops.KubeDNSConfig, out map[string]interface{}) {
	out["cache_max_size"] = func(in int) interface{} {
		return FlattenInt(int(in))
	}(in.CacheMaxSize)
	out["cache_max_concurrent"] = func(in int) interface{} {
		return FlattenInt(int(in))
	}(in.CacheMaxConcurrent)
	out["core_dns_image"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.CoreDNSImage)
	out["cpa_image"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.CPAImage)
	out["domain"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Domain)
	out["external_core_file"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.ExternalCoreFile)
	out["image"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Image)
	out["replicas"] = func(in int) interface{} {
		return FlattenInt(int(in))
	}(in.Replicas)
	out["provider"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Provider)
	out["server_ip"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.ServerIP)
	out["stub_domains"] = func(in map[string][]string) interface{} {
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
	}(in.StubDomains)
	out["upstream_nameservers"] = func(in []string) interface{} {
		return func(in []string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.UpstreamNameservers)
	out["memory_request"] = func(in *resource.Quantity) interface{} {
		return func(in *resource.Quantity) interface{} {
			if in == nil {
				return nil
			}
			return func(in resource.Quantity) interface{} {
				return FlattenQuantity(in)
			}(*in)
		}(in)
	}(in.MemoryRequest)
	out["cpu_request"] = func(in *resource.Quantity) interface{} {
		return func(in *resource.Quantity) interface{} {
			if in == nil {
				return nil
			}
			return func(in resource.Quantity) interface{} {
				return FlattenQuantity(in)
			}(*in)
		}(in)
	}(in.CPURequest)
	out["memory_limit"] = func(in *resource.Quantity) interface{} {
		return func(in *resource.Quantity) interface{} {
			if in == nil {
				return nil
			}
			return func(in resource.Quantity) interface{} {
				return FlattenQuantity(in)
			}(*in)
		}(in)
	}(in.MemoryLimit)
	out["node_local_dns"] = func(in *kops.NodeLocalDNSConfig) interface{} {
		return func(in *kops.NodeLocalDNSConfig) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.NodeLocalDNSConfig) interface{} {
				return func(in kops.NodeLocalDNSConfig) []map[string]interface{} {
					return []map[string]interface{}{FlattenDataSourceNodeLocalDNSConfig(in)}
				}(in)
			}(*in)
		}(in)
	}(in.NodeLocalDNS)
}

func FlattenDataSourceKubeDNSConfig(in kops.KubeDNSConfig) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceKubeDNSConfigInto(in, out)
	return out
}
