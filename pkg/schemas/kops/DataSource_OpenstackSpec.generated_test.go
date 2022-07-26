package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceOpenstackSpec(t *testing.T) {
	_default := kops.OpenstackSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.OpenstackSpec
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
			got := ExpandDataSourceOpenstackSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceOpenstackSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceOpenstackSpecInto(t *testing.T) {
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
		in kops.OpenstackSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.OpenstackSpec{},
			},
			want: _default,
		},
		{
			name: "Loadbalancer - default",
			args: args{
				in: func() kops.OpenstackSpec {
					subject := kops.OpenstackSpec{}
					subject.Loadbalancer = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Monitor - default",
			args: args{
				in: func() kops.OpenstackSpec {
					subject := kops.OpenstackSpec{}
					subject.Monitor = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Router - default",
			args: args{
				in: func() kops.OpenstackSpec {
					subject := kops.OpenstackSpec{}
					subject.Router = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "BlockStorage - default",
			args: args{
				in: func() kops.OpenstackSpec {
					subject := kops.OpenstackSpec{}
					subject.BlockStorage = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "InsecureSkipVerify - default",
			args: args{
				in: func() kops.OpenstackSpec {
					subject := kops.OpenstackSpec{}
					subject.InsecureSkipVerify = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Network - default",
			args: args{
				in: func() kops.OpenstackSpec {
					subject := kops.OpenstackSpec{}
					subject.Network = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Metadata - default",
			args: args{
				in: func() kops.OpenstackSpec {
					subject := kops.OpenstackSpec{}
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
			FlattenDataSourceOpenstackSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceOpenstackSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceOpenstackSpec(t *testing.T) {
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
		in kops.OpenstackSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.OpenstackSpec{},
			},
			want: _default,
		},
		{
			name: "Loadbalancer - default",
			args: args{
				in: func() kops.OpenstackSpec {
					subject := kops.OpenstackSpec{}
					subject.Loadbalancer = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Monitor - default",
			args: args{
				in: func() kops.OpenstackSpec {
					subject := kops.OpenstackSpec{}
					subject.Monitor = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Router - default",
			args: args{
				in: func() kops.OpenstackSpec {
					subject := kops.OpenstackSpec{}
					subject.Router = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "BlockStorage - default",
			args: args{
				in: func() kops.OpenstackSpec {
					subject := kops.OpenstackSpec{}
					subject.BlockStorage = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "InsecureSkipVerify - default",
			args: args{
				in: func() kops.OpenstackSpec {
					subject := kops.OpenstackSpec{}
					subject.InsecureSkipVerify = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Network - default",
			args: args{
				in: func() kops.OpenstackSpec {
					subject := kops.OpenstackSpec{}
					subject.Network = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Metadata - default",
			args: args{
				in: func() kops.OpenstackSpec {
					subject := kops.OpenstackSpec{}
					subject.Metadata = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceOpenstackSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceOpenstackSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
