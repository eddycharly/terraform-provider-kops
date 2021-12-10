package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceAzureConfiguration() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"subscription_id":     ComputedString(),
			"tenant_id":           ComputedString(),
			"resource_group_name": ComputedString(),
			"route_table_name":    ComputedString(),
			"admin_user":          ComputedString(),
		},
	}

	return res
}

func ExpandDataSourceAzureConfiguration(in map[string]interface{}) kops.AzureConfiguration {
	if in == nil {
		panic("expand AzureConfiguration failure, in is nil")
	}
	return kops.AzureConfiguration{
		SubscriptionID: func(in interface{}) string /**/ {
			return string(ExpandString(in))
		}(in["subscription_id"]),
		TenantID: func(in interface{}) string /**/ {
			return string(ExpandString(in))
		}(in["tenant_id"]),
		ResourceGroupName: func(in interface{}) string /**/ {
			return string(ExpandString(in))
		}(in["resource_group_name"]),
		RouteTableName: func(in interface{}) string /**/ {
			return string(ExpandString(in))
		}(in["route_table_name"]),
		AdminUser: func(in interface{}) string /**/ {
			return string(ExpandString(in))
		}(in["admin_user"]),
	}
}

func FlattenDataSourceAzureConfigurationInto(in kops.AzureConfiguration, out map[string]interface{}) {
	out["subscription_id"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.SubscriptionID)
	out["tenant_id"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.TenantID)
	out["resource_group_name"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.ResourceGroupName)
	out["route_table_name"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.RouteTableName)
	out["admin_user"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.AdminUser)
}

func FlattenDataSourceAzureConfiguration(in kops.AzureConfiguration) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceAzureConfigurationInto(in, out)
	return out
}
