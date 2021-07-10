package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceLoadBalancer(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.LoadBalancer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandDataSourceLoadBalancer(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandDataSourceLoadBalancer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenDataSourceLoadBalancerInto(t *testing.T) {
	type args struct {
		in  kops.LoadBalancer
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
			FlattenDataSourceLoadBalancerInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenDataSourceLoadBalancer(t *testing.T) {
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
			want: map[string]interface{}{
				"load_balancer_name": nil,
				"target_group_arn":   nil,
			},
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
			want: map[string]interface{}{
				"load_balancer_name": nil,
				"target_group_arn":   nil,
			},
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
			want: map[string]interface{}{
				"load_balancer_name": nil,
				"target_group_arn":   nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenDataSourceLoadBalancer(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenDataSourceLoadBalancer() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
