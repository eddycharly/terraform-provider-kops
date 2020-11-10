package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ProviderConfig() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"state_store":            RequiredString(),
			"aws":                    OptionalStruct(AwsConfig()),
			"rolling_update_options": OptionalStruct(RollingUpdateOptions()),
			"validate_options":       OptionalStruct(ValidateOptions()),
		},
	}
}
