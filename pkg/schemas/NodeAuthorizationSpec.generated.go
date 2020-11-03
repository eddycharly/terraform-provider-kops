package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func NodeAuthorizationSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"node_authorizer": OptionalStruct(NodeAuthorizerSpec()),
		},
	}
}
