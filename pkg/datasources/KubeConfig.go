package datasources

import (
	"github.com/eddycharly/terraform-provider-kops/pkg/api/kube"
	"github.com/eddycharly/terraform-provider-kops/pkg/api/resources"
	"github.com/eddycharly/terraform-provider-kops/pkg/config"
	"github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func KubeConfig() *schema.Resource {
	return &schema.Resource{
		Read: KubeConfigRead,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: schemas.DataSourceDatasourcesKubeConfig().Schema,
	}
}

func KubeConfigRead(d *schema.ResourceData, m interface{}) error {
	clusterName := d.Get("cluster_name").(string)
	kubeConfigBuilder, err := resources.GetKubeConfig(clusterName, config.Clientset(m))
	if err != nil {
		return err
	}
	for k, v := range schemas.FlattenDataSourceKubeConfig(*kube.FromKubeconfigBuilder(kubeConfigBuilder)) {
		d.Set(k, v)
	}
	d.SetId("-")
	return nil
}
