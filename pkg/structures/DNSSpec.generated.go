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
			return kops.DNSType(ExpandString(in))
		}(in["type"]),
	}
}

func FlattenDNSSpecInto(in kops.DNSSpec, out map[string]interface{}) {
	out["type"] = func(in kops.DNSType) interface{} {
		return FlattenString(string(in))
	}(in.Type)
}

func FlattenDNSSpec(in kops.DNSSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDNSSpecInto(in, out)
	return out
}
