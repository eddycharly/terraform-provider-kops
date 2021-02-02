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
			"cluster_name":   RequiredString(),
			"rolling_update": OptionalStruct(ResourceRollingUpdateOptions()),
		},
	}
}

func ExpandResourceClusterUpdater(in map[string]interface{}) resources.ClusterUpdater {
	if in == nil {
		panic("expand ClusterUpdater failure, in is nil")
	}
	return resources.ClusterUpdater{
		ClusterName: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["cluster_name"]),
		RollingUpdate: func(in interface{}) resources.RollingUpdateOptions {
			return func(in interface{}) resources.RollingUpdateOptions {
				if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
					return resources.RollingUpdateOptions{}
				}
				return (ExpandResourceRollingUpdateOptions(in.([]interface{})[0].(map[string]interface{})))
			}(in)
		}(in["rolling_update"]),
	}
}

func FlattenResourceClusterUpdaterInto(in resources.ClusterUpdater, out map[string]interface{}) {
	out["cluster_name"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.ClusterName)
	out["rolling_update"] = func(in resources.RollingUpdateOptions) interface{} {
		return func(in resources.RollingUpdateOptions) []map[string]interface{} {
			return []map[string]interface{}{FlattenResourceRollingUpdateOptions(in)}
		}(in)
	}(in.RollingUpdate)
}

func FlattenResourceClusterUpdater(in resources.ClusterUpdater) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceClusterUpdaterInto(in, out)
	return out
}
