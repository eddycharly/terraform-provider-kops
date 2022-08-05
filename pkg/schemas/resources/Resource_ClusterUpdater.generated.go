package schemas

import (
	"github.com/eddycharly/terraform-provider-kops/pkg/api/resources"
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var _ = Schema

func ResourceClusterUpdater() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"revision":         ComputedInt(),
			"provider_version": ComputedString(),
			"cluster_name":     RequiredString(),
			"keepers":          OptionalMap(String()),
			"apply":            OptionalStruct(ResourceApplyOptions()),
			"rolling_update":   OptionalStruct(ResourceRollingUpdateOptions()),
			"validate":         OptionalStruct(ResourceValidateOptions()),
		},
	}

	return res
}

func ExpandResourceClusterUpdater(in map[string]interface{}) resources.ClusterUpdater {
	if in == nil {
		panic("expand ClusterUpdater failure, in is nil")
	}
	return resources.ClusterUpdater{
		Revision: func(in interface{}) int {
			return int(ExpandInt(in))
		}(in["revision"]),
		ProviderVersion: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["provider_version"]),
		ClusterName: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["cluster_name"]),
		Keepers: func(in interface{}) map[string]string {
			return func(in interface{}) map[string]string {
				if in == nil {
					return nil
				}
				if in, ok := in.(map[string]interface{}); ok {
					if len(in) > 0 {
						out := map[string]string{}
						for key, in := range in {
							out[key] = string(ExpandString(in))
						}
						return out
					}
				}
				return nil
			}(in)
		}(in["keepers"]),
		Apply: func(in interface{}) resources.ApplyOptions {
			return func(in interface{}) resources.ApplyOptions {
				if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
					return ExpandResourceApplyOptions(in[0].(map[string]interface{}))
				}
				return resources.ApplyOptions{}
			}(in)
		}(in["apply"]),
		RollingUpdate: func(in interface{}) resources.RollingUpdateOptions {
			return func(in interface{}) resources.RollingUpdateOptions {
				if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
					return ExpandResourceRollingUpdateOptions(in[0].(map[string]interface{}))
				}
				return resources.RollingUpdateOptions{}
			}(in)
		}(in["rolling_update"]),
		Validate: func(in interface{}) resources.ValidateOptions {
			return func(in interface{}) resources.ValidateOptions {
				if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
					return ExpandResourceValidateOptions(in[0].(map[string]interface{}))
				}
				return resources.ValidateOptions{}
			}(in)
		}(in["validate"]),
	}
}

func FlattenResourceClusterUpdaterInto(in resources.ClusterUpdater, out map[string]interface{}) {
	out["revision"] = func(in int) interface{} {
		return FlattenInt(int(in))
	}(in.Revision)
	out["provider_version"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.ProviderVersion)
	out["cluster_name"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.ClusterName)
	out["keepers"] = func(in map[string]string) interface{} {
		return func(in map[string]string) map[string]interface{} {
			if in == nil {
				return nil
			}
			out := map[string]interface{}{}
			for key, in := range in {
				out[key] = FlattenString(string(in))
			}
			return out
		}(in)
	}(in.Keepers)
	out["apply"] = func(in resources.ApplyOptions) interface{} {
		return func(in resources.ApplyOptions) []interface{} {
			return []interface{}{FlattenResourceApplyOptions(in)}
		}(in)
	}(in.Apply)
	out["rolling_update"] = func(in resources.RollingUpdateOptions) interface{} {
		return func(in resources.RollingUpdateOptions) []interface{} {
			return []interface{}{FlattenResourceRollingUpdateOptions(in)}
		}(in)
	}(in.RollingUpdate)
	out["validate"] = func(in resources.ValidateOptions) interface{} {
		return func(in resources.ValidateOptions) []interface{} {
			return []interface{}{FlattenResourceValidateOptions(in)}
		}(in)
	}(in.Validate)
}

func FlattenResourceClusterUpdater(in resources.ClusterUpdater) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceClusterUpdaterInto(in, out)
	return out
}
