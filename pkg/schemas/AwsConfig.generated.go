package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func AwsConfig() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"profile":     OptionalString(),
			"region":      OptionalString(),
			"assume_role": OptionalStruct(AwsAssumeRole()),
		},
	}
}
