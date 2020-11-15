package structures

import (
	"github.com/eddycharly/terraform-provider-kops/pkg/api"
)

func ExpandAwsAssumeRole(in map[string]interface{}) api.AwsAssumeRole {
	if in == nil {
		panic("expand AwsAssumeRole failure, in is nil")
	}
	return api.AwsAssumeRole{
		RoleArn: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["role_arn"]),
	}
}

func FlattenAwsAssumeRoleInto(in api.AwsAssumeRole, out map[string]interface{}) {
	out["role_arn"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.RoleArn)
}

func FlattenAwsAssumeRole(in api.AwsAssumeRole) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenAwsAssumeRoleInto(in, out)
	return out
}
