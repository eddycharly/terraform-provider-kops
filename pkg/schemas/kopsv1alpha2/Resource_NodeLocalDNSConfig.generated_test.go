package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

func TestExpandResourceNodeLocalDNSConfig(t *testing.T) {
	_default := kopsv1alpha2.NodeLocalDNSConfig{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kopsv1alpha2.NodeLocalDNSConfig
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
		in kopsv1alpha2.NodeLocalDNSConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.NodeLocalDNSConfig{},
			},
			want: _default,
		},
		{
			name: "Enabled - default",
			args: args{
				in: func() kopsv1alpha2.NodeLocalDNSConfig {
					subject := kopsv1alpha2.NodeLocalDNSConfig{}
					subject.Enabled = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ExternalCoreFile - default",
			args: args{
				in: func() kopsv1alpha2.NodeLocalDNSConfig {
					subject := kopsv1alpha2.NodeLocalDNSConfig{}
					subject.ExternalCoreFile = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Image - default",
			args: args{
				in: func() kopsv1alpha2.NodeLocalDNSConfig {
					subject := kopsv1alpha2.NodeLocalDNSConfig{}
					subject.Image = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LocalIp - default",
			args: args{
				in: func() kopsv1alpha2.NodeLocalDNSConfig {
					subject := kopsv1alpha2.NodeLocalDNSConfig{}
					subject.LocalIP = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ForwardToKubeDns - default",
			args: args{
				in: func() kopsv1alpha2.NodeLocalDNSConfig {
					subject := kopsv1alpha2.NodeLocalDNSConfig{}
					subject.ForwardToKubeDNS = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MemoryRequest - default",
			args: args{
				in: func() kopsv1alpha2.NodeLocalDNSConfig {
					subject := kopsv1alpha2.NodeLocalDNSConfig{}
					subject.MemoryRequest = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CpuRequest - default",
			args: args{
				in: func() kopsv1alpha2.NodeLocalDNSConfig {
					subject := kopsv1alpha2.NodeLocalDNSConfig{}
					subject.CPURequest = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PodAnnotations - default",
			args: args{
				in: func() kopsv1alpha2.NodeLocalDNSConfig {
					subject := kopsv1alpha2.NodeLocalDNSConfig{}
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
		in kopsv1alpha2.NodeLocalDNSConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.NodeLocalDNSConfig{},
			},
			want: _default,
		},
		{
			name: "Enabled - default",
			args: args{
				in: func() kopsv1alpha2.NodeLocalDNSConfig {
					subject := kopsv1alpha2.NodeLocalDNSConfig{}
					subject.Enabled = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ExternalCoreFile - default",
			args: args{
				in: func() kopsv1alpha2.NodeLocalDNSConfig {
					subject := kopsv1alpha2.NodeLocalDNSConfig{}
					subject.ExternalCoreFile = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Image - default",
			args: args{
				in: func() kopsv1alpha2.NodeLocalDNSConfig {
					subject := kopsv1alpha2.NodeLocalDNSConfig{}
					subject.Image = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LocalIp - default",
			args: args{
				in: func() kopsv1alpha2.NodeLocalDNSConfig {
					subject := kopsv1alpha2.NodeLocalDNSConfig{}
					subject.LocalIP = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ForwardToKubeDns - default",
			args: args{
				in: func() kopsv1alpha2.NodeLocalDNSConfig {
					subject := kopsv1alpha2.NodeLocalDNSConfig{}
					subject.ForwardToKubeDNS = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MemoryRequest - default",
			args: args{
				in: func() kopsv1alpha2.NodeLocalDNSConfig {
					subject := kopsv1alpha2.NodeLocalDNSConfig{}
					subject.MemoryRequest = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CpuRequest - default",
			args: args{
				in: func() kopsv1alpha2.NodeLocalDNSConfig {
					subject := kopsv1alpha2.NodeLocalDNSConfig{}
					subject.CPURequest = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PodAnnotations - default",
			args: args{
				in: func() kopsv1alpha2.NodeLocalDNSConfig {
					subject := kopsv1alpha2.NodeLocalDNSConfig{}
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
