package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

func DataSourceKopsEtcdBackupSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"backup_store": ComputedString(),
			"image":        ComputedString(),
		},
	}
}

func ExpandDataSourceKopsEtcdBackupSpec(in map[string]interface{}) kops.EtcdBackupSpec {
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

func FlattenDataSourceKopsEtcdBackupSpecInto(in kops.EtcdBackupSpec, out map[string]interface{}) {
	out["backup_store"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.BackupStore)
	out["image"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Image)
}

func FlattenDataSourceKopsEtcdBackupSpec(in kops.EtcdBackupSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceKopsEtcdBackupSpecInto(in, out)
	return out
}
