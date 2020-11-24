package resources

import (
	"encoding/json"
	"log"

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
	_, err := resources.SyncCluster(&cluster, config.Clientset(m), config.RollingUpdateOptions(m), config.ValidateOptions(m))
	if err != nil {
		return err
	}
	d.SetId(cluster.Name)
	return ClusterRead(d, m)
}

func ClusterUpdate(d *schema.ResourceData, m interface{}) error {
	log.Println("************************************* ClusterUpdate")
	cluster := resourcesschema.ExpandResourceCluster(d.Get("").(map[string]interface{}))
	bolB, err := json.Marshal(&cluster)
	if err != nil {
		return err
	}
	log.Println(string(bolB))
	_, err = resources.SyncCluster(&cluster, config.Clientset(m), config.RollingUpdateOptions(m), config.ValidateOptions(m))
	if err != nil {
		log.Println("-------------------------------------- " + err.Error())
		return err
	}
	return ClusterRead(d, m)
}

func ClusterRead(d *schema.ResourceData, m interface{}) error {
	log.Println("************************************* ClusterRead")
	cluster, err := resources.GetCluster(d.Id(), config.Clientset(m))
	if err != nil {
		return err
	}
	flattened := resourcesschema.FlattenResourceCluster(*cluster)
	for key, value := range flattened {
		// if key == "instance_group" {
		// 	raw, _ := d.GetOk("instance_group")
		// 	setSrc := raw.(*schema.Set)
		// 	setDest := value.(*schema.Set)
		// 	for _, x := range setSrc.List() {
		// 		log.Printf("******************************************** %s -> %t", x.(map[string]interface{})["name"], setDest.Contains(x))
		// 	}
		// }
		if err := d.Set(key, value); err != nil {
			log.Printf("######################### %s", err)
			return err
		}
	}
	// before, after := d.GetChange("instance_group")
	// log.Printf("######################### Before : %#v\n", before)
	// log.Printf("######################### After  : %#v\n", after)
	d.SetId(cluster.Name)
	log.Println("************************************* EndClusterRead")
	return nil
}

func ClusterDelete(d *schema.ResourceData, m interface{}) error {
	err := resources.DeleteCluster(d.Id(), config.Clientset(m))
	if err != nil {
		return err
	}
	return nil
}
