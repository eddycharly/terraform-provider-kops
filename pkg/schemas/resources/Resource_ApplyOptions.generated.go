package schemas

import (
	"github.com/eddycharly/terraform-provider-kops/pkg/api/resources"
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var _ = Schema

func ResourceApplyOptions() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"skip":                 OptionalBool(),
			"allow_kops_downgrade": OptionalBool(),
			"lifecycle_overrides":  OptionalMap(String()),
		},
	}

	return res
}

func ExpandResourceApplyOptions(in map[string]interface{}) resources.ApplyOptions {
	if in == nil {
		panic("expand ApplyOptions failure, in is nil")
	}
	return resources.ApplyOptions{
		Skip: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["skip"]),
		AllowKopsDowngrade: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["allow_kops_downgrade"]),
		LifecycleOverrides: func(in interface{}) map[string]string {
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
		}(in["lifecycle_overrides"]),
	}
}

func FlattenResourceApplyOptionsInto(in resources.ApplyOptions, out map[string]interface{}) {
	out["skip"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.Skip)
	out["allow_kops_downgrade"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.AllowKopsDowngrade)
	out["lifecycle_overrides"] = func(in map[string]string) interface{} {
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
	}(in.LifecycleOverrides)
}

func FlattenResourceApplyOptions(in resources.ApplyOptions) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceApplyOptionsInto(in, out)
	return out
}
