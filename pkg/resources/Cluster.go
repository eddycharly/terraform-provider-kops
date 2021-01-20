package resources

import (
	"github.com/eddycharly/terraform-provider-kops/pkg/api/resources"
	"github.com/eddycharly/terraform-provider-kops/pkg/config"
	resourcesschema "github.com/eddycharly/terraform-provider-kops/pkg/schemas/resources"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Cluster() *schema.Resource {
	return &schema.Resource{
		Create: ClusterCreate,
		Read:   ClusterRead,
		Update: ClusterUpdate,
		Delete: ClusterDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: resourcesschema.ResourceCluster().Schema,
	}
}

func ClusterCreate(d *schema.ResourceData, m interface{}) error {
	cluster := resourcesschema.ExpandResourceCluster(d.Get("").(map[string]interface{}))
	if cluster, err := resources.CreateCluster(cluster.Name, cluster.AdminSshKey, cluster.ClusterSpec, config.Clientset(m)); err != nil {
		return err
	} else {
		d.SetId(cluster.Name)
	}
	return ClusterRead(d, m)
}

func ClusterUpdate(d *schema.ResourceData, m interface{}) error {
	cluster := resourcesschema.ExpandResourceCluster(d.Get("").(map[string]interface{}))
	if _, err := resources.UpdateCluster(cluster.Name, cluster.AdminSshKey, cluster.ClusterSpec, config.Clientset(m)); err != nil {
		return err
	}
	return ClusterRead(d, m)
}

func ClusterRead(d *schema.ResourceData, m interface{}) error {
	if cluster, err := resources.GetCluster(d.Id(), config.Clientset(m)); err != nil {
		return err
	} else {
		flattened := resourcesschema.FlattenResourceCluster(*cluster)
		for key, value := range flattened {
			if err := d.Set(key, value); err != nil {
				return err
			}
		}
		d.SetId(cluster.Name)
	}
	return nil
}

func ClusterDelete(d *schema.ResourceData, m interface{}) error {
	if err := resources.DeleteCluster(d.Id(), config.Clientset(m)); err != nil {
		return err
	}
	return nil
}
