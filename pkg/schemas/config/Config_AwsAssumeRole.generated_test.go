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
	_default := map[string]interface{}{
		"role_arn": "",
	}
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
			want: _default,
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
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenConfigAwsAssumeRoleInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenConfigAwsAssumeRole() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenConfigAwsAssumeRole(t *testing.T) {
	_default := map[string]interface{}{
		"role_arn": "",
	}
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
			want: _default,
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
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenConfigAwsAssumeRole(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenConfigAwsAssumeRole() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
