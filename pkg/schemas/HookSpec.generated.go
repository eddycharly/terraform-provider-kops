package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func HookSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name":             RequiredString(),
			"disabled":         OptionalBool(),
			"roles":            OptionalList(String()),
			"requires":         OptionalList(String()),
			"before":           OptionalList(String()),
			"exec_container":   OptionalStruct(ExecContainerAction()),
			"manifest":         OptionalString(),
			"use_raw_manifest": OptionalBool(),
		},
	}
}
