package schemas

import (
	"reflect"

	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

var _ = Schema

func ResourceEtcdManagerSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"image":                   OptionalString(),
			"env":                     OptionalList(ResourceEnvVar()),
			"backup_interval":         OptionalDuration(),
			"discovery_poll_interval": OptionalDuration(),
			"log_level":               OptionalInt(),
		},
	}

	return res
}

func ExpandResourceEtcdManagerSpec(in map[string]interface{}) kopsv1alpha2.EtcdManagerSpec {
	if in == nil {
		panic("expand EtcdManagerSpec failure, in is nil")
	}
	return kopsv1alpha2.EtcdManagerSpec{
		Image: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["image"]),
		Env: func(in interface{}) []kopsv1alpha2.EnvVar {
			return func(in interface{}) []kopsv1alpha2.EnvVar {
				if in == nil {
					return nil
				}
				var out []kopsv1alpha2.EnvVar
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) kopsv1alpha2.EnvVar {
						if in == nil {
							return kopsv1alpha2.EnvVar{}
						}
						return ExpandResourceEnvVar(in.(map[string]interface{}))
					}(in))
				}
				return out
			}(in)
		}(in["env"]),
		BackupInterval: func(in interface{}) *meta.Duration {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *meta.Duration {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in meta.Duration) *meta.Duration {
					return &in
				}(ExpandDuration(in))
			}(in)
		}(in["backup_interval"]),
		DiscoveryPollInterval: func(in interface{}) *meta.Duration {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *meta.Duration {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in meta.Duration) *meta.Duration {
					return &in
				}(ExpandDuration(in))
			}(in)
		}(in["discovery_poll_interval"]),
		LogLevel: func(in interface{}) *int32 {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *int32 {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in int32) *int32 {
					return &in
				}(int32(ExpandInt(in)))
			}(in)
		}(in["log_level"]),
	}
}

func FlattenResourceEtcdManagerSpecInto(in kopsv1alpha2.EtcdManagerSpec, out map[string]interface{}) {
	out["image"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Image)
	out["env"] = func(in []kopsv1alpha2.EnvVar) interface{} {
		return func(in []kopsv1alpha2.EnvVar) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, func(in kopsv1alpha2.EnvVar) interface{} {
					return FlattenResourceEnvVar(in)
				}(in))
			}
			return out
		}(in)
	}(in.Env)
	out["backup_interval"] = func(in *meta.Duration) interface{} {
		return func(in *meta.Duration) interface{} {
			if in == nil {
				return nil
			}
			return func(in meta.Duration) interface{} {
				return FlattenDuration(in)
			}(*in)
		}(in)
	}(in.BackupInterval)
	out["discovery_poll_interval"] = func(in *meta.Duration) interface{} {
		return func(in *meta.Duration) interface{} {
			if in == nil {
				return nil
			}
			return func(in meta.Duration) interface{} {
				return FlattenDuration(in)
			}(*in)
		}(in)
	}(in.DiscoveryPollInterval)
	out["log_level"] = func(in *int32) interface{} {
		return func(in *int32) interface{} {
			if in == nil {
				return nil
			}
			return func(in int32) interface{} {
				return FlattenInt(int(in))
			}(*in)
		}(in)
	}(in.LogLevel)
}

func FlattenResourceEtcdManagerSpec(in kopsv1alpha2.EtcdManagerSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceEtcdManagerSpecInto(in, out)
	return out
}
