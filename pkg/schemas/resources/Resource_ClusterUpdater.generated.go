package schemas

import (
	"github.com/eddycharly/terraform-provider-kops/pkg/api/resources"
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var _ = Schema

func ResourceClusterUpdater() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"revision":       Computed(Int()),
			"cluster_name":   Required(String()),
			"keepers":        Optional(Map(String())),
			"apply":          Optional(Struct(ResourceApplyOptions())),
			"rolling_update": Optional(Struct(ResourceRollingUpdateOptions())),
			"validate":       Optional(Struct(ResourceValidateOptions())),
		},
	}
}

func ExpandResourceClusterUpdater(in map[string]interface{}) resources.ClusterUpdater {
	if in == nil {
		panic("expand ClusterUpdater failure, in is nil")
	}
	out := resources.ClusterUpdater{}
	if in, ok := in["revision"]; ok && in != nil {
		out.Revision = func(in interface{}) int { return int(in.(int)) }(in)
	}
	if in, ok := in["cluster_name"]; ok && in != nil {
		out.ClusterName = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["keepers"]; ok && in != nil {
		out.Keepers = func(in interface{}) map[string]string {
			if in == nil {
				return nil
			}
			out := map[string]string{}
			for key, in := range in.(map[string]interface{}) {
				out[key] = func(in interface{}) string { return string(in.(string)) }(in)
			}
			return out
		}(in)
	}
	if in, ok := in["apply"]; ok && in != nil {
		out.Apply = func(in interface{}) resources.ApplyOptions {
			if in == nil {
				return resources.ApplyOptions{}
			}
			return ExpandResourceApplyOptions(in.(map[string]interface{}))
		}(in)
	}
	if in, ok := in["rolling_update"]; ok && in != nil {
		out.RollingUpdate = func(in interface{}) resources.RollingUpdateOptions {
			if in == nil {
				return resources.RollingUpdateOptions{}
			}
			return ExpandResourceRollingUpdateOptions(in.(map[string]interface{}))
		}(in)
	}
	if in, ok := in["validate"]; ok && in != nil {
		out.Validate = func(in interface{}) resources.ValidateOptions {
			if in == nil {
				return resources.ValidateOptions{}
			}
			return ExpandResourceValidateOptions(in.(map[string]interface{}))
		}(in)
	}
	return out
}

func FlattenResourceClusterUpdaterInto(in resources.ClusterUpdater, out map[string]interface{}) {
	out["revision"] = func(in int) interface{} { return int(in) }(in.Revision)
	out["cluster_name"] = func(in string) interface{} { return string(in) }(in.ClusterName)
	out["keepers"] = func(in map[string]string) interface{} {
		if in == nil {
			return nil
		}
		out := map[string]interface{}{}
		for key, in := range in {
			out[key] = func(in string) interface{} { return string(in) }(in)
		}
		return out
	}(in.Keepers)
	out["apply"] = func(in resources.ApplyOptions) interface{} { return FlattenResourceApplyOptions(in) }(in.Apply)
	out["rolling_update"] = func(in resources.RollingUpdateOptions) interface{} { return FlattenResourceRollingUpdateOptions(in) }(in.RollingUpdate)
	out["validate"] = func(in resources.ValidateOptions) interface{} { return FlattenResourceValidateOptions(in) }(in.Validate)
}

func FlattenResourceClusterUpdater(in resources.ClusterUpdater) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceClusterUpdaterInto(in, out)
	return out
}
