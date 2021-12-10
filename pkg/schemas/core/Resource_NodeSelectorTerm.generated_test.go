package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	v1 "k8s.io/api/core/v1"
)

func TestExpandResourceNodeSelectorTerm(t *testing.T) {
	_default := v1.NodeSelectorTerm{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want v1.NodeSelectorTerm
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"match_expressions": func() []interface{} { return nil }(),
					"match_fields":      func() []interface{} { return nil }(),
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandResourceNodeSelectorTerm(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceNodeSelectorTerm() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceNodeSelectorTermInto(t *testing.T) {
	_default := map[string]interface{}{
		"match_expressions": func() []interface{} { return nil }(),
		"match_fields":      func() []interface{} { return nil }(),
	}
	type args struct {
		in v1.NodeSelectorTerm
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: v1.NodeSelectorTerm{},
			},
			want: _default,
		},
		{
			name: "MatchExpressions - default",
			args: args{
				in: func() v1.NodeSelectorTerm {
					subject := v1.NodeSelectorTerm{}
					subject.MatchExpressions = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MatchFields - default",
			args: args{
				in: func() v1.NodeSelectorTerm {
					subject := v1.NodeSelectorTerm{}
					subject.MatchFields = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenResourceNodeSelectorTermInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceNodeSelectorTerm() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceNodeSelectorTerm(t *testing.T) {
	_default := map[string]interface{}{
		"match_expressions": func() []interface{} { return nil }(),
		"match_fields":      func() []interface{} { return nil }(),
	}
	type args struct {
		in v1.NodeSelectorTerm
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: v1.NodeSelectorTerm{},
			},
			want: _default,
		},
		{
			name: "MatchExpressions - default",
			args: args{
				in: func() v1.NodeSelectorTerm {
					subject := v1.NodeSelectorTerm{}
					subject.MatchExpressions = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MatchFields - default",
			args: args{
				in: func() v1.NodeSelectorTerm {
					subject := v1.NodeSelectorTerm{}
					subject.MatchFields = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceNodeSelectorTerm(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceNodeSelectorTerm() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
