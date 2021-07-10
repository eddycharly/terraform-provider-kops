package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceKubeDNSConfig(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.KubeDNSConfig
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandDataSourceKubeDNSConfig(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandDataSourceKubeDNSConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenDataSourceKubeDNSConfigInto(t *testing.T) {
	type args struct {
		in  kops.KubeDNSConfig
		out map[string]interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			FlattenDataSourceKubeDNSConfigInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenDataSourceKubeDNSConfig(t *testing.T) {
	type args struct {
		in kops.KubeDNSConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.KubeDNSConfig{},
			},
			want: map[string]interface{}{
				"cache_max_size":       0,
				"cache_max_concurrent": 0,
				"core_dns_image":       "",
				"cpa_image":            "",
				"domain":               "",
				"external_core_file":   "",
				"image":                "",
				"replicas":             0,
				"provider":             "",
				"server_ip":            "",
				"stub_domains":         func() map[string]interface{} { return nil }(),
				"upstream_nameservers": func() []interface{} { return nil }(),
				"memory_request":       nil,
				"cpu_request":          nil,
				"memory_limit":         nil,
				"node_local_dns":       nil,
			},
		},
		{
			name: "CacheMaxSize - default",
			args: args{
				in: func() kops.KubeDNSConfig {
					subject := kops.KubeDNSConfig{}
					subject.CacheMaxSize = 0
					return subject
				}(),
			},
			want: map[string]interface{}{
				"cache_max_size":       0,
				"cache_max_concurrent": 0,
				"core_dns_image":       "",
				"cpa_image":            "",
				"domain":               "",
				"external_core_file":   "",
				"image":                "",
				"replicas":             0,
				"provider":             "",
				"server_ip":            "",
				"stub_domains":         func() map[string]interface{} { return nil }(),
				"upstream_nameservers": func() []interface{} { return nil }(),
				"memory_request":       nil,
				"cpu_request":          nil,
				"memory_limit":         nil,
				"node_local_dns":       nil,
			},
		},
		{
			name: "CacheMaxConcurrent - default",
			args: args{
				in: func() kops.KubeDNSConfig {
					subject := kops.KubeDNSConfig{}
					subject.CacheMaxConcurrent = 0
					return subject
				}(),
			},
			want: map[string]interface{}{
				"cache_max_size":       0,
				"cache_max_concurrent": 0,
				"core_dns_image":       "",
				"cpa_image":            "",
				"domain":               "",
				"external_core_file":   "",
				"image":                "",
				"replicas":             0,
				"provider":             "",
				"server_ip":            "",
				"stub_domains":         func() map[string]interface{} { return nil }(),
				"upstream_nameservers": func() []interface{} { return nil }(),
				"memory_request":       nil,
				"cpu_request":          nil,
				"memory_limit":         nil,
				"node_local_dns":       nil,
			},
		},
		{
			name: "CoreDnsImage - default",
			args: args{
				in: func() kops.KubeDNSConfig {
					subject := kops.KubeDNSConfig{}
					subject.CoreDNSImage = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
				"cache_max_size":       0,
				"cache_max_concurrent": 0,
				"core_dns_image":       "",
				"cpa_image":            "",
				"domain":               "",
				"external_core_file":   "",
				"image":                "",
				"replicas":             0,
				"provider":             "",
				"server_ip":            "",
				"stub_domains":         func() map[string]interface{} { return nil }(),
				"upstream_nameservers": func() []interface{} { return nil }(),
				"memory_request":       nil,
				"cpu_request":          nil,
				"memory_limit":         nil,
				"node_local_dns":       nil,
			},
		},
		{
			name: "CPAImage - default",
			args: args{
				in: func() kops.KubeDNSConfig {
					subject := kops.KubeDNSConfig{}
					subject.CPAImage = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
				"cache_max_size":       0,
				"cache_max_concurrent": 0,
				"core_dns_image":       "",
				"cpa_image":            "",
				"domain":               "",
				"external_core_file":   "",
				"image":                "",
				"replicas":             0,
				"provider":             "",
				"server_ip":            "",
				"stub_domains":         func() map[string]interface{} { return nil }(),
				"upstream_nameservers": func() []interface{} { return nil }(),
				"memory_request":       nil,
				"cpu_request":          nil,
				"memory_limit":         nil,
				"node_local_dns":       nil,
			},
		},
		{
			name: "Domain - default",
			args: args{
				in: func() kops.KubeDNSConfig {
					subject := kops.KubeDNSConfig{}
					subject.Domain = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
				"cache_max_size":       0,
				"cache_max_concurrent": 0,
				"core_dns_image":       "",
				"cpa_image":            "",
				"domain":               "",
				"external_core_file":   "",
				"image":                "",
				"replicas":             0,
				"provider":             "",
				"server_ip":            "",
				"stub_domains":         func() map[string]interface{} { return nil }(),
				"upstream_nameservers": func() []interface{} { return nil }(),
				"memory_request":       nil,
				"cpu_request":          nil,
				"memory_limit":         nil,
				"node_local_dns":       nil,
			},
		},
		{
			name: "ExternalCoreFile - default",
			args: args{
				in: func() kops.KubeDNSConfig {
					subject := kops.KubeDNSConfig{}
					subject.ExternalCoreFile = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
				"cache_max_size":       0,
				"cache_max_concurrent": 0,
				"core_dns_image":       "",
				"cpa_image":            "",
				"domain":               "",
				"external_core_file":   "",
				"image":                "",
				"replicas":             0,
				"provider":             "",
				"server_ip":            "",
				"stub_domains":         func() map[string]interface{} { return nil }(),
				"upstream_nameservers": func() []interface{} { return nil }(),
				"memory_request":       nil,
				"cpu_request":          nil,
				"memory_limit":         nil,
				"node_local_dns":       nil,
			},
		},
		{
			name: "Image - default",
			args: args{
				in: func() kops.KubeDNSConfig {
					subject := kops.KubeDNSConfig{}
					subject.Image = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
				"cache_max_size":       0,
				"cache_max_concurrent": 0,
				"core_dns_image":       "",
				"cpa_image":            "",
				"domain":               "",
				"external_core_file":   "",
				"image":                "",
				"replicas":             0,
				"provider":             "",
				"server_ip":            "",
				"stub_domains":         func() map[string]interface{} { return nil }(),
				"upstream_nameservers": func() []interface{} { return nil }(),
				"memory_request":       nil,
				"cpu_request":          nil,
				"memory_limit":         nil,
				"node_local_dns":       nil,
			},
		},
		{
			name: "Replicas - default",
			args: args{
				in: func() kops.KubeDNSConfig {
					subject := kops.KubeDNSConfig{}
					subject.Replicas = 0
					return subject
				}(),
			},
			want: map[string]interface{}{
				"cache_max_size":       0,
				"cache_max_concurrent": 0,
				"core_dns_image":       "",
				"cpa_image":            "",
				"domain":               "",
				"external_core_file":   "",
				"image":                "",
				"replicas":             0,
				"provider":             "",
				"server_ip":            "",
				"stub_domains":         func() map[string]interface{} { return nil }(),
				"upstream_nameservers": func() []interface{} { return nil }(),
				"memory_request":       nil,
				"cpu_request":          nil,
				"memory_limit":         nil,
				"node_local_dns":       nil,
			},
		},
		{
			name: "Provider - default",
			args: args{
				in: func() kops.KubeDNSConfig {
					subject := kops.KubeDNSConfig{}
					subject.Provider = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
				"cache_max_size":       0,
				"cache_max_concurrent": 0,
				"core_dns_image":       "",
				"cpa_image":            "",
				"domain":               "",
				"external_core_file":   "",
				"image":                "",
				"replicas":             0,
				"provider":             "",
				"server_ip":            "",
				"stub_domains":         func() map[string]interface{} { return nil }(),
				"upstream_nameservers": func() []interface{} { return nil }(),
				"memory_request":       nil,
				"cpu_request":          nil,
				"memory_limit":         nil,
				"node_local_dns":       nil,
			},
		},
		{
			name: "ServerIp - default",
			args: args{
				in: func() kops.KubeDNSConfig {
					subject := kops.KubeDNSConfig{}
					subject.ServerIP = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
				"cache_max_size":       0,
				"cache_max_concurrent": 0,
				"core_dns_image":       "",
				"cpa_image":            "",
				"domain":               "",
				"external_core_file":   "",
				"image":                "",
				"replicas":             0,
				"provider":             "",
				"server_ip":            "",
				"stub_domains":         func() map[string]interface{} { return nil }(),
				"upstream_nameservers": func() []interface{} { return nil }(),
				"memory_request":       nil,
				"cpu_request":          nil,
				"memory_limit":         nil,
				"node_local_dns":       nil,
			},
		},
		{
			name: "StubDomains - default",
			args: args{
				in: func() kops.KubeDNSConfig {
					subject := kops.KubeDNSConfig{}
					subject.StubDomains = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"cache_max_size":       0,
				"cache_max_concurrent": 0,
				"core_dns_image":       "",
				"cpa_image":            "",
				"domain":               "",
				"external_core_file":   "",
				"image":                "",
				"replicas":             0,
				"provider":             "",
				"server_ip":            "",
				"stub_domains":         func() map[string]interface{} { return nil }(),
				"upstream_nameservers": func() []interface{} { return nil }(),
				"memory_request":       nil,
				"cpu_request":          nil,
				"memory_limit":         nil,
				"node_local_dns":       nil,
			},
		},
		{
			name: "UpstreamNameservers - default",
			args: args{
				in: func() kops.KubeDNSConfig {
					subject := kops.KubeDNSConfig{}
					subject.UpstreamNameservers = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"cache_max_size":       0,
				"cache_max_concurrent": 0,
				"core_dns_image":       "",
				"cpa_image":            "",
				"domain":               "",
				"external_core_file":   "",
				"image":                "",
				"replicas":             0,
				"provider":             "",
				"server_ip":            "",
				"stub_domains":         func() map[string]interface{} { return nil }(),
				"upstream_nameservers": func() []interface{} { return nil }(),
				"memory_request":       nil,
				"cpu_request":          nil,
				"memory_limit":         nil,
				"node_local_dns":       nil,
			},
		},
		{
			name: "MemoryRequest - default",
			args: args{
				in: func() kops.KubeDNSConfig {
					subject := kops.KubeDNSConfig{}
					subject.MemoryRequest = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"cache_max_size":       0,
				"cache_max_concurrent": 0,
				"core_dns_image":       "",
				"cpa_image":            "",
				"domain":               "",
				"external_core_file":   "",
				"image":                "",
				"replicas":             0,
				"provider":             "",
				"server_ip":            "",
				"stub_domains":         func() map[string]interface{} { return nil }(),
				"upstream_nameservers": func() []interface{} { return nil }(),
				"memory_request":       nil,
				"cpu_request":          nil,
				"memory_limit":         nil,
				"node_local_dns":       nil,
			},
		},
		{
			name: "CPURequest - default",
			args: args{
				in: func() kops.KubeDNSConfig {
					subject := kops.KubeDNSConfig{}
					subject.CPURequest = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"cache_max_size":       0,
				"cache_max_concurrent": 0,
				"core_dns_image":       "",
				"cpa_image":            "",
				"domain":               "",
				"external_core_file":   "",
				"image":                "",
				"replicas":             0,
				"provider":             "",
				"server_ip":            "",
				"stub_domains":         func() map[string]interface{} { return nil }(),
				"upstream_nameservers": func() []interface{} { return nil }(),
				"memory_request":       nil,
				"cpu_request":          nil,
				"memory_limit":         nil,
				"node_local_dns":       nil,
			},
		},
		{
			name: "MemoryLimit - default",
			args: args{
				in: func() kops.KubeDNSConfig {
					subject := kops.KubeDNSConfig{}
					subject.MemoryLimit = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"cache_max_size":       0,
				"cache_max_concurrent": 0,
				"core_dns_image":       "",
				"cpa_image":            "",
				"domain":               "",
				"external_core_file":   "",
				"image":                "",
				"replicas":             0,
				"provider":             "",
				"server_ip":            "",
				"stub_domains":         func() map[string]interface{} { return nil }(),
				"upstream_nameservers": func() []interface{} { return nil }(),
				"memory_request":       nil,
				"cpu_request":          nil,
				"memory_limit":         nil,
				"node_local_dns":       nil,
			},
		},
		{
			name: "NodeLocalDns - default",
			args: args{
				in: func() kops.KubeDNSConfig {
					subject := kops.KubeDNSConfig{}
					subject.NodeLocalDNS = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"cache_max_size":       0,
				"cache_max_concurrent": 0,
				"core_dns_image":       "",
				"cpa_image":            "",
				"domain":               "",
				"external_core_file":   "",
				"image":                "",
				"replicas":             0,
				"provider":             "",
				"server_ip":            "",
				"stub_domains":         func() map[string]interface{} { return nil }(),
				"upstream_nameservers": func() []interface{} { return nil }(),
				"memory_request":       nil,
				"cpu_request":          nil,
				"memory_limit":         nil,
				"node_local_dns":       nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenDataSourceKubeDNSConfig(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenDataSourceKubeDNSConfig() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
