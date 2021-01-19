package resources

import (
	"fmt"
	"strings"

	"github.com/eddycharly/terraform-provider-kops/pkg/api/resources"
	"github.com/eddycharly/terraform-provider-kops/pkg/config"
	resourcesschema "github.com/eddycharly/terraform-provider-kops/pkg/schemas/resources"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func InstanceGroup() *schema.Resource {
	return &schema.Resource{
		Create: InstanceGroupCreate,
		Read:   InstanceGroupRead,
		Update: InstanceGroupUpdate,
		Delete: InstanceGroupDelete,
		Importer: &schema.ResourceImporter{
			State: InstanceGroupImport,
		},
		Schema: resourcesschema.ResourceInstanceGroup().Schema,
	}
}

func InstanceGroupCreate(d *schema.ResourceData, m interface{}) error {
	// cluster := resourcesschema.ExpandResourceCluster(d.Get("").(map[string]interface{}))
	// _, err := resources.SyncCluster(&cluster, config.Clientset(m), config.RollingUpdateOptions(m), config.ValidateOptions(m))
	// if err != nil {
	// 	return err
	// }
	// d.SetId(cluster.Name)
	return ClusterRead(d, m)
}

func InstanceGroupUpdate(d *schema.ResourceData, m interface{}) error {
	// cluster := resourcesschema.ExpandResourceCluster(d.Get("").(map[string]interface{}))
	// _, err := resources.SyncCluster(&cluster, config.Clientset(m), config.RollingUpdateOptions(m), config.ValidateOptions(m))
	// if err != nil {
	// 	return err
	// }
	return ClusterRead(d, m)
}

func InstanceGroupRead(d *schema.ResourceData, m interface{}) error {
	clusterName := d.Get("cluster_name").(string)
	name := d.Get("name").(string)
	instanceGroup, err := resources.GetInstanceGroup(clusterName, name, config.Clientset(m))
	if err != nil {
		return err
	}
	flattened := resourcesschema.FlattenResourceInstanceGroup(*instanceGroup)
	for key, value := range flattened {
		if err := d.Set(key, value); err != nil {
			return err
		}
	}
	d.SetId(instanceGroup.ClusterName + "/" + instanceGroup.Name)
	return nil
}

func InstanceGroupDelete(d *schema.ResourceData, m interface{}) error {
	// err := resources.DeleteCluster(d.Id(), config.Clientset(m))
	// if err != nil {
	// 	return err
	// }
	return nil
}

func InstanceGroupImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	parts := strings.Split(d.Id(), "/")
	if len(parts) != 2 {
		return []*schema.ResourceData{}, fmt.Errorf("Unexpected format for import: %s. Use 'cluster name/instance group name'", d.Id())
	}
	instanceGroup, err := resources.GetInstanceGroup(parts[0], parts[1], config.Clientset(m))
	if err != nil {
		return []*schema.ResourceData{}, err
	}
	flattened := resourcesschema.FlattenResourceInstanceGroup(*instanceGroup)
	for key, value := range flattened {
		if err := d.Set(key, value); err != nil {
			return []*schema.ResourceData{}, err
		}
	}
	d.SetId(instanceGroup.ClusterName + "/" + instanceGroup.Name)
	return []*schema.ResourceData{d}, nil
}
