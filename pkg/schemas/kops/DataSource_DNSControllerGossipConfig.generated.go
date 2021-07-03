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
			"protocol":  Computed(Ptr(String())),
			"listen":    Computed(Ptr(String())),
			"secret":    Computed(Ptr(String())),
			"secondary": Computed(Ptr(Struct(DataSourceDNSControllerGossipConfigSecondary()))),
			"seed":      Computed(Ptr(String())),
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
