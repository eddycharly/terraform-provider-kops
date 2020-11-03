package structures

import (
	"reflect"

	"github.com/eddycharly/terraform-provider-kops/pkg/api"
	"k8s.io/kops/pkg/apis/kops"
)

func ExpandCluster(in map[string]interface{}) api.Cluster {
	if in == nil {
		panic("expand Cluster failure, in is nil")
	}
	return api.Cluster{
		Name: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["name"]),
		Channel: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["channel"]),
		ConfigBase: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["config_base"]),
		CloudProvider: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["cloud_provider"]),
		ContainerRuntime: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["container_runtime"]),
		KubernetesVersion: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["kubernetes_version"]),
		Subnet: func(in interface{}) []kops.ClusterSubnetSpec {
			value := func(in interface{}) []kops.ClusterSubnetSpec {
				var out []kops.ClusterSubnetSpec
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) kops.ClusterSubnetSpec {
						if in == nil {
							return kops.ClusterSubnetSpec{}
						}
						return (ExpandClusterSubnetSpec(in.(map[string]interface{})))
					}(in))
				}
				return out
			}(in)
			return value
		}(in["subnet"]),
		Project: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["project"]),
		MasterPublicName: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["master_public_name"]),
		MasterInternalName: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["master_internal_name"]),
		NetworkCIDR: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["network_cidr"]),
		AdditionalNetworkCIDRs: func(in interface{}) []string {
			value := func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
			return value
		}(in["additional_network_cidrs"]),
		NetworkID: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["network_id"]),
		Topology: func(in interface{}) *kops.TopologySpec {
			value := func(in interface{}) *kops.TopologySpec {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in kops.TopologySpec) *kops.TopologySpec {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(func(in interface{}) kops.TopologySpec {
					if in.([]interface{})[0] == nil {
						return kops.TopologySpec{}
					}
					return (ExpandTopologySpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
				return tmp
			}(in)
			return value
		}(in["topology"]),
		SecretStore: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["secret_store"]),
		KeyStore: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["key_store"]),
		ConfigStore: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["config_store"]),
		DNSZone: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["dns_zone"]),
		ClusterDNSDomain: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["cluster_dns_domain"]),
		ServiceClusterIPRange: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["service_cluster_ip_range"]),
		PodCIDR: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["pod_cidr"]),
		NonMasqueradeCIDR: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["non_masquerade_cidr"]),
		SSHAccess: func(in interface{}) []string {
			value := func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
			return value
		}(in["ssh_access"]),
		NodePortAccess: func(in interface{}) []string {
			value := func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
			return value
		}(in["node_port_access"]),
		SSHKeyName: func(in interface{}) *string {
			value := func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in string) *string {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(string(ExpandString(in)))
				return tmp
			}(in)
			return value
		}(in["ssh_key_name"]),
		KubernetesAPIAccess: func(in interface{}) []string {
			value := func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
			return value
		}(in["kubernetes_api_access"]),
		IsolateMasters: func(in interface{}) *bool {
			value := func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in bool) *bool {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(bool(ExpandBool(in)))
				return tmp
			}(in)
			return value
		}(in["isolate_masters"]),
		UpdatePolicy: func(in interface{}) *string {
			value := func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in string) *string {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(string(ExpandString(in)))
				return tmp
			}(in)
			return value
		}(in["update_policy"]),
		ExternalPolicies: func(in interface{}) *map[string][]string {
			value := func(in interface{}) *map[string][]string {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in map[string][]string) *map[string][]string {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(func(in interface{}) map[string][]string {
					if in == nil {
						return nil
					}
					out := map[string][]string{}
					for key, in := range in.(map[string]interface{}) {
						out[key] = func(in interface{}) []string {
							var out []string
							for _, in := range in.([]interface{}) {
								out = append(out, string(ExpandString(in)))
							}
							return out
						}(in)
					}
					return out
				}(in))
				return tmp
			}(in)
			return value
		}(in["external_policies"]),
		AdditionalPolicies: func(in interface{}) *map[string]string {
			value := func(in interface{}) *map[string]string {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in map[string]string) *map[string]string {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(func(in interface{}) map[string]string {
					if in == nil {
						return nil
					}
					out := map[string]string{}
					for key, in := range in.(map[string]interface{}) {
						out[key] = string(ExpandString(in))
					}
					return out
				}(in))
				return tmp
			}(in)
			return value
		}(in["additional_policies"]),
		EtcdCluster: func(in interface{}) []*kops.EtcdClusterSpec {
			value := func(in interface{}) []*kops.EtcdClusterSpec {
				var out []*kops.EtcdClusterSpec
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) *kops.EtcdClusterSpec {
						if in == nil {
							return nil
						}
						if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
							return nil
						}
						tmp := func(in kops.EtcdClusterSpec) *kops.EtcdClusterSpec {
							if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
								return nil
							}
							return &in
						}(func(in interface{}) kops.EtcdClusterSpec {
							if in == nil {
								return kops.EtcdClusterSpec{}
							}
							return (ExpandEtcdClusterSpec(in.(map[string]interface{})))
						}(in))
						return tmp
					}(in))
				}
				return out
			}(in)
			return value
		}(in["etcd_cluster"]),
		Networking: func(in interface{}) *kops.NetworkingSpec {
			value := func(in interface{}) *kops.NetworkingSpec {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in kops.NetworkingSpec) *kops.NetworkingSpec {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(func(in interface{}) kops.NetworkingSpec {
					if in.([]interface{})[0] == nil {
						return kops.NetworkingSpec{}
					}
					return (ExpandNetworkingSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
				return tmp
			}(in)
			return value
		}(in["networking"]),
		API: func(in interface{}) *kops.AccessSpec {
			value := func(in interface{}) *kops.AccessSpec {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in kops.AccessSpec) *kops.AccessSpec {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(func(in interface{}) kops.AccessSpec {
					if in.([]interface{})[0] == nil {
						return kops.AccessSpec{}
					}
					return (ExpandAccessSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
				return tmp
			}(in)
			return value
		}(in["api"]),
		CloudLabels: func(in interface{}) map[string]string {
			value := func(in interface{}) map[string]string {
				if in == nil {
					return nil
				}
				out := map[string]string{}
				for key, in := range in.(map[string]interface{}) {
					out[key] = string(ExpandString(in))
				}
				return out
			}(in)
			return value
		}(in["cloud_labels"]),
		EncryptionConfig: func(in interface{}) *bool {
			value := func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in bool) *bool {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(bool(ExpandBool(in)))
				return tmp
			}(in)
			return value
		}(in["encryption_config"]),
		DisableSubnetTags: func(in interface{}) bool {
			value := bool(ExpandBool(in))
			return value
		}(in["disable_subnet_tags"]),
		UseHostCertificates: func(in interface{}) *bool {
			value := func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in bool) *bool {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(bool(ExpandBool(in)))
				return tmp
			}(in)
			return value
		}(in["use_host_certificates"]),
		SysctlParameters: func(in interface{}) []string {
			value := func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
			return value
		}(in["sysctl_parameters"]),
		RollingUpdate: func(in interface{}) *kops.RollingUpdate {
			value := func(in interface{}) *kops.RollingUpdate {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in kops.RollingUpdate) *kops.RollingUpdate {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(func(in interface{}) kops.RollingUpdate {
					if in.([]interface{})[0] == nil {
						return kops.RollingUpdate{}
					}
					return (ExpandRollingUpdate(in.([]interface{})[0].(map[string]interface{})))
				}(in))
				return tmp
			}(in)
			return value
		}(in["rolling_update"]),
		InstanceGroup: func(in interface{}) []*api.InstanceGroup {
			value := func(in interface{}) []*api.InstanceGroup {
				var out []*api.InstanceGroup
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) *api.InstanceGroup {
						if in == nil {
							return nil
						}
						if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
							return nil
						}
						tmp := func(in api.InstanceGroup) *api.InstanceGroup {
							if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
								return nil
							}
							return &in
						}(func(in interface{}) api.InstanceGroup {
							if in == nil {
								return api.InstanceGroup{}
							}
							return (ExpandInstanceGroup(in.(map[string]interface{})))
						}(in))
						return tmp
					}(in))
				}
				return out
			}(in)
			return value
		}(in["instance_group"]),
	}
}

