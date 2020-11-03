package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func AuthorizationSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"always_allow": OptionalStruct(AlwaysAllowAuthorizationSpec()),
			"rbac":         OptionalStruct(RBACAuthorizationSpec()),
		},
	}
}
