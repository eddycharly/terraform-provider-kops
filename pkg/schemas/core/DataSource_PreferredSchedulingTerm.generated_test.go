package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	v1 "k8s.io/api/core/v1"
)

func TestExpandDataSourcePreferredSchedulingTerm(t *testing.T) {
	_default := v1.PreferredSchedulingTerm{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want v1.PreferredSchedulingTerm
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"weight":     0,
					"preference": func() []interface{} { return []interface{}{FlattenDataSourceNodeSelectorTerm(v1.NodeSelectorTerm{})} }(),
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandDataSourcePreferredSchedulingTerm(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourcePreferredSchedulingTerm() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourcePreferredSchedulingTermInto(t *testing.T) {
	_default := map[string]interface{}{
		"weight":     0,
		"preference": func() []interface{} { return []interface{}{FlattenDataSourceNodeSelectorTerm(v1.NodeSelectorTerm{})} }(),
	}
	type args struct {
		in v1.PreferredSchedulingTerm
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: v1.PreferredSchedulingTerm{},
			},
			want: _default,
		},
		{
			name: "Weight - default",
			args: args{
				in: func() v1.PreferredSchedulingTerm {
					subject := v1.PreferredSchedulingTerm{}
					subject.Weight = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Preference - default",
			args: args{
				in: func() v1.PreferredSchedulingTerm {
					subject := v1.PreferredSchedulingTerm{}
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
			FlattenDataSourcePreferredSchedulingTermInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourcePreferredSchedulingTerm() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourcePreferredSchedulingTerm(t *testing.T) {
	_default := map[string]interface{}{
		"weight":     0,
		"preference": func() []interface{} { return []interface{}{FlattenDataSourceNodeSelectorTerm(v1.NodeSelectorTerm{})} }(),
	}
	type args struct {
		in v1.PreferredSchedulingTerm
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: v1.PreferredSchedulingTerm{},
			},
			want: _default,
		},
		{
			name: "Weight - default",
			args: args{
				in: func() v1.PreferredSchedulingTerm {
					subject := v1.PreferredSchedulingTerm{}
					subject.Weight = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Preference - default",
			args: args{
				in: func() v1.PreferredSchedulingTerm {
					subject := v1.PreferredSchedulingTerm{}
					subject.Preference = v1.NodeSelectorTerm{}
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourcePreferredSchedulingTerm(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourcePreferredSchedulingTerm() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
