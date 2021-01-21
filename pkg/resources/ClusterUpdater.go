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
		Update: ClusterUpdaterUpdate,
		Delete: ClusterUpdaterDelete,
		Schema: resourcesschema.ResourceClusterUpdater().Schema,
	}
}

func ClusterUpdaterCreate(d *schema.ResourceData, m interface{}) error {
	updater := resourcesschema.ExpandResourceClusterUpdater(d.Get("").(map[string]interface{}))
	if err := updater.Apply(config.Clientset(m), config.RollingUpdateOptions(m)); err != nil {
		return err
	}
	return ClusterUpdaterRead(d, m)
}

func ClusterUpdaterUpdate(d *schema.ResourceData, m interface{}) error {
	updater := resourcesschema.ExpandResourceClusterUpdater(d.Get("").(map[string]interface{}))
	if err := updater.Apply(config.Clientset(m), config.RollingUpdateOptions(m)); err != nil {
		return err
	}
	return ClusterUpdaterRead(d, m)
}

func ClusterUpdaterRead(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}

func ClusterUpdaterDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
