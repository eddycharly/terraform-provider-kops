package datasources

import (
	"github.com/eddycharly/terraform-provider-kops/pkg/api/resources"
	"github.com/eddycharly/terraform-provider-kops/pkg/config"
	kopsschemas "github.com/eddycharly/terraform-provider-kops/pkg/schemas/kops"
	resourceschemas "github.com/eddycharly/terraform-provider-kops/pkg/schemas/resources"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func InstanceGroup() *schema.Resource {
	return &schema.Resource{
		Read: InstanceGroupRead,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: resourceschemas.DataSourceInstanceGroup().Schema,
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
		if err := d.Set(k, v); err != nil {
			return err
		}
	}
	d.SetId("-")
	return nil
}
