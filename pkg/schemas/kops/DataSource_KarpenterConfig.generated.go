package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceKarpenterConfig() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enabled": ComputedBool(),
		},
	}

	return res
}

func ExpandDataSourceKarpenterConfig(in map[string]interface{}) kops.KarpenterConfig {
	if in == nil {
		panic("expand KarpenterConfig failure, in is nil")
	}
	return kops.KarpenterConfig{
		Enabled: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["enabled"]),
	}
}

func FlattenDataSourceKarpenterConfigInto(in kops.KarpenterConfig, out map[string]interface{}) {
	out["enabled"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.Enabled)
}

func FlattenDataSourceKarpenterConfig(in kops.KarpenterConfig) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceKarpenterConfigInto(in, out)
	return out
}
