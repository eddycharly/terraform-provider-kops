package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceDNSControllerGossipConfig() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"protocol":  Computed(Nullable(String())),
			"listen":    Computed(Nullable(String())),
			"secret":    Computed(Nullable(String())),
			"secondary": Computed(Struct(DataSourceDNSControllerGossipConfigSecondary())),
			"seed":      Computed(Nullable(String())),
		},
	}
}

func ExpandDataSourceDNSControllerGossipConfig(in map[string]interface{}) kops.DNSControllerGossipConfig {
	if in == nil {
		panic("expand DNSControllerGossipConfig failure, in is nil")
	}
	out := kops.DNSControllerGossipConfig{}
	if in, ok := in["protocol"]; ok && in != nil {
		out.Protocol = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["listen"]; ok && in != nil {
		out.Listen = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["secret"]; ok && in != nil {
		out.Secret = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["secondary"]; ok && in != nil {
		out.Secondary = func(in interface{}) *kops.DNSControllerGossipConfigSecondary {
			if in == nil {
				return nil
			}
			return func(in kops.DNSControllerGossipConfigSecondary) *kops.DNSControllerGossipConfigSecondary { return &in }(func(in interface{}) kops.DNSControllerGossipConfigSecondary {
				if in == nil {
					return kops.DNSControllerGossipConfigSecondary{}
				}
				return ExpandDataSourceDNSControllerGossipConfigSecondary(in.(map[string]interface{}))
			}(in))
		}(in)
	}
	if in, ok := in["seed"]; ok && in != nil {
		out.Seed = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	return out
}

func FlattenDataSourceDNSControllerGossipConfigInto(in kops.DNSControllerGossipConfig, out map[string]interface{}) {
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
	out["secondary"] = func(in *kops.DNSControllerGossipConfigSecondary) interface{} {
		if in == nil {
			return nil
		}
		return func(in kops.DNSControllerGossipConfigSecondary) interface{} {
			return FlattenDataSourceDNSControllerGossipConfigSecondary(in)
		}(*in)
	}(in.Secondary)
	out["seed"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.Seed)
}

func FlattenDataSourceDNSControllerGossipConfig(in kops.DNSControllerGossipConfig) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceDNSControllerGossipConfigInto(in, out)
	return out
}
