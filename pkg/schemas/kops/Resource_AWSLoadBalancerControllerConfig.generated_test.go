package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceAWSLoadBalancerControllerConfig(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.AWSLoadBalancerControllerConfig
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandResourceAWSLoadBalancerControllerConfig(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandResourceAWSLoadBalancerControllerConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenResourceAWSLoadBalancerControllerConfigInto(t *testing.T) {
	type args struct {
		in  kops.AWSLoadBalancerControllerConfig
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
			FlattenResourceAWSLoadBalancerControllerConfigInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenResourceAWSLoadBalancerControllerConfig(t *testing.T) {
	type args struct {
		in kops.AWSLoadBalancerControllerConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.AWSLoadBalancerControllerConfig{},
			},
			want: map[string]interface{}{
				"enabled": nil,
				"version": nil,
			},
		},
		{
			name: "Enabled - default",
			args: args{
				in: func() kops.AWSLoadBalancerControllerConfig {
					subject := kops.AWSLoadBalancerControllerConfig{}
					subject.Enabled = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"enabled": nil,
				"version": nil,
			},
		},
		{
			name: "Version - default",
			args: args{
				in: func() kops.AWSLoadBalancerControllerConfig {
					subject := kops.AWSLoadBalancerControllerConfig{}
					subject.Version = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"enabled": nil,
				"version": nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenResourceAWSLoadBalancerControllerConfig(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenResourceAWSLoadBalancerControllerConfig() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
