package schemas

import (
	"reflect"

	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

var _ = Schema

func DataSourceSnapshotControllerConfig() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enabled":               ComputedBool(),
			"install_default_class": ComputedBool(),
		},
	}

	return res
}

func ExpandDataSourceSnapshotControllerConfig(in map[string]interface{}) kopsv1alpha2.SnapshotControllerConfig {
	if in == nil {
		panic("expand SnapshotControllerConfig failure, in is nil")
	}
	return kopsv1alpha2.SnapshotControllerConfig{
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
		InstallDefaultClass: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["install_default_class"]),
	}
}

func FlattenDataSourceSnapshotControllerConfigInto(in kopsv1alpha2.SnapshotControllerConfig, out map[string]interface{}) {
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
	out["install_default_class"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.InstallDefaultClass)
}

func FlattenDataSourceSnapshotControllerConfig(in kopsv1alpha2.SnapshotControllerConfig) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceSnapshotControllerConfigInto(in, out)
	return out
}
