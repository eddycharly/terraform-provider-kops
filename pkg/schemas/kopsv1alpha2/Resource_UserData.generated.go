package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

var _ = Schema

func ResourceUserData() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name":    RequiredString(),
			"type":    RequiredString(),
			"content": RequiredString(),
		},
	}

	return res
}

func ExpandResourceUserData(in map[string]interface{}) kopsv1alpha2.UserData {
	if in == nil {
		panic("expand UserData failure, in is nil")
	}
	return kopsv1alpha2.UserData{
		Name: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["name"]),
		Type: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["type"]),
		Content: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["content"]),
	}
}

func FlattenResourceUserDataInto(in kopsv1alpha2.UserData, out map[string]interface{}) {
	out["name"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Name)
	out["type"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Type)
	out["content"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Content)
}

func FlattenResourceUserData(in kopsv1alpha2.UserData) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceUserDataInto(in, out)
	return out
}
