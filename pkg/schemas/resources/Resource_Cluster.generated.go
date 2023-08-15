package schemas

import (
	"context"

	"github.com/eddycharly/terraform-provider-kops/pkg/api/resources"
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	kopsv1alpha2schemas "github.com/eddycharly/terraform-provider-kops/pkg/schemas/kopsv1alpha2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

var _ = Schema

func ResourceCluster() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"channel":                           OptionalString(),
			"addons":                            OptionalList(kopsv1alpha2schemas.ResourceAddonSpec()),
			"config_base":                       OptionalComputedString(),
			"legacy_cloud_provider":             OptionalString(),
			"container_runtime":                 OptionalString(),
			"kubernetes_version":                OptionalString(),
			"subnets":                           OptionalList(kopsv1alpha2schemas.ResourceClusterSubnetSpec()),
			"project":                           OptionalString(),
			"master_public_name":                OptionalString(),
			"master_internal_name":              OptionalString(),
			"network_cidr":                      OptionalString(),
			"additional_network_cidrs":          OptionalList(String()),
			"network_id":                        OptionalString(),
			"topology":                          OptionalStruct(kopsv1alpha2schemas.ResourceTopologySpec()),
			"secret_store":                      OptionalString(),
			"key_store":                         OptionalString(),
			"config_store":                      OptionalString(),
			"dns_zone":                          OptionalString(),
			"additional_sans":                   OptionalList(String()),
			"cluster_dns_domain":                OptionalString(),
			"service_cluster_ip_range":          OptionalString(),
			"pod_cidr":                          OptionalString(),
			"non_masquerade_cidr":               OptionalString(),
			"ssh_access":                        OptionalList(String()),
			"node_port_access":                  OptionalList(String()),
			"egress_proxy":                      OptionalStruct(kopsv1alpha2schemas.ResourceEgressProxySpec()),
			"ssh_key_name":                      OptionalString(),
			"kubernetes_api_access":             OptionalList(String()),
			"isolate_masters":                   OptionalBool(),
			"update_policy":                     OptionalString(),
			"external_policies":                 OptionalComplexMap(List(String())),
			"additional_policies":               OptionalMap(String()),
			"file_assets":                       OptionalList(kopsv1alpha2schemas.ResourceFileAssetSpec()),
			"etcd_cluster":                      RequiredList(kopsv1alpha2schemas.ResourceEtcdClusterSpec()),
			"containerd":                        OptionalStruct(kopsv1alpha2schemas.ResourceContainerdConfig()),
			"docker":                            OptionalStruct(kopsv1alpha2schemas.ResourceDockerConfig()),
			"kube_dns":                          OptionalStruct(kopsv1alpha2schemas.ResourceKubeDNSConfig()),
			"kube_api_server":                   OptionalStruct(kopsv1alpha2schemas.ResourceKubeAPIServerConfig()),
			"kube_controller_manager":           OptionalStruct(kopsv1alpha2schemas.ResourceKubeControllerManagerConfig()),
			"external_cloud_controller_manager": OptionalStruct(kopsv1alpha2schemas.ResourceCloudControllerManagerConfig()),
			"kube_scheduler":                    OptionalStruct(kopsv1alpha2schemas.ResourceKubeSchedulerConfig()),
			"kube_proxy":                        OptionalStruct(kopsv1alpha2schemas.ResourceKubeProxyConfig()),
			"kubelet":                           OptionalStruct(kopsv1alpha2schemas.ResourceKubeletConfigSpec()),
			"control_plane_kubelet":             OptionalStruct(kopsv1alpha2schemas.ResourceKubeletConfigSpec()),
			"cloud_config":                      OptionalStruct(kopsv1alpha2schemas.ResourceCloudConfiguration()),
			"external_dns":                      OptionalStruct(kopsv1alpha2schemas.ResourceExternalDNSConfig()),
			"ntp":                               OptionalStruct(kopsv1alpha2schemas.ResourceNTPConfig()),
			"node_termination_handler":          OptionalStruct(kopsv1alpha2schemas.ResourceNodeTerminationHandlerSpec()),
			"node_problem_detector":             OptionalStruct(kopsv1alpha2schemas.ResourceNodeProblemDetectorConfig()),
			"metrics_server":                    OptionalStruct(kopsv1alpha2schemas.ResourceMetricsServerConfig()),
			"cert_manager":                      OptionalStruct(kopsv1alpha2schemas.ResourceCertManagerConfig()),
			"aws_load_balancer_controller":      OptionalStruct(kopsv1alpha2schemas.ResourceLoadBalancerControllerSpec()),
			"legacy_networking":                 OptionalStruct(kopsv1alpha2schemas.ResourceNetworkingSpec()),
			"networking":                        RequiredStruct(kopsv1alpha2schemas.ResourceNetworkingSpec()),
			"legacy_api":                        OptionalStruct(kopsv1alpha2schemas.ResourceAPISpec()),
			"api":                               OptionalStruct(kopsv1alpha2schemas.ResourceAPISpec()),
			"authentication":                    OptionalStruct(kopsv1alpha2schemas.ResourceAuthenticationSpec()),
			"authorization":                     OptionalStruct(kopsv1alpha2schemas.ResourceAuthorizationSpec()),
			"node_authorization":                OptionalStruct(kopsv1alpha2schemas.ResourceNodeAuthorizationSpec()),
			"cloud_labels":                      OptionalMap(String()),
			"hooks":                             OptionalList(kopsv1alpha2schemas.ResourceHookSpec()),
			"assets":                            OptionalStruct(kopsv1alpha2schemas.ResourceAssets()),
			"iam":                               OptionalComputedStruct(kopsv1alpha2schemas.ResourceIAMSpec()),
			"encryption_config":                 OptionalBool(),
			"tag_subnets":                       OptionalBool(),
			"use_host_certificates":             OptionalBool(),
			"sysctl_parameters":                 OptionalList(String()),
			"rolling_update":                    OptionalStruct(kopsv1alpha2schemas.ResourceRollingUpdate()),
			"cluster_autoscaler":                OptionalStruct(kopsv1alpha2schemas.ResourceClusterAutoscalerConfig()),
			"warm_pool":                         OptionalStruct(kopsv1alpha2schemas.ResourceWarmPoolSpec()),
			"service_account_issuer_discovery":  OptionalStruct(kopsv1alpha2schemas.ResourceServiceAccountIssuerDiscoveryConfig()),
			"snapshot_controller":               OptionalStruct(kopsv1alpha2schemas.ResourceSnapshotControllerConfig()),
			"karpenter":                         OptionalStruct(kopsv1alpha2schemas.ResourceKarpenterConfig()),
			"pod_identity_webhook":              OptionalStruct(kopsv1alpha2schemas.ResourcePodIdentityWebhookSpec()),
			"labels":                            OptionalMap(String()),
			"annotations":                       OptionalMap(String()),
			"name":                              ForceNew(RequiredString()),
			"admin_ssh_key":                     Sensitive(OptionalString()),
			"secrets":                           OptionalStruct(ResourceClusterSecrets()),
			"revision":                          ComputedInt(),
		},
	}
	res.SchemaVersion = 4
	res.StateUpgraders = []schema.StateUpgrader{
		{
			Type: res.CoreConfigSchema().ImpliedType(),
			Upgrade: func(ctx context.Context, rawState map[string]interface{}, meta interface{}) (map[string]interface{}, error) {
				ret := FlattenResourceCluster(ExpandResourceCluster(rawState))
				ret["id"] = rawState["id"]
				return ret, nil
			},
			Version: 0,
		}, {
			Type: res.CoreConfigSchema().ImpliedType(),
			Upgrade: func(ctx context.Context, rawState map[string]interface{}, meta interface{}) (map[string]interface{}, error) {
				ret := FlattenResourceCluster(ExpandResourceCluster(rawState))
				ret["id"] = rawState["id"]
				return ret, nil
			},
			Version: 1,
		}, {
			Type: res.CoreConfigSchema().ImpliedType(),
			Upgrade: func(ctx context.Context, rawState map[string]interface{}, meta interface{}) (map[string]interface{}, error) {
				ret := FlattenResourceCluster(ExpandResourceCluster(rawState))
				ret["id"] = rawState["id"]
				return ret, nil
			},
			Version: 2,
		}, {
			Type: res.CoreConfigSchema().ImpliedType(),
			Upgrade: func(ctx context.Context, rawState map[string]interface{}, meta interface{}) (map[string]interface{}, error) {
				ret := FlattenResourceCluster(ExpandResourceCluster(rawState))
				ret["id"] = rawState["id"]
				return ret, nil
			},
			Version: 3,
		},
	}
	return res
}

