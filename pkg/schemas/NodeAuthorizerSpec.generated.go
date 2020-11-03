package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func NodeAuthorizerSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"authorizer": OptionalString(),
			"features":   OptionalList(String()),
			"image":      OptionalString(),
			"node_url":   OptionalString(),
			"port":       OptionalInt(),
			"interval":   OptionalDuration(),
			"timeout":    OptionalDuration(),
			"token_ttl":  OptionalDuration(),
		},
	}
}
