package schemas

import (
	"reflect"

	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

var _ = Schema

func DataSourceNvidiaGPUConfig() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"driver_package": ComputedString(),
			"enabled":        ComputedBool(),
			"dcgm_exporter":  ComputedStruct(DataSourceDCGMExporterConfig()),
		},
	}

	return res
}

func ExpandDataSourceNvidiaGPUConfig(in map[string]interface{}) kopsv1alpha2.NvidiaGPUConfig {
	if in == nil {
		panic("expand NvidiaGPUConfig failure, in is nil")
	}
	return kopsv1alpha2.NvidiaGPUConfig{
		DriverPackage: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["driver_package"]),
		Enabled: func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in bool) *bool {
					return &in
				}(bool(ExpandBool(in)))
			}(in)
		}(in["enabled"]),
		DCGMExporter: func(in interface{}) *kopsv1alpha2.DCGMExporterConfig {
			return func(in interface{}) *kopsv1alpha2.DCGMExporterConfig {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.DCGMExporterConfig) *kopsv1alpha2.DCGMExporterConfig {
					return &in
				}(func(in interface{}) kopsv1alpha2.DCGMExporterConfig {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandDataSourceDCGMExporterConfig(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.DCGMExporterConfig{}
				}(in))
			}(in)
		}(in["dcgm_exporter"]),
	}
}

func FlattenDataSourceNvidiaGPUConfigInto(in kopsv1alpha2.NvidiaGPUConfig, out map[string]interface{}) {
	out["driver_package"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.DriverPackage)
	out["enabled"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.Enabled)
	out["dcgm_exporter"] = func(in *kopsv1alpha2.DCGMExporterConfig) interface{} {
		return func(in *kopsv1alpha2.DCGMExporterConfig) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.DCGMExporterConfig) interface{} {
				return func(in kopsv1alpha2.DCGMExporterConfig) []interface{} {
					return []interface{}{FlattenDataSourceDCGMExporterConfig(in)}
				}(in)
			}(*in)
		}(in)
	}(in.DCGMExporter)
}

func FlattenDataSourceNvidiaGPUConfig(in kopsv1alpha2.NvidiaGPUConfig) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceNvidiaGPUConfigInto(in, out)
	return out
}
