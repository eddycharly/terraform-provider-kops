package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceOpenstackConfiguration(t *testing.T) {
	_default := kops.OpenstackConfiguration{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.OpenstackConfiguration
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"loadbalancer":         nil,
					"monitor":              nil,
					"router":               nil,
					"block_storage":        nil,
					"insecure_skip_verify": nil,
					"network":              nil,
					"metadata":             nil,
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandResourceOpenstackConfiguration(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceOpenstackConfiguration() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceOpenstackConfigurationInto(t *testing.T) {
	_default := map[string]interface{}{
		"loadbalancer":         nil,
		"monitor":              nil,
		"router":               nil,
		"block_storage":        nil,
		"insecure_skip_verify": nil,
		"network":              nil,
		"metadata":             nil,
	}
	type args struct {
		in kops.OpenstackConfiguration
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.OpenstackConfiguration{},
			},
			want: _default,
		},
		{
			name: "Loadbalancer - default",
			args: args{
				in: func() kops.OpenstackConfiguration {
					subject := kops.OpenstackConfiguration{}
					subject.Loadbalancer = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Monitor - default",
			args: args{
				in: func() kops.OpenstackConfiguration {
					subject := kops.OpenstackConfiguration{}
					subject.Monitor = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Router - default",
			args: args{
				in: func() kops.OpenstackConfiguration {
					subject := kops.OpenstackConfiguration{}
					subject.Router = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "BlockStorage - default",
			args: args{
				in: func() kops.OpenstackConfiguration {
					subject := kops.OpenstackConfiguration{}
					subject.BlockStorage = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "InsecureSkipVerify - default",
			args: args{
				in: func() kops.OpenstackConfiguration {
					subject := kops.OpenstackConfiguration{}
					subject.InsecureSkipVerify = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Network - default",
			args: args{
				in: func() kops.OpenstackConfiguration {
					subject := kops.OpenstackConfiguration{}
					subject.Network = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Metadata - default",
			args: args{
				in: func() kops.OpenstackConfiguration {
					subject := kops.OpenstackConfiguration{}
					subject.Metadata = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenResourceOpenstackConfigurationInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceOpenstackConfiguration() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceOpenstackConfiguration(t *testing.T) {
	_default := map[string]interface{}{
		"loadbalancer":         nil,
		"monitor":              nil,
		"router":               nil,
		"block_storage":        nil,
		"insecure_skip_verify": nil,
		"network":              nil,
		"metadata":             nil,
	}
	type args struct {
		in kops.OpenstackConfiguration
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.OpenstackConfiguration{},
			},
			want: _default,
		},
		{
			name: "Loadbalancer - default",
			args: args{
				in: func() kops.OpenstackConfiguration {
					subject := kops.OpenstackConfiguration{}
					subject.Loadbalancer = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Monitor - default",
			args: args{
				in: func() kops.OpenstackConfiguration {
					subject := kops.OpenstackConfiguration{}
					subject.Monitor = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Router - default",
			args: args{
				in: func() kops.OpenstackConfiguration {
					subject := kops.OpenstackConfiguration{}
					subject.Router = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "BlockStorage - default",
			args: args{
				in: func() kops.OpenstackConfiguration {
					subject := kops.OpenstackConfiguration{}
					subject.BlockStorage = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "InsecureSkipVerify - default",
			args: args{
				in: func() kops.OpenstackConfiguration {
					subject := kops.OpenstackConfiguration{}
					subject.InsecureSkipVerify = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Network - default",
			args: args{
				in: func() kops.OpenstackConfiguration {
					subject := kops.OpenstackConfiguration{}
					subject.Network = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Metadata - default",
			args: args{
				in: func() kops.OpenstackConfiguration {
					subject := kops.OpenstackConfiguration{}
					subject.Metadata = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceOpenstackConfiguration(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceOpenstackConfiguration() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
