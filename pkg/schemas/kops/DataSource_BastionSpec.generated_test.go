package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceBastionSpec(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.BastionSpec
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandDataSourceBastionSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandDataSourceBastionSpec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenDataSourceBastionSpecInto(t *testing.T) {
	type args struct {
		in  kops.BastionSpec
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
			FlattenDataSourceBastionSpecInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenDataSourceBastionSpec(t *testing.T) {
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
			want: map[string]interface{}{
				"bastion_public_name":  "",
				"idle_timeout_seconds": nil,
				"load_balancer":        nil,
			},
		},
		{
			name: "BastionPublicName - default",
			args: args{
				in: func() kops.BastionSpec {
					subject := kops.BastionSpec{}
					subject.BastionPublicName = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
				"bastion_public_name":  "",
				"idle_timeout_seconds": nil,
				"load_balancer":        nil,
			},
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
			want: map[string]interface{}{
				"bastion_public_name":  "",
				"idle_timeout_seconds": nil,
				"load_balancer":        nil,
			},
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
			want: map[string]interface{}{
				"bastion_public_name":  "",
				"idle_timeout_seconds": nil,
				"load_balancer":        nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenDataSourceBastionSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenDataSourceBastionSpec() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
