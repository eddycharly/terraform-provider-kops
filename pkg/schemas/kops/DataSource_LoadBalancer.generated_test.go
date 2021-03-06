package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceLoadBalancer(t *testing.T) {
	_default := kops.LoadBalancer{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.LoadBalancer
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
			got := ExpandDataSourceLoadBalancer(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceLoadBalancer() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceLoadBalancerInto(t *testing.T) {
	_default := map[string]interface{}{
		"load_balancer_name": nil,
		"target_group_arn":   nil,
	}
	type args struct {
		in kops.LoadBalancer
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.LoadBalancer{},
			},
			want: _default,
		},
		{
			name: "LoadBalancerName - default",
			args: args{
				in: func() kops.LoadBalancer {
					subject := kops.LoadBalancer{}
					subject.LoadBalancerName = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "TargetGroupARN - default",
			args: args{
				in: func() kops.LoadBalancer {
					subject := kops.LoadBalancer{}
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
			FlattenDataSourceLoadBalancerInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceLoadBalancer() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceLoadBalancer(t *testing.T) {
	_default := map[string]interface{}{
		"load_balancer_name": nil,
		"target_group_arn":   nil,
	}
	type args struct {
		in kops.LoadBalancer
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.LoadBalancer{},
			},
			want: _default,
		},
		{
			name: "LoadBalancerName - default",
			args: args{
				in: func() kops.LoadBalancer {
					subject := kops.LoadBalancer{}
					subject.LoadBalancerName = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "TargetGroupARN - default",
			args: args{
				in: func() kops.LoadBalancer {
					subject := kops.LoadBalancer{}
					subject.TargetGroupARN = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceLoadBalancer(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceLoadBalancer() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
