package resources

import (
	"fmt"
	"strings"
	"time"

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
	instanceGroup := resourcesschema.ExpandResourceInstanceGroup(d.Get("").(map[string]interface{}))
	if instanceGroup, err := resources.CreateInstanceGroup(instanceGroup.ClusterName, instanceGroup.Name, instanceGroup.InstanceGroupSpec, config.Clientset(m)); err != nil {
		return err
	} else {
		d.SetId(fmt.Sprintf("%s-%s-%d", instanceGroup.ClusterName, instanceGroup.Name, time.Now().UnixNano()))
	}
	return InstanceGroupRead(d, m)
}

func InstanceGroupUpdate(d *schema.ResourceData, m interface{}) error {
	instanceGroup := resourcesschema.ExpandResourceInstanceGroup(d.Get("").(map[string]interface{}))
	if _, err := resources.UpdateInstanceGroup(instanceGroup.ClusterName, instanceGroup.Name, instanceGroup.InstanceGroupSpec, config.Clientset(m)); err != nil {
		return err
	}
	d.SetId(fmt.Sprintf("%s-%s-%d", instanceGroup.ClusterName, instanceGroup.Name, time.Now().UnixNano()))
	return InstanceGroupRead(d, m)
}

func InstanceGroupRead(d *schema.ResourceData, m interface{}) error {
	instanceGroup := resourcesschema.ExpandResourceInstanceGroup(d.Get("").(map[string]interface{}))
	if instanceGroup, err := resources.GetInstanceGroup(instanceGroup.ClusterName, instanceGroup.Name, config.Clientset(m)); err != nil {
		return err
	} else {
		flattened := resourcesschema.FlattenResourceInstanceGroup(*instanceGroup)
		for key, value := range flattened {
			if err := d.Set(key, value); err != nil {
				return err
			}
		}
	}
	return nil
}

func InstanceGroupDelete(d *schema.ResourceData, m interface{}) error {
	instanceGroup := resourcesschema.ExpandResourceInstanceGroup(d.Get("").(map[string]interface{}))
	if err := resources.DeleteInstanceGroup(instanceGroup.ClusterName, instanceGroup.Name, config.Clientset(m)); err != nil {
		return err
	}
	return nil
}

func InstanceGroupImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	if parts := strings.Split(d.Id(), "/"); len(parts) != 2 {
		return []*schema.ResourceData{}, fmt.Errorf("Unexpected id format: %s. Please use 'cluster name/instance group name' format.", d.Id())
	} else {
		if instanceGroup, err := resources.GetInstanceGroup(parts[0], parts[1], config.Clientset(m)); err != nil {
			return []*schema.ResourceData{}, err
		} else {
			flattened := resourcesschema.FlattenResourceInstanceGroup(*instanceGroup)
			for key, value := range flattened {
				if err := d.Set(key, value); err != nil {
					return []*schema.ResourceData{}, err
				}
			}
			d.SetId(fmt.Sprintf("%s-%s-%d", instanceGroup.ClusterName, instanceGroup.Name, time.Now().UnixNano()))
		}
	}
	return []*schema.ResourceData{d}, nil
}
