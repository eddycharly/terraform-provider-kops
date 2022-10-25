package schemas

import (
	"reflect"
	"sort"

	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	coreschemas "github.com/eddycharly/terraform-provider-kops/pkg/schemas/core"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	core "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceKubeDNSConfig() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cache_max_size":       OptionalInt(),
			"cache_max_concurrent": OptionalInt(),
			"tolerations":          OptionalList(coreschemas.ResourceToleration()),
			"affinity":             OptionalStruct(coreschemas.ResourceAffinity()),
			"core_dns_image":       OptionalString(),
			"cpa_image":            OptionalString(),
			"domain":               OptionalString(),
			"external_core_file":   OptionalString(),
			"provider":             OptionalString(),
			"server_ip":            OptionalString(),
			"stub_domains":         OptionalComplexMap(List(String())),
			"upstream_nameservers": OptionalList(String()),
			"memory_request":       OptionalQuantity(),
			"cpu_request":          OptionalQuantity(),
			"memory_limit":         OptionalQuantity(),
			"node_local_dns":       OptionalStruct(ResourceNodeLocalDNSConfig()),
		},
	}

	return res
}

func ExpandResourceKubeDNSConfig(in map[string]interface{}) kops.KubeDNSConfig {
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
		Tolerations: func(in interface{}) []core.Toleration {
			return func(in interface{}) []core.Toleration {
				if in == nil {
					return nil
				}
				var out []core.Toleration
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) core.Toleration {
						if in == nil {
							return core.Toleration{}
						}
						return coreschemas.ExpandResourceToleration(in.(map[string]interface{}))
					}(in))
				}
				return out
			}(in)
		}(in["tolerations"]),
		Affinity: func(in interface{}) *core.Affinity {
			return func(in interface{}) *core.Affinity {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in core.Affinity) *core.Affinity {
					return &in
				}(func(in interface{}) core.Affinity {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return coreschemas.ExpandResourceAffinity(in[0].(map[string]interface{}))
					}
					return core.Affinity{}
				}(in))
			}(in)
		}(in["affinity"]),
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
				if in, ok := in.([]interface{}); ok {
					if len(in) > 0 {
						out := map[string][]string{}
						for _, in := range in {
							if in, ok := in.(map[string]interface{}); ok {
								key := ExpandString(in["key"])
								value := func(in interface{}) []string {
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
								}(in["value"])
								out[key] = value
							}
						}
						return out
					}
				}
				return nil
			}(in)
		}(in["stub_domains"]),
		UpstreamNameservers: func(in interface{}) []string {
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
		}(in["upstream_nameservers"]),
		MemoryRequest: func(in interface{}) *resource.Quantity {
			if in == nil {
				return nil
			}
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
			if in == nil {
				return nil
			}
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
			if in == nil {
				return nil
			}
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
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandResourceNodeLocalDNSConfig(in[0].(map[string]interface{}))
					}
					return kops.NodeLocalDNSConfig{}
				}(in))
			}(in)
		}(in["node_local_dns"]),
	}
}

func FlattenResourceKubeDNSConfigInto(in kops.KubeDNSConfig, out map[string]interface{}) {
	out["cache_max_size"] = func(in int) interface{} {
		return FlattenInt(int(in))
	}(in.CacheMaxSize)
	out["cache_max_concurrent"] = func(in int) interface{} {
		return FlattenInt(int(in))
	}(in.CacheMaxConcurrent)
	out["tolerations"] = func(in []core.Toleration) interface{} {
		return func(in []core.Toleration) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, func(in core.Toleration) interface{} {
					return coreschemas.FlattenResourceToleration(in)
				}(in))
			}
			return out
		}(in)
	}(in.Tolerations)
	out["affinity"] = func(in *core.Affinity) interface{} {
		return func(in *core.Affinity) interface{} {
			if in == nil {
				return nil
			}
			return func(in core.Affinity) interface{} {
				return func(in core.Affinity) []interface{} {
					return []interface{}{coreschemas.FlattenResourceAffinity(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Affinity)
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
	out["provider"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Provider)
	out["server_ip"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.ServerIP)
	out["stub_domains"] = func(in map[string][]string) interface{} {
		return func(in map[string][]string) []interface{} {
			if in == nil {
				return nil
			}
			keys := make([]string, 0, len(in))
			for key := range in {
				keys = append(keys, key)
			}
			sort.SliceStable(keys, func(i, j int) bool {
				return keys[i] > keys[j]
			})
			var out []interface{}
			for _, key := range keys {
				in := in[key]
				out = append(out, map[string]interface{}{
					"key": key,
					"value": func(in []string) []interface{} {
						var out []interface{}
						for _, in := range in {
							out = append(out, FlattenString(string(in)))
						}
						return out
					}(in),
				})
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
				return func(in kops.NodeLocalDNSConfig) []interface{} {
					return []interface{}{FlattenResourceNodeLocalDNSConfig(in)}
				}(in)
			}(*in)
		}(in)
	}(in.NodeLocalDNS)
}

func FlattenResourceKubeDNSConfig(in kops.KubeDNSConfig) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceKubeDNSConfigInto(in, out)
	return out
}
