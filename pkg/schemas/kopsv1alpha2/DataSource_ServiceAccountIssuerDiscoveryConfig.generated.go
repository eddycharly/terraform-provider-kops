package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

var _ = Schema

func DataSourceServiceAccountIssuerDiscoveryConfig() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"discovery_store":          ComputedString(),
			"enable_aws_oidc_provider": ComputedBool(),
			"additional_audiences":     ComputedList(String()),
		},
	}

	return res
}

func ExpandDataSourceServiceAccountIssuerDiscoveryConfig(in map[string]interface{}) kopsv1alpha2.ServiceAccountIssuerDiscoveryConfig {
	if in == nil {
		panic("expand ServiceAccountIssuerDiscoveryConfig failure, in is nil")
	}
	return kopsv1alpha2.ServiceAccountIssuerDiscoveryConfig{
		DiscoveryStore: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["discovery_store"]),
		EnableAWSOIDCProvider: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["enable_aws_oidc_provider"]),
		AdditionalAudiences: func(in interface{}) []string {
			return func(in interface{}) []string {
				if in == nil {
					return nil
				}
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["additional_audiences"]),
	}
}

func FlattenDataSourceServiceAccountIssuerDiscoveryConfigInto(in kopsv1alpha2.ServiceAccountIssuerDiscoveryConfig, out map[string]interface{}) {
	out["discovery_store"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.DiscoveryStore)
	out["enable_aws_oidc_provider"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.EnableAWSOIDCProvider)
	out["additional_audiences"] = func(in []string) interface{} {
		return func(in []string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.AdditionalAudiences)
}

func FlattenDataSourceServiceAccountIssuerDiscoveryConfig(in kopsv1alpha2.ServiceAccountIssuerDiscoveryConfig) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceServiceAccountIssuerDiscoveryConfigInto(in, out)
	return out
}
