package structures

import (
	"k8s.io/kops/pkg/apis/kops"
)

func ExpandTopologySpec(in map[string]interface{}) kops.TopologySpec {
	if in == nil {
		panic("expand TopologySpec failure, in is nil")
	}
	return kops.TopologySpec{
		Masters: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["masters"]),
		Nodes: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["nodes"]),
		Bastion: func(in interface{}) *kops.BastionSpec {
			value := func(in interface{}) *kops.BastionSpec {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in kops.BastionSpec) *kops.BastionSpec {
					return &in
				}(func(in interface{}) kops.BastionSpec {
					if in.([]interface{})[0] == nil {
						return kops.BastionSpec{}
					}
					return (ExpandBastionSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
			return value
		}(in["bastion"]),
		DNS: func(in interface{}) *kops.DNSSpec {
			value := func(in interface{}) *kops.DNSSpec {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				return func(in kops.DNSSpec) *kops.DNSSpec {
					return &in
				}(func(in interface{}) kops.DNSSpec {
					if in.([]interface{})[0] == nil {
						return kops.DNSSpec{}
					}
					return (ExpandDNSSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
			return value
		}(in["dns"]),
	}
}

func FlattenTopologySpec(in kops.TopologySpec) map[string]interface{} {
	return map[string]interface{}{
		"masters": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.Masters),
		"nodes": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.Nodes),
		"bastion": func(in *kops.BastionSpec) interface{} {
			value := func(in *kops.BastionSpec) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.BastionSpec) interface{} {
					return func(in kops.BastionSpec) []map[string]interface{} {
						return []map[string]interface{}{FlattenBastionSpec(in)}
					}(in)
				}(*in)
			}(in)
			return value
		}(in.Bastion),
		"dns": func(in *kops.DNSSpec) interface{} {
			value := func(in *kops.DNSSpec) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.DNSSpec) interface{} {
					return func(in kops.DNSSpec) []map[string]interface{} {
						return []map[string]interface{}{FlattenDNSSpec(in)}
					}(in)
				}(*in)
			}(in)
			return value
		}(in.DNS),
	}
}
