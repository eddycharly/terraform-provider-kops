package schemas

import (
	"reflect"

	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

var _ = Schema

func ResourceRunc() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"version":  OptionalString(),
			"packages": OptionalStruct(ResourcePackagesConfig()),
		},
	}

	return res
}

func ExpandResourceRunc(in map[string]interface{}) kopsv1alpha2.Runc {
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
						return ExpandResourcePackagesConfig(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.PackagesConfig{}
				}(in))
			}(in)
		}(in["packages"]),
	}
}

func FlattenResourceRuncInto(in kopsv1alpha2.Runc, out map[string]interface{}) {
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
					return []interface{}{FlattenResourcePackagesConfig(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Packages)
}

func FlattenResourceRunc(in kopsv1alpha2.Runc) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceRuncInto(in, out)
	return out
}
