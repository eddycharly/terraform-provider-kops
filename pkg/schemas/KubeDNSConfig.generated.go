package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func KubeDNSConfig() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cache_max_size":       OptionalInt(),
			"cache_max_concurrent": OptionalInt(),
			"core_dns_image":       OptionalString(),
			"domain":               OptionalString(),
			"external_core_file":   OptionalString(),
			"image":                OptionalString(),
			"replicas":             OptionalInt(),
			"provider":             OptionalString(),
			"server_ip":            OptionalString(),
			"stub_domains":         OptionalMap(List(String())),
			"upstream_nameservers": OptionalList(String()),
			"memory_request":       OptionalQuantity(),
			"cpu_request":          OptionalQuantity(),
			"memory_limit":         OptionalQuantity(),
			"node_local_dns":       OptionalStruct(NodeLocalDNSConfig()),
		},
	}
}
