package schemas

import (
	"reflect"

	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
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

func ExpandDataSourceNvidiaGPUConfig(in map[string]interface{}) kops.NvidiaGPUConfig {
	if in == nil {
		panic("expand NvidiaGPUConfig failure, in is nil")
	}
	return kops.NvidiaGPUConfig{
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
		DCGMExporter: func(in interface{}) *kops.DCGMExporterConfig {
			return func(in interface{}) *kops.DCGMExporterConfig {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.DCGMExporterConfig) *kops.DCGMExporterConfig {
					return &in
				}(func(in interface{}) kops.DCGMExporterConfig {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandDataSourceDCGMExporterConfig(in[0].(map[string]interface{}))
					}
					return kops.DCGMExporterConfig{}
				}(in))
			}(in)
		}(in["dcgm_exporter"]),
	}
}

func FlattenDataSourceNvidiaGPUConfigInto(in kops.NvidiaGPUConfig, out map[string]interface{}) {
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
	out["dcgm_exporter"] = func(in *kops.DCGMExporterConfig) interface{} {
		return func(in *kops.DCGMExporterConfig) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.DCGMExporterConfig) interface{} {
				return func(in kops.DCGMExporterConfig) []interface{} {
					return []interface{}{FlattenDataSourceDCGMExporterConfig(in)}
				}(in)
			}(*in)
		}(in)
	}(in.DCGMExporter)
}

func FlattenDataSourceNvidiaGPUConfig(in kops.NvidiaGPUConfig) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceNvidiaGPUConfigInto(in, out)
	return out
}
