package schemas

import (
	"reflect"

	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceOpenstackMetadata() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"config_drive": ComputedBool(),
		},
	}

	return res
}

func ExpandDataSourceOpenstackMetadata(in map[string]interface{}) kops.OpenstackMetadata {
	if in == nil {
		panic("expand OpenstackMetadata failure, in is nil")
	}
	return kops.OpenstackMetadata{
		ConfigDrive: func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in bool) *bool {
					return &in
				}(bool(ExpandBool(in)))
			}(in)
		}(in["config_drive"]),
	}
}

func FlattenDataSourceOpenstackMetadataInto(in kops.OpenstackMetadata, out map[string]interface{}) {
	out["config_drive"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.ConfigDrive)
}

func FlattenDataSourceOpenstackMetadata(in kops.OpenstackMetadata) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceOpenstackMetadataInto(in, out)
	return out
}
