package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
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

func ExpandDataSourceServiceAccountExternalPermission(in map[string]interface{}) kopsv1alpha2.ServiceAccountExternalPermission {
	if in == nil {
		panic("expand ServiceAccountExternalPermission failure, in is nil")
	}
	return kopsv1alpha2.ServiceAccountExternalPermission{
		Name: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["name"]),
		Namespace: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["namespace"]),
		AWS: func(in interface{}) *kopsv1alpha2.AWSPermission {
			return func(in interface{}) *kopsv1alpha2.AWSPermission {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kopsv1alpha2.AWSPermission) *kopsv1alpha2.AWSPermission {
					return &in
				}(func(in interface{}) kopsv1alpha2.AWSPermission {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandDataSourceAWSPermission(in[0].(map[string]interface{}))
					}
					return kopsv1alpha2.AWSPermission{}
				}(in))
			}(in)
		}(in["aws"]),
	}
}

func FlattenDataSourceServiceAccountExternalPermissionInto(in kopsv1alpha2.ServiceAccountExternalPermission, out map[string]interface{}) {
	out["name"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Name)
	out["namespace"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Namespace)
	out["aws"] = func(in *kopsv1alpha2.AWSPermission) interface{} {
		return func(in *kopsv1alpha2.AWSPermission) interface{} {
			if in == nil {
				return nil
			}
			return func(in kopsv1alpha2.AWSPermission) interface{} {
				return func(in kopsv1alpha2.AWSPermission) []interface{} {
					return []interface{}{FlattenDataSourceAWSPermission(in)}
				}(in)
			}(*in)
		}(in)
	}(in.AWS)
}

func FlattenDataSourceServiceAccountExternalPermission(in kopsv1alpha2.ServiceAccountExternalPermission) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceServiceAccountExternalPermissionInto(in, out)
	return out
}
