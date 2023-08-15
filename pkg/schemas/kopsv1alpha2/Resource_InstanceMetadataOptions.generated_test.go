package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

func TestExpandResourceInstanceMetadataOptions(t *testing.T) {
	_default := kopsv1alpha2.InstanceMetadataOptions{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kopsv1alpha2.InstanceMetadataOptions
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"http_put_response_hop_limit": nil,
					"http_tokens":                 nil,
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandResourceInstanceMetadataOptions(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceInstanceMetadataOptions() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceInstanceMetadataOptionsInto(t *testing.T) {
	_default := map[string]interface{}{
		"http_put_response_hop_limit": nil,
		"http_tokens":                 nil,
	}
	type args struct {
		in kopsv1alpha2.InstanceMetadataOptions
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.InstanceMetadataOptions{},
			},
			want: _default,
		},
		{
			name: "HTTPPutResponseHopLimit - default",
			args: args{
				in: func() kopsv1alpha2.InstanceMetadataOptions {
					subject := kopsv1alpha2.InstanceMetadataOptions{}
					subject.HTTPPutResponseHopLimit = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "HTTPTokens - default",
			args: args{
				in: func() kopsv1alpha2.InstanceMetadataOptions {
					subject := kopsv1alpha2.InstanceMetadataOptions{}
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
			FlattenResourceInstanceMetadataOptionsInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceInstanceMetadataOptions() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceInstanceMetadataOptions(t *testing.T) {
	_default := map[string]interface{}{
		"http_put_response_hop_limit": nil,
		"http_tokens":                 nil,
	}
	type args struct {
		in kopsv1alpha2.InstanceMetadataOptions
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.InstanceMetadataOptions{},
			},
			want: _default,
		},
		{
			name: "HTTPPutResponseHopLimit - default",
			args: args{
				in: func() kopsv1alpha2.InstanceMetadataOptions {
					subject := kopsv1alpha2.InstanceMetadataOptions{}
					subject.HTTPPutResponseHopLimit = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "HTTPTokens - default",
			args: args{
				in: func() kopsv1alpha2.InstanceMetadataOptions {
					subject := kopsv1alpha2.InstanceMetadataOptions{}
					subject.HTTPTokens = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceInstanceMetadataOptions(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceInstanceMetadataOptions() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
