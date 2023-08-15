package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

var _ = Schema

func ResourceDCGMExporterConfig() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enabled": OptionalBool(),
		},
	}

	return res
}

func ExpandResourceDCGMExporterConfig(in map[string]interface{}) kopsv1alpha2.DCGMExporterConfig {
	if in == nil {
		panic("expand DCGMExporterConfig failure, in is nil")
	}
	return kopsv1alpha2.DCGMExporterConfig{
		Enabled: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["enabled"]),
	}
}

func FlattenResourceDCGMExporterConfigInto(in kopsv1alpha2.DCGMExporterConfig, out map[string]interface{}) {
	out["enabled"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.Enabled)
}

func FlattenResourceDCGMExporterConfig(in kopsv1alpha2.DCGMExporterConfig) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceDCGMExporterConfigInto(in, out)
	return out
}
