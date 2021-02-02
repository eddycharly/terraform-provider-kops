package resources

import (
	"fmt"
	"time"

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
			State: ClusterImport,
		},
		Schema: resourcesschema.ResourceCluster().Schema,
	}
}

func ClusterCreate(d *schema.ResourceData, m interface{}) error {
	cluster := resourcesschema.ExpandResourceCluster(d.Get("").(map[string]interface{}))
	if cluster, err := resources.CreateCluster(cluster.Name, cluster.AdminSshKey, cluster.ClusterSpec, config.Clientset(m)); err != nil {
		return err
	} else {
		d.SetId(fmt.Sprintf("%s-%d", cluster.Name, time.Now().UnixNano()))
	}
	return ClusterRead(d, m)
}

func ClusterUpdate(d *schema.ResourceData, m interface{}) error {
	cluster := resourcesschema.ExpandResourceCluster(d.Get("").(map[string]interface{}))
	if _, err := resources.UpdateCluster(cluster.Name, cluster.AdminSshKey, cluster.ClusterSpec, config.Clientset(m)); err != nil {
		return err
	}
	d.SetId(fmt.Sprintf("%s-%d", cluster.Name, time.Now().UnixNano()))
	return ClusterRead(d, m)
}

func ClusterRead(d *schema.ResourceData, m interface{}) error {
	cluster := resourcesschema.ExpandResourceCluster(d.Get("").(map[string]interface{}))
	if cluster, err := resources.GetCluster(cluster.Name, config.Clientset(m)); err != nil {
		return err
	} else {
		flattened := resourcesschema.FlattenResourceCluster(*cluster)
		for key, value := range flattened {
			if err := d.Set(key, value); err != nil {
				return err
			}
		}
	}
	return nil
}

func ClusterDelete(d *schema.ResourceData, m interface{}) error {
	cluster := resourcesschema.ExpandResourceCluster(d.Get("").(map[string]interface{}))
	if err := resources.DeleteCluster(cluster.Name, config.Clientset(m)); err != nil {
		return err
	}
	return nil
}

func ClusterImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	if cluster, err := resources.GetCluster(d.Id(), config.Clientset(m)); err != nil {
		return []*schema.ResourceData{}, err
	} else {
		flattened := resourcesschema.FlattenResourceCluster(*cluster)
		for key, value := range flattened {
			if err := d.Set(key, value); err != nil {
				return []*schema.ResourceData{}, err
			}
		}
		d.SetId(fmt.Sprintf("%s-%d", cluster.Name, time.Now().UnixNano()))
	}
	return []*schema.ResourceData{d}, nil
}
