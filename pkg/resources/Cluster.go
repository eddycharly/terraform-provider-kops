package resources

import (
	"context"

	"github.com/eddycharly/terraform-provider-kops/pkg/api"
	"github.com/eddycharly/terraform-provider-kops/pkg/config"
	"github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/eddycharly/terraform-provider-kops/pkg/structures"
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
		Schema: schemas.Cluster().Schema,
	}
}

func ClusterCreate(d *schema.ResourceData, m interface{}) error {
	cluster := structures.ExpandCluster(d.Get("").(map[string]interface{}))
	clientset := m.(*config.Config).Clientset
	_, err := api.SyncCluster(cluster, clientset)
	if err != nil {
		return err
	}
	d.SetId(cluster.Name)
	return ClusterRead(d, m)
}

func ClusterUpdate(d *schema.ResourceData, m interface{}) error {
	cluster := structures.ExpandCluster(d.Get("").(map[string]interface{}))
	clientset := m.(*config.Config).Clientset
	_, err := api.SyncCluster(cluster, clientset)
	if err != nil {
		return err
	}
	return ClusterRead(d, m)
}

func ClusterRead(d *schema.ResourceData, m interface{}) error {
	clientset := m.(*config.Config).Clientset
	cluster, err := api.GetCluster(d.Id(), clientset)
	if err != nil {
		return err
	}
	flattened := structures.FlattenCluster(*cluster)
	for key, value := range flattened {
		if err := d.Set(key, value); err != nil {
			return err
		}
	}
	return nil
}

func ClusterDelete(d *schema.ResourceData, m interface{}) error {
	clientset := m.(*config.Config).Clientset
	cluster, err := clientset.GetCluster(context.Background(), d.Id())
	if err != nil {
		return err
	}
	return clientset.DeleteCluster(context.Background(), cluster)
}

func ClusterExists(d *schema.ResourceData, m interface{}) (bool, error) {
	clientset := m.(*config.Config).Clientset
	return api.ClusterExists(d.Id(), clientset)
}
