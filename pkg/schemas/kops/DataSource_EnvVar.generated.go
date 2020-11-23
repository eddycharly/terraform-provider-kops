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
			"name":  ComputedString(),
			"value": ComputedString(),
		},
	}
}

func ExpandDataSourceEnvVar(in map[string]interface{}) kops.EnvVar {
	if in == nil {
		panic("expand EnvVar failure, in is nil")
	}
	return kops.EnvVar{
		Name: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["name"]),
		Value: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["value"]),
	}
}

func FlattenDataSourceEnvVarInto(in kops.EnvVar, out map[string]interface{}) {
	out["name"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Name)
	out["value"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Value)
}

func FlattenDataSourceEnvVar(in kops.EnvVar) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceEnvVarInto(in, out)
	return out
}
