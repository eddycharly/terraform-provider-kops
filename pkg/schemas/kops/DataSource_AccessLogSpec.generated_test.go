package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceAccessLogSpec(t *testing.T) {
	_default := kops.AccessLogSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.AccessLogSpec
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
		in kops.AccessLogSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.AccessLogSpec{},
			},
			want: _default,
		},
		{
			name: "Interval - default",
			args: args{
				in: func() kops.AccessLogSpec {
					subject := kops.AccessLogSpec{}
					subject.Interval = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Bucket - default",
			args: args{
				in: func() kops.AccessLogSpec {
					subject := kops.AccessLogSpec{}
					subject.Bucket = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "BucketPrefix - default",
			args: args{
				in: func() kops.AccessLogSpec {
					subject := kops.AccessLogSpec{}
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
		in kops.AccessLogSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.AccessLogSpec{},
			},
			want: _default,
		},
		{
			name: "Interval - default",
			args: args{
				in: func() kops.AccessLogSpec {
					subject := kops.AccessLogSpec{}
					subject.Interval = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Bucket - default",
			args: args{
				in: func() kops.AccessLogSpec {
					subject := kops.AccessLogSpec{}
					subject.Bucket = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "BucketPrefix - default",
			args: args{
				in: func() kops.AccessLogSpec {
					subject := kops.AccessLogSpec{}
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
