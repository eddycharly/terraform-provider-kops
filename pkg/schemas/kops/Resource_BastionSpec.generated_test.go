package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceBastionSpec(t *testing.T) {
	_default := kops.BastionSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.BastionSpec
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"bastion_public_name":  "",
					"idle_timeout_seconds": nil,
					"load_balancer":        nil,
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandResourceBastionSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceBastionSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceBastionSpecInto(t *testing.T) {
	_default := map[string]interface{}{
		"bastion_public_name":  "",
		"idle_timeout_seconds": nil,
		"load_balancer":        nil,
	}
	type args struct {
		in kops.BastionSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.BastionSpec{},
			},
			want: _default,
		},
		{
			name: "BastionPublicName - default",
			args: args{
				in: func() kops.BastionSpec {
					subject := kops.BastionSpec{}
					subject.PublicName = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "IdleTimeoutSeconds - default",
			args: args{
				in: func() kops.BastionSpec {
					subject := kops.BastionSpec{}
					subject.IdleTimeoutSeconds = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LoadBalancer - default",
			args: args{
				in: func() kops.BastionSpec {
					subject := kops.BastionSpec{}
					subject.LoadBalancer = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenResourceBastionSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceBastionSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceBastionSpec(t *testing.T) {
	_default := map[string]interface{}{
		"bastion_public_name":  "",
		"idle_timeout_seconds": nil,
		"load_balancer":        nil,
	}
	type args struct {
		in kops.BastionSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.BastionSpec{},
			},
			want: _default,
		},
		{
			name: "BastionPublicName - default",
			args: args{
				in: func() kops.BastionSpec {
					subject := kops.BastionSpec{}
					subject.PublicName = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "IdleTimeoutSeconds - default",
			args: args{
				in: func() kops.BastionSpec {
					subject := kops.BastionSpec{}
					subject.IdleTimeoutSeconds = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LoadBalancer - default",
			args: args{
				in: func() kops.BastionSpec {
					subject := kops.BastionSpec{}
					subject.LoadBalancer = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceBastionSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceBastionSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
