package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

var _ = Schema

func DataSourcePodIdentityWebhookSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enabled":  ComputedBool(),
			"replicas": ComputedInt(),
		},
	}

	return res
}

func ExpandDataSourcePodIdentityWebhookSpec(in map[string]interface{}) kopsv1alpha2.PodIdentityWebhookSpec {
	if in == nil {
		panic("expand PodIdentityWebhookSpec failure, in is nil")
	}
	return kopsv1alpha2.PodIdentityWebhookSpec{
		Enabled: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["enabled"]),
		Replicas: func(in interface{}) int {
			return int(ExpandInt(in))
		}(in["replicas"]),
	}
}

func FlattenDataSourcePodIdentityWebhookSpecInto(in kopsv1alpha2.PodIdentityWebhookSpec, out map[string]interface{}) {
	out["enabled"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.Enabled)
	out["replicas"] = func(in int) interface{} {
		return FlattenInt(int(in))
	}(in.Replicas)
}

func FlattenDataSourcePodIdentityWebhookSpec(in kopsv1alpha2.PodIdentityWebhookSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourcePodIdentityWebhookSpecInto(in, out)
	return out
}
