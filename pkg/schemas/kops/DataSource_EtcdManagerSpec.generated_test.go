package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceEtcdManagerSpec(t *testing.T) {
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
			if got := ExpandDataSourceEtcdManagerSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandDataSourceEtcdManagerSpec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenDataSourceEtcdManagerSpecInto(t *testing.T) {
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
			FlattenDataSourceEtcdManagerSpecInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenDataSourceEtcdManagerSpec(t *testing.T) {
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
			if got := FlattenDataSourceEtcdManagerSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenDataSourceEtcdManagerSpec() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
