package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceDNSControllerGossipConfigSecondary() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"protocol": Optional(Nullable(String())),
			"listen":   Optional(Nullable(String())),
			"secret":   Optional(Nullable(String())),
			"seed":     Optional(Nullable(String())),
		},
	}
}

func ExpandResourceDNSControllerGossipConfigSecondary(in map[string]interface{}) kops.DNSControllerGossipConfigSecondary {
	if in == nil {
		panic("expand DNSControllerGossipConfigSecondary failure, in is nil")
	}
	out := kops.DNSControllerGossipConfigSecondary{}
	if in, ok := in["protocol"]; ok && in != nil {
		out.Protocol = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["listen"]; ok && in != nil {
		out.Listen = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["secret"]; ok && in != nil {
		out.Secret = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["seed"]; ok && in != nil {
		out.Seed = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	return out
}

func FlattenResourceDNSControllerGossipConfigSecondaryInto(in kops.DNSControllerGossipConfigSecondary, out map[string]interface{}) {
	out["protocol"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.Protocol)
	out["listen"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.Listen)
	out["secret"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.Secret)
	out["seed"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.Seed)
}

func FlattenResourceDNSControllerGossipConfigSecondary(in kops.DNSControllerGossipConfigSecondary) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceDNSControllerGossipConfigSecondaryInto(in, out)
	return out
}
