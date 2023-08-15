package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

func TestExpandResourceLoadBalancerSpec(t *testing.T) {
	_default := kopsv1alpha2.LoadBalancerSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kopsv1alpha2.LoadBalancerSpec
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"load_balancer_name": nil,
					"target_group_arn":   nil,
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandResourceLoadBalancerSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceLoadBalancerSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceLoadBalancerSpecInto(t *testing.T) {
	_default := map[string]interface{}{
		"load_balancer_name": nil,
		"target_group_arn":   nil,
	}
	type args struct {
		in kopsv1alpha2.LoadBalancerSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.LoadBalancerSpec{},
			},
			want: _default,
		},
		{
			name: "LoadBalancerName - default",
			args: args{
				in: func() kopsv1alpha2.LoadBalancerSpec {
					subject := kopsv1alpha2.LoadBalancerSpec{}
					subject.LoadBalancerName = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "TargetGroupARN - default",
			args: args{
				in: func() kopsv1alpha2.LoadBalancerSpec {
					subject := kopsv1alpha2.LoadBalancerSpec{}
					subject.TargetGroupARN = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenResourceLoadBalancerSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceLoadBalancerSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceLoadBalancerSpec(t *testing.T) {
	_default := map[string]interface{}{
		"load_balancer_name": nil,
		"target_group_arn":   nil,
	}
	type args struct {
		in kopsv1alpha2.LoadBalancerSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.LoadBalancerSpec{},
			},
			want: _default,
		},
		{
			name: "LoadBalancerName - default",
			args: args{
				in: func() kopsv1alpha2.LoadBalancerSpec {
					subject := kopsv1alpha2.LoadBalancerSpec{}
					subject.LoadBalancerName = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "TargetGroupARN - default",
			args: args{
				in: func() kopsv1alpha2.LoadBalancerSpec {
					subject := kopsv1alpha2.LoadBalancerSpec{}
					subject.TargetGroupARN = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceLoadBalancerSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceLoadBalancerSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
