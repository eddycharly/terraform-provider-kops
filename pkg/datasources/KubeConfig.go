package datasources

import (
	"context"

	"github.com/eddycharly/terraform-provider-kops/pkg/config"
	datasourcesschemas "github.com/eddycharly/terraform-provider-kops/pkg/schemas/datasources"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func KubeConfig() *schema.Resource {
	return &schema.Resource{
		ReadContext: KubeConfigRead,
		Schema:      datasourcesschemas.DataSourceKubeConfig().Schema,
	}
}

func KubeConfigRead(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	in := datasourcesschemas.ExpandDataSourceKubeConfig(d.Get("").(map[string]interface{}))
	err := in.GetKubeConfig(config.Clientset(m))
	if err != nil {
		return diag.FromErr(err)
	}
	for k, v := range datasourcesschemas.FlattenDataSourceKubeConfig(in) {
		if err := d.Set(k, v); err != nil {
			return diag.FromErr(err)
		}
	}
	d.SetId("-")
	return nil
}
