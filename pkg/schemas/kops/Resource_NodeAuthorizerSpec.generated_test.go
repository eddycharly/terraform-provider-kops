package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceNodeAuthorizerSpec(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.NodeAuthorizerSpec
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandResourceNodeAuthorizerSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandResourceNodeAuthorizerSpec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenResourceNodeAuthorizerSpecInto(t *testing.T) {
	type args struct {
		in  kops.NodeAuthorizerSpec
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
			FlattenResourceNodeAuthorizerSpecInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenResourceNodeAuthorizerSpec(t *testing.T) {
	type args struct {
		in kops.NodeAuthorizerSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.NodeAuthorizerSpec{},
			},
			want: map[string]interface{}{
				"authorizer": "",
				"features":   func() []interface{} { return nil }(),
				"image":      "",
				"node_url":   "",
				"port":       0,
				"interval":   nil,
				"timeout":    nil,
				"token_ttl":  nil,
			},
		},
		{
			name: "Authorizer - default",
			args: args{
				in: func() kops.NodeAuthorizerSpec {
					subject := kops.NodeAuthorizerSpec{}
					subject.Authorizer = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
				"authorizer": "",
				"features":   func() []interface{} { return nil }(),
				"image":      "",
				"node_url":   "",
				"port":       0,
				"interval":   nil,
				"timeout":    nil,
				"token_ttl":  nil,
			},
		},
		{
			name: "Features - default",
			args: args{
				in: func() kops.NodeAuthorizerSpec {
					subject := kops.NodeAuthorizerSpec{}
					subject.Features = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"authorizer": "",
				"features":   func() []interface{} { return nil }(),
				"image":      "",
				"node_url":   "",
				"port":       0,
				"interval":   nil,
				"timeout":    nil,
				"token_ttl":  nil,
			},
		},
		{
			name: "Image - default",
			args: args{
				in: func() kops.NodeAuthorizerSpec {
					subject := kops.NodeAuthorizerSpec{}
					subject.Image = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
				"authorizer": "",
				"features":   func() []interface{} { return nil }(),
				"image":      "",
				"node_url":   "",
				"port":       0,
				"interval":   nil,
				"timeout":    nil,
				"token_ttl":  nil,
			},
		},
		{
			name: "NodeURL - default",
			args: args{
				in: func() kops.NodeAuthorizerSpec {
					subject := kops.NodeAuthorizerSpec{}
					subject.NodeURL = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
				"authorizer": "",
				"features":   func() []interface{} { return nil }(),
				"image":      "",
				"node_url":   "",
				"port":       0,
				"interval":   nil,
				"timeout":    nil,
				"token_ttl":  nil,
			},
		},
		{
			name: "Port - default",
			args: args{
				in: func() kops.NodeAuthorizerSpec {
					subject := kops.NodeAuthorizerSpec{}
					subject.Port = 0
					return subject
				}(),
			},
			want: map[string]interface{}{
				"authorizer": "",
				"features":   func() []interface{} { return nil }(),
				"image":      "",
				"node_url":   "",
				"port":       0,
				"interval":   nil,
				"timeout":    nil,
				"token_ttl":  nil,
			},
		},
		{
			name: "Interval - default",
			args: args{
				in: func() kops.NodeAuthorizerSpec {
					subject := kops.NodeAuthorizerSpec{}
					subject.Interval = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"authorizer": "",
				"features":   func() []interface{} { return nil }(),
				"image":      "",
				"node_url":   "",
				"port":       0,
				"interval":   nil,
				"timeout":    nil,
				"token_ttl":  nil,
			},
		},
		{
			name: "Timeout - default",
			args: args{
				in: func() kops.NodeAuthorizerSpec {
					subject := kops.NodeAuthorizerSpec{}
					subject.Timeout = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"authorizer": "",
				"features":   func() []interface{} { return nil }(),
				"image":      "",
				"node_url":   "",
				"port":       0,
				"interval":   nil,
				"timeout":    nil,
				"token_ttl":  nil,
			},
		},
		{
			name: "TokenTTL - default",
			args: args{
				in: func() kops.NodeAuthorizerSpec {
					subject := kops.NodeAuthorizerSpec{}
					subject.TokenTTL = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"authorizer": "",
				"features":   func() []interface{} { return nil }(),
				"image":      "",
				"node_url":   "",
				"port":       0,
				"interval":   nil,
				"timeout":    nil,
				"token_ttl":  nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenResourceNodeAuthorizerSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenResourceNodeAuthorizerSpec() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
