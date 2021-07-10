package schemas

import (
	"reflect"
	"testing"

	"github.com/eddycharly/terraform-provider-kops/pkg/api/config"
	"github.com/google/go-cmp/cmp"
)

func TestExpandConfigAwsAssumeRole(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want config.AwsAssumeRole
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandConfigAwsAssumeRole(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandConfigAwsAssumeRole() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenConfigAwsAssumeRoleInto(t *testing.T) {
	type args struct {
		in  config.AwsAssumeRole
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
			FlattenConfigAwsAssumeRoleInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenConfigAwsAssumeRole(t *testing.T) {
	type args struct {
		in config.AwsAssumeRole
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: config.AwsAssumeRole{},
			},
			want: map[string]interface{}{
				"role_arn": "",
			},
		},
		{
			name: "RoleArn - default",
			args: args{
				in: func() config.AwsAssumeRole {
					subject := config.AwsAssumeRole{}
					subject.RoleArn = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
				"role_arn": "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenConfigAwsAssumeRole(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenConfigAwsAssumeRole() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