func FlattenCluster(in api.Cluster) map[string]interface{} {
	return map[string]interface{}{
		"name": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.Name),
		"channel": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.Channel),
		"config_base": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.ConfigBase),
		"cloud_provider": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.CloudProvider),
		"container_runtime": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.ContainerRuntime),
		"kubernetes_version": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.KubernetesVersion),
		"subnet": func(in []kops.ClusterSubnetSpec) interface{} {
			value := func(in []kops.ClusterSubnetSpec) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, func(in kops.ClusterSubnetSpec) interface{} {
						return FlattenClusterSubnetSpec(in)
					}(in))
				}
				return out
			}(in)
			return value
		}(in.Subnet),
		"project": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.Project),
		"master_public_name": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.MasterPublicName),
		"master_internal_name": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.MasterInternalName),
		"network_cidr": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.NetworkCIDR),
		"additional_network_cidrs": func(in []string) interface{} {
			value := func(in []string) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, FlattenString(string(in)))
				}
				return out
			}(in)
			return value
		}(in.AdditionalNetworkCIDRs),
		"network_id": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.NetworkID),
		"topology": func(in *kops.TopologySpec) interface{} {
			value := func(in *kops.TopologySpec) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.TopologySpec) interface{} {
					return func(in kops.TopologySpec) []map[string]interface{} {
						return []map[string]interface{}{FlattenTopologySpec(in)}
					}(in)
				}(*in)
			}(in)
			return value
		}(in.Topology),
		"secret_store": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.SecretStore),
		"key_store": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.KeyStore),
		"config_store": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.ConfigStore),
		"dns_zone": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.DNSZone),
		"cluster_dns_domain": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.ClusterDNSDomain),
		"service_cluster_ip_range": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.ServiceClusterIPRange),
		"pod_cidr": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.PodCIDR),
		"non_masquerade_cidr": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.NonMasqueradeCIDR),
		"ssh_access": func(in []string) interface{} {
			value := func(in []string) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, FlattenString(string(in)))
				}
				return out
			}(in)
			return value
		}(in.SSHAccess),
		"node_port_access": func(in []string) interface{} {
			value := func(in []string) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, FlattenString(string(in)))
				}
				return out
			}(in)
			return value
		}(in.NodePortAccess),
		"ssh_key_name": func(in *string) interface{} {
			value := func(in *string) interface{} {
				if in == nil {
					return nil
				}
				return func(in string) interface{} {
					return FlattenString(string(in))
				}(*in)
			}(in)
			return value
		}(in.SSHKeyName),
		"kubernetes_api_access": func(in []string) interface{} {
			value := func(in []string) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, FlattenString(string(in)))
				}
				return out
			}(in)
			return value
		}(in.KubernetesAPIAccess),
		"isolate_masters": func(in *bool) interface{} {
			value := func(in *bool) interface{} {
				if in == nil {
					return nil
				}
				return func(in bool) interface{} {
					return FlattenBool(bool(in))
				}(*in)
			}(in)
			return value
		}(in.IsolateMasters),
		"update_policy": func(in *string) interface{} {
			value := func(in *string) interface{} {
				if in == nil {
					return nil
				}
				return func(in string) interface{} {
					return FlattenString(string(in))
				}(*in)
			}(in)
			return value
		}(in.UpdatePolicy),
		"external_policies": func(in *map[string][]string) interface{} {
			value := func(in *map[string][]string) interface{} {
				if in == nil {
					return nil
				}
				return func(in map[string][]string) interface{} {
					return func(in map[string][]string) map[string]interface{} {
						if in == nil {
							return nil
						}
						// TODO
						return nil
					}(in)
				}(*in)
			}(in)
			return value
		}(in.ExternalPolicies),
		"additional_policies": func(in *map[string]string) interface{} {
			value := func(in *map[string]string) interface{} {
				if in == nil {
					return nil
				}
				return func(in map[string]string) interface{} {
					return func(in map[string]string) map[string]interface{} {
						if in == nil {
							return nil
						}
						// TODO
						return nil
					}(in)
				}(*in)
			}(in)
			return value
		}(in.AdditionalPolicies),
		"etcd_cluster": func(in []*kops.EtcdClusterSpec) interface{} {
			value := func(in []*kops.EtcdClusterSpec) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, func(in *kops.EtcdClusterSpec) interface{} {
						if in == nil {
							return nil
						}
						return func(in kops.EtcdClusterSpec) interface{} {
							return func(in kops.EtcdClusterSpec) interface{} {
								return FlattenEtcdClusterSpec(in)
							}(in)
						}(*in)
					}(in))
				}
				return out
			}(in)
			return value
		}(in.EtcdCluster),
		"networking": func(in *kops.NetworkingSpec) interface{} {
			value := func(in *kops.NetworkingSpec) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.NetworkingSpec) interface{} {
					return func(in kops.NetworkingSpec) []map[string]interface{} {
						return []map[string]interface{}{FlattenNetworkingSpec(in)}
					}(in)
				}(*in)
			}(in)
			return value
		}(in.Networking),
		"api": func(in *kops.AccessSpec) interface{} {
			value := func(in *kops.AccessSpec) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.AccessSpec) interface{} {
					return func(in kops.AccessSpec) []map[string]interface{} {
						return []map[string]interface{}{FlattenAccessSpec(in)}
					}(in)
				}(*in)
			}(in)
			return value
		}(in.API),
		"cloud_labels": func(in map[string]string) interface{} {
			value := func(in map[string]string) map[string]interface{} {
				if in == nil {
					return nil
				}
				// TODO
				return nil
			}(in)
			return value
		}(in.CloudLabels),
		"encryption_config": func(in *bool) interface{} {
			value := func(in *bool) interface{} {
				if in == nil {
					return nil
				}
				return func(in bool) interface{} {
					return FlattenBool(bool(in))
				}(*in)
			}(in)
			return value
		}(in.EncryptionConfig),
		"disable_subnet_tags": func(in bool) interface{} {
			value := FlattenBool(bool(in))
			return value
		}(in.DisableSubnetTags),
		"use_host_certificates": func(in *bool) interface{} {
			value := func(in *bool) interface{} {
				if in == nil {
					return nil
				}
				return func(in bool) interface{} {
					return FlattenBool(bool(in))
				}(*in)
			}(in)
			return value
		}(in.UseHostCertificates),
		"sysctl_parameters": func(in []string) interface{} {
			value := func(in []string) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, FlattenString(string(in)))
				}
				return out
			}(in)
			return value
		}(in.SysctlParameters),
		"rolling_update": func(in *kops.RollingUpdate) interface{} {
			value := func(in *kops.RollingUpdate) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.RollingUpdate) interface{} {
					return func(in kops.RollingUpdate) []map[string]interface{} {
						return []map[string]interface{}{FlattenRollingUpdate(in)}
					}(in)
				}(*in)
			}(in)
			return value
		}(in.RollingUpdate),
		"instance_group": func(in []*api.InstanceGroup) interface{} {
			value := func(in []*api.InstanceGroup) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, func(in *api.InstanceGroup) interface{} {
						if in == nil {
							return nil
						}
						return func(in api.InstanceGroup) interface{} {
							return func(in api.InstanceGroup) interface{} {
								return FlattenInstanceGroup(in)
							}(in)
						}(*in)
					}(in))
				}
				return out
			}(in)
			return value
		}(in.InstanceGroup),
	}
}
