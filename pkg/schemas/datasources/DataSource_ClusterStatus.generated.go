package schemas

import (
	"github.com/eddycharly/terraform-provider-kops/pkg/api/datasources"
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var _ = Schema

func DataSourceClusterStatus() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cluster_name":    RequiredString(),
			"exists":          ComputedBool(),
			"is_valid":        ComputedBool(),
			"needs_update":    ComputedBool(),
			"instance_groups": ComputedList(String()),
		},
	}

	return res
}

func ExpandDataSourceClusterStatus(in map[string]interface{}) datasources.ClusterStatus {
	if in == nil {
		panic("expand ClusterStatus failure, in is nil")
	}
	return datasources.ClusterStatus{
		ClusterName: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["cluster_name"]),
		Exists: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["exists"]),
		IsValid: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["is_valid"]),
		NeedsUpdate: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["needs_update"]),
		InstanceGroups: func(in interface{}) []string {
			return func(in interface{}) []string {
				if in == nil {
					return nil
				}
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["instance_groups"]),
	}
}

func FlattenDataSourceClusterStatusInto(in datasources.ClusterStatus, out map[string]interface{}) {
	out["cluster_name"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.ClusterName)
	out["exists"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.Exists)
	out["is_valid"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.IsValid)
	out["needs_update"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.NeedsUpdate)
	out["instance_groups"] = func(in []string) interface{} {
		return func(in []string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.InstanceGroups)
}

func FlattenDataSourceClusterStatus(in datasources.ClusterStatus) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceClusterStatusInto(in, out)
	return out
}
