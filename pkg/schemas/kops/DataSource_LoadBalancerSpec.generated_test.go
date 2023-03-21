package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceLoadBalancerSpec(t *testing.T) {
	_default := kops.LoadBalancerSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.LoadBalancerSpec
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
			got := ExpandDataSourceLoadBalancerSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceLoadBalancerSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceLoadBalancerSpecInto(t *testing.T) {
	_default := map[string]interface{}{
		"load_balancer_name": nil,
		"target_group_arn":   nil,
	}
	type args struct {
		in kops.LoadBalancerSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.LoadBalancerSpec{},
			},
			want: _default,
		},
		{
			name: "LoadBalancerName - default",
			args: args{
				in: func() kops.LoadBalancerSpec {
					subject := kops.LoadBalancerSpec{}
					subject.LoadBalancerName = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "TargetGroupARN - default",
			args: args{
				in: func() kops.LoadBalancerSpec {
					subject := kops.LoadBalancerSpec{}
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
			FlattenDataSourceLoadBalancerSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceLoadBalancerSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceLoadBalancerSpec(t *testing.T) {
	_default := map[string]interface{}{
		"load_balancer_name": nil,
		"target_group_arn":   nil,
	}
	type args struct {
		in kops.LoadBalancerSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.LoadBalancerSpec{},
			},
			want: _default,
		},
		{
			name: "LoadBalancerName - default",
			args: args{
				in: func() kops.LoadBalancerSpec {
					subject := kops.LoadBalancerSpec{}
					subject.LoadBalancerName = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "TargetGroupARN - default",
			args: args{
				in: func() kops.LoadBalancerSpec {
					subject := kops.LoadBalancerSpec{}
					subject.TargetGroupARN = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceLoadBalancerSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceLoadBalancerSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
