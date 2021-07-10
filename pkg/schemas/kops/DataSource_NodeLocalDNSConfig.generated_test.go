package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceNodeLocalDNSConfig(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.NodeLocalDNSConfig
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandDataSourceNodeLocalDNSConfig(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandDataSourceNodeLocalDNSConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenDataSourceNodeLocalDNSConfigInto(t *testing.T) {
	_default := map[string]interface{}{
		"enabled":             nil,
		"local_ip":            "",
		"forward_to_kube_dns": nil,
		"memory_request":      nil,
		"cpu_request":         nil,
	}
	type args struct {
		in kops.NodeLocalDNSConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.NodeLocalDNSConfig{},
			},
			want: _default,
		},
		{
			name: "Enabled - default",
			args: args{
				in: func() kops.NodeLocalDNSConfig {
					subject := kops.NodeLocalDNSConfig{}
					subject.Enabled = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LocalIp - default",
			args: args{
				in: func() kops.NodeLocalDNSConfig {
					subject := kops.NodeLocalDNSConfig{}
					subject.LocalIP = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ForwardToKubeDns - default",
			args: args{
				in: func() kops.NodeLocalDNSConfig {
					subject := kops.NodeLocalDNSConfig{}
					subject.ForwardToKubeDNS = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MemoryRequest - default",
			args: args{
				in: func() kops.NodeLocalDNSConfig {
					subject := kops.NodeLocalDNSConfig{}
					subject.MemoryRequest = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CPURequest - default",
			args: args{
				in: func() kops.NodeLocalDNSConfig {
					subject := kops.NodeLocalDNSConfig{}
					subject.CPURequest = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenDataSourceNodeLocalDNSConfigInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceNodeLocalDNSConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceNodeLocalDNSConfig(t *testing.T) {
	_default := map[string]interface{}{
		"enabled":             nil,
		"local_ip":            "",
		"forward_to_kube_dns": nil,
		"memory_request":      nil,
		"cpu_request":         nil,
	}
	type args struct {
		in kops.NodeLocalDNSConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.NodeLocalDNSConfig{},
			},
			want: _default,
		},
		{
			name: "Enabled - default",
			args: args{
				in: func() kops.NodeLocalDNSConfig {
					subject := kops.NodeLocalDNSConfig{}
					subject.Enabled = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LocalIp - default",
			args: args{
				in: func() kops.NodeLocalDNSConfig {
					subject := kops.NodeLocalDNSConfig{}
					subject.LocalIP = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ForwardToKubeDns - default",
			args: args{
				in: func() kops.NodeLocalDNSConfig {
					subject := kops.NodeLocalDNSConfig{}
					subject.ForwardToKubeDNS = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MemoryRequest - default",
			args: args{
				in: func() kops.NodeLocalDNSConfig {
					subject := kops.NodeLocalDNSConfig{}
					subject.MemoryRequest = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CPURequest - default",
			args: args{
				in: func() kops.NodeLocalDNSConfig {
					subject := kops.NodeLocalDNSConfig{}
					subject.CPURequest = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceNodeLocalDNSConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceNodeLocalDNSConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
