package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

func TestExpandResourceOpenstackSpec(t *testing.T) {
	_default := kopsv1alpha2.OpenstackSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kopsv1alpha2.OpenstackSpec
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
			got := ExpandResourceOpenstackSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceOpenstackSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceOpenstackSpecInto(t *testing.T) {
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
		in kopsv1alpha2.OpenstackSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.OpenstackSpec{},
			},
			want: _default,
		},
		{
			name: "Loadbalancer - default",
			args: args{
				in: func() kopsv1alpha2.OpenstackSpec {
					subject := kopsv1alpha2.OpenstackSpec{}
					subject.Loadbalancer = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Monitor - default",
			args: args{
				in: func() kopsv1alpha2.OpenstackSpec {
					subject := kopsv1alpha2.OpenstackSpec{}
					subject.Monitor = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Router - default",
			args: args{
				in: func() kopsv1alpha2.OpenstackSpec {
					subject := kopsv1alpha2.OpenstackSpec{}
					subject.Router = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "BlockStorage - default",
			args: args{
				in: func() kopsv1alpha2.OpenstackSpec {
					subject := kopsv1alpha2.OpenstackSpec{}
					subject.BlockStorage = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "InsecureSkipVerify - default",
			args: args{
				in: func() kopsv1alpha2.OpenstackSpec {
					subject := kopsv1alpha2.OpenstackSpec{}
					subject.InsecureSkipVerify = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Network - default",
			args: args{
				in: func() kopsv1alpha2.OpenstackSpec {
					subject := kopsv1alpha2.OpenstackSpec{}
					subject.Network = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Metadata - default",
			args: args{
				in: func() kopsv1alpha2.OpenstackSpec {
					subject := kopsv1alpha2.OpenstackSpec{}
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
			FlattenResourceOpenstackSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceOpenstackSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceOpenstackSpec(t *testing.T) {
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
		in kopsv1alpha2.OpenstackSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.OpenstackSpec{},
			},
			want: _default,
		},
		{
			name: "Loadbalancer - default",
			args: args{
				in: func() kopsv1alpha2.OpenstackSpec {
					subject := kopsv1alpha2.OpenstackSpec{}
					subject.Loadbalancer = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Monitor - default",
			args: args{
				in: func() kopsv1alpha2.OpenstackSpec {
					subject := kopsv1alpha2.OpenstackSpec{}
					subject.Monitor = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Router - default",
			args: args{
				in: func() kopsv1alpha2.OpenstackSpec {
					subject := kopsv1alpha2.OpenstackSpec{}
					subject.Router = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "BlockStorage - default",
			args: args{
				in: func() kopsv1alpha2.OpenstackSpec {
					subject := kopsv1alpha2.OpenstackSpec{}
					subject.BlockStorage = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "InsecureSkipVerify - default",
			args: args{
				in: func() kopsv1alpha2.OpenstackSpec {
					subject := kopsv1alpha2.OpenstackSpec{}
					subject.InsecureSkipVerify = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Network - default",
			args: args{
				in: func() kopsv1alpha2.OpenstackSpec {
					subject := kopsv1alpha2.OpenstackSpec{}
					subject.Network = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Metadata - default",
			args: args{
				in: func() kopsv1alpha2.OpenstackSpec {
					subject := kopsv1alpha2.OpenstackSpec{}
					subject.Metadata = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceOpenstackSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceOpenstackSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
