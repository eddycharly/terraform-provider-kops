package structures

import (
	"reflect"

	"k8s.io/apimachinery/pkg/api/resource"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/kops/pkg/apis/kops"
)

func ExpandEtcdClusterSpec(in map[string]interface{}) kops.EtcdClusterSpec {
	if in == nil {
		panic("expand EtcdClusterSpec failure, in is nil")
	}
	return kops.EtcdClusterSpec{
		Name: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["name"]),
		Provider: func(in interface{}) kops.EtcdProviderType {
			value := kops.EtcdProviderType(ExpandString(in))
			return value
		}(in["provider"]),
		Members: func(in interface{}) []*kops.EtcdMemberSpec {
			value := func(in interface{}) []*kops.EtcdMemberSpec {
				var out []*kops.EtcdMemberSpec
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) *kops.EtcdMemberSpec {
						if in == nil {
							return nil
						}
						if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
							return nil
						}
						return func(in kops.EtcdMemberSpec) *kops.EtcdMemberSpec {
							return &in
						}(func(in interface{}) kops.EtcdMemberSpec {
							if in == nil {
								return kops.EtcdMemberSpec{}
							}
							return (ExpandEtcdMemberSpec(in.(map[string]interface{})))
						}(in))
					}(in))
				}
				return out
			}(in)
			return value
		}(in["members"]),
		EnableEtcdTLS: func(in interface{}) bool {
			value := bool(ExpandBool(in))
			return value
		}(in["enable_etcd_tls"]),
		EnableTLSAuth: func(in interface{}) bool {
			value := bool(ExpandBool(in))
			return value
		}(in["enable_tls_auth"]),
		Version: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["version"]),
		LeaderElectionTimeout: func(in interface{}) *v1.Duration {
			value := func(in interface{}) *v1.Duration {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in v1.Duration) *v1.Duration {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(ExpandDuration(in))
			}(in)
			return value
		}(in["leader_election_timeout"]),
		HeartbeatInterval: func(in interface{}) *v1.Duration {
			value := func(in interface{}) *v1.Duration {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in v1.Duration) *v1.Duration {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(ExpandDuration(in))
			}(in)
			return value
		}(in["heartbeat_interval"]),
		Image: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["image"]),
		Backups: func(in interface{}) *kops.EtcdBackupSpec {
			value := func(in interface{}) *kops.EtcdBackupSpec {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in kops.EtcdBackupSpec) *kops.EtcdBackupSpec {
					return &in
				}(func(in interface{}) kops.EtcdBackupSpec {
					if in.([]interface{})[0] == nil {
						return kops.EtcdBackupSpec{}
					}
					return (ExpandEtcdBackupSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
			return value
		}(in["backups"]),
		Manager: func(in interface{}) *kops.EtcdManagerSpec {
			value := func(in interface{}) *kops.EtcdManagerSpec {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in kops.EtcdManagerSpec) *kops.EtcdManagerSpec {
					return &in
				}(func(in interface{}) kops.EtcdManagerSpec {
					if in.([]interface{})[0] == nil {
						return kops.EtcdManagerSpec{}
					}
					return (ExpandEtcdManagerSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
			return value
		}(in["manager"]),
		MemoryRequest: func(in interface{}) *resource.Quantity {
			value := func(in interface{}) *resource.Quantity {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in resource.Quantity) *resource.Quantity {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(ExpandQuantity(in))
			}(in)
			return value
		}(in["memory_request"]),
		CPURequest: func(in interface{}) *resource.Quantity {
			value := func(in interface{}) *resource.Quantity {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in resource.Quantity) *resource.Quantity {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(ExpandQuantity(in))
			}(in)
			return value
		}(in["cpu_request"]),
	}
}

func FlattenEtcdClusterSpec(in kops.EtcdClusterSpec) map[string]interface{} {
	return map[string]interface{}{
		"name": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.Name),
		"provider": func(in kops.EtcdProviderType) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.Provider),
		"members": func(in []*kops.EtcdMemberSpec) interface{} {
			value := func(in []*kops.EtcdMemberSpec) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, func(in *kops.EtcdMemberSpec) interface{} {
						if in == nil {
							return nil
						}
						return func(in kops.EtcdMemberSpec) interface{} {
							return func(in kops.EtcdMemberSpec) interface{} {
								return FlattenEtcdMemberSpec(in)
							}(in)
						}(*in)
					}(in))
				}
				return out
			}(in)
			return value
		}(in.Members),
		"enable_etcd_tls": func(in bool) interface{} {
			value := FlattenBool(bool(in))
			return value
		}(in.EnableEtcdTLS),
		"enable_tls_auth": func(in bool) interface{} {
			value := FlattenBool(bool(in))
			return value
		}(in.EnableTLSAuth),
		"version": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.Version),
		"leader_election_timeout": func(in *v1.Duration) interface{} {
			value := func(in *v1.Duration) interface{} {
				if in == nil {
					return nil
				}
				return func(in v1.Duration) interface{} {
					return FlattenDuration(in)
				}(*in)
			}(in)
			return value
		}(in.LeaderElectionTimeout),
		"heartbeat_interval": func(in *v1.Duration) interface{} {
			value := func(in *v1.Duration) interface{} {
				if in == nil {
					return nil
				}
				return func(in v1.Duration) interface{} {
					return FlattenDuration(in)
				}(*in)
			}(in)
			return value
		}(in.HeartbeatInterval),
		"image": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.Image),
		"backups": func(in *kops.EtcdBackupSpec) interface{} {
			value := func(in *kops.EtcdBackupSpec) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.EtcdBackupSpec) interface{} {
					return func(in kops.EtcdBackupSpec) []map[string]interface{} {
						return []map[string]interface{}{FlattenEtcdBackupSpec(in)}
					}(in)
				}(*in)
			}(in)
			return value
		}(in.Backups),
		"manager": func(in *kops.EtcdManagerSpec) interface{} {
			value := func(in *kops.EtcdManagerSpec) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.EtcdManagerSpec) interface{} {
					return func(in kops.EtcdManagerSpec) []map[string]interface{} {
						return []map[string]interface{}{FlattenEtcdManagerSpec(in)}
					}(in)
				}(*in)
			}(in)
			return value
		}(in.Manager),
		"memory_request": func(in *resource.Quantity) interface{} {
			value := func(in *resource.Quantity) interface{} {
				if in == nil {
					return nil
				}
				return func(in resource.Quantity) interface{} {
					return FlattenQuantity(in)
				}(*in)
			}(in)
			return value
		}(in.MemoryRequest),
		"cpu_request": func(in *resource.Quantity) interface{} {
			value := func(in *resource.Quantity) interface{} {
				if in == nil {
					return nil
				}
				return func(in resource.Quantity) interface{} {
					return FlattenQuantity(in)
				}(*in)
			}(in)
			return value
		}(in.CPURequest),
	}
}
