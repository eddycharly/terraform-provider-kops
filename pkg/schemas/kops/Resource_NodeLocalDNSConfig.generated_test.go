package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceNodeLocalDNSConfig(t *testing.T) {
	_default := kops.NodeLocalDNSConfig{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.NodeLocalDNSConfig
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"enabled":             nil,
					"external_core_file":  "",
					"image":               nil,
					"local_ip":            "",
					"forward_to_kube_dns": nil,
					"memory_request":      nil,
					"cpu_request":         nil,
					"pod_annotations":     func() map[string]interface{} { return nil }(),
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandResourceNodeLocalDNSConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceNodeLocalDNSConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceNodeLocalDNSConfigInto(t *testing.T) {
	_default := map[string]interface{}{
		"enabled":             nil,
		"external_core_file":  "",
		"image":               nil,
		"local_ip":            "",
		"forward_to_kube_dns": nil,
		"memory_request":      nil,
		"cpu_request":         nil,
		"pod_annotations":     func() map[string]interface{} { return nil }(),
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
			name: "ExternalCoreFile - default",
			args: args{
				in: func() kops.NodeLocalDNSConfig {
					subject := kops.NodeLocalDNSConfig{}
					subject.ExternalCoreFile = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Image - default",
			args: args{
				in: func() kops.NodeLocalDNSConfig {
					subject := kops.NodeLocalDNSConfig{}
					subject.Image = nil
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
			name: "CpuRequest - default",
			args: args{
				in: func() kops.NodeLocalDNSConfig {
					subject := kops.NodeLocalDNSConfig{}
					subject.CPURequest = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PodAnnotations - default",
			args: args{
				in: func() kops.NodeLocalDNSConfig {
					subject := kops.NodeLocalDNSConfig{}
					subject.PodAnnotations = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenResourceNodeLocalDNSConfigInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceNodeLocalDNSConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceNodeLocalDNSConfig(t *testing.T) {
	_default := map[string]interface{}{
		"enabled":             nil,
		"external_core_file":  "",
		"image":               nil,
		"local_ip":            "",
		"forward_to_kube_dns": nil,
		"memory_request":      nil,
		"cpu_request":         nil,
		"pod_annotations":     func() map[string]interface{} { return nil }(),
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
			name: "ExternalCoreFile - default",
			args: args{
				in: func() kops.NodeLocalDNSConfig {
					subject := kops.NodeLocalDNSConfig{}
					subject.ExternalCoreFile = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Image - default",
			args: args{
				in: func() kops.NodeLocalDNSConfig {
					subject := kops.NodeLocalDNSConfig{}
					subject.Image = nil
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
			name: "CpuRequest - default",
			args: args{
				in: func() kops.NodeLocalDNSConfig {
					subject := kops.NodeLocalDNSConfig{}
					subject.CPURequest = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PodAnnotations - default",
			args: args{
				in: func() kops.NodeLocalDNSConfig {
					subject := kops.NodeLocalDNSConfig{}
					subject.PodAnnotations = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceNodeLocalDNSConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceNodeLocalDNSConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
