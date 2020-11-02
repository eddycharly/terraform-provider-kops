package structures

import (
	"log"

	"k8s.io/kops/pkg/apis/kops"
)

func ExpandEtcdBackupSpec(in map[string]interface{}) kops.EtcdBackupSpec {
	if in == nil {
		panic("expand EtcdBackupSpec failure, in is nil")
	}
	return kops.EtcdBackupSpec{
		BackupStore: func(in interface{}) string {
			value := string(ExpandString(in))
			log.Printf("%s - %#v", "backup_store", value)
			return value
		}(in["backup_store"]),
		Image: func(in interface{}) string {
			value := string(ExpandString(in))
			log.Printf("%s - %#v", "image", value)
			return value
		}(in["image"]),
	}
}

func FlattenEtcdBackupSpec(in kops.EtcdBackupSpec) map[string]interface{} {
	return map[string]interface{}{
		"backup_store": func(in string) interface{} {
			value := FlattenString(string(in))
			log.Printf("%s - %v", "backup_store", value)
			return value
		}(in.BackupStore),
		"image": func(in string) interface{} {
			value := FlattenString(string(in))
			log.Printf("%s - %v", "image", value)
			return value
		}(in.Image),
	}
}
