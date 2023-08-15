package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

var _ = Schema

func ResourceKarpenterConfig() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enabled": OptionalBool(),
		},
	}

	return res
}

func ExpandResourceKarpenterConfig(in map[string]interface{}) kopsv1alpha2.KarpenterConfig {
	if in == nil {
		panic("expand KarpenterConfig failure, in is nil")
	}
	return kopsv1alpha2.KarpenterConfig{
		Enabled: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["enabled"]),
	}
}

func FlattenResourceKarpenterConfigInto(in kopsv1alpha2.KarpenterConfig, out map[string]interface{}) {
	out["enabled"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.Enabled)
}

func FlattenResourceKarpenterConfig(in kopsv1alpha2.KarpenterConfig) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceKarpenterConfigInto(in, out)
	return out
}
