package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceGossipConfigSecondary() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"protocol": Computed(Nullable(String())),
			"listen":   Computed(Nullable(String())),
			"secret":   Computed(Nullable(String())),
		},
	}
}

func ExpandDataSourceGossipConfigSecondary(in map[string]interface{}) kops.GossipConfigSecondary {
	if in == nil {
		panic("expand GossipConfigSecondary failure, in is nil")
	}
	out := kops.GossipConfigSecondary{}
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
	return out
}

func FlattenDataSourceGossipConfigSecondaryInto(in kops.GossipConfigSecondary, out map[string]interface{}) {
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
}

func FlattenDataSourceGossipConfigSecondary(in kops.GossipConfigSecondary) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceGossipConfigSecondaryInto(in, out)
	return out
}
