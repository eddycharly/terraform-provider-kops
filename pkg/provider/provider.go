package provider

import (
	"github.com/eddycharly/terraform-provider-kops/pkg/config"
	"github.com/eddycharly/terraform-provider-kops/pkg/resources"
	"github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func NewProvider() *schema.Provider {
	return &schema.Provider{
		Schema: schemas.ConfigProvider().Schema,
		// DataSourcesMap: map[string]*schema.Resource{
		// 	"kops_cluster": datasources.Cluster(),
		// },
		ResourcesMap: map[string]*schema.Resource{
			"kops_cluster": resources.Cluster(),
		},
		ConfigureFunc: config.ConfigureProvider,
	}
}
