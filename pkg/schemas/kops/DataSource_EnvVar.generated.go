package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceEnvVar() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name":  Computed(String()),
			"value": Computed(String()),
		},
	}
}

func ExpandDataSourceEnvVar(in map[string]interface{}) kops.EnvVar {
	if in == nil {
		panic("expand EnvVar failure, in is nil")
	}
	out := kops.EnvVar{}
	if in, ok := in["name"]; ok && in != nil {
		out.Name = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["value"]; ok && in != nil {
		out.Value = func(in interface{}) string { return string(in.(string)) }(in)
	}
	return out
}

func FlattenDataSourceEnvVarInto(in kops.EnvVar, out map[string]interface{}) {
	out["name"] = func(in string) interface{} { return string(in) }(in.Name)
	out["value"] = func(in string) interface{} { return string(in) }(in.Value)
}

func FlattenDataSourceEnvVar(in kops.EnvVar) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceEnvVarInto(in, out)
	return out
}
