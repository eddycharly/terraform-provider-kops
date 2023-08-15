package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

func TestExpandDataSourceAWSAuthenticationIdentityMappingSpec(t *testing.T) {
	_default := kopsv1alpha2.AWSAuthenticationIdentityMappingSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kopsv1alpha2.AWSAuthenticationIdentityMappingSpec
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"arn":      "",
					"username": "",
					"groups":   func() []interface{} { return nil }(),
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandDataSourceAWSAuthenticationIdentityMappingSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceAWSAuthenticationIdentityMappingSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceAWSAuthenticationIdentityMappingSpecInto(t *testing.T) {
	_default := map[string]interface{}{
		"arn":      "",
		"username": "",
		"groups":   func() []interface{} { return nil }(),
	}
	type args struct {
		in kopsv1alpha2.AWSAuthenticationIdentityMappingSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.AWSAuthenticationIdentityMappingSpec{},
			},
			want: _default,
		},
		{
			name: "ARN - default",
			args: args{
				in: func() kopsv1alpha2.AWSAuthenticationIdentityMappingSpec {
					subject := kopsv1alpha2.AWSAuthenticationIdentityMappingSpec{}
					subject.ARN = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Username - default",
			args: args{
				in: func() kopsv1alpha2.AWSAuthenticationIdentityMappingSpec {
					subject := kopsv1alpha2.AWSAuthenticationIdentityMappingSpec{}
					subject.Username = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Groups - default",
			args: args{
				in: func() kopsv1alpha2.AWSAuthenticationIdentityMappingSpec {
					subject := kopsv1alpha2.AWSAuthenticationIdentityMappingSpec{}
					subject.Groups = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenDataSourceAWSAuthenticationIdentityMappingSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceAWSAuthenticationIdentityMappingSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceAWSAuthenticationIdentityMappingSpec(t *testing.T) {
	_default := map[string]interface{}{
		"arn":      "",
		"username": "",
		"groups":   func() []interface{} { return nil }(),
	}
	type args struct {
		in kopsv1alpha2.AWSAuthenticationIdentityMappingSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.AWSAuthenticationIdentityMappingSpec{},
			},
			want: _default,
		},
		{
			name: "ARN - default",
			args: args{
				in: func() kopsv1alpha2.AWSAuthenticationIdentityMappingSpec {
					subject := kopsv1alpha2.AWSAuthenticationIdentityMappingSpec{}
					subject.ARN = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Username - default",
			args: args{
				in: func() kopsv1alpha2.AWSAuthenticationIdentityMappingSpec {
					subject := kopsv1alpha2.AWSAuthenticationIdentityMappingSpec{}
					subject.Username = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Groups - default",
			args: args{
				in: func() kopsv1alpha2.AWSAuthenticationIdentityMappingSpec {
					subject := kopsv1alpha2.AWSAuthenticationIdentityMappingSpec{}
					subject.Groups = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceAWSAuthenticationIdentityMappingSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceAWSAuthenticationIdentityMappingSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
