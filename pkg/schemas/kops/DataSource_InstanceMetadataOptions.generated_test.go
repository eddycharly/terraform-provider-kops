package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceInstanceMetadataOptions(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.InstanceMetadataOptions
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandDataSourceInstanceMetadataOptions(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandDataSourceInstanceMetadataOptions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenDataSourceInstanceMetadataOptionsInto(t *testing.T) {
	_default := map[string]interface{}{
		"http_put_response_hop_limit": nil,
		"http_tokens":                 nil,
	}
	type args struct {
		in kops.InstanceMetadataOptions
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.InstanceMetadataOptions{},
			},
			want: _default,
		},
		{
			name: "HTTPPutResponseHopLimit - default",
			args: args{
				in: func() kops.InstanceMetadataOptions {
					subject := kops.InstanceMetadataOptions{}
					subject.HTTPPutResponseHopLimit = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "HTTPTokens - default",
			args: args{
				in: func() kops.InstanceMetadataOptions {
					subject := kops.InstanceMetadataOptions{}
					subject.HTTPTokens = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenDataSourceInstanceMetadataOptionsInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceInstanceMetadataOptions() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceInstanceMetadataOptions(t *testing.T) {
	_default := map[string]interface{}{
		"http_put_response_hop_limit": nil,
		"http_tokens":                 nil,
	}
	type args struct {
		in kops.InstanceMetadataOptions
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.InstanceMetadataOptions{},
			},
			want: _default,
		},
		{
			name: "HTTPPutResponseHopLimit - default",
			args: args{
				in: func() kops.InstanceMetadataOptions {
					subject := kops.InstanceMetadataOptions{}
					subject.HTTPPutResponseHopLimit = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "HTTPTokens - default",
			args: args{
				in: func() kops.InstanceMetadataOptions {
					subject := kops.InstanceMetadataOptions{}
					subject.HTTPTokens = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceInstanceMetadataOptions(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceInstanceMetadataOptions() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
