package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceExternalDNSConfig(t *testing.T) {
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
			if got := ExpandResourceExternalDNSConfig(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandResourceExternalDNSConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenResourceExternalDNSConfigInto(t *testing.T) {
	type args struct {
		in  kops.ExternalDNSConfig
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
			FlattenResourceExternalDNSConfigInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenResourceExternalDNSConfig(t *testing.T) {
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
			want: map[string]interface{}{
				"disable":         false,
				"watch_ingress":   nil,
				"watch_namespace": "",
			},
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
			want: map[string]interface{}{
				"disable":         false,
				"watch_ingress":   nil,
				"watch_namespace": "",
			},
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
			want: map[string]interface{}{
				"disable":         false,
				"watch_ingress":   nil,
				"watch_namespace": "",
			},
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
			want: map[string]interface{}{
				"disable":         false,
				"watch_ingress":   nil,
				"watch_namespace": "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenResourceExternalDNSConfig(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenResourceExternalDNSConfig() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
