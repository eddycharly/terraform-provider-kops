package schemas

import (
	"reflect"

	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

var _ = Schema

func DataSourceRunc() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"version":  ComputedString(),
			"packages": ComputedStruct(DataSourcePackagesConfig()),
		},
	}

	return res
}

func ExpandDataSourceRunc(in map[string]interface{}) kopsv1alpha2.Runc {
	if in == nil {
		panic("expand Runc failure, in is nil")
	}
	return kopsv1alpha2.Runc{
		Version: func(in interface{}) *string {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in string) *string {
					return &in
				}(string(ExpandString(in)))
			}(in)
		}(in["version"]),
		Packages: func(in interface{}) *kopsv1alpha2.PackagesConfig {
			return func(in interface{}) *kopsv1alpha2.PackagesConfig {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.PackagesConfig) *kopsv1alpha2.PackagesConfig {
					return &in
				}(func(in interface{}) kopsv1alpha2.PackagesConfig {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandDataSourcePackagesConfig(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.PackagesConfig{}
				}(in))
			}(in)
		}(in["packages"]),
	}
}

func FlattenDataSourceRuncInto(in kopsv1alpha2.Runc, out map[string]interface{}) {
	out["version"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.Version)
	out["packages"] = func(in *kopsv1alpha2.PackagesConfig) interface{} {
		return func(in *kopsv1alpha2.PackagesConfig) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.PackagesConfig) interface{} {
				return func(in kopsv1alpha2.PackagesConfig) []interface{} {
					return []interface{}{FlattenDataSourcePackagesConfig(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Packages)
}

func FlattenDataSourceRunc(in kopsv1alpha2.Runc) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceRuncInto(in, out)
	return out
}
