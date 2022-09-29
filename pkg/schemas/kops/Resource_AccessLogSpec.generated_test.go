package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceAccessLogSpec(t *testing.T) {
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
			got := ExpandResourceAccessLogSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceAccessLogSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceAccessLogSpecInto(t *testing.T) {
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
			FlattenResourceAccessLogSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceAccessLogSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceAccessLogSpec(t *testing.T) {
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
			got := FlattenResourceAccessLogSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceAccessLogSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
