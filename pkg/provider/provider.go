package provider

import (
	"github.com/eddycharly/terraform-provider-kops/pkg/config"
	"github.com/eddycharly/terraform-provider-kops/pkg/datasources"
	"github.com/eddycharly/terraform-provider-kops/pkg/resources"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func NewProvider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"state_store": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("KOPS_STATE_STORE", nil),
				Description: "Location of state storage.",
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"kops_cluster": datasources.Cluster(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"kops_cluster": resources.Cluster(),
		},
		ConfigureFunc: config.ConfigureProvider,
	}
}
