package resources

import (
	"context"
	"fmt"
	"strings"

	"github.com/eddycharly/terraform-provider-kops/pkg/api/resources"
	"github.com/eddycharly/terraform-provider-kops/pkg/api/utils"
	"github.com/eddycharly/terraform-provider-kops/pkg/config"
	"github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	resourcesschema "github.com/eddycharly/terraform-provider-kops/pkg/schemas/resources"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func InstanceGroup() *schema.Resource {
	return &schema.Resource{
		Create:        InstanceGroupCreate,
		Read:          InstanceGroupRead,
		Update:        InstanceGroupUpdate,
		Delete:        InstanceGroupDelete,
		CustomizeDiff: schemas.CustomizeDiffRevision,
		Importer:      &schema.ResourceImporter{StateContext: InstanceGroupImport},
		Schema:        resourcesschema.ResourceInstanceGroup().Schema,
	}
}

func InstanceGroupCreate(d *schema.ResourceData, m interface{}) error {
	in := resourcesschema.ExpandResourceInstanceGroup(d.Get("").(map[string]interface{}))
	if instanceGroup, err := resources.CreateInstanceGroup(in.ClusterName, in.Name, in.InstanceGroupSpec, config.Clientset(m)); err != nil {
		return err
	} else {
		d.SetId(fmt.Sprintf("%s/%s", instanceGroup.ClusterName, instanceGroup.Name))
		return InstanceGroupRead(d, m)
	}
}

func InstanceGroupUpdate(d *schema.ResourceData, m interface{}) error {
	in := resourcesschema.ExpandResourceInstanceGroup(d.Get("").(map[string]interface{}))
	if instanceGroup, err := resources.UpdateInstanceGroup(in.ClusterName, in.Name, in.InstanceGroupSpec, config.Clientset(m)); err != nil {
		return err
	} else {
		d.SetId(fmt.Sprintf("%s/%s", instanceGroup.ClusterName, instanceGroup.Name))
		return InstanceGroupRead(d, m)
	}
}

func InstanceGroupRead(d *schema.ResourceData, m interface{}) error {
	in := resourcesschema.ExpandResourceInstanceGroup(d.Get("").(map[string]interface{}))
	if instanceGroup, err := resources.GetInstanceGroup(in.ClusterName, in.Name, config.Clientset(m)); err != nil {
		return err
	} else {
		flattened := resourcesschema.FlattenResourceInstanceGroup(*instanceGroup)
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

func InstanceGroupDelete(d *schema.ResourceData, m interface{}) error {
	in := resourcesschema.ExpandResourceInstanceGroup(d.Get("").(map[string]interface{}))
	if err := utils.InstanceGroupDelete(config.Clientset(m), in.ClusterName, in.Name); err != nil {
		return err
	}
	return nil
}

func InstanceGroupImport(_ context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
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
			d.SetId(fmt.Sprintf("%s/%s", instanceGroup.ClusterName, instanceGroup.Name))
		}
	}
	return []*schema.ResourceData{d}, nil
}
