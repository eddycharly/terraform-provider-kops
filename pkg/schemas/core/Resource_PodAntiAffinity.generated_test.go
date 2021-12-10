package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	core "k8s.io/api/core/v1"
)

func TestExpandResourcePodAntiAffinity(t *testing.T) {
	_default := core.PodAntiAffinity{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want core.PodAntiAffinity
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"required_during_scheduling_ignored_during_execution":  func() []interface{} { return nil }(),
					"preferred_during_scheduling_ignored_during_execution": func() []interface{} { return nil }(),
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandResourcePodAntiAffinity(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourcePodAntiAffinity() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourcePodAntiAffinityInto(t *testing.T) {
	_default := map[string]interface{}{
		"required_during_scheduling_ignored_during_execution":  func() []interface{} { return nil }(),
		"preferred_during_scheduling_ignored_during_execution": func() []interface{} { return nil }(),
	}
	type args struct {
		in core.PodAntiAffinity
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: core.PodAntiAffinity{},
			},
			want: _default,
		},
		{
			name: "RequiredDuringSchedulingIgnoredDuringExecution - default",
			args: args{
				in: func() core.PodAntiAffinity {
					subject := core.PodAntiAffinity{}
					subject.RequiredDuringSchedulingIgnoredDuringExecution = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PreferredDuringSchedulingIgnoredDuringExecution - default",
			args: args{
				in: func() core.PodAntiAffinity {
					subject := core.PodAntiAffinity{}
					subject.PreferredDuringSchedulingIgnoredDuringExecution = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenResourcePodAntiAffinityInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourcePodAntiAffinity() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourcePodAntiAffinity(t *testing.T) {
	_default := map[string]interface{}{
		"required_during_scheduling_ignored_during_execution":  func() []interface{} { return nil }(),
		"preferred_during_scheduling_ignored_during_execution": func() []interface{} { return nil }(),
	}
	type args struct {
		in core.PodAntiAffinity
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: core.PodAntiAffinity{},
			},
			want: _default,
		},
		{
			name: "RequiredDuringSchedulingIgnoredDuringExecution - default",
			args: args{
				in: func() core.PodAntiAffinity {
					subject := core.PodAntiAffinity{}
					subject.RequiredDuringSchedulingIgnoredDuringExecution = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PreferredDuringSchedulingIgnoredDuringExecution - default",
			args: args{
				in: func() core.PodAntiAffinity {
					subject := core.PodAntiAffinity{}
					subject.PreferredDuringSchedulingIgnoredDuringExecution = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourcePodAntiAffinity(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourcePodAntiAffinity() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
