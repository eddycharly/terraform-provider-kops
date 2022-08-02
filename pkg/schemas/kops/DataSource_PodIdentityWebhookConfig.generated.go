package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourcePodIdentityWebhookConfig() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enabled":  ComputedBool(),
			"replicas": ComputedInt(),
		},
	}

	return res
}

func ExpandDataSourcePodIdentityWebhookConfig(in map[string]interface{}) kops.PodIdentityWebhookConfig {
	if in == nil {
		panic("expand PodIdentityWebhookConfig failure, in is nil")
	}
	return kops.PodIdentityWebhookConfig{
		Enabled: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["enabled"]),
		Replicas: func(in interface{}) int {
			return int(ExpandInt(in))
		}(in["replicas"]),
	}
}

func FlattenDataSourcePodIdentityWebhookConfigInto(in kops.PodIdentityWebhookConfig, out map[string]interface{}) {
	out["enabled"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.Enabled)
	out["replicas"] = func(in int) interface{} {
		return FlattenInt(int(in))
	}(in.Replicas)
}

func FlattenDataSourcePodIdentityWebhookConfig(in kops.PodIdentityWebhookConfig) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourcePodIdentityWebhookConfigInto(in, out)
	return out
}
