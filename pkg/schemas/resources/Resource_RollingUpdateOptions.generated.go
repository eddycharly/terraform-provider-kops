package schemas

import (
	"github.com/eddycharly/terraform-provider-kops/pkg/api/resources"
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	utilsschemas "github.com/eddycharly/terraform-provider-kops/pkg/schemas/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var _ = Schema

func ResourceRollingUpdateOptions() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"skip":                Optional(Bool()),
			"master_interval":     Optional(Ptr(Duration())),
			"node_interval":       Optional(Ptr(Duration())),
			"bastion_interval":    Optional(Ptr(Duration())),
			"fail_on_drain_error": Optional(Bool()),
			"fail_on_validate":    Optional(Bool()),
			"post_drain_delay":    Optional(Ptr(Duration())),
			"validation_timeout":  Optional(Ptr(Duration())),
			"validate_count":      Optional(Ptr(Int())),
			"cloud_only":          Optional(Bool()),
		},
	}
}

func ExpandResourceRollingUpdateOptions(in map[string]interface{}) resources.RollingUpdateOptions {
	if in == nil {
		panic("expand RollingUpdateOptions failure, in is nil")
	}
	out := resources.RollingUpdateOptions{}
	if in, ok := in["skip"]; ok && in != nil {
		out.Skip = func(in interface{}) bool { return in.(bool) }(in)
	}
	out.RollingUpdateOptions = utilsschemas.ExpandResourceRollingUpdateOptions(in)
	return out
}
