package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceAssets() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"container_registry": Computed(Nullable(String())),
			"file_repository":    Computed(Nullable(String())),
			"container_proxy":    Computed(Nullable(String())),
		},
	}
}

func ExpandDataSourceAssets(in map[string]interface{}) kops.Assets {
	if in == nil {
		panic("expand Assets failure, in is nil")
	}
	out := kops.Assets{}
	if in, ok := in["container_registry"]; ok && in != nil {
		out.ContainerRegistry = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["file_repository"]; ok && in != nil {
		out.FileRepository = func(in interface{}) *string {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && in != nil && len(in) == 1 {
				return func(in string) *string { return &in }(func(in interface{}) string { return string(in.(string)) }(in[0].(map[string]interface{})["value"]))
			}
			return nil
		}(in)
	}
	if in, ok := in["container_proxy"]; ok && in != nil {
		out.ContainerProxy = func(in interface{}) *string {
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

func FlattenDataSourceAssetsInto(in kops.Assets, out map[string]interface{}) {
	out["container_registry"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.ContainerRegistry)
	out["file_repository"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.FileRepository)
	out["container_proxy"] = func(in *string) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in string) interface{} { return string(in) }(*in)}
	}(in.ContainerProxy)
}

func FlattenDataSourceAssets(in kops.Assets) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceAssetsInto(in, out)
	return out
}
