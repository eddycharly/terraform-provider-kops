package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func FileAssetSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name":       StringRequired(),
			"path":       StringRequired(),
			"roles":      ListRequired(String()),
			"content":    StringRequired(),
			"is_base_64": BoolOptionalComputed(),
		},
	}
}
