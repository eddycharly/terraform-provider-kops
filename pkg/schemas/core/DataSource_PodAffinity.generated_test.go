package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	core "k8s.io/api/core/v1"
)

func TestExpandDataSourcePodAffinity(t *testing.T) {
	_default := core.PodAffinity{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want core.PodAffinity
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
			got := ExpandDataSourcePodAffinity(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourcePodAffinity() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourcePodAffinityInto(t *testing.T) {
	_default := map[string]interface{}{
		"required_during_scheduling_ignored_during_execution":  func() []interface{} { return nil }(),
		"preferred_during_scheduling_ignored_during_execution": func() []interface{} { return nil }(),
	}
	type args struct {
		in core.PodAffinity
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: core.PodAffinity{},
			},
			want: _default,
		},
		{
			name: "RequiredDuringSchedulingIgnoredDuringExecution - default",
			args: args{
				in: func() core.PodAffinity {
					subject := core.PodAffinity{}
					subject.RequiredDuringSchedulingIgnoredDuringExecution = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PreferredDuringSchedulingIgnoredDuringExecution - default",
			args: args{
				in: func() core.PodAffinity {
					subject := core.PodAffinity{}
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
			FlattenDataSourcePodAffinityInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourcePodAffinity() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourcePodAffinity(t *testing.T) {
	_default := map[string]interface{}{
		"required_during_scheduling_ignored_during_execution":  func() []interface{} { return nil }(),
		"preferred_during_scheduling_ignored_during_execution": func() []interface{} { return nil }(),
	}
	type args struct {
		in core.PodAffinity
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: core.PodAffinity{},
			},
			want: _default,
		},
		{
			name: "RequiredDuringSchedulingIgnoredDuringExecution - default",
			args: args{
				in: func() core.PodAffinity {
					subject := core.PodAffinity{}
					subject.RequiredDuringSchedulingIgnoredDuringExecution = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PreferredDuringSchedulingIgnoredDuringExecution - default",
			args: args{
				in: func() core.PodAffinity {
					subject := core.PodAffinity{}
					subject.PreferredDuringSchedulingIgnoredDuringExecution = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourcePodAffinity(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourcePodAffinity() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
