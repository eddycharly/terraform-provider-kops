package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceExternalDNSConfig(t *testing.T) {
	_default := kops.ExternalDNSConfig{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.ExternalDNSConfig
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"watch_ingress":   nil,
					"watch_namespace": "",
					"provider":        "",
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandResourceExternalDNSConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceExternalDNSConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceExternalDNSConfigInto(t *testing.T) {
	_default := map[string]interface{}{
		"watch_ingress":   nil,
		"watch_namespace": "",
		"provider":        "",
	}
	type args struct {
		in kops.ExternalDNSConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.ExternalDNSConfig{},
			},
			want: _default,
		},
		{
			name: "WatchIngress - default",
			args: args{
				in: func() kops.ExternalDNSConfig {
					subject := kops.ExternalDNSConfig{}
					subject.WatchIngress = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "WatchNamespace - default",
			args: args{
				in: func() kops.ExternalDNSConfig {
					subject := kops.ExternalDNSConfig{}
					subject.WatchNamespace = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Provider - default",
			args: args{
				in: func() kops.ExternalDNSConfig {
					subject := kops.ExternalDNSConfig{}
					subject.Provider = ""
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenResourceExternalDNSConfigInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceExternalDNSConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceExternalDNSConfig(t *testing.T) {
	_default := map[string]interface{}{
		"watch_ingress":   nil,
		"watch_namespace": "",
		"provider":        "",
	}
	type args struct {
		in kops.ExternalDNSConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.ExternalDNSConfig{},
			},
			want: _default,
		},
		{
			name: "WatchIngress - default",
			args: args{
				in: func() kops.ExternalDNSConfig {
					subject := kops.ExternalDNSConfig{}
					subject.WatchIngress = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "WatchNamespace - default",
			args: args{
				in: func() kops.ExternalDNSConfig {
					subject := kops.ExternalDNSConfig{}
					subject.WatchNamespace = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Provider - default",
			args: args{
				in: func() kops.ExternalDNSConfig {
					subject := kops.ExternalDNSConfig{}
					subject.Provider = ""
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceExternalDNSConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceExternalDNSConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
