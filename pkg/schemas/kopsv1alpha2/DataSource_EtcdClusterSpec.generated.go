package schemas

import (
	"reflect"

	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/apimachinery/pkg/api/resource"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/kops/pkg/apis/kops/v1alpha2"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
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

func ExpandDataSourceEtcdClusterSpec(in map[string]interface{}) kopsv1alpha2.EtcdClusterSpec {
	if in == nil {
		panic("expand EtcdClusterSpec failure, in is nil")
	}
	return kopsv1alpha2.EtcdClusterSpec{
		Name: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["name"]),
		Provider: func(in interface{}) kopsv1alpha2.EtcdProviderType {
			return v1alpha2.EtcdProviderType(ExpandString(in))
		}(in["provider"]),
		Members: func(in interface{}) []kopsv1alpha2.EtcdMemberSpec {
			return func(in interface{}) []kopsv1alpha2.EtcdMemberSpec {
				if in == nil {
					return nil
				}
				var out []kopsv1alpha2.EtcdMemberSpec
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) kopsv1alpha2.EtcdMemberSpec {
						if in == nil {
							return kopsv1alpha2.EtcdMemberSpec{}
						}
						return ExpandDataSourceEtcdMemberSpec(in.(map[string]interface{}))
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
		LeaderElectionTimeout: func(in interface{}) *meta.Duration {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *meta.Duration {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in meta.Duration) *meta.Duration {
					return &in
				}(ExpandDuration(in))
			}(in)
		}(in["leader_election_timeout"]),
		HeartbeatInterval: func(in interface{}) *meta.Duration {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *meta.Duration {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in meta.Duration) *meta.Duration {
					return &in
				}(ExpandDuration(in))
			}(in)
		}(in["heartbeat_interval"]),
		Image: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["image"]),
		Backups: func(in interface{}) *kopsv1alpha2.EtcdBackupSpec {
			return func(in interface{}) *kopsv1alpha2.EtcdBackupSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.EtcdBackupSpec) *kopsv1alpha2.EtcdBackupSpec {
					return &in
				}(func(in interface{}) kopsv1alpha2.EtcdBackupSpec {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandDataSourceEtcdBackupSpec(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.EtcdBackupSpec{}
				}(in))
			}(in)
		}(in["backups"]),
		Manager: func(in interface{}) *kopsv1alpha2.EtcdManagerSpec {
			return func(in interface{}) *kopsv1alpha2.EtcdManagerSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.EtcdManagerSpec) *kopsv1alpha2.EtcdManagerSpec {
					return &in
				}(func(in interface{}) kopsv1alpha2.EtcdManagerSpec {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandDataSourceEtcdManagerSpec(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.EtcdManagerSpec{}
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

func FlattenDataSourceEtcdClusterSpecInto(in kopsv1alpha2.EtcdClusterSpec, out map[string]interface{}) {
	out["name"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Name)
	out["provider"] = func(in kopsv1alpha2.EtcdProviderType) interface{} {
		return FlattenString(string(in))
	}(in.Provider)
	out["member"] = func(in []kopsv1alpha2.EtcdMemberSpec) interface{} {
		return func(in []kopsv1alpha2.EtcdMemberSpec) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, func(in kopsv1alpha2.EtcdMemberSpec) interface{} {
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
	out["leader_election_timeout"] = func(in *meta.Duration) interface{} {
		return func(in *meta.Duration) interface{} {
			if in == nil {
				return nil
			}
			return func(in meta.Duration) interface{} {
				return FlattenDuration(in)
			}(*in)
		}(in)
	}(in.LeaderElectionTimeout)
	out["heartbeat_interval"] = func(in *meta.Duration) interface{} {
		return func(in *meta.Duration) interface{} {
			if in == nil {
				return nil
			}
			return func(in meta.Duration) interface{} {
				return FlattenDuration(in)
			}(*in)
		}(in)
	}(in.HeartbeatInterval)
	out["image"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Image)
	out["backups"] = func(in *kopsv1alpha2.EtcdBackupSpec) interface{} {
		return func(in *kopsv1alpha2.EtcdBackupSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.EtcdBackupSpec) interface{} {
				return func(in kopsv1alpha2.EtcdBackupSpec) []interface{} {
					return []interface{}{FlattenDataSourceEtcdBackupSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Backups)
	out["manager"] = func(in *kopsv1alpha2.EtcdManagerSpec) interface{} {
		return func(in *kopsv1alpha2.EtcdManagerSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.EtcdManagerSpec) interface{} {
				return func(in kopsv1alpha2.EtcdManagerSpec) []interface{} {
					return []interface{}{FlattenDataSourceEtcdManagerSpec(in)}
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

func FlattenDataSourceEtcdClusterSpec(in kopsv1alpha2.EtcdClusterSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceEtcdClusterSpecInto(in, out)
	return out
}
