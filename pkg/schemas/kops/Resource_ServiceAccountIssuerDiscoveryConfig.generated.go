package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceServiceAccountIssuerDiscoveryConfig() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"discovery_store":          OptionalString(),
			"enable_aws_oidc_provider": OptionalBool(),
		},
	}

	return res
}

func ExpandResourceServiceAccountIssuerDiscoveryConfig(in map[string]interface{}) kops.ServiceAccountIssuerDiscoveryConfig {
	if in == nil {
		panic("expand ServiceAccountIssuerDiscoveryConfig failure, in is nil")
	}
	return kops.ServiceAccountIssuerDiscoveryConfig{
		DiscoveryStore: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["discovery_store"]),
		EnableAWSOIDCProvider: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["enable_aws_oidc_provider"]),
	}
}

func FlattenResourceServiceAccountIssuerDiscoveryConfigInto(in kops.ServiceAccountIssuerDiscoveryConfig, out map[string]interface{}) {
	out["discovery_store"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.DiscoveryStore)
	out["enable_aws_oidc_provider"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.EnableAWSOIDCProvider)
}

func FlattenResourceServiceAccountIssuerDiscoveryConfig(in kops.ServiceAccountIssuerDiscoveryConfig) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceServiceAccountIssuerDiscoveryConfigInto(in, out)
	return out
}
