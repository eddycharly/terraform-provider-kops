package provider

import (
	"github.com/eddycharly/terraform-provider-kops/pkg/config"
	"github.com/eddycharly/terraform-provider-kops/pkg/datasources"
	"github.com/eddycharly/terraform-provider-kops/pkg/resources"
	configschemas "github.com/eddycharly/terraform-provider-kops/pkg/schemas/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func NewProvider() *schema.Provider {
	return &schema.Provider{
		Schema: configschemas.ConfigProvider().Schema,
		DataSourcesMap: map[string]*schema.Resource{
			"kops_cluster":        datasources.Cluster(),
			"kops_cluster_status": datasources.ClusterStatus(),
			"kops_instance_group": datasources.InstanceGroup(),
			"kops_kube_config":    datasources.KubeConfig(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"kops_cluster":         resources.Cluster(),
			"kops_cluster_updater": resources.ClusterUpdater(),
			"kops_dockerconfig":    resources.DockerConfig(),
			"kops_instance_group":  resources.InstanceGroup(),
		},
		ConfigureContextFunc: config.ConfigureProvider,
	}
}