func ExpandResourceCluster(in map[string]interface{}) resources.Cluster {
	if in == nil {
		panic("expand Cluster failure, in is nil")
	}
	return resources.Cluster{
		ClusterSpec: func(in interface{}) kopsv1alpha2.ClusterSpec {
			return kopsv1alpha2schemas.ExpandResourceClusterSpec(in.(map[string]interface{}))
		}(in),
		Labels: func(in interface{}) map[string]string {
			return func(in interface{}) map[string]string {
				if in == nil {
					return nil
				}
				if in, ok := in.(map[string]interface{}); ok {
					if len(in) > 0 {
						out := map[string]string{}
						for key, in := range in {
							out[key] = string(ExpandString(in))
						}
						return out
					}
				}
				return nil
			}(in)
		}(in["labels"]),
		Annotations: func(in interface{}) map[string]string {
			return func(in interface{}) map[string]string {
				if in == nil {
					return nil
				}
				if in, ok := in.(map[string]interface{}); ok {
					if len(in) > 0 {
						out := map[string]string{}
						for key, in := range in {
							out[key] = string(ExpandString(in))
						}
						return out
					}
				}
				return nil
			}(in)
		}(in["annotations"]),
		Name: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["name"]),
		AdminSshKey: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["admin_ssh_key"]),
		Secrets: func(in interface{}) *resources.ClusterSecrets {
			return func(in interface{}) *resources.ClusterSecrets {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in resources.ClusterSecrets) *resources.ClusterSecrets {
					return &in
				}(func(in interface{}) resources.ClusterSecrets {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandResourceClusterSecrets(in[0].(map[string]interface{}))
					}
					return resources.ClusterSecrets{}
				}(in))
			}(in)
		}(in["secrets"]),
		Revision: func(in interface{}) int {
			return int(ExpandInt(in))
		}(in["revision"]),
	}
}

