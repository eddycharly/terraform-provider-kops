package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceCloudProviderSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"aws":       OptionalStruct(ResourceAWSSpec()),
			"azure":     OptionalStruct(ResourceAzureSpec()),
			"do":        OptionalStruct(ResourceDOSpec()),
			"gce":       OptionalStruct(ResourceGCESpec()),
			"hetzner":   OptionalStruct(ResourceHetznerSpec()),
			"openstack": OptionalStruct(ResourceOpenstackSpec()),
			"scaleway":  OptionalStruct(ResourceScalewaySpec()),
		},
	}

	return res
}

func ExpandResourceCloudProviderSpec(in map[string]interface{}) kops.CloudProviderSpec {
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
						return ExpandResourceAWSSpec(in[0].(map[string]interface{}))
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
						return ExpandResourceAzureSpec(in[0].(map[string]interface{}))
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
						return ExpandResourceDOSpec(in[0].(map[string]interface{}))
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
						return ExpandResourceGCESpec(in[0].(map[string]interface{}))
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
						return ExpandResourceHetznerSpec(in[0].(map[string]interface{}))
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
						return ExpandResourceOpenstackSpec(in[0].(map[string]interface{}))
					}
					return kops.OpenstackSpec{}
				}(in))
			}(in)
		}(in["openstack"]),
		Scaleway: func(in interface{}) *kops.ScalewaySpec {
			return func(in interface{}) *kops.ScalewaySpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.ScalewaySpec) *kops.ScalewaySpec {
					return &in
				}(func(in interface{}) kops.ScalewaySpec {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandResourceScalewaySpec(in[0].(map[string]interface{}))
					}
					return kops.ScalewaySpec{}
				}(in))
			}(in)
		}(in["scaleway"]),
	}
}

func FlattenResourceCloudProviderSpecInto(in kops.CloudProviderSpec, out map[string]interface{}) {
	out["aws"] = func(in *kops.AWSSpec) interface{} {
		return func(in *kops.AWSSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.AWSSpec) interface{} {
				return func(in kops.AWSSpec) []interface{} {
					return []interface{}{FlattenResourceAWSSpec(in)}
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
					return []interface{}{FlattenResourceAzureSpec(in)}
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
					return []interface{}{FlattenResourceDOSpec(in)}
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
					return []interface{}{FlattenResourceGCESpec(in)}
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
					return []interface{}{FlattenResourceHetznerSpec(in)}
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
					return []interface{}{FlattenResourceOpenstackSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Openstack)
	out["scaleway"] = func(in *kops.ScalewaySpec) interface{} {
		return func(in *kops.ScalewaySpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.ScalewaySpec) interface{} {
				return func(in kops.ScalewaySpec) []interface{} {
					return []interface{}{FlattenResourceScalewaySpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Scaleway)
}

func FlattenResourceCloudProviderSpec(in kops.CloudProviderSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceCloudProviderSpecInto(in, out)
	return out
}
