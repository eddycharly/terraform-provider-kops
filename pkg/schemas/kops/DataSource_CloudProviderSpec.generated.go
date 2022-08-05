package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceCloudProviderSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"aws":       ComputedStruct(DataSourceAWSSpec()),
			"azure":     ComputedStruct(DataSourceAzureSpec()),
			"do":        ComputedStruct(DataSourceDOSpec()),
			"gce":       ComputedStruct(DataSourceGCESpec()),
			"hetzner":   ComputedStruct(DataSourceHetznerSpec()),
			"openstack": ComputedStruct(DataSourceOpenstackSpec()),
		},
	}

	return res
}

func ExpandDataSourceCloudProviderSpec(in map[string]interface{}) kops.CloudProviderSpec {
	if in == nil {
		panic("expand CloudProviderSpec failure, in is nil")
	}
	return kops.CloudProviderSpec{
		AWS: func(in interface{}) *kops.AWSSpec {
			return func(in interface{}) *kops.AWSSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.AWSSpec) *kops.AWSSpec {
					return &in
				}(func(in interface{}) kops.AWSSpec {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandDataSourceAWSSpec(in[0].(map[string]interface{}))
					}
					return kops.AWSSpec{}
				}(in))
			}(in)
		}(in["aws"]),
		Azure: func(in interface{}) *kops.AzureSpec {
			return func(in interface{}) *kops.AzureSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.AzureSpec) *kops.AzureSpec {
					return &in
				}(func(in interface{}) kops.AzureSpec {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandDataSourceAzureSpec(in[0].(map[string]interface{}))
					}
					return kops.AzureSpec{}
				}(in))
			}(in)
		}(in["azure"]),
		DO: func(in interface{}) *kops.DOSpec {
			return func(in interface{}) *kops.DOSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.DOSpec) *kops.DOSpec {
					return &in
				}(func(in interface{}) kops.DOSpec {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandDataSourceDOSpec(in[0].(map[string]interface{}))
					}
					return kops.DOSpec{}
				}(in))
			}(in)
		}(in["do"]),
		GCE: func(in interface{}) *kops.GCESpec {
			return func(in interface{}) *kops.GCESpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.GCESpec) *kops.GCESpec {
					return &in
				}(func(in interface{}) kops.GCESpec {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandDataSourceGCESpec(in[0].(map[string]interface{}))
					}
					return kops.GCESpec{}
				}(in))
			}(in)
		}(in["gce"]),
		Hetzner: func(in interface{}) *kops.HetznerSpec {
			return func(in interface{}) *kops.HetznerSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.HetznerSpec) *kops.HetznerSpec {
					return &in
				}(func(in interface{}) kops.HetznerSpec {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandDataSourceHetznerSpec(in[0].(map[string]interface{}))
					}
					return kops.HetznerSpec{}
				}(in))
			}(in)
		}(in["hetzner"]),
		Openstack: func(in interface{}) *kops.OpenstackSpec {
			return func(in interface{}) *kops.OpenstackSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.OpenstackSpec) *kops.OpenstackSpec {
					return &in
				}(func(in interface{}) kops.OpenstackSpec {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandDataSourceOpenstackSpec(in[0].(map[string]interface{}))
					}
					return kops.OpenstackSpec{}
				}(in))
			}(in)
		}(in["openstack"]),
	}
}

func FlattenDataSourceCloudProviderSpecInto(in kops.CloudProviderSpec, out map[string]interface{}) {
	out["aws"] = func(in *kops.AWSSpec) interface{} {
		return func(in *kops.AWSSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.AWSSpec) interface{} {
				return func(in kops.AWSSpec) []interface{} {
					return []interface{}{FlattenDataSourceAWSSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.AWS)
	out["azure"] = func(in *kops.AzureSpec) interface{} {
		return func(in *kops.AzureSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.AzureSpec) interface{} {
				return func(in kops.AzureSpec) []interface{} {
					return []interface{}{FlattenDataSourceAzureSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Azure)
	out["do"] = func(in *kops.DOSpec) interface{} {
		return func(in *kops.DOSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.DOSpec) interface{} {
				return func(in kops.DOSpec) []interface{} {
					return []interface{}{FlattenDataSourceDOSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.DO)
	out["gce"] = func(in *kops.GCESpec) interface{} {
		return func(in *kops.GCESpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.GCESpec) interface{} {
				return func(in kops.GCESpec) []interface{} {
					return []interface{}{FlattenDataSourceGCESpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.GCE)
	out["hetzner"] = func(in *kops.HetznerSpec) interface{} {
		return func(in *kops.HetznerSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.HetznerSpec) interface{} {
				return func(in kops.HetznerSpec) []interface{} {
					return []interface{}{FlattenDataSourceHetznerSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Hetzner)
	out["openstack"] = func(in *kops.OpenstackSpec) interface{} {
		return func(in *kops.OpenstackSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.OpenstackSpec) interface{} {
				return func(in kops.OpenstackSpec) []interface{} {
					return []interface{}{FlattenDataSourceOpenstackSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Openstack)
}

func FlattenDataSourceCloudProviderSpec(in kops.CloudProviderSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceCloudProviderSpecInto(in, out)
	return out
}
