package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceServiceAccountExternalPermission() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name":      ComputedString(),
			"namespace": ComputedString(),
			"aws":       ComputedStruct(DataSourceAWSPermission()),
		},
	}

	return res
}

func ExpandDataSourceServiceAccountExternalPermission(in map[string]interface{}) kops.ServiceAccountExternalPermission {
	if in == nil {
		panic("expand ServiceAccountExternalPermission failure, in is nil")
	}
	return kops.ServiceAccountExternalPermission{
		Name: func(in interface{}) string /**/ {
			return string(ExpandString(in))
		}(in["name"]),
		Namespace: func(in interface{}) string /**/ {
			return string(ExpandString(in))
		}(in["namespace"]),
		AWS: func(in interface{}) *kops.AWSPermission /*k8s.io/kops/pkg/apis/kops*/ {
			return func(in interface{}) *kops.AWSPermission {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.AWSPermission) *kops.AWSPermission {
					return &in
				}(func(in interface{}) kops.AWSPermission {
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return kops.AWSPermission{}
					}
					return (ExpandDataSourceAWSPermission(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["aws"]),
	}
}

func FlattenDataSourceServiceAccountExternalPermissionInto(in kops.ServiceAccountExternalPermission, out map[string]interface{}) {
	out["name"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Name)
	out["namespace"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Namespace)
	out["aws"] = func(in *kops.AWSPermission) interface{} {
		return func(in *kops.AWSPermission) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.AWSPermission) interface{} {
				return func(in kops.AWSPermission) []interface{} {
					return []interface{}{FlattenDataSourceAWSPermission(in)}
				}(in)
			}(*in)
		}(in)
	}(in.AWS)
}

func FlattenDataSourceServiceAccountExternalPermission(in kops.ServiceAccountExternalPermission) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceServiceAccountExternalPermissionInto(in, out)
	return out
}
