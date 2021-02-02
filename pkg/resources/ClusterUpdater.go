package resources

import (
	"fmt"
	"time"

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
	in := resourcesschema.ExpandResourceClusterUpdater(d.Get("").(map[string]interface{}))
	if err := in.UpdateCluster(config.Clientset(m)); err != nil {
		return err
	}
	d.SetId(fmt.Sprintf("%s-%d", in.ClusterName, time.Now().UnixNano()))
	return nil
}

func ClusterUpdaterRead(d *schema.ResourceData, m interface{}) error {
	return nil
}
