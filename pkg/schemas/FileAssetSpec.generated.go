package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func FileAssetSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name":       RequiredString(),
			"path":       RequiredString(),
			"roles":      OptionalList(String()),
			"content":    RequiredString(),
			"is_base_64": OptionalBool(),
		},
	}
}
