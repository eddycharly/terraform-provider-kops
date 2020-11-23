package datasources

import (
	"github.com/eddycharly/terraform-provider-kops/pkg/api/resources"
	"github.com/eddycharly/terraform-provider-kops/pkg/config"
	resourceschemas "github.com/eddycharly/terraform-provider-kops/pkg/schemas/resources"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Cluster() *schema.Resource {
	return &schema.Resource{
		Read: ClusterRead,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
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
		d.Set(k, v)
	}
	d.SetId("-")
	return nil
}
