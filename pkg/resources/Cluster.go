package resources

import (
	"github.com/eddycharly/terraform-provider-kops/pkg/api/resources"
	"github.com/eddycharly/terraform-provider-kops/pkg/config"
	"github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Cluster() *schema.Resource {
	return &schema.Resource{
		Create: ClusterCreate,
		Read:   ClusterRead,
		Update: ClusterUpdate,
		Delete: ClusterDelete,
		Exists: ClusterExists,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: schemas.ResourceCluster().Schema,
	}
}

func ClusterCreate(d *schema.ResourceData, m interface{}) error {
	cluster := schemas.ExpandResourceCluster(d.Get("").(map[string]interface{}))
	_, err := resources.SyncCluster(&cluster, config.Clientset(m), config.RollingUpdateOptions(m), config.ValidateOptions(m))
	if err != nil {
		return err
	}
	d.SetId(cluster.Name)
	return ClusterRead(d, m)
}

func ClusterUpdate(d *schema.ResourceData, m interface{}) error {
	cluster := schemas.ExpandResourceCluster(d.Get("").(map[string]interface{}))
	_, err := resources.SyncCluster(&cluster, config.Clientset(m), config.RollingUpdateOptions(m), config.ValidateOptions(m))
	if err != nil {
		return err
	}
	return ClusterRead(d, m)
}

func ClusterRead(d *schema.ResourceData, m interface{}) error {
	cluster, err := resources.GetCluster(d.Id(), config.Clientset(m))
	if err != nil {
		return err
	}
	flattened := schemas.FlattenResourceCluster(*cluster)
	for key, value := range flattened {
		if err := d.Set(key, value); err != nil {
			return err
		}
	}
	return nil
}

func ClusterDelete(d *schema.ResourceData, m interface{}) error {
	err := resources.DeleteCluster(d.Id(), config.Clientset(m))
	if err != nil {
		return err
	}
	return nil
}

func ClusterExists(d *schema.ResourceData, m interface{}) (bool, error) {
	return resources.ClusterExists(d.Id(), config.Clientset(m))
}
