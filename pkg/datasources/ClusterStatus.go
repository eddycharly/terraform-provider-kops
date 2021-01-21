package datasources

import (
	"github.com/eddycharly/terraform-provider-kops/pkg/api/datasources"
	"github.com/eddycharly/terraform-provider-kops/pkg/config"
	datasourcesschemas "github.com/eddycharly/terraform-provider-kops/pkg/schemas/datasources"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ClusterStatus() *schema.Resource {
	return &schema.Resource{
		Read:   ClusterStatusRead,
		Schema: datasourcesschemas.DataSourceClusterStatus().Schema,
	}
}

func ClusterStatusRead(d *schema.ResourceData, m interface{}) error {
	clusterName := d.Get("cluster_name").(string)
	clusterStatus, err := datasources.GetClusterStatus(clusterName, config.Clientset(m))
	if err != nil {
		return err
	}
	for k, v := range datasourcesschemas.FlattenDataSourceClusterStatus(*clusterStatus) {
		if err := d.Set(k, v); err != nil {
			return err
		}
	}
	d.SetId("-")
	return nil
}
