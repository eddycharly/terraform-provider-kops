package datasources

import (
	"context"

	"github.com/eddycharly/terraform-provider-kops/pkg/api/resources"
	"github.com/eddycharly/terraform-provider-kops/pkg/config"
	resourceschemas "github.com/eddycharly/terraform-provider-kops/pkg/schemas/resources"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Cluster(providerVersion string) *schema.Resource {
	res := resourceschemas.DataSourceCluster()
	return &schema.Resource{
		ReadContext:    ClusterRead(providerVersion),
		Schema:         res.Schema,
		SchemaVersion:  res.SchemaVersion,
		StateUpgraders: res.StateUpgraders,
	}
}

func ClusterRead(providerVersion string) func(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return func(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		clusterName := d.Get("name").(string)
		cluster, err := resources.GetCluster(clusterName, config.Clientset(m))
		if err != nil {
			return diag.FromErr(err)
		}
		cluster.ProviderVersion = providerVersion
		for k, v := range resourceschemas.FlattenDataSourceCluster(*cluster) {
			if err := d.Set(k, v); err != nil {
				return diag.FromErr(err)
			}
		}
		d.SetId("-")
		return nil
	}
}
