package schemas

import (
	"reflect"

	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceNodeAuthorizerSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"authorizer": ComputedString(),
			"features":   ComputedList(String()),
			"image":      ComputedString(),
			"node_url":   ComputedString(),
			"port":       ComputedInt(),
			"interval":   ComputedDuration(),
			"timeout":    ComputedDuration(),
			"token_ttl":  ComputedDuration(),
		},
	}
}

func ExpandDataSourceNodeAuthorizerSpec(in map[string]interface{}) kops.NodeAuthorizerSpec {
	if in == nil {
		panic("expand NodeAuthorizerSpec failure, in is nil")
	}
	return kops.NodeAuthorizerSpec{
		Authorizer: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["authorizer"]),
		Features: func(in interface{}) []string {
			return func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["features"]),
		Image: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["image"]),
		NodeURL: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["node_url"]),
		Port: func(in interface{}) int {
			return int(ExpandInt(in))
		}(in["port"]),
		Interval: func(in interface{}) *v1.Duration {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *v1.Duration {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in v1.Duration) *v1.Duration {
					return &in
				}(ExpandDuration(in))
			}(in)
		}(in["interval"]),
		Timeout: func(in interface{}) *v1.Duration {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *v1.Duration {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in v1.Duration) *v1.Duration {
					return &in
				}(ExpandDuration(in))
			}(in)
		}(in["timeout"]),
		TokenTTL: func(in interface{}) *v1.Duration {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *v1.Duration {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in v1.Duration) *v1.Duration {
					return &in
				}(ExpandDuration(in))
			}(in)
		}(in["token_ttl"]),
	}
}

func FlattenDataSourceNodeAuthorizerSpecInto(in kops.NodeAuthorizerSpec, out map[string]interface{}) {
	out["authorizer"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Authorizer)
	out["features"] = func(in []string) interface{} {
		return func(in []string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.Features)
	out["image"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Image)
	out["node_url"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.NodeURL)
	out["port"] = func(in int) interface{} {
		return FlattenInt(int(in))
	}(in.Port)
	out["interval"] = func(in *v1.Duration) interface{} {
		return func(in *v1.Duration) interface{} {
			if in == nil {
				return nil
			}
			return func(in v1.Duration) interface{} {
				return FlattenDuration(in)
			}(*in)
		}(in)
	}(in.Interval)
	out["timeout"] = func(in *v1.Duration) interface{} {
		return func(in *v1.Duration) interface{} {
			if in == nil {
				return nil
			}
			return func(in v1.Duration) interface{} {
				return FlattenDuration(in)
			}(*in)
		}(in)
	}(in.Timeout)
	out["token_ttl"] = func(in *v1.Duration) interface{} {
		return func(in *v1.Duration) interface{} {
			if in == nil {
				return nil
			}
			return func(in v1.Duration) interface{} {
				return FlattenDuration(in)
			}(*in)
		}(in)
	}(in.TokenTTL)
}

func FlattenDataSourceNodeAuthorizerSpec(in kops.NodeAuthorizerSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceNodeAuthorizerSpecInto(in, out)
	return out
}
