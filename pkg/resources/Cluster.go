package resources

import (
	"github.com/eddycharly/terraform-provider-kops/pkg/api/resources"
	"github.com/eddycharly/terraform-provider-kops/pkg/config"
	"github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	resourcesschema "github.com/eddycharly/terraform-provider-kops/pkg/schemas/resources"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Cluster() *schema.Resource {
	return &schema.Resource{
		Create:        ClusterCreate,
		Read:          ClusterRead,
		Update:        ClusterUpdate,
		Delete:        ClusterDelete,
		CustomizeDiff: schemas.CustomizeDiffRevision,
		Importer:      &schema.ResourceImporter{StateContext: schema.ImportStatePassthroughContext},
		Schema:        resourcesschema.ResourceCluster().Schema,
	}
}

func ClusterCreate(d *schema.ResourceData, m interface{}) error {
	in := resourcesschema.ExpandResourceCluster(d.Get("").(map[string]interface{}))
	if cluster, err := resources.CreateCluster(in.Name, in.AdminSshKey, in.ClusterSpec, config.Clientset(m)); err != nil {
		return err
	} else {
		d.SetId(cluster.Name)
		return ClusterRead(d, m)
	}
}

func ClusterUpdate(d *schema.ResourceData, m interface{}) error {
	in := resourcesschema.ExpandResourceCluster(d.Get("").(map[string]interface{}))
	if cluster, err := resources.UpdateCluster(in.Name, in.AdminSshKey, in.ClusterSpec, config.Clientset(m)); err != nil {
		return err
	} else {
		d.SetId(cluster.Name)
		return ClusterRead(d, m)
	}
}

func ClusterRead(d *schema.ResourceData, m interface{}) error {
	in := resourcesschema.ExpandResourceCluster(d.Get("").(map[string]interface{}))
	if cluster, err := resources.GetCluster(in.Name, config.Clientset(m)); err != nil {
		return err
	} else {
		flattened := resourcesschema.FlattenResourceCluster(*cluster)
		for key, value := range flattened {
			if key != "revision" {
				if err := d.Set(key, value); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func ClusterDelete(d *schema.ResourceData, m interface{}) error {
	in := resourcesschema.ExpandResourceCluster(d.Get("").(map[string]interface{}))
	if err := resources.DeleteCluster(in.Name, config.Clientset(m)); err != nil {
		return err
	}
	return nil
}
