package resources

import (
	"context"

	"github.com/eddycharly/terraform-provider-kops/pkg/api/resources"
	"github.com/eddycharly/terraform-provider-kops/pkg/config"
	"github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	resourcesschema "github.com/eddycharly/terraform-provider-kops/pkg/schemas/resources"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Cluster(providerVersion string) *schema.Resource {
	res := resourcesschema.ResourceCluster()
	return &schema.Resource{
		CreateContext:  ClusterCreate(providerVersion),
		ReadContext:    ClusterRead,
		UpdateContext:  ClusterUpdate(providerVersion),
		DeleteContext:  ClusterDelete,
		CustomizeDiff:  schemas.CustomizeDiffRevision(providerVersion),
		Schema:         res.Schema,
		SchemaVersion:  res.SchemaVersion,
		StateUpgraders: res.StateUpgraders,
		Importer: &schema.ResourceImporter{
			StateContext: ClusterImport(providerVersion),
		},
	}
}

func ClusterCreate(providerVersion string) func(c context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return func(c context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		in := resourcesschema.ExpandResourceCluster(d.Get("").(map[string]interface{}))
		if cluster, err := resources.CreateCluster(in.Name, in.Labels, in.Annotations, in.AdminSshKey, in.Secrets, in.ClusterSpec, config.Clientset(m)); err != nil {
			return diag.FromErr(err)
		} else {
			d.SetId(cluster.Name)
			return ClusterRead(c, d, m)
		}
	}
}

func ClusterUpdate(providerVersion string) func(c context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return func(c context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		in := resourcesschema.ExpandResourceCluster(d.Get("").(map[string]interface{}))
		if cluster, err := resources.UpdateCluster(in.Name, in.Labels, in.Annotations, in.AdminSshKey, in.Secrets, in.ClusterSpec, config.Clientset(m)); err != nil {
			return diag.FromErr(err)
		} else {
			d.SetId(cluster.Name)
			return ClusterRead(c, d, m)
		}
	}
}

func ClusterRead(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	in := resourcesschema.ExpandResourceCluster(d.Get("").(map[string]interface{}))
	if cluster, err := resources.GetCluster(in.Name, config.Clientset(m)); err != nil {
		return diag.FromErr(err)
	} else {
		flattened := resourcesschema.FlattenResourceCluster(*cluster)
		for key, value := range flattened {
			switch key {
			case "revision":
			case "provider_version":
				continue
			default:
				if err := d.Set(key, value); err != nil {
					return diag.FromErr(err)
				}
			}
		}
	}
	return nil
}

func ClusterDelete(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	in := resourcesschema.ExpandResourceCluster(d.Get("").(map[string]interface{}))
	if err := resources.DeleteCluster(in.Name, config.Clientset(m)); err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func ClusterImport(providerVersion string) func(_ context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	return func(_ context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
		if cluster, err := resources.GetCluster(d.Id(), config.Clientset(m)); err != nil {
			return []*schema.ResourceData{}, err
		} else {
			cluster.ProviderVersion = providerVersion
			flattened := resourcesschema.FlattenResourceCluster(*cluster)
			for key, value := range flattened {
				switch key {
				case "revision":
				case "provider_version":
					continue
				default:
					if err := d.Set(key, value); err != nil {
						return []*schema.ResourceData{}, err
					}
				}
			}
			return []*schema.ResourceData{d}, nil
		}
	}
}
