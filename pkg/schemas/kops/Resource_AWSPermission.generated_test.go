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
	_default := map[string]interface{}{
		"policy_ar_ns":  func() []interface{} { return nil }(),
		"inline_policy": "",
	}
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
			want: _default,
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
			want: _default,
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
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenResourceAWSPermissionInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceAWSPermission() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceAWSPermission(t *testing.T) {
	_default := map[string]interface{}{
		"policy_ar_ns":  func() []interface{} { return nil }(),
		"inline_policy": "",
	}
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
			want: _default,
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
			want: _default,
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
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceAWSPermission(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceAWSPermission() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
