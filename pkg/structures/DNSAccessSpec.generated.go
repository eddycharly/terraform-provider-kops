package structures

import (
	"k8s.io/kops/pkg/apis/kops"
)

func ExpandDNSAccessSpec(in map[string]interface{}) kops.DNSAccessSpec {
	if in == nil {
		panic("expand DNSAccessSpec failure, in is nil")
	}
	return kops.DNSAccessSpec{}
}

func FlattenDNSAccessSpec(in kops.DNSAccessSpec) map[string]interface{} {
	return map[string]interface{}{}
}
