package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceCertManagerConfig() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enabled":        Computed(Nullable(Bool())),
			"managed":        Computed(Nullable(Bool())),
			"image":          Computed(Nullable(String())),
			"default_issuer": Computed(Nullable(String())),
		},
	}
}

func ExpandDataSourceCertManagerConfig(in map[string]interface{}) kops.CertManagerConfig {
	if in == nil {
		panic("expand CertManagerConfig failure, in is nil")
	}
	out := kops.CertManagerConfig{}
	if in, ok := in["enabled"]; ok && in != nil {
		out.Enabled = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["managed"]; ok && in != nil {
		out.Managed = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["image"]; ok && in != nil {
		out.Image = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["default_issuer"]; ok && in != nil {
		out.DefaultIssuer = func(in interface{}) *string {
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

func FlattenDataSourceCertManagerConfigInto(in kops.CertManagerConfig, out map[string]interface{}) {
	out["enabled"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in bool) interface{} { return in }(*in)}
	}(in.Enabled)
	out["managed"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in bool) interface{} { return in }(*in)}
	}(in.Managed)
	out["image"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.Image)
	out["default_issuer"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.DefaultIssuer)
}

func FlattenDataSourceCertManagerConfig(in kops.CertManagerConfig) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceCertManagerConfigInto(in, out)
	return out
}
