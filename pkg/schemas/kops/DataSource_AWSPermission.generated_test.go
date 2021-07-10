package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceAWSPermission(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.AWSPermission
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandDataSourceAWSPermission(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandDataSourceAWSPermission() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenDataSourceAWSPermissionInto(t *testing.T) {
	type args struct {
		in  kops.AWSPermission
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
			FlattenDataSourceAWSPermissionInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenDataSourceAWSPermission(t *testing.T) {
	type args struct {
		in kops.AWSPermission
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.AWSPermission{},
			},
			want: map[string]interface{}{
				"policy_ar_ns":  func() []interface{} { return nil }(),
				"inline_policy": "",
			},
		},
		{
			name: "PolicyARNs - default",
			args: args{
				in: func() kops.AWSPermission {
					subject := kops.AWSPermission{}
					subject.PolicyARNs = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"policy_ar_ns":  func() []interface{} { return nil }(),
				"inline_policy": "",
			},
		},
		{
			name: "InlinePolicy - default",
			args: args{
				in: func() kops.AWSPermission {
					subject := kops.AWSPermission{}
					subject.InlinePolicy = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
				"policy_ar_ns":  func() []interface{} { return nil }(),
				"inline_policy": "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenDataSourceAWSPermission(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenDataSourceAWSPermission() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
