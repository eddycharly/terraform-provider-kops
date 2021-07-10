package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceOpenstackConfiguration(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.OpenstackConfiguration
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandDataSourceOpenstackConfiguration(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandDataSourceOpenstackConfiguration() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenDataSourceOpenstackConfigurationInto(t *testing.T) {
	type args struct {
		in  kops.OpenstackConfiguration
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
			FlattenDataSourceOpenstackConfigurationInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenDataSourceOpenstackConfiguration(t *testing.T) {
	type args struct {
		in kops.OpenstackConfiguration
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.OpenstackConfiguration{},
			},
			want: map[string]interface{}{
				"loadbalancer":         nil,
				"monitor":              nil,
				"router":               nil,
				"block_storage":        nil,
				"insecure_skip_verify": nil,
				"network":              nil,
			},
		},
		{
			name: "Loadbalancer - default",
			args: args{
				in: func() kops.OpenstackConfiguration {
					subject := kops.OpenstackConfiguration{}
					subject.Loadbalancer = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"loadbalancer":         nil,
				"monitor":              nil,
				"router":               nil,
				"block_storage":        nil,
				"insecure_skip_verify": nil,
				"network":              nil,
			},
		},
		{
			name: "Monitor - default",
			args: args{
				in: func() kops.OpenstackConfiguration {
					subject := kops.OpenstackConfiguration{}
					subject.Monitor = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"loadbalancer":         nil,
				"monitor":              nil,
				"router":               nil,
				"block_storage":        nil,
				"insecure_skip_verify": nil,
				"network":              nil,
			},
		},
		{
			name: "Router - default",
			args: args{
				in: func() kops.OpenstackConfiguration {
					subject := kops.OpenstackConfiguration{}
					subject.Router = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"loadbalancer":         nil,
				"monitor":              nil,
				"router":               nil,
				"block_storage":        nil,
				"insecure_skip_verify": nil,
				"network":              nil,
			},
		},
		{
			name: "BlockStorage - default",
			args: args{
				in: func() kops.OpenstackConfiguration {
					subject := kops.OpenstackConfiguration{}
					subject.BlockStorage = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"loadbalancer":         nil,
				"monitor":              nil,
				"router":               nil,
				"block_storage":        nil,
				"insecure_skip_verify": nil,
				"network":              nil,
			},
		},
		{
			name: "InsecureSkipVerify - default",
			args: args{
				in: func() kops.OpenstackConfiguration {
					subject := kops.OpenstackConfiguration{}
					subject.InsecureSkipVerify = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"loadbalancer":         nil,
				"monitor":              nil,
				"router":               nil,
				"block_storage":        nil,
				"insecure_skip_verify": nil,
				"network":              nil,
			},
		},
		{
			name: "Network - default",
			args: args{
				in: func() kops.OpenstackConfiguration {
					subject := kops.OpenstackConfiguration{}
					subject.Network = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"loadbalancer":         nil,
				"monitor":              nil,
				"router":               nil,
				"block_storage":        nil,
				"insecure_skip_verify": nil,
				"network":              nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenDataSourceOpenstackConfiguration(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenDataSourceOpenstackConfiguration() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
