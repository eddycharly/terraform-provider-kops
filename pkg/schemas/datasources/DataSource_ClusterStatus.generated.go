package schemas

import (
	"github.com/eddycharly/terraform-provider-kops/pkg/api/datasources"
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var _ = Schema

func DataSourceClusterStatus() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cluster_name": RequiredString(),
			"exists":       ComputedBool(),
			"is_valid":     ComputedBool(),
			"needs_update": ComputedBool(),
		},
	}
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
}

func FlattenDataSourceClusterStatus(in datasources.ClusterStatus) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceClusterStatusInto(in, out)
	return out
}
