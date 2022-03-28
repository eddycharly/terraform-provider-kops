package schemas

import (
	"reflect"

	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceGCPPDCSIDriver() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enabled": OptionalBool(),
		},
	}

	return res
}

func ExpandResourceGCPPDCSIDriver(in map[string]interface{}) kops.GCPPDCSIDriver {
	if in == nil {
		panic("expand GCPPDCSIDriver failure, in is nil")
	}
	return kops.GCPPDCSIDriver{
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
	}
}

func FlattenResourceGCPPDCSIDriverInto(in kops.GCPPDCSIDriver, out map[string]interface{}) {
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
}

func FlattenResourceGCPPDCSIDriver(in kops.GCPPDCSIDriver) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceGCPPDCSIDriverInto(in, out)
	return out
}
