package resources

import (
	"context"

	"github.com/eddycharly/terraform-provider-kops/pkg/config"
	"github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	resourcesschema "github.com/eddycharly/terraform-provider-kops/pkg/schemas/resources"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ClusterUpdater() *schema.Resource {
	return &schema.Resource{
		CreateWithoutTimeout: ClusterUpdaterCreateOrUpdate,
		ReadContext:          schema.NoopContext,
		UpdateWithoutTimeout: ClusterUpdaterCreateOrUpdate,
		DeleteWithoutTimeout: ClusterUpdaterDelete,
		CustomizeDiff:        schemas.CustomizeDiffRevision,
		Schema:               resourcesschema.ResourceClusterUpdater().Schema,
	}
}

func ClusterUpdaterCreateOrUpdate(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	in := resourcesschema.ExpandResourceClusterUpdater(d.Get("").(map[string]interface{}))
	if err := in.UpdateCluster(config.Clientset(m)); err != nil {
		return diag.FromErr(err)
	}
	d.SetId(in.ClusterName)
	return nil
}

func ClusterUpdaterDelete(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return diag.FromErr(schema.RemoveFromState(d, m))
}
