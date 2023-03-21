package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceGCESpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"project":         ComputedString(),
			"service_account": ComputedString(),
			"pd_csi_driver":   ComputedStruct(DataSourcePDCSIDriver()),
		},
	}

	return res
}

func ExpandDataSourceGCESpec(in map[string]interface{}) kops.GCESpec {
	if in == nil {
		panic("expand GCESpec failure, in is nil")
	}
	return kops.GCESpec{
		Project: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["project"]),
		ServiceAccount: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["service_account"]),
		PDCSIDriver: func(in interface{}) *kops.PDCSIDriver {
			return func(in interface{}) *kops.PDCSIDriver {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.PDCSIDriver) *kops.PDCSIDriver {
					return &in
				}(func(in interface{}) kops.PDCSIDriver {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandDataSourcePDCSIDriver(in[0].(map[string]interface{}))
					}
					return kops.PDCSIDriver{}
				}(in))
			}(in)
		}(in["pd_csi_driver"]),
	}
}

func FlattenDataSourceGCESpecInto(in kops.GCESpec, out map[string]interface{}) {
	out["project"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Project)
	out["service_account"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.ServiceAccount)
	out["pd_csi_driver"] = func(in *kops.PDCSIDriver) interface{} {
		return func(in *kops.PDCSIDriver) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.PDCSIDriver) interface{} {
				return func(in kops.PDCSIDriver) []interface{} {
					return []interface{}{FlattenDataSourcePDCSIDriver(in)}
				}(in)
			}(*in)
		}(in)
	}(in.PDCSIDriver)
}

func FlattenDataSourceGCESpec(in kops.GCESpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceGCESpecInto(in, out)
	return out
}
