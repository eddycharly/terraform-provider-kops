package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func AwsAssumeRole() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"role_arn": OptionalString(),
		},
	}
}
