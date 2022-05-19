package schemas

import (
	"context"

	"github.com/eddycharly/terraform-provider-kops/pkg/api/resources"
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	kopsschemas "github.com/eddycharly/terraform-provider-kops/pkg/schemas/kops"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceCluster() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"channel":                           ComputedString(),
			"addons":                            ComputedList(kopsschemas.DataSourceAddonSpec()),
			"config_base":                       ComputedString(),
			"cloud_provider":                    ComputedString(),
			"container_runtime":                 ComputedString(),
			"kubernetes_version":                ComputedString(),
			"subnet":                            ComputedList(kopsschemas.DataSourceClusterSubnetSpec()),
			"project":                           ComputedString(),
			"master_public_name":                ComputedString(),
			"master_internal_name":              ComputedString(),
			"network_cidr":                      ComputedString(),
			"additional_network_cidrs":          ComputedList(String()),
			"network_id":                        ComputedString(),
			"topology":                          ComputedStruct(kopsschemas.DataSourceTopologySpec()),
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
			"egress_proxy":                      ComputedStruct(kopsschemas.DataSourceEgressProxySpec()),
			"ssh_key_name":                      ComputedString(),
			"kubernetes_api_access":             ComputedList(String()),
			"isolate_masters":                   ComputedBool(),
			"update_policy":                     ComputedString(),
			"external_policies":                 ComputedComplexMap(List(String())),
			"additional_policies":               ComputedMap(String()),
			"file_assets":                       ComputedList(kopsschemas.DataSourceFileAssetSpec()),
			"etcd_cluster":                      ComputedList(kopsschemas.DataSourceEtcdClusterSpec()),
			"containerd":                        ComputedStruct(kopsschemas.DataSourceContainerdConfig()),
			"docker":                            ComputedStruct(kopsschemas.DataSourceDockerConfig()),
			"kube_dns":                          ComputedStruct(kopsschemas.DataSourceKubeDNSConfig()),
			"kube_api_server":                   ComputedStruct(kopsschemas.DataSourceKubeAPIServerConfig()),
			"kube_controller_manager":           ComputedStruct(kopsschemas.DataSourceKubeControllerManagerConfig()),
			"external_cloud_controller_manager": ComputedStruct(kopsschemas.DataSourceCloudControllerManagerConfig()),
			"kube_scheduler":                    ComputedStruct(kopsschemas.DataSourceKubeSchedulerConfig()),
			"kube_proxy":                        ComputedStruct(kopsschemas.DataSourceKubeProxyConfig()),
			"kubelet":                           ComputedStruct(kopsschemas.DataSourceKubeletConfigSpec()),
			"master_kubelet":                    ComputedStruct(kopsschemas.DataSourceKubeletConfigSpec()),
			"cloud_config":                      ComputedStruct(kopsschemas.DataSourceCloudConfiguration()),
			"external_dns":                      ComputedStruct(kopsschemas.DataSourceExternalDNSConfig()),
			"ntp":                               ComputedStruct(kopsschemas.DataSourceNTPConfig()),
			"node_termination_handler":          ComputedStruct(kopsschemas.DataSourceNodeTerminationHandlerConfig()),
			"node_problem_detector":             ComputedStruct(kopsschemas.DataSourceNodeProblemDetectorConfig()),
			"metrics_server":                    ComputedStruct(kopsschemas.DataSourceMetricsServerConfig()),
			"cert_manager":                      ComputedStruct(kopsschemas.DataSourceCertManagerConfig()),
			"aws_load_balancer_controller":      ComputedStruct(kopsschemas.DataSourceAWSLoadBalancerControllerConfig()),
			"networking":                        ComputedStruct(kopsschemas.DataSourceNetworkingSpec()),
			"api":                               ComputedStruct(kopsschemas.DataSourceAccessSpec()),
			"authentication":                    ComputedStruct(kopsschemas.DataSourceAuthenticationSpec()),
			"authorization":                     ComputedStruct(kopsschemas.DataSourceAuthorizationSpec()),
			"node_authorization":                ComputedStruct(kopsschemas.DataSourceNodeAuthorizationSpec()),
			"cloud_labels":                      ComputedMap(String()),
			"hooks":                             ComputedList(kopsschemas.DataSourceHookSpec()),
			"assets":                            ComputedStruct(kopsschemas.DataSourceAssets()),
			"iam":                               ComputedStruct(kopsschemas.DataSourceIAMSpec()),
			"encryption_config":                 ComputedBool(),
			"tag_subnets":                       ComputedBool(),
			"use_host_certificates":             ComputedBool(),
			"sysctl_parameters":                 ComputedList(String()),
			"rolling_update":                    ComputedStruct(kopsschemas.DataSourceRollingUpdate()),
			"cluster_autoscaler":                ComputedStruct(kopsschemas.DataSourceClusterAutoscalerConfig()),
			"warm_pool":                         ComputedStruct(kopsschemas.DataSourceWarmPoolSpec()),
			"service_account_issuer_discovery":  ComputedStruct(kopsschemas.DataSourceServiceAccountIssuerDiscoveryConfig()),
			"snapshot_controller":               ComputedStruct(kopsschemas.DataSourceSnapshotControllerConfig()),
			"pod_identity_webhook":              ComputedStruct(kopsschemas.DataSourcePodIdentityWebhookConfig()),
			"labels":                            ComputedMap(String()),
			"annotations":                       ComputedMap(String()),
			"name":                              RequiredString(),
			"admin_ssh_key":                     ComputedString(),
			"secrets":                           ComputedStruct(DataSourceClusterSecrets()),
		},
	}
	res.SchemaVersion = 2
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
		},
	}
	return res
}

func ExpandDataSourceCluster(in map[string]interface{}) resources.Cluster {
	if in == nil {
		panic("expand Cluster failure, in is nil")
	}
	return resources.Cluster{
		ClusterSpec: func(in interface{}) kops.ClusterSpec {
			return kopsschemas.ExpandDataSourceClusterSpec(in.(map[string]interface{}))
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
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return resources.ClusterSecrets{}
					}
					return (ExpandDataSourceClusterSecrets(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["secrets"]),
	}
}

func FlattenDataSourceClusterInto(in resources.Cluster, out map[string]interface{}) {
	kopsschemas.FlattenDataSourceClusterSpecInto(in.ClusterSpec, out)
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
