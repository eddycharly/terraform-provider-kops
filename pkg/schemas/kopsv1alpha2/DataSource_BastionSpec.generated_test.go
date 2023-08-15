package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

func TestExpandDataSourceBastionSpec(t *testing.T) {
	_default := kopsv1alpha2.BastionSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kopsv1alpha2.BastionSpec
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
			got := ExpandDataSourceBastionSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceBastionSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceBastionSpecInto(t *testing.T) {
	_default := map[string]interface{}{
		"bastion_public_name":  "",
		"idle_timeout_seconds": nil,
		"load_balancer":        nil,
	}
	type args struct {
		in kopsv1alpha2.BastionSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.BastionSpec{},
			},
			want: _default,
		},
		{
			name: "BastionPublicName - default",
			args: args{
				in: func() kopsv1alpha2.BastionSpec {
					subject := kopsv1alpha2.BastionSpec{}
					subject.PublicName = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "IdleTimeoutSeconds - default",
			args: args{
				in: func() kopsv1alpha2.BastionSpec {
					subject := kopsv1alpha2.BastionSpec{}
					subject.IdleTimeoutSeconds = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LoadBalancer - default",
			args: args{
				in: func() kopsv1alpha2.BastionSpec {
					subject := kopsv1alpha2.BastionSpec{}
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
			FlattenDataSourceBastionSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceBastionSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceBastionSpec(t *testing.T) {
	_default := map[string]interface{}{
		"bastion_public_name":  "",
		"idle_timeout_seconds": nil,
		"load_balancer":        nil,
	}
	type args struct {
		in kopsv1alpha2.BastionSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.BastionSpec{},
			},
			want: _default,
		},
		{
			name: "BastionPublicName - default",
			args: args{
				in: func() kopsv1alpha2.BastionSpec {
					subject := kopsv1alpha2.BastionSpec{}
					subject.PublicName = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "IdleTimeoutSeconds - default",
			args: args{
				in: func() kopsv1alpha2.BastionSpec {
					subject := kopsv1alpha2.BastionSpec{}
					subject.IdleTimeoutSeconds = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LoadBalancer - default",
			args: args{
				in: func() kopsv1alpha2.BastionSpec {
					subject := kopsv1alpha2.BastionSpec{}
					subject.LoadBalancer = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceBastionSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceBastionSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
