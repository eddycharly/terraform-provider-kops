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
			"enabled":                           Computed(Nullable(Bool())),
			"enable_spot_interruption_draining": Computed(Nullable(Bool())),
			"enable_scheduled_event_draining":   Computed(Nullable(Bool())),
			"enable_prometheus_metrics":         Computed(Nullable(Bool())),
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

func FlattenDataSourceNodeTerminationHandlerConfigInto(in kops.NodeTerminationHandlerConfig, out map[string]interface{}) {
	out["enabled"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in bool) interface{} { return in }(*in)}
	}(in.Enabled)
	out["enable_spot_interruption_draining"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in bool) interface{} { return in }(*in)}
	}(in.EnableSpotInterruptionDraining)
	out["enable_scheduled_event_draining"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in bool) interface{} { return in }(*in)}
	}(in.EnableScheduledEventDraining)
	out["enable_prometheus_metrics"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return map[string]interface{}{"value": func(in bool) interface{} { return in }(*in)}
	}(in.EnablePrometheusMetrics)
}

func FlattenDataSourceNodeTerminationHandlerConfig(in kops.NodeTerminationHandlerConfig) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceNodeTerminationHandlerConfigInto(in, out)
	return out
}
