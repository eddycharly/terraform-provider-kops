package datasources

import (
	"context"

	"github.com/eddycharly/terraform-provider-kops/pkg/api/resources"
	"github.com/eddycharly/terraform-provider-kops/pkg/config"
	resourceschemas "github.com/eddycharly/terraform-provider-kops/pkg/schemas/resources"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func InstanceGroup(providerVersion string) *schema.Resource {
	res := resourceschemas.DataSourceInstanceGroup()
	return &schema.Resource{
		ReadContext:    InstanceGroupRead(providerVersion),
		Schema:         res.Schema,
		SchemaVersion:  res.SchemaVersion,
		StateUpgraders: res.StateUpgraders,
	}
}

func InstanceGroupRead(providerVersion string) func(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return func(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		clusterName := d.Get("cluster_name").(string)
		name := d.Get("name").(string)
		instanceGroup, err := resources.GetInstanceGroup(clusterName, name, config.Clientset(m))
		if err != nil {
			return diag.FromErr(err)
		}
		instanceGroup.ProviderVersion = providerVersion
		for k, v := range resourceschemas.FlattenDataSourceInstanceGroup(*instanceGroup) {
			if err := d.Set(k, v); err != nil {
				return diag.FromErr(err)
			}
		}
		d.SetId("-")
		return nil
	}
}
