package datasources

import (
	"github.com/eddycharly/terraform-provider-kops/pkg/api/resources"
	"github.com/eddycharly/terraform-provider-kops/pkg/config"
	resourceschemas "github.com/eddycharly/terraform-provider-kops/pkg/schemas/resources"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Cluster() *schema.Resource {
	return &schema.Resource{
		Read:   ClusterRead,
		Schema: resourceschemas.DataSourceCluster().Schema,
	}
}

func ClusterRead(d *schema.ResourceData, m interface{}) error {
	clusterName := d.Get("name").(string)
	cluster, err := resources.GetCluster(clusterName, config.Clientset(m))
	if err != nil {
		return err
	}
	for k, v := range resourceschemas.FlattenDataSourceCluster(*cluster) {
		if err := d.Set(k, v); err != nil {
			return err
		}
	}
	d.SetId("-")
	return nil
}
