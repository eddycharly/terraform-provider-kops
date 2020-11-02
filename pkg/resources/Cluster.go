package resources

import (
	"context"
	"log"

	"github.com/eddycharly/terraform-provider-kops/pkg/api"
	"github.com/eddycharly/terraform-provider-kops/pkg/config"
	"github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/eddycharly/terraform-provider-kops/pkg/structures"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/apimachinery/pkg/api/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/kops/pkg/apis/kops"
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
	log.Println(d.Get(""))

	cluster := structures.ExpandCluster(d.Get("").(map[string]interface{}))
	clientset := m.(*config.Config).Clientset
	err := cluster.Create(clientset)
	if err != nil {
		return err
	}
	d.SetId(cluster.Name)
	return ClusterRead(d, m)
}

func ClusterUpdate(d *schema.ResourceData, m interface{}) error {
	cluster := structures.ExpandCluster(d.Get("").(map[string]interface{}))
	clientset := m.(*config.Config).Clientset
	err := cluster.Update(clientset)
	if err != nil {
		return err
	}
	return ClusterRead(d, m)
}

func ClusterRead(d *schema.ResourceData, m interface{}) error {
	cluster, instanceGroups, err := getClusterAndInstanceGroups(d, m)
	if err != nil {
		return err
	}
	flattened := structures.FlattenCluster(api.FromKopsCluster(*cluster, instanceGroups...))
	for key, value := range flattened {
		if err := d.Set(key, value); err != nil {
			return err
		}
	}
	return nil
}

// func clusterUpdate(d *schema.ResourceData, m interface{}) error {
// 	log.Println("*******************************************")
// 	log.Println("* UPDATE                                  *")
// 	log.Println("*******************************************")
// 	if ok, _ := clusterExists(d, m); !ok {
// 		d.SetId("")
// 		return nil
// 	}

// 	clientset := m.(*config.Config).Clientset

// 	_, err := clientset.UpdateCluster(
// 		context.Background(),
// 		&kops.Cluster{
// 			ObjectMeta: structures.ExpandObjectMeta(sectionData(d, "metadata")),
// 			Spec:       structures.ExpandClusterSpec(sectionData(d, "spec")),
// 		}, nil)

// 	if err != nil {
// 		return err
// 	}

// 	return clusterRead(d, m)
// }

func ClusterDelete(d *schema.ResourceData, m interface{}) error {
	clientset := m.(*config.Config).Clientset
	cluster, _, err := getClusterAndInstanceGroups(d, m)
	if err != nil {
		return err
	}
	return clientset.DeleteCluster(context.Background(), cluster)
}

func ClusterExists(d *schema.ResourceData, m interface{}) (bool, error) {
	_, _, err := getClusterAndInstanceGroups(d, m)
	if err != nil {
		if errors.IsNotFound(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func getClusterAndInstanceGroups(d *schema.ResourceData, m interface{}) (*kops.Cluster, []kops.InstanceGroup, error) {
	clientset := m.(*config.Config).Clientset
	cluster, err := clientset.GetCluster(context.Background(), d.Id())
	if err != nil {
		return nil, nil, err
	}
	ig := clientset.InstanceGroupsFor(cluster)
	instanceGroups, err := ig.List(context.Background(), v1.ListOptions{})
	if err != nil {
		return nil, nil, err
	}
	return cluster, instanceGroups.Items, err
}
