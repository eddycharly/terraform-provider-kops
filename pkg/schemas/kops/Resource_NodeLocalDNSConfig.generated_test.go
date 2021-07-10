package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceNodeLocalDNSConfig(t *testing.T) {
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
			if got := ExpandResourceNodeLocalDNSConfig(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandResourceNodeLocalDNSConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenResourceNodeLocalDNSConfigInto(t *testing.T) {
	type args struct {
		in  kops.NodeLocalDNSConfig
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
			FlattenResourceNodeLocalDNSConfigInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenResourceNodeLocalDNSConfig(t *testing.T) {
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
			want: map[string]interface{}{
				"enabled":             nil,
				"local_ip":            "",
				"forward_to_kube_dns": nil,
				"memory_request":      nil,
				"cpu_request":         nil,
			},
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
			want: map[string]interface{}{
				"enabled":             nil,
				"local_ip":            "",
				"forward_to_kube_dns": nil,
				"memory_request":      nil,
				"cpu_request":         nil,
			},
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
			want: map[string]interface{}{
				"enabled":             nil,
				"local_ip":            "",
				"forward_to_kube_dns": nil,
				"memory_request":      nil,
				"cpu_request":         nil,
			},
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
			want: map[string]interface{}{
				"enabled":             nil,
				"local_ip":            "",
				"forward_to_kube_dns": nil,
				"memory_request":      nil,
				"cpu_request":         nil,
			},
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
			want: map[string]interface{}{
				"enabled":             nil,
				"local_ip":            "",
				"forward_to_kube_dns": nil,
				"memory_request":      nil,
				"cpu_request":         nil,
			},
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
			want: map[string]interface{}{
				"enabled":             nil,
				"local_ip":            "",
				"forward_to_kube_dns": nil,
				"memory_request":      nil,
				"cpu_request":         nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenResourceNodeLocalDNSConfig(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenResourceNodeLocalDNSConfig() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
