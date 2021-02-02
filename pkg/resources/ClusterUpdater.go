package resources

import (
	"github.com/eddycharly/terraform-provider-kops/pkg/config"
	resourcesschema "github.com/eddycharly/terraform-provider-kops/pkg/schemas/resources"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ClusterUpdater() *schema.Resource {
	return &schema.Resource{
		Create: ClusterUpdaterCreate,
		Read:   ClusterUpdaterRead,
		Delete: schema.RemoveFromState,
		Schema: resourcesschema.ResourceClusterUpdater().Schema,
	}
}

func ClusterUpdaterCreate(d *schema.ResourceData, m interface{}) error {
	updater := resourcesschema.ExpandResourceClusterUpdater(d.Get("").(map[string]interface{}))
	if err := updater.UpdateCluster(config.Clientset(m)); err != nil {
		return err
	}
	d.SetId(updater.ClusterName)
	return nil
}

func ClusterUpdaterRead(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}
