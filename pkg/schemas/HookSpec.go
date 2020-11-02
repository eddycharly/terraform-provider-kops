package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func HookSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name":             StringRequired(),
			"disabled":         BoolOptionalComputed(),
			"roles":            ListOptional(String()),
			"requires":         ListOptional(String()),
			"before":           ListOptional(String()),
			"exec_container":   StructOptionalComputed(ExecContainerAction()),
			"manifest":         StringOptionalComputed(),
			"use_raw_manifest": BoolOptionalComputed(),
		},
	}
}
