package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

func TestExpandDataSourceExternalDNSConfig(t *testing.T) {
	_default := kopsv1alpha2.ExternalDNSConfig{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kopsv1alpha2.ExternalDNSConfig
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"disable":         false,
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
			got := ExpandDataSourceExternalDNSConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceExternalDNSConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceExternalDNSConfigInto(t *testing.T) {
	_default := map[string]interface{}{
		"disable":         false,
		"watch_ingress":   nil,
		"watch_namespace": "",
		"provider":        "",
	}
	type args struct {
		in kopsv1alpha2.ExternalDNSConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.ExternalDNSConfig{},
			},
			want: _default,
		},
		{
			name: "Disable - default",
			args: args{
				in: func() kopsv1alpha2.ExternalDNSConfig {
					subject := kopsv1alpha2.ExternalDNSConfig{}
					subject.Disable = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "WatchIngress - default",
			args: args{
				in: func() kopsv1alpha2.ExternalDNSConfig {
					subject := kopsv1alpha2.ExternalDNSConfig{}
					subject.WatchIngress = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "WatchNamespace - default",
			args: args{
				in: func() kopsv1alpha2.ExternalDNSConfig {
					subject := kopsv1alpha2.ExternalDNSConfig{}
					subject.WatchNamespace = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Provider - default",
			args: args{
				in: func() kopsv1alpha2.ExternalDNSConfig {
					subject := kopsv1alpha2.ExternalDNSConfig{}
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
			FlattenDataSourceExternalDNSConfigInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceExternalDNSConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceExternalDNSConfig(t *testing.T) {
	_default := map[string]interface{}{
		"disable":         false,
		"watch_ingress":   nil,
		"watch_namespace": "",
		"provider":        "",
	}
	type args struct {
		in kopsv1alpha2.ExternalDNSConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.ExternalDNSConfig{},
			},
			want: _default,
		},
		{
			name: "Disable - default",
			args: args{
				in: func() kopsv1alpha2.ExternalDNSConfig {
					subject := kopsv1alpha2.ExternalDNSConfig{}
					subject.Disable = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "WatchIngress - default",
			args: args{
				in: func() kopsv1alpha2.ExternalDNSConfig {
					subject := kopsv1alpha2.ExternalDNSConfig{}
					subject.WatchIngress = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "WatchNamespace - default",
			args: args{
				in: func() kopsv1alpha2.ExternalDNSConfig {
					subject := kopsv1alpha2.ExternalDNSConfig{}
					subject.WatchNamespace = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Provider - default",
			args: args{
				in: func() kopsv1alpha2.ExternalDNSConfig {
					subject := kopsv1alpha2.ExternalDNSConfig{}
					subject.Provider = ""
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceExternalDNSConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceExternalDNSConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
