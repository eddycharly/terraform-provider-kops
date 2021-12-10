package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	core "k8s.io/api/core/v1"
)

func TestExpandResourcePodAffinityTerm(t *testing.T) {
	_default := core.PodAffinityTerm{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want core.PodAffinityTerm
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"label_selector":     nil,
					"namespaces":         func() []interface{} { return nil }(),
					"topology_key":       "",
					"namespace_selector": nil,
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandResourcePodAffinityTerm(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourcePodAffinityTerm() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourcePodAffinityTermInto(t *testing.T) {
	_default := map[string]interface{}{
		"label_selector":     nil,
		"namespaces":         func() []interface{} { return nil }(),
		"topology_key":       "",
		"namespace_selector": nil,
	}
	type args struct {
		in core.PodAffinityTerm
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: core.PodAffinityTerm{},
			},
			want: _default,
		},
		{
			name: "LabelSelector - default",
			args: args{
				in: func() core.PodAffinityTerm {
					subject := core.PodAffinityTerm{}
					subject.LabelSelector = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Namespaces - default",
			args: args{
				in: func() core.PodAffinityTerm {
					subject := core.PodAffinityTerm{}
					subject.Namespaces = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "TopologyKey - default",
			args: args{
				in: func() core.PodAffinityTerm {
					subject := core.PodAffinityTerm{}
					subject.TopologyKey = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NamespaceSelector - default",
			args: args{
				in: func() core.PodAffinityTerm {
					subject := core.PodAffinityTerm{}
					subject.NamespaceSelector = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenResourcePodAffinityTermInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourcePodAffinityTerm() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourcePodAffinityTerm(t *testing.T) {
	_default := map[string]interface{}{
		"label_selector":     nil,
		"namespaces":         func() []interface{} { return nil }(),
		"topology_key":       "",
		"namespace_selector": nil,
	}
	type args struct {
		in core.PodAffinityTerm
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: core.PodAffinityTerm{},
			},
			want: _default,
		},
		{
			name: "LabelSelector - default",
			args: args{
				in: func() core.PodAffinityTerm {
					subject := core.PodAffinityTerm{}
					subject.LabelSelector = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Namespaces - default",
			args: args{
				in: func() core.PodAffinityTerm {
					subject := core.PodAffinityTerm{}
					subject.Namespaces = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "TopologyKey - default",
			args: args{
				in: func() core.PodAffinityTerm {
					subject := core.PodAffinityTerm{}
					subject.TopologyKey = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NamespaceSelector - default",
			args: args{
				in: func() core.PodAffinityTerm {
					subject := core.PodAffinityTerm{}
					subject.NamespaceSelector = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourcePodAffinityTerm(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourcePodAffinityTerm() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
