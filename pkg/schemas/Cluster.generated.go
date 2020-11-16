package schemas

import (
	"github.com/eddycharly/terraform-provider-kops/pkg/api"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

func Cluster() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name":                              RequiredString(),
			"admin_ssh_key":                     Sensitive(RequiredString()),
			"channel":                           OptionalString(),
			"addons":                            OptionalList(AddonSpec()),
			"config_base":                       OptionalComputedString(),
			"cloud_provider":                    RequiredString(),
			"container_runtime":                 OptionalString(),
			"kubernetes_version":                OptionalString(),
			"subnet":                            RequiredList(ClusterSubnetSpec()),
			"project":                           OptionalString(),
			"master_public_name":                OptionalComputedString(),
			"master_internal_name":              OptionalComputedString(),
			"network_cidr":                      OptionalComputedString(),
			"additional_network_cidrs":          OptionalList(String()),
			"network_id":                        RequiredString(),
			"topology":                          RequiredStruct(TopologySpec()),
			"secret_store":                      OptionalString(),
			"key_store":                         OptionalString(),
			"config_store":                      OptionalString(),
			"dns_zone":                          OptionalString(),
			"additional_sans":                   OptionalList(String()),
			"cluster_dns_domain":                OptionalString(),
			"service_cluster_ip_range":          OptionalString(),
			"pod_cidr":                          OptionalString(),
			"non_masquerade_cidr":               OptionalComputedString(),
			"ssh_access":                        OptionalList(String()),
			"node_port_access":                  OptionalList(String()),
			"egress_proxy":                      OptionalStruct(EgressProxySpec()),
			"ssh_key_name":                      OptionalString(),
			"kubernetes_api_access":             OptionalList(String()),
			"isolate_masters":                   OptionalBool(),
			"update_policy":                     OptionalString(),
			"external_policies":                 OptionalMap(List(String())),
			"additional_policies":               OptionalMap(String()),
			"file_assets":                       OptionalList(FileAssetSpec()),
			"etcd_cluster":                      RequiredList(EtcdClusterSpec()),
			"containerd":                        OptionalStruct(ContainerdConfig()),
			"docker":                            OptionalStruct(DockerConfig()),
			"kube_dns":                          OptionalStruct(KubeDNSConfig()),
			"kube_api_server":                   OptionalStruct(KubeAPIServerConfig()),
			"kube_controller_manager":           OptionalStruct(KubeControllerManagerConfig()),
			"external_cloud_controller_manager": OptionalStruct(CloudControllerManagerConfig()),
			"kube_scheduler":                    OptionalStruct(KubeSchedulerConfig()),
			"kube_proxy":                        OptionalStruct(KubeProxyConfig()),
			"kubelet":                           OptionalStruct(KubeletConfigSpec()),
			"master_kubelet":                    OptionalStruct(KubeletConfigSpec()),
			"cloud_config":                      OptionalStruct(CloudConfiguration()),
			"external_dns":                      OptionalStruct(ExternalDNSConfig()),
			"networking":                        RequiredStruct(NetworkingSpec()),
			"api":                               OptionalStruct(AccessSpec()),
			"authentication":                    OptionalStruct(AuthenticationSpec()),
			"authorization":                     OptionalStruct(AuthorizationSpec()),
			"node_authorization":                OptionalStruct(NodeAuthorizationSpec()),
			"cloud_labels":                      OptionalMap(String()),
			"hooks":                             OptionalList(HookSpec()),
			"assets":                            OptionalStruct(Assets()),
			"iam":                               OptionalComputedStruct(IAMSpec()),
			"encryption_config":                 OptionalBool(),
			"disable_subnet_tags":               OptionalBool(),
			"use_host_certificates":             OptionalBool(),
			"sysctl_parameters":                 OptionalList(String()),
			"rolling_update":                    OptionalStruct(RollingUpdate()),
			"instance_group":                    RequiredList(InstanceGroup()),
			"kube_config":                       Sensitive(ComputedStruct(KubeConfig())),
		},
	}
}

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
