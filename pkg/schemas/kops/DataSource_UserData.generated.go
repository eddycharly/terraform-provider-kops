package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceUserData() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name":    Computed(String()),
			"type":    Computed(String()),
			"content": Computed(String()),
		},
	}
}

func ExpandDataSourceUserData(in map[string]interface{}) kops.UserData {
	if in == nil {
		panic("expand UserData failure, in is nil")
	}
	out := kops.UserData{}
	if in, ok := in["name"]; ok && in != nil {
		out.Name = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["type"]; ok && in != nil {
		out.Type = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["content"]; ok && in != nil {
		out.Content = func(in interface{}) string { return string(in.(string)) }(in)
	}
	return out
}
