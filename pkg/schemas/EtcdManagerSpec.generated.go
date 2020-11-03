package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func EtcdManagerSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"image": RequiredString(),
			"env":   OptionalList(EnvVar()),
		},
	}
}
