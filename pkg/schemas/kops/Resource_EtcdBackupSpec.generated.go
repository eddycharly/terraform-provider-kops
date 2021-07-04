package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceEtcdBackupSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"backup_store": Required(String()),
			"image":        Required(String()),
		},
	}
}

func ExpandResourceEtcdBackupSpec(in map[string]interface{}) kops.EtcdBackupSpec {
	if in == nil {
		panic("expand EtcdBackupSpec failure, in is nil")
	}
	out := kops.EtcdBackupSpec{}
	if in, ok := in["backup_store"]; ok && in != nil {
		out.BackupStore = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["image"]; ok && in != nil {
		out.Image = func(in interface{}) string { return string(in.(string)) }(in)
	}
	return out
}

func FlattenResourceEtcdBackupSpecInto(in kops.EtcdBackupSpec, out map[string]interface{}) {
	out["backup_store"] = func(in string) interface{} { return string(in) }(in.BackupStore)
	out["image"] = func(in string) interface{} { return string(in) }(in.Image)
}

func FlattenResourceEtcdBackupSpec(in kops.EtcdBackupSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceEtcdBackupSpecInto(in, out)
	return out
}
