package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestExpandResourceLabelSelector(t *testing.T) {
	_default := meta.LabelSelector{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want meta.LabelSelector
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
			got := ExpandResourceLabelSelector(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceLabelSelector() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceLabelSelectorInto(t *testing.T) {
	_default := map[string]interface{}{
		"match_labels":      func() map[string]interface{} { return nil }(),
		"match_expressions": func() []interface{} { return nil }(),
	}
	type args struct {
		in meta.LabelSelector
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: meta.LabelSelector{},
			},
			want: _default,
		},
		{
			name: "MatchLabels - default",
			args: args{
				in: func() meta.LabelSelector {
					subject := meta.LabelSelector{}
					subject.MatchLabels = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MatchExpressions - default",
			args: args{
				in: func() meta.LabelSelector {
					subject := meta.LabelSelector{}
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
			FlattenResourceLabelSelectorInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceLabelSelector() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceLabelSelector(t *testing.T) {
	_default := map[string]interface{}{
		"match_labels":      func() map[string]interface{} { return nil }(),
		"match_expressions": func() []interface{} { return nil }(),
	}
	type args struct {
		in meta.LabelSelector
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: meta.LabelSelector{},
			},
			want: _default,
		},
		{
			name: "MatchLabels - default",
			args: args{
				in: func() meta.LabelSelector {
					subject := meta.LabelSelector{}
					subject.MatchLabels = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MatchExpressions - default",
			args: args{
				in: func() meta.LabelSelector {
					subject := meta.LabelSelector{}
					subject.MatchExpressions = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceLabelSelector(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceLabelSelector() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
