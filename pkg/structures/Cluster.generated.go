package structures

import (
	"github.com/eddycharly/terraform-provider-kops/pkg/api"
	"k8s.io/kops/pkg/apis/kops"
)

func ExpandCluster(in map[string]interface{}) api.Cluster {
	if in == nil {
		panic("expand Cluster failure, in is nil")
	}
	return api.Cluster{
		Name: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["name"]),
		AdminSshKey: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["admin_ssh_key"]),
		ClusterSpec: func(in interface{}) kops.ClusterSpec {
			return func(in interface{}) kops.ClusterSpec {
				if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
					return kops.ClusterSpec{}
				}
				return (ExpandClusterSpec(in.([]interface{})[0].(map[string]interface{})))
			}(in)
		}(in),
		InstanceGroup: func(in interface{}) []*api.InstanceGroup {
			return func(in interface{}) []*api.InstanceGroup {
				var out []*api.InstanceGroup
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) *api.InstanceGroup {
						if in == nil {
							return nil
						}
						if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
							return nil
						}
						return func(in api.InstanceGroup) *api.InstanceGroup {
							return &in
						}(func(in interface{}) api.InstanceGroup {
							if in == nil {
								return api.InstanceGroup{}
							}
							return (ExpandInstanceGroup(in.(map[string]interface{})))
						}(in))
					}(in))
				}
				return out
			}(in)
		}(in["instance_group"]),
		KubeConfig: func(in interface{}) *api.KubeConfig {
			return func(in interface{}) *api.KubeConfig {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in api.KubeConfig) *api.KubeConfig {
					return &in
				}(func(in interface{}) api.KubeConfig {
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return api.KubeConfig{}
					}
					return (ExpandKubeConfig(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["kube_config"]),
	}
}

func FlattenClusterInto(in api.Cluster, out map[string]interface{}) {
	out["name"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Name)
	out["admin_ssh_key"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.AdminSshKey)
	FlattenClusterSpecInto(in.ClusterSpec, out)
	out["instance_group"] = func(in []*api.InstanceGroup) interface{} {
		return func(in []*api.InstanceGroup) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, func(in *api.InstanceGroup) interface{} {
					if in == nil {
						return nil
					}
					return func(in api.InstanceGroup) interface{} {
						return func(in api.InstanceGroup) interface{} {
							return FlattenInstanceGroup(in)
						}(in)
					}(*in)
				}(in))
			}
			return out
		}(in)
	}(in.InstanceGroup)
	out["kube_config"] = func(in *api.KubeConfig) interface{} {
		return func(in *api.KubeConfig) interface{} {
			if in == nil {
				return nil
			}
			return func(in api.KubeConfig) interface{} {
				return func(in api.KubeConfig) []map[string]interface{} {
					return []map[string]interface{}{FlattenKubeConfig(in)}
				}(in)
			}(*in)
		}(in)
	}(in.KubeConfig)
}

func FlattenCluster(in api.Cluster) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenClusterInto(in, out)
	return out
}
