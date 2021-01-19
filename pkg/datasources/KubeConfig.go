package datasources

import (
	"github.com/eddycharly/terraform-provider-kops/pkg/api/kube"
	"github.com/eddycharly/terraform-provider-kops/pkg/config"
	datasourcesschemas "github.com/eddycharly/terraform-provider-kops/pkg/schemas/datasources"
	kubeschemas "github.com/eddycharly/terraform-provider-kops/pkg/schemas/kube"
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
	kubeConfigBuilder, err := kube.GetConfig(clusterName, config.Clientset(m))
	if err != nil {
		return err
	}
	for k, v := range kubeschemas.FlattenResourceConfig(*kubeConfigBuilder) {
		if err := d.Set(k, v); err != nil {
			return err
		}
	}
	d.SetId("-")
	return nil
}
