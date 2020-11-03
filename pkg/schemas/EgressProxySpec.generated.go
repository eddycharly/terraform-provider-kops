package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func EgressProxySpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"http_proxy":     RequiredStruct(HTTPProxy()),
			"proxy_excludes": OptionalString(),
		},
	}
}
