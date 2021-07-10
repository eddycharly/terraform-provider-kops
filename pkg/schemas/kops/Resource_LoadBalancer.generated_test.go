package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceLoadBalancer(t *testing.T) {
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
			if got := ExpandResourceLoadBalancer(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandResourceLoadBalancer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenResourceLoadBalancerInto(t *testing.T) {
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
			FlattenResourceLoadBalancerInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenResourceLoadBalancer(t *testing.T) {
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
			if got := FlattenResourceLoadBalancer(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenResourceLoadBalancer() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
