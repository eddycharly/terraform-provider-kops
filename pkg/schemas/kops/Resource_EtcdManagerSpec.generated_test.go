package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceEtcdManagerSpec(t *testing.T) {
	_default := kops.EtcdManagerSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.EtcdManagerSpec
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"image":                   "",
					"env":                     func() []interface{} { return nil }(),
					"discovery_poll_interval": nil,
					"log_level":               nil,
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandResourceEtcdManagerSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceEtcdManagerSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceEtcdManagerSpecInto(t *testing.T) {
	_default := map[string]interface{}{
		"image":                   "",
		"env":                     func() []interface{} { return nil }(),
		"discovery_poll_interval": nil,
		"log_level":               nil,
	}
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenResourceEtcdManagerSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceEtcdManagerSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceEtcdManagerSpec(t *testing.T) {
	_default := map[string]interface{}{
		"image":                   "",
		"env":                     func() []interface{} { return nil }(),
		"discovery_poll_interval": nil,
		"log_level":               nil,
	}
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceEtcdManagerSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceEtcdManagerSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
