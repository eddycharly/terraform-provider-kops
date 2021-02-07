package datasources

import (
	"context"

	"github.com/eddycharly/terraform-provider-kops/pkg/config"
	datasourcesschemas "github.com/eddycharly/terraform-provider-kops/pkg/schemas/datasources"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ClusterStatus() *schema.Resource {
	return &schema.Resource{
		ReadContext: ClusterStatusRead,
		Schema:      datasourcesschemas.DataSourceClusterStatus().Schema,
	}
}

func ClusterStatusRead(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	in := datasourcesschemas.ExpandDataSourceClusterStatus(d.Get("").(map[string]interface{}))
	err := in.GetClusterStatus(config.Clientset(m))
	if err != nil {
		return diag.FromErr(err)
	}
	for k, v := range datasourcesschemas.FlattenDataSourceClusterStatus(in) {
		if err := d.Set(k, v); err != nil {
			return diag.FromErr(err)
		}
	}
	d.SetId("-")
	return nil
}
