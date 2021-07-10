package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceEtcdManagerSpec(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.EtcdManagerSpec
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandResourceEtcdManagerSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandResourceEtcdManagerSpec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenResourceEtcdManagerSpecInto(t *testing.T) {
	type args struct {
		in  kops.EtcdManagerSpec
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
			FlattenResourceEtcdManagerSpecInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenResourceEtcdManagerSpec(t *testing.T) {
	type args struct {
		in kops.EtcdManagerSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.EtcdManagerSpec{},
			},
			want: map[string]interface{}{
				"image":                   "",
				"env":                     func() []interface{} { return nil }(),
				"discovery_poll_interval": nil,
				"log_level":               nil,
			},
		},
		{
			name: "Image - default",
			args: args{
				in: func() kops.EtcdManagerSpec {
					subject := kops.EtcdManagerSpec{}
					subject.Image = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
				"image":                   "",
				"env":                     func() []interface{} { return nil }(),
				"discovery_poll_interval": nil,
				"log_level":               nil,
			},
		},
		{
			name: "Env - default",
			args: args{
				in: func() kops.EtcdManagerSpec {
					subject := kops.EtcdManagerSpec{}
					subject.Env = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"image":                   "",
				"env":                     func() []interface{} { return nil }(),
				"discovery_poll_interval": nil,
				"log_level":               nil,
			},
		},
		{
			name: "DiscoveryPollInterval - default",
			args: args{
				in: func() kops.EtcdManagerSpec {
					subject := kops.EtcdManagerSpec{}
					subject.DiscoveryPollInterval = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"image":                   "",
				"env":                     func() []interface{} { return nil }(),
				"discovery_poll_interval": nil,
				"log_level":               nil,
			},
		},
		{
			name: "LogLevel - default",
			args: args{
				in: func() kops.EtcdManagerSpec {
					subject := kops.EtcdManagerSpec{}
					subject.LogLevel = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"image":                   "",
				"env":                     func() []interface{} { return nil }(),
				"discovery_poll_interval": nil,
				"log_level":               nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenResourceEtcdManagerSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenResourceEtcdManagerSpec() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