func FlattenResourceClusterInto(in resources.Cluster, out map[string]interface{}) {
	kopsv1alpha2schemas.FlattenResourceClusterSpecInto(in.ClusterSpec, out)
	out["labels"] = func(in map[string]string) interface{} {
		return func(in map[string]string) map[string]interface{} {
			if in == nil {
				return nil
			}
			out := map[string]interface{}{}
			for key, in := range in {
				out[key] = FlattenString(string(in))
			}
			return out
		}(in)
	}(in.Labels)
	out["annotations"] = func(in map[string]string) interface{} {
		return func(in map[string]string) map[string]interface{} {
			if in == nil {
				return nil
			}
			out := map[string]interface{}{}
			for key, in := range in {
				out[key] = FlattenString(string(in))
			}
			return out
		}(in)
	}(in.Annotations)
	out["name"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Name)
	out["admin_ssh_key"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.AdminSshKey)
	out["secrets"] = func(in *resources.ClusterSecrets) interface{} {
		return func(in *resources.ClusterSecrets) interface{} {
			if in == nil {
				return nil
			}
			return func(in resources.ClusterSecrets) interface{} {
				return func(in resources.ClusterSecrets) []interface{} {
					return []interface{}{FlattenResourceClusterSecrets(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Secrets)
	out["revision"] = func(in int) interface{} {
		return FlattenInt(int(in))
	}(in.Revision)
}

func FlattenResourceCluster(in resources.Cluster) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceClusterInto(in, out)
	return out
}
