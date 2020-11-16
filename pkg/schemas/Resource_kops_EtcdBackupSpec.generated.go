package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

func ResourceKopsEtcdBackupSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"backup_store": RequiredString(),
			"image":        RequiredString(),
		},
	}
}

func ExpandResourceKopsEtcdBackupSpec(in map[string]interface{}) kops.EtcdBackupSpec {
	if in == nil {
		panic("expand EtcdBackupSpec failure, in is nil")
	}
	return kops.EtcdBackupSpec{
		BackupStore: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["backup_store"]),
		Image: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["image"]),
	}
}

func FlattenResourceKopsEtcdBackupSpecInto(in kops.EtcdBackupSpec, out map[string]interface{}) {
	out["backup_store"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.BackupStore)
	out["image"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Image)
}

func FlattenResourceKopsEtcdBackupSpec(in kops.EtcdBackupSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceKopsEtcdBackupSpecInto(in, out)
	return out
}
