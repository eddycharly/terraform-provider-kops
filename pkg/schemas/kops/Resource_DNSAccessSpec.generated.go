package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceDNSAccessSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{},
	}
}

func ExpandResourceDNSAccessSpec(in map[string]interface{}) kops.DNSAccessSpec {
	if in == nil {
		panic("expand DNSAccessSpec failure, in is nil")
	}
	return kops.DNSAccessSpec{}
}

func FlattenResourceDNSAccessSpecInto(in kops.DNSAccessSpec, out map[string]interface{}) {
}

func FlattenResourceDNSAccessSpec(in kops.DNSAccessSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceDNSAccessSpecInto(in, out)
	return out
}
