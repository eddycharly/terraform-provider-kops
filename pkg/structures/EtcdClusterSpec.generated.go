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
			return string(ExpandString(in))
		}(in["name"]),
		Provider: func(in interface{}) kops.EtcdProviderType {
			return kops.EtcdProviderType(ExpandString(in))
		}(in["provider"]),
		Members: func(in interface{}) []*kops.EtcdMemberSpec {
			return func(in interface{}) []*kops.EtcdMemberSpec {
				var out []*kops.EtcdMemberSpec
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) *kops.EtcdMemberSpec {
						if in == nil {
							return nil
						}
						if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
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
		}(in["members"]),
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
					return (ExpandEtcdBackupSpec(in.([]interface{})[0].(map[string]interface{})))
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
					return (ExpandEtcdManagerSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["manager"]),
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
	}
}

func FlattenEtcdClusterSpec(in kops.EtcdClusterSpec) map[string]interface{} {
	return map[string]interface{}{
		"name": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.Name),
		"provider": func(in kops.EtcdProviderType) interface{} {
			return FlattenString(string(in))
		}(in.Provider),
		"members": func(in []*kops.EtcdMemberSpec) interface{} {
			return func(in []*kops.EtcdMemberSpec) []interface{} {
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
		}(in.Members),
		"enable_etcd_tls": func(in bool) interface{} {
			return FlattenBool(bool(in))
		}(in.EnableEtcdTLS),
		"enable_tls_auth": func(in bool) interface{} {
			return FlattenBool(bool(in))
		}(in.EnableTLSAuth),
		"version": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.Version),
		"leader_election_timeout": func(in *v1.Duration) interface{} {
			return func(in *v1.Duration) interface{} {
				if in == nil {
					return nil
				}
				return func(in v1.Duration) interface{} {
					return FlattenDuration(in)
				}(*in)
			}(in)
		}(in.LeaderElectionTimeout),
		"heartbeat_interval": func(in *v1.Duration) interface{} {
			return func(in *v1.Duration) interface{} {
				if in == nil {
					return nil
				}
				return func(in v1.Duration) interface{} {
					return FlattenDuration(in)
				}(*in)
			}(in)
		}(in.HeartbeatInterval),
		"image": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.Image),
		"backups": func(in *kops.EtcdBackupSpec) interface{} {
			return func(in *kops.EtcdBackupSpec) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.EtcdBackupSpec) interface{} {
					return func(in kops.EtcdBackupSpec) []map[string]interface{} {
						return []map[string]interface{}{FlattenEtcdBackupSpec(in)}
					}(in)
				}(*in)
			}(in)
		}(in.Backups),
		"manager": func(in *kops.EtcdManagerSpec) interface{} {
			return func(in *kops.EtcdManagerSpec) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.EtcdManagerSpec) interface{} {
					return func(in kops.EtcdManagerSpec) []map[string]interface{} {
						return []map[string]interface{}{FlattenEtcdManagerSpec(in)}
					}(in)
				}(*in)
			}(in)
		}(in.Manager),
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
	}
}
