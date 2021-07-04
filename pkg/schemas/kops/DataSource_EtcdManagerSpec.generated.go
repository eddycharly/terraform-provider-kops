package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceEtcdManagerSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"image":                   Computed(String()),
			"env":                     Computed(List(DataSourceEnvVar())),
			"discovery_poll_interval": Computed(Nullable(String())),
			"log_level":               Computed(Nullable(Int())),
		},
	}
}

func ExpandDataSourceEtcdManagerSpec(in map[string]interface{}) kops.EtcdManagerSpec {
	if in == nil {
		panic("expand EtcdManagerSpec failure, in is nil")
	}
	out := kops.EtcdManagerSpec{}
	if in, ok := in["image"]; ok && in != nil {
		out.Image = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["env"]; ok && in != nil {
		out.Env = func(in interface{}) []kops.EnvVar {
			var out []kops.EnvVar
			for _, in := range in.([]interface{}) {
				out = append(out, func(in interface{}) kops.EnvVar { return ExpandDataSourceEnvVar(in.(map[string]interface{})) }(in))
			}
			return out
		}(in)
	}
	if in, ok := in["discovery_poll_interval"]; ok && in != nil {
		out.DiscoveryPollInterval = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["log_level"]; ok && in != nil {
		out.LogLevel = func(in interface{}) *int32 {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in int32) *int32 { return &in }(func(in interface{}) int32 { return int32(in.(int)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	return out
}

func FlattenDataSourceEtcdManagerSpecInto(in kops.EtcdManagerSpec, out map[string]interface{}) {
	out["image"] = func(in string) interface{} { return string(in) }(in.Image)
	out["env"] = func(in []kops.EnvVar) interface{} {
		var out []interface{}
		for _, in := range in {
			out = append(out, func(in kops.EnvVar) interface{} { return FlattenDataSourceEnvVar(in) }(in))
		}
		return out
	}(in.Env)
	out["discovery_poll_interval"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.DiscoveryPollInterval)
	out["log_level"] = func(in *int32) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in int32) interface{} { return int(in) }(*in)}
	}(in.LogLevel)
}

func FlattenDataSourceEtcdManagerSpec(in kops.EtcdManagerSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceEtcdManagerSpecInto(in, out)
	return out
}
