package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceKubeDNSConfig(t *testing.T) {
	_default := kops.KubeDNSConfig{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.KubeDNSConfig
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"cache_max_size":       0,
					"cache_max_concurrent": 0,
					"tolerations":          func() []interface{} { return nil }(),
					"affinity":             nil,
					"core_dns_image":       "",
					"cpa_image":            "",
					"domain":               "",
					"external_core_file":   "",
					"provider":             "",
					"server_ip":            "",
					"stub_domains":         func() []interface{} { return nil }(),
					"upstream_nameservers": func() []interface{} { return nil }(),
					"memory_request":       nil,
					"cpu_request":          nil,
					"memory_limit":         nil,
					"node_local_dns":       nil,
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandResourceKubeDNSConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceKubeDNSConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceKubeDNSConfigInto(t *testing.T) {
	_default := map[string]interface{}{
		"cache_max_size":       0,
		"cache_max_concurrent": 0,
		"tolerations":          func() []interface{} { return nil }(),
		"affinity":             nil,
		"core_dns_image":       "",
		"cpa_image":            "",
		"domain":               "",
		"external_core_file":   "",
		"provider":             "",
		"server_ip":            "",
		"stub_domains":         func() []interface{} { return nil }(),
		"upstream_nameservers": func() []interface{} { return nil }(),
		"memory_request":       nil,
		"cpu_request":          nil,
		"memory_limit":         nil,
		"node_local_dns":       nil,
	}
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
			want: _default,
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
			want: _default,
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
			want: _default,
		},
		{
			name: "Tolerations - default",
			args: args{
				in: func() kops.KubeDNSConfig {
					subject := kops.KubeDNSConfig{}
					subject.Tolerations = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Affinity - default",
			args: args{
				in: func() kops.KubeDNSConfig {
					subject := kops.KubeDNSConfig{}
					subject.Affinity = nil
					return subject
				}(),
			},
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
		},
		{
			name: "CpuRequest - default",
			args: args{
				in: func() kops.KubeDNSConfig {
					subject := kops.KubeDNSConfig{}
					subject.CPURequest = nil
					return subject
				}(),
			},
			want: _default,
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
			want: _default,
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
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenResourceKubeDNSConfigInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceKubeDNSConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceKubeDNSConfig(t *testing.T) {
	_default := map[string]interface{}{
		"cache_max_size":       0,
		"cache_max_concurrent": 0,
		"tolerations":          func() []interface{} { return nil }(),
		"affinity":             nil,
		"core_dns_image":       "",
		"cpa_image":            "",
		"domain":               "",
		"external_core_file":   "",
		"provider":             "",
		"server_ip":            "",
		"stub_domains":         func() []interface{} { return nil }(),
		"upstream_nameservers": func() []interface{} { return nil }(),
		"memory_request":       nil,
		"cpu_request":          nil,
		"memory_limit":         nil,
		"node_local_dns":       nil,
	}
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
			want: _default,
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
			want: _default,
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
			want: _default,
		},
		{
			name: "Tolerations - default",
			args: args{
				in: func() kops.KubeDNSConfig {
					subject := kops.KubeDNSConfig{}
					subject.Tolerations = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Affinity - default",
			args: args{
				in: func() kops.KubeDNSConfig {
					subject := kops.KubeDNSConfig{}
					subject.Affinity = nil
					return subject
				}(),
			},
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
		},
		{
			name: "CpuRequest - default",
			args: args{
				in: func() kops.KubeDNSConfig {
					subject := kops.KubeDNSConfig{}
					subject.CPURequest = nil
					return subject
				}(),
			},
			want: _default,
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
			want: _default,
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
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceKubeDNSConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceKubeDNSConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
