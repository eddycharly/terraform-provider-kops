package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
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

func ExpandResourceKarpenterConfig(in map[string]interface{}) kops.KarpenterConfig {
	if in == nil {
		panic("expand KarpenterConfig failure, in is nil")
	}
	return kops.KarpenterConfig{
		Enabled: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["enabled"]),
	}
}

func FlattenResourceKarpenterConfigInto(in kops.KarpenterConfig, out map[string]interface{}) {
	out["enabled"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.Enabled)
}

func FlattenResourceKarpenterConfig(in kops.KarpenterConfig) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceKarpenterConfigInto(in, out)
	return out
}
