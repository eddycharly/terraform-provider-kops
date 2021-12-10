package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	core "k8s.io/api/core/v1"
)

func TestExpandResourceAffinity(t *testing.T) {
	_default := core.Affinity{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want core.Affinity
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"node_affinity":     nil,
					"pod_affinity":      nil,
					"pod_anti_affinity": nil,
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandResourceAffinity(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceAffinity() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceAffinityInto(t *testing.T) {
	_default := map[string]interface{}{
		"node_affinity":     nil,
		"pod_affinity":      nil,
		"pod_anti_affinity": nil,
	}
	type args struct {
		in core.Affinity
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: core.Affinity{},
			},
			want: _default,
		},
		{
			name: "NodeAffinity - default",
			args: args{
				in: func() core.Affinity {
					subject := core.Affinity{}
					subject.NodeAffinity = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PodAffinity - default",
			args: args{
				in: func() core.Affinity {
					subject := core.Affinity{}
					subject.PodAffinity = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PodAntiAffinity - default",
			args: args{
				in: func() core.Affinity {
					subject := core.Affinity{}
					subject.PodAntiAffinity = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenResourceAffinityInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceAffinity() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceAffinity(t *testing.T) {
	_default := map[string]interface{}{
		"node_affinity":     nil,
		"pod_affinity":      nil,
		"pod_anti_affinity": nil,
	}
	type args struct {
		in core.Affinity
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: core.Affinity{},
			},
			want: _default,
		},
		{
			name: "NodeAffinity - default",
			args: args{
				in: func() core.Affinity {
					subject := core.Affinity{}
					subject.NodeAffinity = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PodAffinity - default",
			args: args{
				in: func() core.Affinity {
					subject := core.Affinity{}
					subject.PodAffinity = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PodAntiAffinity - default",
			args: args{
				in: func() core.Affinity {
					subject := core.Affinity{}
					subject.PodAntiAffinity = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceAffinity(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceAffinity() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
