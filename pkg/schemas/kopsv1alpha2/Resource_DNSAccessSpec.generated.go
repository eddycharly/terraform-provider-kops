package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

var _ = Schema

func ResourceDNSAccessSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{},
	}

	return res
}

func ExpandResourceDNSAccessSpec(in map[string]interface{}) kopsv1alpha2.DNSAccessSpec {
	if in == nil {
		panic("expand DNSAccessSpec failure, in is nil")
	}
	return kopsv1alpha2.DNSAccessSpec{}
}

func FlattenResourceDNSAccessSpecInto(in kopsv1alpha2.DNSAccessSpec, out map[string]interface{}) {
}

func FlattenResourceDNSAccessSpec(in kopsv1alpha2.DNSAccessSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceDNSAccessSpecInto(in, out)
	return out
}
