package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceAzureConfiguration() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"subscription_id":     Optional(String()),
			"tenant_id":           Optional(String()),
			"resource_group_name": Optional(String()),
			"route_table_name":    Optional(String()),
			"admin_user":          Optional(String()),
		},
	}
}

func ExpandResourceAzureConfiguration(in map[string]interface{}) kops.AzureConfiguration {
	if in == nil {
		panic("expand AzureConfiguration failure, in is nil")
	}
	out := kops.AzureConfiguration{}
	if in, ok := in["subscription_id"]; ok && in != nil {
		out.SubscriptionID = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["tenant_id"]; ok && in != nil {
		out.TenantID = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["resource_group_name"]; ok && in != nil {
		out.ResourceGroupName = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["route_table_name"]; ok && in != nil {
		out.RouteTableName = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["admin_user"]; ok && in != nil {
		out.AdminUser = func(in interface{}) string { return string(in.(string)) }(in)
	}
	return out
}
