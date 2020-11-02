package structures

import (
	"k8s.io/kops/pkg/apis/kops"
)

func ExpandDNSSpec(in map[string]interface{}) kops.DNSSpec {
	if in == nil {
		panic("expand DNSSpec failure, in is nil")
	}
	return kops.DNSSpec{
		Type: func(in interface{}) kops.DNSType {
			value := kops.DNSType(ExpandString(in))
			return value
		}(in["type"]),
	}
}

func FlattenDNSSpec(in kops.DNSSpec) map[string]interface{} {
	return map[string]interface{}{
		"type": func(in kops.DNSType) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.Type),
	}
}
