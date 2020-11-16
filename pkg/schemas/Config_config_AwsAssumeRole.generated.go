package schemas

import (
	"github.com/eddycharly/terraform-provider-kops/pkg/api/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ConfigConfigAwsAssumeRole() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"role_arn": OptionalString(),
		},
	}
}

func ExpandConfigConfigAwsAssumeRole(in map[string]interface{}) config.AwsAssumeRole {
	if in == nil {
		panic("expand AwsAssumeRole failure, in is nil")
	}
	return config.AwsAssumeRole{
		RoleArn: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["role_arn"]),
	}
}

func FlattenConfigConfigAwsAssumeRoleInto(in config.AwsAssumeRole, out map[string]interface{}) {
	out["role_arn"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.RoleArn)
}

func FlattenConfigConfigAwsAssumeRole(in config.AwsAssumeRole) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenConfigConfigAwsAssumeRoleInto(in, out)
	return out
}
