package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceNodeTerminationHandlerConfig() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enabled":                           Computed(Ptr(Bool())),
			"enable_spot_interruption_draining": Computed(Ptr(Bool())),
			"enable_scheduled_event_draining":   Computed(Ptr(Bool())),
			"enable_prometheus_metrics":         Computed(Ptr(Bool())),
		},
	}
}

func ExpandDataSourceNodeTerminationHandlerConfig(in map[string]interface{}) kops.NodeTerminationHandlerConfig {
	if in == nil {
		panic("expand NodeTerminationHandlerConfig failure, in is nil")
	}
	out := kops.NodeTerminationHandlerConfig{}
	if in, ok := in["enabled"]; ok && in != nil {
		out.Enabled = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["enable_spot_interruption_draining"]; ok && in != nil {
		out.EnableSpotInterruptionDraining = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["enable_scheduled_event_draining"]; ok && in != nil {
		out.EnableScheduledEventDraining = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	if in, ok := in["enable_prometheus_metrics"]; ok && in != nil {
		out.EnablePrometheusMetrics = func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			return func(in bool) *bool { return &in }(func(in interface{}) bool { return in.(bool) }(in.(map[string]interface{})["value"]))
		}(in)
	}
	return out
}
