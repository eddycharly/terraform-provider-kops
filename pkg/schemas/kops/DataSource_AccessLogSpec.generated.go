package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceAccessLogSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"interval":      ComputedInt(),
			"bucket":        ComputedString(),
			"bucket_prefix": ComputedString(),
		},
	}

	return res
}

func ExpandDataSourceAccessLogSpec(in map[string]interface{}) kops.AccessLogSpec {
	if in == nil {
		panic("expand AccessLogSpec failure, in is nil")
	}
	return kops.AccessLogSpec{
		Interval: func(in interface{}) int {
			return int(ExpandInt(in))
		}(in["interval"]),
		Bucket: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["bucket"]),
		BucketPrefix: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["bucket_prefix"]),
	}
}

func FlattenDataSourceAccessLogSpecInto(in kops.AccessLogSpec, out map[string]interface{}) {
	out["interval"] = func(in int) interface{} {
		return FlattenInt(int(in))
	}(in.Interval)
	out["bucket"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Bucket)
	out["bucket_prefix"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.BucketPrefix)
}

func FlattenDataSourceAccessLogSpec(in kops.AccessLogSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceAccessLogSpecInto(in, out)
	return out
}
