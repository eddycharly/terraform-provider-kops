package structures

import (
	"github.com/eddycharly/terraform-provider-kops/pkg/api"
)

func ExpandAwsConfig(in map[string]interface{}) api.AwsConfig {
	if in == nil {
		panic("expand AwsConfig failure, in is nil")
	}
	return api.AwsConfig{
		Profile: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["profile"]),
	}
}

func FlattenAwsConfig(in api.AwsConfig) map[string]interface{} {
	return map[string]interface{}{
		"profile": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.Profile),
	}
}
