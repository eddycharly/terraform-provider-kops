package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceAWSPermission(t *testing.T) {
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
			if got := ExpandResourceAWSPermission(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandResourceAWSPermission() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenResourceAWSPermissionInto(t *testing.T) {
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
			FlattenResourceAWSPermissionInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenResourceAWSPermission(t *testing.T) {
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
			if got := FlattenResourceAWSPermission(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenResourceAWSPermission() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
