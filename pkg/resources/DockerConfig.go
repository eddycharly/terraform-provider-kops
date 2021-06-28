package resources

import (
	"context"

	"github.com/eddycharly/terraform-provider-kops/pkg/api/resources"
	"github.com/eddycharly/terraform-provider-kops/pkg/config"
	resourcesschema "github.com/eddycharly/terraform-provider-kops/pkg/schemas/resources"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DockerConfig() *schema.Resource {
	return &schema.Resource{
		CreateContext: DockerConfigCreate,
		ReadContext:   DockerConfigRead,
		UpdateContext: DockerConfigUpdate,
		DeleteContext: DockerConfigDelete,
		Schema:        resourcesschema.ResourceDockerConfig().Schema,
	}
}

func DockerConfigCreate(c context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	in := resourcesschema.ExpandResourceDockerConfig(d.Get("").(map[string]interface{}))
	if dockerConfig, err := resources.CreateOrUpdateDockerConfig(in.ClusterName, in.DockerConfig, config.Clientset(m)); err != nil {
		return diag.FromErr(err)
	} else {
		d.SetId(dockerConfig.ClusterName)
		return DockerConfigRead(c, d, m)
	}
}

func DockerConfigUpdate(c context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	in := resourcesschema.ExpandResourceDockerConfig(d.Get("").(map[string]interface{}))
	if dockerConfig, err := resources.CreateOrUpdateDockerConfig(in.ClusterName, in.DockerConfig, config.Clientset(m)); err != nil {
		return diag.FromErr(err)
	} else {
		d.SetId(dockerConfig.ClusterName)
		return DockerConfigRead(c, d, m)
	}
}

func DockerConfigRead(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	in := resourcesschema.ExpandResourceDockerConfig(d.Get("").(map[string]interface{}))
	if dockerConfig, err := resources.GetDockerConfig(in.ClusterName, config.Clientset(m)); err != nil {
		return diag.FromErr(err)
	} else {
		flattened := resourcesschema.FlattenResourceDockerConfig(*dockerConfig)
		for key, value := range flattened {
			if err := d.Set(key, value); err != nil {
				return diag.FromErr(err)
			}
		}
	}
	return nil
}

func DockerConfigDelete(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	in := resourcesschema.ExpandResourceDockerConfig(d.Get("").(map[string]interface{}))
	if err := resources.DeleteDockerConfig(in.ClusterName, config.Clientset(m)); err != nil {
		return diag.FromErr(err)
	}
	return nil
}

// func ClusterImport(_ context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
// 	if cluster, err := resources.GetCluster(d.Id(), config.Clientset(m)); err != nil {
// 		return []*schema.ResourceData{}, err
// 	} else {
// 		flattened := resourcesschema.FlattenResourceCluster(*cluster)
// 		for key, value := range flattened {
// 			if key != "revision" {
// 				if err := d.Set(key, value); err != nil {
// 					return []*schema.ResourceData{}, err
// 				}
// 			}
// 		}
// 		return []*schema.ResourceData{d}, nil
// 	}
// }
