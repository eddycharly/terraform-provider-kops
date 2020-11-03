package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func AuthenticationSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"kopeio": OptionalStruct(KopeioAuthenticationSpec()),
			"aws":    OptionalStruct(AwsAuthenticationSpec()),
		},
	}
}
