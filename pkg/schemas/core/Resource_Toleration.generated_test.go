package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	v1 "k8s.io/api/core/v1"
)

func TestExpandResourceToleration(t *testing.T) {
	_default := v1.Toleration{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want v1.Toleration
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
			got := ExpandResourceToleration(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceToleration() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceTolerationInto(t *testing.T) {
	_default := map[string]interface{}{
		"key":                "",
		"operator":           "",
		"value":              "",
		"effect":             "",
		"toleration_seconds": nil,
	}
	type args struct {
		in v1.Toleration
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: v1.Toleration{},
			},
			want: _default,
		},
		{
			name: "Key - default",
			args: args{
				in: func() v1.Toleration {
					subject := v1.Toleration{}
					subject.Key = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Operator - default",
			args: args{
				in: func() v1.Toleration {
					subject := v1.Toleration{}
					subject.Operator = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Value - default",
			args: args{
				in: func() v1.Toleration {
					subject := v1.Toleration{}
					subject.Value = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Effect - default",
			args: args{
				in: func() v1.Toleration {
					subject := v1.Toleration{}
					subject.Effect = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "TolerationSeconds - default",
			args: args{
				in: func() v1.Toleration {
					subject := v1.Toleration{}
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
			FlattenResourceTolerationInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceToleration() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceToleration(t *testing.T) {
	_default := map[string]interface{}{
		"key":                "",
		"operator":           "",
		"value":              "",
		"effect":             "",
		"toleration_seconds": nil,
	}
	type args struct {
		in v1.Toleration
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: v1.Toleration{},
			},
			want: _default,
		},
		{
			name: "Key - default",
			args: args{
				in: func() v1.Toleration {
					subject := v1.Toleration{}
					subject.Key = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Operator - default",
			args: args{
				in: func() v1.Toleration {
					subject := v1.Toleration{}
					subject.Operator = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Value - default",
			args: args{
				in: func() v1.Toleration {
					subject := v1.Toleration{}
					subject.Value = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Effect - default",
			args: args{
				in: func() v1.Toleration {
					subject := v1.Toleration{}
					subject.Effect = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "TolerationSeconds - default",
			args: args{
				in: func() v1.Toleration {
					subject := v1.Toleration{}
					subject.TolerationSeconds = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceToleration(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceToleration() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
