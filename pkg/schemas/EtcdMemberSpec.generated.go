package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func EtcdMemberSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name":             RequiredString(),
			"instance_group":   RequiredString(),
			"volume_type":      OptionalString(),
			"volume_iops":      OptionalInt(),
			"volume_size":      OptionalInt(),
			"kms_key_id":       OptionalString(),
			"encrypted_volume": OptionalBool(),
		},
	}
}
