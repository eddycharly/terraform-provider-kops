package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	v1 "k8s.io/client-go/applyconfigurations/meta/v1"
)

func TestExpandDataSourceLabelSelector(t *testing.T) {
	_default := v1.LabelSelector{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want v1.LabelSelector
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"match_labels":      func() map[string]interface{} { return nil }(),
					"match_expressions": func() []interface{} { return nil }(),
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandDataSourceLabelSelector(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceLabelSelector() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceLabelSelectorInto(t *testing.T) {
	_default := map[string]interface{}{
		"match_labels":      func() map[string]interface{} { return nil }(),
		"match_expressions": func() []interface{} { return nil }(),
	}
	type args struct {
		in v1.LabelSelector
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: v1.LabelSelector{},
			},
			want: _default,
		},
		{
			name: "MatchLabels - default",
			args: args{
				in: func() v1.LabelSelector {
					subject := v1.LabelSelector{}
					subject.MatchLabels = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MatchExpressions - default",
			args: args{
				in: func() v1.LabelSelector {
					subject := v1.LabelSelector{}
					subject.MatchExpressions = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenDataSourceLabelSelectorInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceLabelSelector() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceLabelSelector(t *testing.T) {
	_default := map[string]interface{}{
		"match_labels":      func() map[string]interface{} { return nil }(),
		"match_expressions": func() []interface{} { return nil }(),
	}
	type args struct {
		in v1.LabelSelector
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: v1.LabelSelector{},
			},
			want: _default,
		},
		{
			name: "MatchLabels - default",
			args: args{
				in: func() v1.LabelSelector {
					subject := v1.LabelSelector{}
					subject.MatchLabels = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MatchExpressions - default",
			args: args{
				in: func() v1.LabelSelector {
					subject := v1.LabelSelector{}
					subject.MatchExpressions = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceLabelSelector(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceLabelSelector() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
