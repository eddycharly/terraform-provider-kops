package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceNodeAuthorizerSpec(t *testing.T) {
	_default := kops.NodeAuthorizerSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.NodeAuthorizerSpec
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
			got := ExpandDataSourceNodeAuthorizerSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceNodeAuthorizerSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceNodeAuthorizerSpecInto(t *testing.T) {
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenDataSourceNodeAuthorizerSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceNodeAuthorizerSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceNodeAuthorizerSpec(t *testing.T) {
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceNodeAuthorizerSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceNodeAuthorizerSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
