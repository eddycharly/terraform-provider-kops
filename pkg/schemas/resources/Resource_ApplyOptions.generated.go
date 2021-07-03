package schemas

import (
	"github.com/eddycharly/terraform-provider-kops/pkg/api/resources"
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var _ = Schema

func ResourceApplyOptions() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"skip":                 Optional(Bool()),
			"allow_kops_downgrade": Optional(Bool()),
		},
	}
}

func ExpandResourceApplyOptions(in map[string]interface{}) resources.ApplyOptions {
	if in == nil {
		panic("expand ApplyOptions failure, in is nil")
	}
	out := resources.ApplyOptions{}
	if in, ok := in["skip"]; ok && in != nil {
		out.Skip = func(in interface{}) bool { return in.(bool) }(in)
	}
	if in, ok := in["allow_kops_downgrade"]; ok && in != nil {
		out.AllowKopsDowngrade = func(in interface{}) bool { return in.(bool) }(in)
	}
	return out
}
