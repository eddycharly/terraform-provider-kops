package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/apimachinery/pkg/api/resource"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceEtcdClusterSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name":                    Computed(String()),
			"provider":                Computed(String()),
			"member":                  Computed(List(DataSourceEtcdMemberSpec())),
			"enable_etcd_tls":         Computed(Bool()),
			"enable_tls_auth":         Computed(Bool()),
			"version":                 Computed(String()),
			"leader_election_timeout": Computed(Nullable(Duration())),
			"heartbeat_interval":      Computed(Nullable(Duration())),
			"image":                   Computed(String()),
			"backups":                 Computed(Struct(DataSourceEtcdBackupSpec())),
			"manager":                 Computed(Struct(DataSourceEtcdManagerSpec())),
			"memory_request":          Computed(Nullable(Quantity())),
			"cpu_request":             Computed(Nullable(Quantity())),
		},
	}
}

func ExpandDataSourceEtcdClusterSpec(in map[string]interface{}) kops.EtcdClusterSpec {
	if in == nil {
		panic("expand EtcdClusterSpec failure, in is nil")
	}
	out := kops.EtcdClusterSpec{}
	if in, ok := in["name"]; ok && in != nil {
		out.Name = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["provider"]; ok && in != nil {
		out.Provider = func(in interface{}) kops.EtcdProviderType { return kops.EtcdProviderType(in.(string)) }(in)
	}
	if in, ok := in["member"]; ok && in != nil {
		out.Members = func(in interface{}) []kops.EtcdMemberSpec {
			var out []kops.EtcdMemberSpec
			for _, in := range in.([]interface{}) {
				out = append(out, func(in interface{}) kops.EtcdMemberSpec {
					if in == nil {
						return kops.EtcdMemberSpec{}
					}
					return ExpandDataSourceEtcdMemberSpec(in.(map[string]interface{}))
				}(in))
			}
			return out
		}(in)
	}
	if in, ok := in["enable_etcd_tls"]; ok && in != nil {
		out.EnableEtcdTLS = func(in interface{}) bool { return in.(bool) }(in)
	}
	if in, ok := in["enable_tls_auth"]; ok && in != nil {
		out.EnableTLSAuth = func(in interface{}) bool { return in.(bool) }(in)
	}
	if in, ok := in["version"]; ok && in != nil {
		out.Version = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["leader_election_timeout"]; ok && in != nil {
		out.LeaderElectionTimeout = func(in interface{}) *v1.Duration {
			if in == nil {
				return nil
			}
			return func(in v1.Duration) *v1.Duration { return &in }(func(in interface{}) v1.Duration { return ExpandDuration(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["heartbeat_interval"]; ok && in != nil {
		out.HeartbeatInterval = func(in interface{}) *v1.Duration {
			if in == nil {
				return nil
			}
			return func(in v1.Duration) *v1.Duration { return &in }(func(in interface{}) v1.Duration { return ExpandDuration(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["image"]; ok && in != nil {
		out.Image = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["backups"]; ok && in != nil {
		out.Backups = func(in interface{}) *kops.EtcdBackupSpec {
			if in == nil {
				return nil
			}
			return func(in kops.EtcdBackupSpec) *kops.EtcdBackupSpec { return &in }(func(in interface{}) kops.EtcdBackupSpec {
				if in == nil {
					return kops.EtcdBackupSpec{}
				}
				return ExpandDataSourceEtcdBackupSpec(in.(map[string]interface{}))
			}(in))
		}(in)
	}
	if in, ok := in["manager"]; ok && in != nil {
		out.Manager = func(in interface{}) *kops.EtcdManagerSpec {
			if in == nil {
				return nil
			}
			return func(in kops.EtcdManagerSpec) *kops.EtcdManagerSpec { return &in }(func(in interface{}) kops.EtcdManagerSpec {
				if in == nil {
					return kops.EtcdManagerSpec{}
				}
				return ExpandDataSourceEtcdManagerSpec(in.(map[string]interface{}))
			}(in))
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
	return out
}

func FlattenDataSourceEtcdClusterSpecInto(in kops.EtcdClusterSpec, out map[string]interface{}) {
	out["name"] = func(in string) interface{} { return string(in) }(in.Name)
	out["provider"] = func(in kops.EtcdProviderType) interface{} { return string(in) }(in.Provider)
	out["member"] = func(in []kops.EtcdMemberSpec) interface{} {
		var out []interface{}
		for _, in := range in {
			out = append(out, func(in kops.EtcdMemberSpec) interface{} { return FlattenDataSourceEtcdMemberSpec(in) }(in))
		}
		return out
	}(in.Members)
	out["enable_etcd_tls"] = func(in bool) interface{} { return in }(in.EnableEtcdTLS)
	out["enable_tls_auth"] = func(in bool) interface{} { return in }(in.EnableTLSAuth)
	out["version"] = func(in string) interface{} { return string(in) }(in.Version)
	out["leader_election_timeout"] = func(in *v1.Duration) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in v1.Duration) interface{} { return FlattenDuration(in) }(*in)}
	}(in.LeaderElectionTimeout)
	out["heartbeat_interval"] = func(in *v1.Duration) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in v1.Duration) interface{} { return FlattenDuration(in) }(*in)}
	}(in.HeartbeatInterval)
	out["image"] = func(in string) interface{} { return string(in) }(in.Image)
	out["backups"] = func(in *kops.EtcdBackupSpec) interface{} {
		if in == nil {
			return nil
		}
		return func(in kops.EtcdBackupSpec) interface{} { return FlattenDataSourceEtcdBackupSpec(in) }(*in)
	}(in.Backups)
	out["manager"] = func(in *kops.EtcdManagerSpec) interface{} {
		if in == nil {
			return nil
		}
		return func(in kops.EtcdManagerSpec) interface{} { return FlattenDataSourceEtcdManagerSpec(in) }(*in)
	}(in.Manager)
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
}

func FlattenDataSourceEtcdClusterSpec(in kops.EtcdClusterSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceEtcdClusterSpecInto(in, out)
	return out
}
