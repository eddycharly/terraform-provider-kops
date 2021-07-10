package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceExternalDNSConfig(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.ExternalDNSConfig
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandDataSourceExternalDNSConfig(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandDataSourceExternalDNSConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenDataSourceExternalDNSConfigInto(t *testing.T) {
	_default := map[string]interface{}{
		"disable":         false,
		"watch_ingress":   nil,
		"watch_namespace": "",
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
			name: "Disable - default",
			args: args{
				in: func() kops.ExternalDNSConfig {
					subject := kops.ExternalDNSConfig{}
					subject.Disable = false
					return subject
				}(),
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
			name: "Disable - default",
			args: args{
				in: func() kops.ExternalDNSConfig {
					subject := kops.ExternalDNSConfig{}
					subject.Disable = false
					return subject
				}(),
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
