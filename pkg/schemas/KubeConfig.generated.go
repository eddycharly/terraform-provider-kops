package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func KubeConfig() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"server":            ComputedString(),
			"context":           ComputedString(),
			"namespace":         ComputedString(),
			"kube_bearer_token": Sensitive(ComputedString()),
			"kube_user":         Sensitive(ComputedString()),
			"kube_password":     Sensitive(ComputedString()),
			"ca_cert":           Sensitive(ComputedString()),
			"client_cert":       Sensitive(ComputedString()),
			"client_key":        Sensitive(ComputedString()),
		},
	}
}
