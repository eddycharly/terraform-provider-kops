package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	v1 "k8s.io/api/core/v1"
)

func TestExpandDataSourcePodAffinityTerm(t *testing.T) {
	_default := v1.PodAffinityTerm{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want v1.PodAffinityTerm
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
			got := ExpandDataSourcePodAffinityTerm(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourcePodAffinityTerm() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourcePodAffinityTermInto(t *testing.T) {
	_default := map[string]interface{}{
		"label_selector":     nil,
		"namespaces":         func() []interface{} { return nil }(),
		"topology_key":       "",
		"namespace_selector": nil,
	}
	type args struct {
		in v1.PodAffinityTerm
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: v1.PodAffinityTerm{},
			},
			want: _default,
		},
		{
			name: "LabelSelector - default",
			args: args{
				in: func() v1.PodAffinityTerm {
					subject := v1.PodAffinityTerm{}
					subject.LabelSelector = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Namespaces - default",
			args: args{
				in: func() v1.PodAffinityTerm {
					subject := v1.PodAffinityTerm{}
					subject.Namespaces = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "TopologyKey - default",
			args: args{
				in: func() v1.PodAffinityTerm {
					subject := v1.PodAffinityTerm{}
					subject.TopologyKey = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NamespaceSelector - default",
			args: args{
				in: func() v1.PodAffinityTerm {
					subject := v1.PodAffinityTerm{}
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
			FlattenDataSourcePodAffinityTermInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourcePodAffinityTerm() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourcePodAffinityTerm(t *testing.T) {
	_default := map[string]interface{}{
		"label_selector":     nil,
		"namespaces":         func() []interface{} { return nil }(),
		"topology_key":       "",
		"namespace_selector": nil,
	}
	type args struct {
		in v1.PodAffinityTerm
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: v1.PodAffinityTerm{},
			},
			want: _default,
		},
		{
			name: "LabelSelector - default",
			args: args{
				in: func() v1.PodAffinityTerm {
					subject := v1.PodAffinityTerm{}
					subject.LabelSelector = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Namespaces - default",
			args: args{
				in: func() v1.PodAffinityTerm {
					subject := v1.PodAffinityTerm{}
					subject.Namespaces = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "TopologyKey - default",
			args: args{
				in: func() v1.PodAffinityTerm {
					subject := v1.PodAffinityTerm{}
					subject.TopologyKey = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NamespaceSelector - default",
			args: args{
				in: func() v1.PodAffinityTerm {
					subject := v1.PodAffinityTerm{}
					subject.NamespaceSelector = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourcePodAffinityTerm(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourcePodAffinityTerm() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
