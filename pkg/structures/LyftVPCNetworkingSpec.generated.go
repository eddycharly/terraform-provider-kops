package structures

import (
	"log"

	"k8s.io/kops/pkg/apis/kops"
)

func ExpandLyftVPCNetworkingSpec(in map[string]interface{}) kops.LyftVPCNetworkingSpec {
	if in == nil {
		panic("expand LyftVPCNetworkingSpec failure, in is nil")
	}
	return kops.LyftVPCNetworkingSpec{
		SubnetTags: func(in interface{}) map[string]string {
			value := func(in interface{}) map[string]string {
				if in == nil {
					return nil
				}
				out := map[string]string{}
				for key, in := range in.(map[string]interface{}) {
					out[key] = string(ExpandString(in))
				}
				return out
			}(in)
			log.Printf("%s - %#v", "subnet_tags", value)
			return value
		}(in["subnet_tags"]),
	}
}

func FlattenLyftVPCNetworkingSpec(in kops.LyftVPCNetworkingSpec) map[string]interface{} {
	return map[string]interface{}{
		"subnet_tags": func(in map[string]string) interface{} {
			value := func(in map[string]string) map[string]interface{} {
				if in == nil {
					return nil
				}
				// TODO
				return nil
			}(in)
			log.Printf("%s - %v", "subnet_tags", value)
			return value
		}(in.SubnetTags),
	}
}
