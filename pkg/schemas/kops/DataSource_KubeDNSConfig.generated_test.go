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
	_default := map[string]interface{}{
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
			name: "Image - default",
			args: args{
				in: func() kops.KubeDNSConfig {
					subject := kops.KubeDNSConfig{}
					subject.Image = ""
					return subject
				}(),
			},
			want: _default,
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
			name: "CPURequest - default",
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
			FlattenDataSourceKubeDNSConfigInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceKubeDNSConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceKubeDNSConfig(t *testing.T) {
	_default := map[string]interface{}{
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
			name: "Image - default",
			args: args{
				in: func() kops.KubeDNSConfig {
					subject := kops.KubeDNSConfig{}
					subject.Image = ""
					return subject
				}(),
			},
			want: _default,
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
			name: "CPURequest - default",
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
			got := FlattenDataSourceKubeDNSConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceKubeDNSConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
