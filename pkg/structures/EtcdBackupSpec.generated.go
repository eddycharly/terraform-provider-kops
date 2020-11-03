package structures

import (
	"k8s.io/kops/pkg/apis/kops"
)

func ExpandEtcdBackupSpec(in map[string]interface{}) kops.EtcdBackupSpec {
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

func FlattenEtcdBackupSpec(in kops.EtcdBackupSpec) map[string]interface{} {
	return map[string]interface{}{
		"backup_store": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.BackupStore),
		"image": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.Image),
	}
}
