package schemas

import (
	"reflect"

	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

var _ = Schema

func ResourceIAMSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"legacy":                   OptionalBool(),
			"allow_container_registry": OptionalBool(),
			"permissions_boundary":     OptionalString(),
			"use_service_account_external_permissions": OptionalBool(),
			"service_account_external_permissions":     OptionalList(ResourceServiceAccountExternalPermission()),
		},
	}

	return res
}

func ExpandResourceIAMSpec(in map[string]interface{}) kopsv1alpha2.IAMSpec {
	if in == nil {
		panic("expand IAMSpec failure, in is nil")
	}
	return kopsv1alpha2.IAMSpec{
		Legacy: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["legacy"]),
		AllowContainerRegistry: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["allow_container_registry"]),
		PermissionsBoundary: func(in interface{}) *string {
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
		}(in["permissions_boundary"]),
		UseServiceAccountExternalPermissions: func(in interface{}) *bool {
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
		}(in["use_service_account_external_permissions"]),
		ServiceAccountExternalPermissions: func(in interface{}) []kopsv1alpha2.ServiceAccountExternalPermission {
			return func(in interface{}) []kopsv1alpha2.ServiceAccountExternalPermission {
				if in == nil {
					return nil
				}
				var out []kopsv1alpha2.ServiceAccountExternalPermission
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) kopsv1alpha2.ServiceAccountExternalPermission {
						if in == nil {
							return kopsv1alpha2.ServiceAccountExternalPermission{}
						}
						return ExpandResourceServiceAccountExternalPermission(in.(map[string]interface{}))
					}(in))
				}
				return out
			}(in)
		}(in["service_account_external_permissions"]),
	}
}

func FlattenResourceIAMSpecInto(in kopsv1alpha2.IAMSpec, out map[string]interface{}) {
	out["legacy"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.Legacy)
	out["allow_container_registry"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.AllowContainerRegistry)
	out["permissions_boundary"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.PermissionsBoundary)
	out["use_service_account_external_permissions"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.UseServiceAccountExternalPermissions)
	out["service_account_external_permissions"] = func(in []kopsv1alpha2.ServiceAccountExternalPermission) interface{} {
		return func(in []kopsv1alpha2.ServiceAccountExternalPermission) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, func(in kopsv1alpha2.ServiceAccountExternalPermission) interface{} {
					return FlattenResourceServiceAccountExternalPermission(in)
				}(in))
			}
			return out
		}(in)
	}(in.ServiceAccountExternalPermissions)
}

func FlattenResourceIAMSpec(in kopsv1alpha2.IAMSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceIAMSpecInto(in, out)
	return out
}
