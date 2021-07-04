package schemas

import (
	"github.com/eddycharly/terraform-provider-kops/pkg/api/resources"
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	utilsschemas "github.com/eddycharly/terraform-provider-kops/pkg/schemas/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var _ = Schema

func ResourceValidateOptions() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"skip":          Optional(Bool()),
			"timeout":       Optional(Nullable(Duration())),
			"poll_interval": Optional(Nullable(Duration())),
		},
	}
}

func ExpandResourceValidateOptions(in map[string]interface{}) resources.ValidateOptions {
	if in == nil {
		panic("expand ValidateOptions failure, in is nil")
	}
	out := resources.ValidateOptions{}
	if in, ok := in["skip"]; ok && in != nil {
		out.Skip = func(in interface{}) bool { return in.(bool) }(in)
	}
	out.ValidateOptions = utilsschemas.ExpandResourceValidateOptions(in)
	return out
}

func FlattenResourceValidateOptionsInto(in resources.ValidateOptions, out map[string]interface{}) {
	out["skip"] = func(in bool) interface{} { return in }(in.Skip)
	utilsschemas.FlattenResourceValidateOptionsInto(in.ValidateOptions, out)
}

func FlattenResourceValidateOptions(in resources.ValidateOptions) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceValidateOptionsInto(in, out)
	return out
}
