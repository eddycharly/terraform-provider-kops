package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

var _ = Schema

func ResourceEnvVar() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name":  RequiredString(),
			"value": OptionalString(),
		},
	}

	return res
}

func ExpandResourceEnvVar(in map[string]interface{}) kopsv1alpha2.EnvVar {
	if in == nil {
		panic("expand EnvVar failure, in is nil")
	}
	return kopsv1alpha2.EnvVar{
		Name: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["name"]),
		Value: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["value"]),
	}
}

func FlattenResourceEnvVarInto(in kopsv1alpha2.EnvVar, out map[string]interface{}) {
	out["name"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Name)
	out["value"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Value)
}

func FlattenResourceEnvVar(in kopsv1alpha2.EnvVar) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceEnvVarInto(in, out)
	return out
}
