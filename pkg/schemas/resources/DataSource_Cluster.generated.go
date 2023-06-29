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

func DataSourceCluster() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"channel":                           ComputedString(),
			"addons":                            ComputedList(kopsv1alpha2schemas.DataSourceAddonSpec()),
			"config_base":                       ComputedString(),
			"legacy_cloud_provider":             ComputedString(),
			"container_runtime":                 ComputedString(),
			"kubernetes_version":                ComputedString(),
			"subnets":                           ComputedList(kopsv1alpha2schemas.DataSourceClusterSubnetSpec()),
			"project":                           ComputedString(),
			"master_public_name":                ComputedString(),
			"master_internal_name":              ComputedString(),
			"network_cidr":                      ComputedString(),
			"additional_network_cidrs":          ComputedList(String()),
			"network_id":                        ComputedString(),
			"topology":                          ComputedStruct(kopsv1alpha2schemas.DataSourceTopologySpec()),
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
			"egress_proxy":                      ComputedStruct(kopsv1alpha2schemas.DataSourceEgressProxySpec()),
			"ssh_key_name":                      ComputedString(),
			"kubernetes_api_access":             ComputedList(String()),
			"isolate_masters":                   ComputedBool(),
			"update_policy":                     ComputedString(),
			"external_policies":                 ComputedComplexMap(List(String())),
			"additional_policies":               ComputedMap(String()),
			"file_assets":                       ComputedList(kopsv1alpha2schemas.DataSourceFileAssetSpec()),
			"etcd_cluster":                      ComputedList(kopsv1alpha2schemas.DataSourceEtcdClusterSpec()),
			"containerd":                        ComputedStruct(kopsv1alpha2schemas.DataSourceContainerdConfig()),
			"docker":                            ComputedStruct(kopsv1alpha2schemas.DataSourceDockerConfig()),
			"kube_dns":                          ComputedStruct(kopsv1alpha2schemas.DataSourceKubeDNSConfig()),
			"kube_api_server":                   ComputedStruct(kopsv1alpha2schemas.DataSourceKubeAPIServerConfig()),
			"kube_controller_manager":           ComputedStruct(kopsv1alpha2schemas.DataSourceKubeControllerManagerConfig()),
			"external_cloud_controller_manager": ComputedStruct(kopsv1alpha2schemas.DataSourceCloudControllerManagerConfig()),
			"kube_scheduler":                    ComputedStruct(kopsv1alpha2schemas.DataSourceKubeSchedulerConfig()),
			"kube_proxy":                        ComputedStruct(kopsv1alpha2schemas.DataSourceKubeProxyConfig()),
			"kubelet":                           ComputedStruct(kopsv1alpha2schemas.DataSourceKubeletConfigSpec()),
			"control_plane_kubelet":             ComputedStruct(kopsv1alpha2schemas.DataSourceKubeletConfigSpec()),
			"cloud_config":                      ComputedStruct(kopsv1alpha2schemas.DataSourceCloudConfiguration()),
			"external_dns":                      ComputedStruct(kopsv1alpha2schemas.DataSourceExternalDNSConfig()),
			"ntp":                               ComputedStruct(kopsv1alpha2schemas.DataSourceNTPConfig()),
			"node_termination_handler":          ComputedStruct(kopsv1alpha2schemas.DataSourceNodeTerminationHandlerSpec()),
			"node_problem_detector":             ComputedStruct(kopsv1alpha2schemas.DataSourceNodeProblemDetectorConfig()),
			"metrics_server":                    ComputedStruct(kopsv1alpha2schemas.DataSourceMetricsServerConfig()),
			"cert_manager":                      ComputedStruct(kopsv1alpha2schemas.DataSourceCertManagerConfig()),
			"aws_load_balancer_controller":      ComputedStruct(kopsv1alpha2schemas.DataSourceLoadBalancerControllerSpec()),
			"legacy_networking":                 ComputedStruct(kopsv1alpha2schemas.DataSourceNetworkingSpec()),
			"networking":                        ComputedStruct(kopsv1alpha2schemas.DataSourceNetworkingSpec()),
			"legacy_api":                        ComputedStruct(kopsv1alpha2schemas.DataSourceAPISpec()),
			"api":                               ComputedStruct(kopsv1alpha2schemas.DataSourceAPISpec()),
			"authentication":                    ComputedStruct(kopsv1alpha2schemas.DataSourceAuthenticationSpec()),
			"authorization":                     ComputedStruct(kopsv1alpha2schemas.DataSourceAuthorizationSpec()),
			"node_authorization":                ComputedStruct(kopsv1alpha2schemas.DataSourceNodeAuthorizationSpec()),
			"cloud_labels":                      ComputedMap(String()),
			"hooks":                             ComputedList(kopsv1alpha2schemas.DataSourceHookSpec()),
			"assets":                            ComputedStruct(kopsv1alpha2schemas.DataSourceAssets()),
			"iam":                               ComputedStruct(kopsv1alpha2schemas.DataSourceIAMSpec()),
			"encryption_config":                 ComputedBool(),
			"tag_subnets":                       ComputedBool(),
			"use_host_certificates":             ComputedBool(),
			"sysctl_parameters":                 ComputedList(String()),
			"rolling_update":                    ComputedStruct(kopsv1alpha2schemas.DataSourceRollingUpdate()),
			"cluster_autoscaler":                ComputedStruct(kopsv1alpha2schemas.DataSourceClusterAutoscalerConfig()),
			"warm_pool":                         ComputedStruct(kopsv1alpha2schemas.DataSourceWarmPoolSpec()),
			"service_account_issuer_discovery":  ComputedStruct(kopsv1alpha2schemas.DataSourceServiceAccountIssuerDiscoveryConfig()),
			"snapshot_controller":               ComputedStruct(kopsv1alpha2schemas.DataSourceSnapshotControllerConfig()),
			"karpenter":                         ComputedStruct(kopsv1alpha2schemas.DataSourceKarpenterConfig()),
			"pod_identity_webhook":              ComputedStruct(kopsv1alpha2schemas.DataSourcePodIdentityWebhookSpec()),
			"labels":                            ComputedMap(String()),
			"annotations":                       ComputedMap(String()),
			"name":                              RequiredString(),
			"admin_ssh_key":                     ComputedString(),
			"secrets":                           ComputedStruct(DataSourceClusterSecrets()),
		},
	}
	res.SchemaVersion = 4
	res.StateUpgraders = []schema.StateUpgrader{
		{
			Type: res.CoreConfigSchema().ImpliedType(),
			Upgrade: func(ctx context.Context, rawState map[string]interface{}, meta interface{}) (map[string]interface{}, error) {
				ret := FlattenDataSourceCluster(ExpandDataSourceCluster(rawState))
				ret["id"] = rawState["id"]
				return ret, nil
			},
			Version: 0,
		}, {
			Type: res.CoreConfigSchema().ImpliedType(),
			Upgrade: func(ctx context.Context, rawState map[string]interface{}, meta interface{}) (map[string]interface{}, error) {
				ret := FlattenDataSourceCluster(ExpandDataSourceCluster(rawState))
				ret["id"] = rawState["id"]
				return ret, nil
			},
			Version: 1,
		}, {
			Type: res.CoreConfigSchema().ImpliedType(),
			Upgrade: func(ctx context.Context, rawState map[string]interface{}, meta interface{}) (map[string]interface{}, error) {
				ret := FlattenDataSourceCluster(ExpandDataSourceCluster(rawState))
				ret["id"] = rawState["id"]
				return ret, nil
			},
			Version: 2,
		}, {
			Type: res.CoreConfigSchema().ImpliedType(),
			Upgrade: func(ctx context.Context, rawState map[string]interface{}, meta interface{}) (map[string]interface{}, error) {
				ret := FlattenDataSourceCluster(ExpandDataSourceCluster(rawState))
				ret["id"] = rawState["id"]
				return ret, nil
			},
			Version: 3,
		},
	}
	return res
}

func ExpandDataSourceCluster(in map[string]interface{}) resources.Cluster {
	if in == nil {
		panic("expand Cluster failure, in is nil")
	}
	return resources.Cluster{
		ClusterSpec: func(in interface{}) kopsv1alpha2.ClusterSpec {
			return kopsv1alpha2schemas.ExpandDataSourceClusterSpec(in.(map[string]interface{}))
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
						return ExpandDataSourceClusterSecrets(in[0].(map[string]interface{}))
					}
					return resources.ClusterSecrets{}
				}(in))
			}(in)
		}(in["secrets"]),
	}
}

func FlattenDataSourceClusterInto(in resources.Cluster, out map[string]interface{}) {
	kopsv1alpha2schemas.FlattenDataSourceClusterSpecInto(in.ClusterSpec, out)
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
					return []interface{}{FlattenDataSourceClusterSecrets(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Secrets)
}

func FlattenDataSourceCluster(in resources.Cluster) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceClusterInto(in, out)
	return out
}
