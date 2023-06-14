package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceAWSSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ebs_csi_driver":           ComputedStruct(DataSourceEBSCSIDriverSpec()),
			"node_termination_handler": ComputedStruct(DataSourceNodeTerminationHandlerSpec()),
			"load_balancer_controller": ComputedStruct(DataSourceLoadBalancerControllerSpec()),
			"pod_identity_webhook":     ComputedStruct(DataSourcePodIdentityWebhookSpec()),
			"warm_pool":                ComputedStruct(DataSourceWarmPoolSpec()),
		},
	}

	return res
}

func ExpandDataSourceAWSSpec(in map[string]interface{}) kops.AWSSpec {
	if in == nil {
		panic("expand AWSSpec failure, in is nil")
	}
	return kops.AWSSpec{
		EBSCSIDriver: func(in interface{}) *kops.EBSCSIDriverSpec {
			return func(in interface{}) *kops.EBSCSIDriverSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.EBSCSIDriverSpec) *kops.EBSCSIDriverSpec {
					return &in
				}(func(in interface{}) kops.EBSCSIDriverSpec {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandDataSourceEBSCSIDriverSpec(in[0].(map[string]interface{}))
					}
					return kops.EBSCSIDriverSpec{}
				}(in))
			}(in)
		}(in["ebs_csi_driver"]),
		NodeTerminationHandler: func(in interface{}) *kops.NodeTerminationHandlerSpec {
			return func(in interface{}) *kops.NodeTerminationHandlerSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.NodeTerminationHandlerSpec) *kops.NodeTerminationHandlerSpec {
					return &in
				}(func(in interface{}) kops.NodeTerminationHandlerSpec {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandDataSourceNodeTerminationHandlerSpec(in[0].(map[string]interface{}))
					}
					return kops.NodeTerminationHandlerSpec{}
				}(in))
			}(in)
		}(in["node_termination_handler"]),
		LoadBalancerController: func(in interface{}) *kops.LoadBalancerControllerSpec {
			return func(in interface{}) *kops.LoadBalancerControllerSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.LoadBalancerControllerSpec) *kops.LoadBalancerControllerSpec {
					return &in
				}(func(in interface{}) kops.LoadBalancerControllerSpec {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandDataSourceLoadBalancerControllerSpec(in[0].(map[string]interface{}))
					}
					return kops.LoadBalancerControllerSpec{}
				}(in))
			}(in)
		}(in["load_balancer_controller"]),
		PodIdentityWebhook: func(in interface{}) *kops.PodIdentityWebhookSpec {
			return func(in interface{}) *kops.PodIdentityWebhookSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.PodIdentityWebhookSpec) *kops.PodIdentityWebhookSpec {
					return &in
				}(func(in interface{}) kops.PodIdentityWebhookSpec {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandDataSourcePodIdentityWebhookSpec(in[0].(map[string]interface{}))
					}
					return kops.PodIdentityWebhookSpec{}
				}(in))
			}(in)
		}(in["pod_identity_webhook"]),
		WarmPool: func(in interface{}) *kops.WarmPoolSpec {
			return func(in interface{}) *kops.WarmPoolSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.WarmPoolSpec) *kops.WarmPoolSpec {
					return &in
				}(func(in interface{}) kops.WarmPoolSpec {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandDataSourceWarmPoolSpec(in[0].(map[string]interface{}))
					}
					return kops.WarmPoolSpec{}
				}(in))
			}(in)
		}(in["warm_pool"]),
	}
}

func FlattenDataSourceAWSSpecInto(in kops.AWSSpec, out map[string]interface{}) {
	out["ebs_csi_driver"] = func(in *kops.EBSCSIDriverSpec) interface{} {
		return func(in *kops.EBSCSIDriverSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.EBSCSIDriverSpec) interface{} {
				return func(in kops.EBSCSIDriverSpec) []interface{} {
					return []interface{}{FlattenDataSourceEBSCSIDriverSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.EBSCSIDriver)
	out["node_termination_handler"] = func(in *kops.NodeTerminationHandlerSpec) interface{} {
		return func(in *kops.NodeTerminationHandlerSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.NodeTerminationHandlerSpec) interface{} {
				return func(in kops.NodeTerminationHandlerSpec) []interface{} {
					return []interface{}{FlattenDataSourceNodeTerminationHandlerSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.NodeTerminationHandler)
	out["load_balancer_controller"] = func(in *kops.LoadBalancerControllerSpec) interface{} {
		return func(in *kops.LoadBalancerControllerSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.LoadBalancerControllerSpec) interface{} {
				return func(in kops.LoadBalancerControllerSpec) []interface{} {
					return []interface{}{FlattenDataSourceLoadBalancerControllerSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.LoadBalancerController)
	out["pod_identity_webhook"] = func(in *kops.PodIdentityWebhookSpec) interface{} {
		return func(in *kops.PodIdentityWebhookSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.PodIdentityWebhookSpec) interface{} {
				return func(in kops.PodIdentityWebhookSpec) []interface{} {
					return []interface{}{FlattenDataSourcePodIdentityWebhookSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.PodIdentityWebhook)
	out["warm_pool"] = func(in *kops.WarmPoolSpec) interface{} {
		return func(in *kops.WarmPoolSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.WarmPoolSpec) interface{} {
				return func(in kops.WarmPoolSpec) []interface{} {
					return []interface{}{FlattenDataSourceWarmPoolSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.WarmPool)
}

func FlattenDataSourceAWSSpec(in kops.AWSSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceAWSSpecInto(in, out)
	return out
}
