package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	core "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
)

func TestExpandResourcePreferredSchedulingTerm(t *testing.T) {
	_default := core.PreferredSchedulingTerm{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want core.PreferredSchedulingTerm
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"weight":     0,
					"preference": func() []interface{} { return []interface{}{FlattenResourceNodeSelectorTerm(core.NodeSelectorTerm{})} }(),
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandResourcePreferredSchedulingTerm(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourcePreferredSchedulingTerm() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourcePreferredSchedulingTermInto(t *testing.T) {
	_default := map[string]interface{}{
		"weight":     0,
		"preference": func() []interface{} { return []interface{}{FlattenResourceNodeSelectorTerm(core.NodeSelectorTerm{})} }(),
	}
	type args struct {
		in core.PreferredSchedulingTerm
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: core.PreferredSchedulingTerm{},
			},
			want: _default,
		},
		{
			name: "Weight - default",
			args: args{
				in: func() core.PreferredSchedulingTerm {
					subject := core.PreferredSchedulingTerm{}
					subject.Weight = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Preference - default",
			args: args{
				in: func() core.PreferredSchedulingTerm {
					subject := core.PreferredSchedulingTerm{}
					subject.Preference = v1.NodeSelectorTerm{}
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenResourcePreferredSchedulingTermInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourcePreferredSchedulingTerm() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourcePreferredSchedulingTerm(t *testing.T) {
	_default := map[string]interface{}{
		"weight":     0,
		"preference": func() []interface{} { return []interface{}{FlattenResourceNodeSelectorTerm(core.NodeSelectorTerm{})} }(),
	}
	type args struct {
		in core.PreferredSchedulingTerm
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: core.PreferredSchedulingTerm{},
			},
			want: _default,
		},
		{
			name: "Weight - default",
			args: args{
				in: func() core.PreferredSchedulingTerm {
					subject := core.PreferredSchedulingTerm{}
					subject.Weight = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Preference - default",
			args: args{
				in: func() core.PreferredSchedulingTerm {
					subject := core.PreferredSchedulingTerm{}
					subject.Preference = v1.NodeSelectorTerm{}
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourcePreferredSchedulingTerm(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourcePreferredSchedulingTerm() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
