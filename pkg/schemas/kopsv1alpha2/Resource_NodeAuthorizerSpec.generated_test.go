package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

func TestExpandResourceNodeAuthorizerSpec(t *testing.T) {
	_default := kopsv1alpha2.NodeAuthorizerSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kopsv1alpha2.NodeAuthorizerSpec
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
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
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandResourceNodeAuthorizerSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceNodeAuthorizerSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceNodeAuthorizerSpecInto(t *testing.T) {
	_default := map[string]interface{}{
		"authorizer": "",
		"features":   func() []interface{} { return nil }(),
		"image":      "",
		"node_url":   "",
		"port":       0,
		"interval":   nil,
		"timeout":    nil,
		"token_ttl":  nil,
	}
	type args struct {
		in kopsv1alpha2.NodeAuthorizerSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.NodeAuthorizerSpec{},
			},
			want: _default,
		},
		{
			name: "Authorizer - default",
			args: args{
				in: func() kopsv1alpha2.NodeAuthorizerSpec {
					subject := kopsv1alpha2.NodeAuthorizerSpec{}
					subject.Authorizer = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Features - default",
			args: args{
				in: func() kopsv1alpha2.NodeAuthorizerSpec {
					subject := kopsv1alpha2.NodeAuthorizerSpec{}
					subject.Features = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Image - default",
			args: args{
				in: func() kopsv1alpha2.NodeAuthorizerSpec {
					subject := kopsv1alpha2.NodeAuthorizerSpec{}
					subject.Image = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NodeURL - default",
			args: args{
				in: func() kopsv1alpha2.NodeAuthorizerSpec {
					subject := kopsv1alpha2.NodeAuthorizerSpec{}
					subject.NodeURL = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Port - default",
			args: args{
				in: func() kopsv1alpha2.NodeAuthorizerSpec {
					subject := kopsv1alpha2.NodeAuthorizerSpec{}
					subject.Port = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Interval - default",
			args: args{
				in: func() kopsv1alpha2.NodeAuthorizerSpec {
					subject := kopsv1alpha2.NodeAuthorizerSpec{}
					subject.Interval = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Timeout - default",
			args: args{
				in: func() kopsv1alpha2.NodeAuthorizerSpec {
					subject := kopsv1alpha2.NodeAuthorizerSpec{}
					subject.Timeout = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "TokenTTL - default",
			args: args{
				in: func() kopsv1alpha2.NodeAuthorizerSpec {
					subject := kopsv1alpha2.NodeAuthorizerSpec{}
					subject.TokenTTL = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenResourceNodeAuthorizerSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceNodeAuthorizerSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceNodeAuthorizerSpec(t *testing.T) {
	_default := map[string]interface{}{
		"authorizer": "",
		"features":   func() []interface{} { return nil }(),
		"image":      "",
		"node_url":   "",
		"port":       0,
		"interval":   nil,
		"timeout":    nil,
		"token_ttl":  nil,
	}
	type args struct {
		in kopsv1alpha2.NodeAuthorizerSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.NodeAuthorizerSpec{},
			},
			want: _default,
		},
		{
			name: "Authorizer - default",
			args: args{
				in: func() kopsv1alpha2.NodeAuthorizerSpec {
					subject := kopsv1alpha2.NodeAuthorizerSpec{}
					subject.Authorizer = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Features - default",
			args: args{
				in: func() kopsv1alpha2.NodeAuthorizerSpec {
					subject := kopsv1alpha2.NodeAuthorizerSpec{}
					subject.Features = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Image - default",
			args: args{
				in: func() kopsv1alpha2.NodeAuthorizerSpec {
					subject := kopsv1alpha2.NodeAuthorizerSpec{}
					subject.Image = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NodeURL - default",
			args: args{
				in: func() kopsv1alpha2.NodeAuthorizerSpec {
					subject := kopsv1alpha2.NodeAuthorizerSpec{}
					subject.NodeURL = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Port - default",
			args: args{
				in: func() kopsv1alpha2.NodeAuthorizerSpec {
					subject := kopsv1alpha2.NodeAuthorizerSpec{}
					subject.Port = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Interval - default",
			args: args{
				in: func() kopsv1alpha2.NodeAuthorizerSpec {
					subject := kopsv1alpha2.NodeAuthorizerSpec{}
					subject.Interval = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Timeout - default",
			args: args{
				in: func() kopsv1alpha2.NodeAuthorizerSpec {
					subject := kopsv1alpha2.NodeAuthorizerSpec{}
					subject.Timeout = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "TokenTTL - default",
			args: args{
				in: func() kopsv1alpha2.NodeAuthorizerSpec {
					subject := kopsv1alpha2.NodeAuthorizerSpec{}
					subject.TokenTTL = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceNodeAuthorizerSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceNodeAuthorizerSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
