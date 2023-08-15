package schemas

import (
	"reflect"

	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

var _ = Schema

func ResourceGossipConfigSecondary() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"protocol": OptionalString(),
			"listen":   OptionalString(),
			"secret":   OptionalString(),
		},
	}

	return res
}

func ExpandResourceGossipConfigSecondary(in map[string]interface{}) kopsv1alpha2.GossipConfigSecondary {
	if in == nil {
		panic("expand GossipConfigSecondary failure, in is nil")
	}
	return kopsv1alpha2.GossipConfigSecondary{
		Protocol: func(in interface{}) *string {
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
		}(in["protocol"]),
		Listen: func(in interface{}) *string {
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
		}(in["listen"]),
		Secret: func(in interface{}) *string {
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
		}(in["secret"]),
	}
}

func FlattenResourceGossipConfigSecondaryInto(in kopsv1alpha2.GossipConfigSecondary, out map[string]interface{}) {
	out["protocol"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.Protocol)
	out["listen"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.Listen)
	out["secret"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.Secret)
}

func FlattenResourceGossipConfigSecondary(in kopsv1alpha2.GossipConfigSecondary) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceGossipConfigSecondaryInto(in, out)
	return out
}
