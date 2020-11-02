package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func EtcdMemberSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name":             StringRequired(),
			"instance_group":   StringRequired(),
			"volume_type":      StringEnumOptionalComputed(),
			"volume_iops":      IntOptionalComputed(),
			"volume_size":      IntOptionalComputed(),
			"kms_key_id":       StringOptionalComputed(),
			"encrypted_volume": BoolOptionalComputed(),
		},
	}
}
