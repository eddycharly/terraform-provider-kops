package schemas

import (
	"reflect"

	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/apimachinery/pkg/api/resource"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceEtcdClusterSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name":                    ComputedString(),
			"provider":                ComputedString(),
			"member":                  ComputedList(DataSourceEtcdMemberSpec()),
			"enable_etcd_tls":         ComputedBool(),
			"enable_tls_auth":         ComputedBool(),
			"version":                 ComputedString(),
			"leader_election_timeout": ComputedDuration(),
			"heartbeat_interval":      ComputedDuration(),
			"image":                   ComputedString(),
			"backups":                 ComputedStruct(DataSourceEtcdBackupSpec()),
			"manager":                 ComputedStruct(DataSourceEtcdManagerSpec()),
			"memory_request":          ComputedQuantity(),
			"cpu_request":             ComputedQuantity(),
		},
	}

	return res
}

func ExpandDataSourceEtcdClusterSpec(in map[string]interface{}) kops.EtcdClusterSpec {
	if in == nil {
		panic("expand EtcdClusterSpec failure, in is nil")
	}
	return kops.EtcdClusterSpec{
		Name: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["name"]),
		Provider: func(in interface{}) kops.EtcdProviderType {
			return kops.EtcdProviderType(ExpandString(in))
		}(in["provider"]),
		Members: func(in interface{}) []kops.EtcdMemberSpec {
			return func(in interface{}) []kops.EtcdMemberSpec {
				if in == nil {
					return nil
				}
				var out []kops.EtcdMemberSpec
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) kops.EtcdMemberSpec {
						if in == nil {
							return kops.EtcdMemberSpec{}
						}
						return (ExpandDataSourceEtcdMemberSpec(in.(map[string]interface{})))
					}(in))
				}
				return out
			}(in)
		}(in["member"]),
		EnableEtcdTLS: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["enable_etcd_tls"]),
		EnableTLSAuth: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["enable_tls_auth"]),
		Version: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["version"]),
		LeaderElectionTimeout: func(in interface{}) *v1.Duration {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *v1.Duration {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in v1.Duration) *v1.Duration {
					return &in
				}(ExpandDuration(in))
			}(in)
		}(in["leader_election_timeout"]),
		HeartbeatInterval: func(in interface{}) *v1.Duration {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *v1.Duration {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in v1.Duration) *v1.Duration {
					return &in
				}(ExpandDuration(in))
			}(in)
		}(in["heartbeat_interval"]),
		Image: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["image"]),
		Backups: func(in interface{}) *kops.EtcdBackupSpec {
			return func(in interface{}) *kops.EtcdBackupSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.EtcdBackupSpec) *kops.EtcdBackupSpec {
					return &in
				}(func(in interface{}) kops.EtcdBackupSpec {
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return kops.EtcdBackupSpec{}
					}
					return (ExpandDataSourceEtcdBackupSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["backups"]),
		Manager: func(in interface{}) *kops.EtcdManagerSpec {
			return func(in interface{}) *kops.EtcdManagerSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.EtcdManagerSpec) *kops.EtcdManagerSpec {
					return &in
				}(func(in interface{}) kops.EtcdManagerSpec {
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return kops.EtcdManagerSpec{}
					}
					return (ExpandDataSourceEtcdManagerSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["manager"]),
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
	}
}

func FlattenDataSourceEtcdClusterSpecInto(in kops.EtcdClusterSpec, out map[string]interface{}) {
	out["name"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Name)
	out["provider"] = func(in kops.EtcdProviderType) interface{} {
		return FlattenString(string(in))
	}(in.Provider)
	out["member"] = func(in []kops.EtcdMemberSpec) interface{} {
		return func(in []kops.EtcdMemberSpec) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, func(in kops.EtcdMemberSpec) interface{} {
					return FlattenDataSourceEtcdMemberSpec(in)
				}(in))
			}
			return out
		}(in)
	}(in.Members)
	out["enable_etcd_tls"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.EnableEtcdTLS)
	out["enable_tls_auth"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.EnableTLSAuth)
	out["version"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Version)
	out["leader_election_timeout"] = func(in *v1.Duration) interface{} {
		return func(in *v1.Duration) interface{} {
			if in == nil {
				return nil
			}
			return func(in v1.Duration) interface{} {
				return FlattenDuration(in)
			}(*in)
		}(in)
	}(in.LeaderElectionTimeout)
	out["heartbeat_interval"] = func(in *v1.Duration) interface{} {
		return func(in *v1.Duration) interface{} {
			if in == nil {
				return nil
			}
			return func(in v1.Duration) interface{} {
				return FlattenDuration(in)
			}(*in)
		}(in)
	}(in.HeartbeatInterval)
	out["image"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Image)
	out["backups"] = func(in *kops.EtcdBackupSpec) interface{} {
		return func(in *kops.EtcdBackupSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.EtcdBackupSpec) interface{} {
				return func(in kops.EtcdBackupSpec) []map[string]interface{} {
					return []map[string]interface{}{FlattenDataSourceEtcdBackupSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Backups)
	out["manager"] = func(in *kops.EtcdManagerSpec) interface{} {
		return func(in *kops.EtcdManagerSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.EtcdManagerSpec) interface{} {
				return func(in kops.EtcdManagerSpec) []map[string]interface{} {
					return []map[string]interface{}{FlattenDataSourceEtcdManagerSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Manager)
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
}

func FlattenDataSourceEtcdClusterSpec(in kops.EtcdClusterSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceEtcdClusterSpecInto(in, out)
	return out
}
