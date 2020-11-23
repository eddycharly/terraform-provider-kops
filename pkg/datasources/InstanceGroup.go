package datasources

import (
	"github.com/eddycharly/terraform-provider-kops/pkg/api/resources"
	"github.com/eddycharly/terraform-provider-kops/pkg/config"
	datasourcesschemas "github.com/eddycharly/terraform-provider-kops/pkg/schemas/datasources"
	kopsschemas "github.com/eddycharly/terraform-provider-kops/pkg/schemas/kops"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func InstanceGroup() *schema.Resource {
	return &schema.Resource{
		Read: InstanceGroupRead,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: datasourcesschemas.DataSourceInstanceGroup().Schema,
	}
}

func InstanceGroupRead(d *schema.ResourceData, m interface{}) error {
	clusterName := d.Get("cluster_name").(string)
	name := d.Get("name").(string)
	instanceGroup, err := resources.GetInstanceGroup(clusterName, name, config.Clientset(m))
	if err != nil {
		return err
	}
	for k, v := range kopsschemas.FlattenDataSourceInstanceGroupSpec(instanceGroup.Spec) {
		d.Set(k, v)
	}
	d.SetId("-")
	return nil
}
