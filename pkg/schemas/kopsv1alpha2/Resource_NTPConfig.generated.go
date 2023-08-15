package schemas

import (
	"reflect"

	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

var _ = Schema

func ResourceNTPConfig() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"managed": OptionalBool(),
		},
	}

	return res
}

func ExpandResourceNTPConfig(in map[string]interface{}) kopsv1alpha2.NTPConfig {
	if in == nil {
		panic("expand NTPConfig failure, in is nil")
	}
	return kopsv1alpha2.NTPConfig{
		Managed: func(in interface{}) *bool {
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
		}(in["managed"]),
	}
}

func FlattenResourceNTPConfigInto(in kopsv1alpha2.NTPConfig, out map[string]interface{}) {
	out["managed"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.Managed)
}

func FlattenResourceNTPConfig(in kopsv1alpha2.NTPConfig) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceNTPConfigInto(in, out)
	return out
}
