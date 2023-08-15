package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

var _ = Schema

func DataSourceAzureSpec() *schema.Resource {
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

func ExpandDataSourceAzureSpec(in map[string]interface{}) kopsv1alpha2.AzureSpec {
	if in == nil {
		panic("expand AzureSpec failure, in is nil")
	}
	return kopsv1alpha2.AzureSpec{
		SubscriptionID: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["subscription_id"]),
		TenantID: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["tenant_id"]),
		ResourceGroupName: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["resource_group_name"]),
		RouteTableName: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["route_table_name"]),
		AdminUser: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["admin_user"]),
	}
}

func FlattenDataSourceAzureSpecInto(in kopsv1alpha2.AzureSpec, out map[string]interface{}) {
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

func FlattenDataSourceAzureSpec(in kopsv1alpha2.AzureSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceAzureSpecInto(in, out)
	return out
}
