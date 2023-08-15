package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

var _ = Schema

func ResourceEtcdBackupSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"backup_store": RequiredString(),
			"image":        RequiredString(),
		},
	}

	return res
}

func ExpandResourceEtcdBackupSpec(in map[string]interface{}) kopsv1alpha2.EtcdBackupSpec {
	if in == nil {
		panic("expand EtcdBackupSpec failure, in is nil")
	}
	return kopsv1alpha2.EtcdBackupSpec{
		BackupStore: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["backup_store"]),
		Image: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["image"]),
	}
}

func FlattenResourceEtcdBackupSpecInto(in kopsv1alpha2.EtcdBackupSpec, out map[string]interface{}) {
	out["backup_store"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.BackupStore)
	out["image"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Image)
}

func FlattenResourceEtcdBackupSpec(in kopsv1alpha2.EtcdBackupSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceEtcdBackupSpecInto(in, out)
	return out
}
