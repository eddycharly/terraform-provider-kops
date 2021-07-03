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
			"cluster_name":    Required(String()),
			"exists":          Computed(Bool()),
			"is_valid":        Computed(Bool()),
			"needs_update":    Computed(Bool()),
			"instance_groups": Computed(List(String())),
		},
	}
}

func ExpandDataSourceClusterStatus(in map[string]interface{}) datasources.ClusterStatus {
	if in == nil {
		panic("expand ClusterStatus failure, in is nil")
	}
	out := datasources.ClusterStatus{}
	if in, ok := in["cluster_name"]; ok && in != nil {
		out.ClusterName = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["exists"]; ok && in != nil {
		out.Exists = func(in interface{}) bool { return in.(bool) }(in)
	}
	if in, ok := in["is_valid"]; ok && in != nil {
		out.IsValid = func(in interface{}) bool { return in.(bool) }(in)
	}
	if in, ok := in["needs_update"]; ok && in != nil {
		out.NeedsUpdate = func(in interface{}) bool { return in.(bool) }(in)
	}
	if in, ok := in["instance_groups"]; ok && in != nil {
		out.InstanceGroups = func(in interface{}) []string {
			var out []string
			for _, in := range in.([]interface{}) {
				out = append(out, func(in interface{}) string { return string(in.(string)) }(in))
			}
			return out
		}(in)
	}
	return out
}
