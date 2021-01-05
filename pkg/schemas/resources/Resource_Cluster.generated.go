package schemas

import (
	"github.com/eddycharly/terraform-provider-kops/pkg/api/kube"
	"github.com/eddycharly/terraform-provider-kops/pkg/api/resources"
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	kopsschemas "github.com/eddycharly/terraform-provider-kops/pkg/schemas/kops"
	kubeschemas "github.com/eddycharly/terraform-provider-kops/pkg/schemas/kube"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceCluster() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name":                              RequiredString(),
			"admin_ssh_key":                     Sensitive(RequiredString()),
			"channel":                           OptionalString(),
			"addons":                            OptionalList(kopsschemas.ResourceAddonSpec()),
			"config_base":                       OptionalComputedString(),
			"cloud_provider":                    RequiredString(),
			"container_runtime":                 OptionalString(),
			"kubernetes_version":                OptionalString(),
			"subnet":                            RequiredList(kopsschemas.ResourceClusterSubnetSpec()),
			"project":                           OptionalString(),
			"master_public_name":                OptionalComputedString(),
			"master_internal_name":              OptionalComputedString(),
			"network_cidr":                      OptionalComputedString(),
			"additional_network_cidrs":          OptionalList(String()),
			"network_id":                        RequiredString(),
			"topology":                          RequiredStruct(kopsschemas.ResourceTopologySpec()),
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
			"egress_proxy":                      OptionalStruct(kopsschemas.ResourceEgressProxySpec()),
			"ssh_key_name":                      OptionalString(),
			"kubernetes_api_access":             OptionalList(String()),
			"isolate_masters":                   OptionalBool(),
			"update_policy":                     OptionalString(),
			"external_policies":                 OptionalMap(List(String())),
			"additional_policies":               OptionalMap(String()),
			"file_assets":                       OptionalList(kopsschemas.ResourceFileAssetSpec()),
			"etcd_cluster":                      RequiredList(kopsschemas.ResourceEtcdClusterSpec()),
			"containerd":                        OptionalStruct(kopsschemas.ResourceContainerdConfig()),
			"docker":                            OptionalStruct(kopsschemas.ResourceDockerConfig()),
			"kube_dns":                          OptionalStruct(kopsschemas.ResourceKubeDNSConfig()),
			"kube_api_server":                   OptionalStruct(kopsschemas.ResourceKubeAPIServerConfig()),
			"kube_controller_manager":           OptionalStruct(kopsschemas.ResourceKubeControllerManagerConfig()),
			"external_cloud_controller_manager": OptionalStruct(kopsschemas.ResourceCloudControllerManagerConfig()),
			"kube_scheduler":                    OptionalStruct(kopsschemas.ResourceKubeSchedulerConfig()),
			"kube_proxy":                        OptionalStruct(kopsschemas.ResourceKubeProxyConfig()),
			"kubelet":                           OptionalStruct(kopsschemas.ResourceKubeletConfigSpec()),
			"master_kubelet":                    OptionalStruct(kopsschemas.ResourceKubeletConfigSpec()),
			"cloud_config":                      OptionalStruct(kopsschemas.ResourceCloudConfiguration()),
			"external_dns":                      OptionalStruct(kopsschemas.ResourceExternalDNSConfig()),
			"networking":                        RequiredStruct(kopsschemas.ResourceNetworkingSpec()),
			"api":                               OptionalStruct(kopsschemas.ResourceAccessSpec()),
			"authentication":                    OptionalStruct(kopsschemas.ResourceAuthenticationSpec()),
			"authorization":                     OptionalStruct(kopsschemas.ResourceAuthorizationSpec()),
			"node_authorization":                OptionalStruct(kopsschemas.ResourceNodeAuthorizationSpec()),
			"cloud_labels":                      OptionalMap(String()),
			"hooks":                             OptionalList(kopsschemas.ResourceHookSpec()),
			"assets":                            OptionalStruct(kopsschemas.ResourceAssets()),
			"iam":                               OptionalComputedStruct(kopsschemas.ResourceIAMSpec()),
			"encryption_config":                 OptionalBool(),
			"disable_subnet_tags":               OptionalBool(),
			"use_host_certificates":             OptionalBool(),
			"sysctl_parameters":                 OptionalList(String()),
			"rolling_update":                    OptionalStruct(kopsschemas.ResourceRollingUpdate()),
			"instance_group":                    RequiredSetList(ResourceInstanceGroup()),
			"kube_config":                       Sensitive(ComputedStruct(kubeschemas.ResourceConfig())),
		},
	}
}

func ExpandResourceCluster(in map[string]interface{}) resources.Cluster {
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
			return kopsschemas.ExpandResourceClusterSpec(in.(map[string]interface{}))
		}(in),
		InstanceGroup: func(in interface{}) []*resources.InstanceGroup {
			return func(in interface{}) []*resources.InstanceGroup {
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
								return (ExpandResourceInstanceGroup(in.(map[string]interface{})))
							}(in))
						}(in))
					}
					return out
				}(in)
			}(in.(*schema.Set).List())
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
					return (kubeschemas.ExpandResourceConfig(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["kube_config"]),
	}
}

func FlattenResourceClusterInto(in resources.Cluster, out map[string]interface{}) {
	out["name"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Name)
	out["admin_ssh_key"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.AdminSshKey)
	kopsschemas.FlattenResourceClusterSpecInto(in.ClusterSpec, out)
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
							return FlattenResourceInstanceGroup(in)
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
					return []map[string]interface{}{kubeschemas.FlattenResourceConfig(in)}
				}(in)
			}(*in)
		}(in)
	}(in.KubeConfig)
}

func FlattenResourceCluster(in resources.Cluster) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceClusterInto(in, out)
	return out
}
