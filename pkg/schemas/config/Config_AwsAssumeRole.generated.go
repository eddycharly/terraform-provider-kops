package schemas

import (
	"github.com/eddycharly/terraform-provider-kops/pkg/api/config"
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var _ = Schema

func ConfigAwsAssumeRole() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"role_arn": Optional(String()),
		},
	}
}

func ExpandConfigAwsAssumeRole(in map[string]interface{}) config.AwsAssumeRole {
	if in == nil {
		panic("expand AwsAssumeRole failure, in is nil")
	}
	out := config.AwsAssumeRole{}
	if in, ok := in["role_arn"]; ok && in != nil {
		out.RoleArn = func(in interface{}) string { return string(in.(string)) }(in)
	}
	return out
}
