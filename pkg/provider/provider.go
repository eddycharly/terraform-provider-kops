package provider

import (
	"github.com/eddycharly/terraform-provider-kops/pkg/config"
	"github.com/eddycharly/terraform-provider-kops/pkg/datasources"
	"github.com/eddycharly/terraform-provider-kops/pkg/resources"
	configschemas "github.com/eddycharly/terraform-provider-kops/pkg/schemas/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func NewProvider(providerVersion string) func() *schema.Provider {
	return func() *schema.Provider {
		return &schema.Provider{
			Schema: configschemas.ConfigProvider().Schema,
			DataSourcesMap: map[string]*schema.Resource{
				"kops_cluster":        datasources.Cluster(providerVersion),
				"kops_cluster_status": datasources.ClusterStatus(),
				"kops_instance_group": datasources.InstanceGroup(providerVersion),
				"kops_kube_config":    datasources.KubeConfig(),
			},
			ResourcesMap: map[string]*schema.Resource{
				"kops_cluster":         resources.Cluster(providerVersion),
				"kops_cluster_updater": resources.ClusterUpdater(providerVersion),
				"kops_instance_group":  resources.InstanceGroup(providerVersion),
			},
			ConfigureContextFunc: config.ConfigureProvider,
		}
	}
}
