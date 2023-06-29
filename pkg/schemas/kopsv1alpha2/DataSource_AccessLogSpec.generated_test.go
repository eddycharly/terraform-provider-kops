package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

func TestExpandDataSourceAccessLogSpec(t *testing.T) {
	_default := kopsv1alpha2.AccessLogSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kopsv1alpha2.AccessLogSpec
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"interval":      0,
					"bucket":        nil,
					"bucket_prefix": nil,
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandDataSourceAccessLogSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceAccessLogSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceAccessLogSpecInto(t *testing.T) {
	_default := map[string]interface{}{
		"interval":      0,
		"bucket":        nil,
		"bucket_prefix": nil,
	}
	type args struct {
		in kopsv1alpha2.AccessLogSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.AccessLogSpec{},
			},
			want: _default,
		},
		{
			name: "Interval - default",
			args: args{
				in: func() kopsv1alpha2.AccessLogSpec {
					subject := kopsv1alpha2.AccessLogSpec{}
					subject.Interval = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Bucket - default",
			args: args{
				in: func() kopsv1alpha2.AccessLogSpec {
					subject := kopsv1alpha2.AccessLogSpec{}
					subject.Bucket = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "BucketPrefix - default",
			args: args{
				in: func() kopsv1alpha2.AccessLogSpec {
					subject := kopsv1alpha2.AccessLogSpec{}
					subject.BucketPrefix = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenDataSourceAccessLogSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceAccessLogSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceAccessLogSpec(t *testing.T) {
	_default := map[string]interface{}{
		"interval":      0,
		"bucket":        nil,
		"bucket_prefix": nil,
	}
	type args struct {
		in kopsv1alpha2.AccessLogSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.AccessLogSpec{},
			},
			want: _default,
		},
		{
			name: "Interval - default",
			args: args{
				in: func() kopsv1alpha2.AccessLogSpec {
					subject := kopsv1alpha2.AccessLogSpec{}
					subject.Interval = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Bucket - default",
			args: args{
				in: func() kopsv1alpha2.AccessLogSpec {
					subject := kopsv1alpha2.AccessLogSpec{}
					subject.Bucket = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "BucketPrefix - default",
			args: args{
				in: func() kopsv1alpha2.AccessLogSpec {
					subject := kopsv1alpha2.AccessLogSpec{}
					subject.BucketPrefix = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceAccessLogSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceAccessLogSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
