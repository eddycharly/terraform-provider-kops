package datasources

import (
	"github.com/eddycharly/terraform-provider-kops/pkg/api/datasources"
	"github.com/eddycharly/terraform-provider-kops/pkg/config"
	datasourcesschemas "github.com/eddycharly/terraform-provider-kops/pkg/schemas/datasources"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func KubeConfig() *schema.Resource {
	return &schema.Resource{
		Read:   KubeConfigRead,
		Schema: datasourcesschemas.DataSourceKubeConfig().Schema,
	}
}

func KubeConfigRead(d *schema.ResourceData, m interface{}) error {
	clusterName := d.Get("cluster_name").(string)
	kubeConfig, err := datasources.GetConfig(config.Clientset(m), clusterName)
	if err != nil {
		return err
	}
	for k, v := range datasourcesschemas.FlattenDataSourceKubeConfig(*kubeConfig) {
		if err := d.Set(k, v); err != nil {
			return err
		}
	}
	d.SetId("-")
	return nil
}
