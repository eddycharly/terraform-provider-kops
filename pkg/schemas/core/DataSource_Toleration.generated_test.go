package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	core "k8s.io/api/core/v1"
)

func TestExpandDataSourceToleration(t *testing.T) {
	_default := core.Toleration{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want core.Toleration
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"key":                "",
					"operator":           "",
					"value":              "",
					"effect":             "",
					"toleration_seconds": nil,
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandDataSourceToleration(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceToleration() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceTolerationInto(t *testing.T) {
	_default := map[string]interface{}{
		"key":                "",
		"operator":           "",
		"value":              "",
		"effect":             "",
		"toleration_seconds": nil,
	}
	type args struct {
		in core.Toleration
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: core.Toleration{},
			},
			want: _default,
		},
		{
			name: "Key - default",
			args: args{
				in: func() core.Toleration {
					subject := core.Toleration{}
					subject.Key = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Operator - default",
			args: args{
				in: func() core.Toleration {
					subject := core.Toleration{}
					subject.Operator = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Value - default",
			args: args{
				in: func() core.Toleration {
					subject := core.Toleration{}
					subject.Value = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Effect - default",
			args: args{
				in: func() core.Toleration {
					subject := core.Toleration{}
					subject.Effect = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "TolerationSeconds - default",
			args: args{
				in: func() core.Toleration {
					subject := core.Toleration{}
					subject.TolerationSeconds = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenDataSourceTolerationInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceToleration() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceToleration(t *testing.T) {
	_default := map[string]interface{}{
		"key":                "",
		"operator":           "",
		"value":              "",
		"effect":             "",
		"toleration_seconds": nil,
	}
	type args struct {
		in core.Toleration
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: core.Toleration{},
			},
			want: _default,
		},
		{
			name: "Key - default",
			args: args{
				in: func() core.Toleration {
					subject := core.Toleration{}
					subject.Key = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Operator - default",
			args: args{
				in: func() core.Toleration {
					subject := core.Toleration{}
					subject.Operator = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Value - default",
			args: args{
				in: func() core.Toleration {
					subject := core.Toleration{}
					subject.Value = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Effect - default",
			args: args{
				in: func() core.Toleration {
					subject := core.Toleration{}
					subject.Effect = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "TolerationSeconds - default",
			args: args{
				in: func() core.Toleration {
					subject := core.Toleration{}
					subject.TolerationSeconds = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceToleration(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceToleration() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
