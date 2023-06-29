package schemas

import (
	"reflect"

	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

var _ = Schema

func ResourcePDCSIDriver() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enabled": OptionalBool(),
		},
	}

	return res
}

func ExpandResourcePDCSIDriver(in map[string]interface{}) kopsv1alpha2.PDCSIDriver {
	if in == nil {
		panic("expand PDCSIDriver failure, in is nil")
	}
	return kopsv1alpha2.PDCSIDriver{
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

func FlattenResourcePDCSIDriverInto(in kopsv1alpha2.PDCSIDriver, out map[string]interface{}) {
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

func FlattenResourcePDCSIDriver(in kopsv1alpha2.PDCSIDriver) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourcePDCSIDriverInto(in, out)
	return out
}
