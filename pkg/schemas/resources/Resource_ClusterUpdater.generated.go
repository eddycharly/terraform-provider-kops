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
			"cluster_name": RequiredString(),
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
	}
}

func FlattenResourceClusterUpdaterInto(in resources.ClusterUpdater, out map[string]interface{}) {
	out["cluster_name"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.ClusterName)
}

func FlattenResourceClusterUpdater(in resources.ClusterUpdater) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceClusterUpdaterInto(in, out)
	return out
}
