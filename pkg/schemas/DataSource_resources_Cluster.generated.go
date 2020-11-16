package schemas

import (
	"github.com/eddycharly/terraform-provider-kops/pkg/api/kube"
	"github.com/eddycharly/terraform-provider-kops/pkg/api/resources"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

func DataSourceResourcesCluster() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name":                              RequiredString(),
			"admin_ssh_key":                     ComputedString(),
			"channel":                           ComputedString(),
			"addons":                            ComputedList(DataSourceKopsAddonSpec()),
			"config_base":                       ComputedString(),
			"cloud_provider":                    ComputedString(),
			"container_runtime":                 ComputedString(),
			"kubernetes_version":                ComputedString(),
			"subnets":                           ComputedList(DataSourceKopsClusterSubnetSpec()),
			"project":                           ComputedString(),
			"master_public_name":                ComputedString(),
			"master_internal_name":              ComputedString(),
			"network_cidr":                      ComputedString(),
			"additional_network_cidrs":          ComputedList(String()),
			"network_id":                        ComputedString(),
			"topology":                          ComputedStruct(DataSourceKopsTopologySpec()),
			"secret_store":                      ComputedString(),
			"key_store":                         ComputedString(),
			"config_store":                      ComputedString(),
			"dns_zone":                          ComputedString(),
			"additional_sans":                   ComputedList(String()),
			"cluster_dns_domain":                ComputedString(),
			"service_cluster_ip_range":          ComputedString(),
			"pod_cidr":                          ComputedString(),
			"non_masquerade_cidr":               ComputedString(),
			"ssh_access":                        ComputedList(String()),
			"node_port_access":                  ComputedList(String()),
			"egress_proxy":                      ComputedStruct(DataSourceKopsEgressProxySpec()),
			"ssh_key_name":                      ComputedString(),
			"kubernetes_api_access":             ComputedList(String()),
			"isolate_masters":                   ComputedBool(),
			"update_policy":                     ComputedString(),
			"external_policies":                 ComputedMap(List(String())),
			"additional_policies":               ComputedMap(String()),
			"file_assets":                       ComputedList(DataSourceKopsFileAssetSpec()),
			"etcd_clusters":                     ComputedList(DataSourceKopsEtcdClusterSpec()),
			"containerd":                        ComputedStruct(DataSourceKopsContainerdConfig()),
			"docker":                            ComputedStruct(DataSourceKopsDockerConfig()),
			"kube_dns":                          ComputedStruct(DataSourceKopsKubeDNSConfig()),
			"kube_api_server":                   ComputedStruct(DataSourceKopsKubeAPIServerConfig()),
			"kube_controller_manager":           ComputedStruct(DataSourceKopsKubeControllerManagerConfig()),
			"external_cloud_controller_manager": ComputedStruct(DataSourceKopsCloudControllerManagerConfig()),
			"kube_scheduler":                    ComputedStruct(DataSourceKopsKubeSchedulerConfig()),
			"kube_proxy":                        ComputedStruct(DataSourceKopsKubeProxyConfig()),
			"kubelet":                           ComputedStruct(DataSourceKopsKubeletConfigSpec()),
			"master_kubelet":                    ComputedStruct(DataSourceKopsKubeletConfigSpec()),
			"cloud_config":                      ComputedStruct(DataSourceKopsCloudConfiguration()),
			"external_dns":                      ComputedStruct(DataSourceKopsExternalDNSConfig()),
			"networking":                        ComputedStruct(DataSourceKopsNetworkingSpec()),
			"api":                               ComputedStruct(DataSourceKopsAccessSpec()),
			"authentication":                    ComputedStruct(DataSourceKopsAuthenticationSpec()),
			"authorization":                     ComputedStruct(DataSourceKopsAuthorizationSpec()),
			"node_authorization":                ComputedStruct(DataSourceKopsNodeAuthorizationSpec()),
			"cloud_labels":                      ComputedMap(String()),
			"hooks":                             ComputedList(DataSourceKopsHookSpec()),
			"assets":                            ComputedStruct(DataSourceKopsAssets()),
			"iam":                               ComputedStruct(DataSourceKopsIAMSpec()),
			"encryption_config":                 ComputedBool(),
			"disable_subnet_tags":               ComputedBool(),
			"use_host_certificates":             ComputedBool(),
			"sysctl_parameters":                 ComputedList(String()),
			"rolling_update":                    ComputedStruct(DataSourceKopsRollingUpdate()),
			"instance_group":                    ComputedList(DataSourceResourcesInstanceGroup()),
			"kube_config":                       ComputedStruct(DataSourceKubeConfig()),
		},
	}
}

func ExpandDataSourceResourcesCluster(in map[string]interface{}) resources.Cluster {
	if in == nil {
		panic("expand Cluster failure, in is nil")
	}
	return resources.Cluster{
		Name: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["name"]),
		AdminSshKey: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["admin_ssh_key"]),
		ClusterSpec: func(in interface{}) kops.ClusterSpec {
			return ExpandDataSourceKopsClusterSpec(in.(map[string]interface{}))
		}(in),
		InstanceGroup: func(in interface{}) []*resources.InstanceGroup {
			return func(in interface{}) []*resources.InstanceGroup {
				var out []*resources.InstanceGroup
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) *resources.InstanceGroup {
						if in == nil {
							return nil
						}
						if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
							return nil
						}
						return func(in resources.InstanceGroup) *resources.InstanceGroup {
							return &in
						}(func(in interface{}) resources.InstanceGroup {
							if in == nil {
								return resources.InstanceGroup{}
							}
							return (ExpandDataSourceResourcesInstanceGroup(in.(map[string]interface{})))
						}(in))
					}(in))
				}
				return out
			}(in)
		}(in["instance_group"]),
		KubeConfig: func(in interface{}) *kube.Config {
			return func(in interface{}) *kube.Config {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kube.Config) *kube.Config {
					return &in
				}(func(in interface{}) kube.Config {
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return kube.Config{}
					}
					return (ExpandDataSourceKubeConfig(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["kube_config"]),
	}
}

func FlattenDataSourceResourcesClusterInto(in resources.Cluster, out map[string]interface{}) {
	out["name"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Name)
	out["admin_ssh_key"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.AdminSshKey)
	FlattenDataSourceKopsClusterSpecInto(in.ClusterSpec, out)
	out["instance_group"] = func(in []*resources.InstanceGroup) interface{} {
		return func(in []*resources.InstanceGroup) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, func(in *resources.InstanceGroup) interface{} {
					if in == nil {
						return nil
					}
					return func(in resources.InstanceGroup) interface{} {
						return func(in resources.InstanceGroup) interface{} {
							return FlattenDataSourceResourcesInstanceGroup(in)
						}(in)
					}(*in)
				}(in))
			}
			return out
		}(in)
	}(in.InstanceGroup)
	out["kube_config"] = func(in *kube.Config) interface{} {
		return func(in *kube.Config) interface{} {
			if in == nil {
				return nil
			}
			return func(in kube.Config) interface{} {
				return func(in kube.Config) []map[string]interface{} {
					return []map[string]interface{}{FlattenDataSourceKubeConfig(in)}
				}(in)
			}(*in)
		}(in)
	}(in.KubeConfig)
}

func FlattenDataSourceResourcesCluster(in resources.Cluster) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceResourcesClusterInto(in, out)
	return out
}
