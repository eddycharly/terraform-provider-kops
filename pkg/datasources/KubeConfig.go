package datasources

import (
	"log"

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
		Schema: schemas.DataSourceKubeConfig().Schema,
	}
}

func KubeConfigRead(d *schema.ResourceData, m interface{}) error {
	kubeConfig := schemas.ExpandDataSourceKubeConfig(d.Get("").(map[string]interface{}))
	kubeConfigBuilder, err := resources.GetKubeConfig(kubeConfig.ClusterName, config.Clientset(m))
	if err != nil {
		return err
	}
	for k, v := range schemas.FlattenDataSourceConfig(*kube.FromKubeconfigBuilder(kubeConfigBuilder)) {
		log.Printf("%s - %s\n", k, v)
		d.Set(k, v)
	}
	d.SetId(kubeConfig.ClusterName)
	return nil
}
