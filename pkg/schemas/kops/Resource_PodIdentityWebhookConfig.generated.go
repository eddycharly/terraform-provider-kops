package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourcePodIdentityWebhookConfig() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enabled": OptionalBool(),
		},
	}

	return res
}

func ExpandResourcePodIdentityWebhookConfig(in map[string]interface{}) kops.PodIdentityWebhookConfig {
	if in == nil {
		panic("expand PodIdentityWebhookConfig failure, in is nil")
	}
	return kops.PodIdentityWebhookConfig{
		Enabled: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["enabled"]),
	}
}

func FlattenResourcePodIdentityWebhookConfigInto(in kops.PodIdentityWebhookConfig, out map[string]interface{}) {
	out["enabled"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.Enabled)
}

func FlattenResourcePodIdentityWebhookConfig(in kops.PodIdentityWebhookConfig) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourcePodIdentityWebhookConfigInto(in, out)
	return out
}
