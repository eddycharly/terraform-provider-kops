package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceInstanceMetadataOptions(t *testing.T) {
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
			if got := ExpandResourceInstanceMetadataOptions(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandResourceInstanceMetadataOptions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenResourceInstanceMetadataOptionsInto(t *testing.T) {
	type args struct {
		in  kops.InstanceMetadataOptions
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
			FlattenResourceInstanceMetadataOptionsInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenResourceInstanceMetadataOptions(t *testing.T) {
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
			want: map[string]interface{}{
				"http_put_response_hop_limit": nil,
				"http_tokens":                 nil,
			},
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
			want: map[string]interface{}{
				"http_put_response_hop_limit": nil,
				"http_tokens":                 nil,
			},
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
			want: map[string]interface{}{
				"http_put_response_hop_limit": nil,
				"http_tokens":                 nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenResourceInstanceMetadataOptions(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenResourceInstanceMetadataOptions() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
