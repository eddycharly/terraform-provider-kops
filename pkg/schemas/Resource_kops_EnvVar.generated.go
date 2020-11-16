package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

func ResourceKopsEnvVar() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name":  RequiredString(),
			"value": OptionalString(),
		},
	}
}

func ExpandResourceKopsEnvVar(in map[string]interface{}) kops.EnvVar {
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

func FlattenResourceKopsEnvVarInto(in kops.EnvVar, out map[string]interface{}) {
	out["name"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Name)
	out["value"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Value)
}

func FlattenResourceKopsEnvVar(in kops.EnvVar) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceKopsEnvVarInto(in, out)
	return out
}
