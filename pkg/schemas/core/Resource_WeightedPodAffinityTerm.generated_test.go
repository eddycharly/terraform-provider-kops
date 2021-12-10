package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	v1 "k8s.io/api/core/v1"
)

func TestExpandResourceWeightedPodAffinityTerm(t *testing.T) {
	_default := v1.WeightedPodAffinityTerm{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want v1.WeightedPodAffinityTerm
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"weight":            0,
					"pod_affinity_term": func() []interface{} { return []interface{}{FlattenResourcePodAffinityTerm(v1.PodAffinityTerm{})} }(),
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandResourceWeightedPodAffinityTerm(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceWeightedPodAffinityTerm() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceWeightedPodAffinityTermInto(t *testing.T) {
	_default := map[string]interface{}{
		"weight":            0,
		"pod_affinity_term": func() []interface{} { return []interface{}{FlattenResourcePodAffinityTerm(v1.PodAffinityTerm{})} }(),
	}
	type args struct {
		in v1.WeightedPodAffinityTerm
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: v1.WeightedPodAffinityTerm{},
			},
			want: _default,
		},
		{
			name: "Weight - default",
			args: args{
				in: func() v1.WeightedPodAffinityTerm {
					subject := v1.WeightedPodAffinityTerm{}
					subject.Weight = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PodAffinityTerm - default",
			args: args{
				in: func() v1.WeightedPodAffinityTerm {
					subject := v1.WeightedPodAffinityTerm{}
					subject.PodAffinityTerm = v1.PodAffinityTerm{}
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenResourceWeightedPodAffinityTermInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceWeightedPodAffinityTerm() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceWeightedPodAffinityTerm(t *testing.T) {
	_default := map[string]interface{}{
		"weight":            0,
		"pod_affinity_term": func() []interface{} { return []interface{}{FlattenResourcePodAffinityTerm(v1.PodAffinityTerm{})} }(),
	}
	type args struct {
		in v1.WeightedPodAffinityTerm
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: v1.WeightedPodAffinityTerm{},
			},
			want: _default,
		},
		{
			name: "Weight - default",
			args: args{
				in: func() v1.WeightedPodAffinityTerm {
					subject := v1.WeightedPodAffinityTerm{}
					subject.Weight = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PodAffinityTerm - default",
			args: args{
				in: func() v1.WeightedPodAffinityTerm {
					subject := v1.WeightedPodAffinityTerm{}
					subject.PodAffinityTerm = v1.PodAffinityTerm{}
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceWeightedPodAffinityTerm(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceWeightedPodAffinityTerm() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
