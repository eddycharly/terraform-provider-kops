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
			"master_interval":     Optional(Nullable(Duration())),
			"node_interval":       Optional(Nullable(Duration())),
			"bastion_interval":    Optional(Nullable(Duration())),
			"fail_on_drain_error": Optional(Bool()),
			"fail_on_validate":    Optional(Bool()),
			"post_drain_delay":    Optional(Nullable(Duration())),
			"validation_timeout":  Optional(Nullable(Duration())),
			"validate_count":      Optional(Nullable(Int())),
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

func FlattenResourceRollingUpdateOptionsInto(in resources.RollingUpdateOptions, out map[string]interface{}) {
	out["skip"] = func(in bool) interface{} { return in }(in.Skip)
	utilsschemas.FlattenResourceRollingUpdateOptionsInto(in.RollingUpdateOptions, out)
}

func FlattenResourceRollingUpdateOptions(in resources.RollingUpdateOptions) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceRollingUpdateOptionsInto(in, out)
	return out
}
