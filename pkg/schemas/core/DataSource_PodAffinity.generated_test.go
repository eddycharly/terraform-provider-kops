package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	v1 "k8s.io/api/core/v1"
)

func TestExpandDataSourcePodAffinity(t *testing.T) {
	_default := v1.PodAffinity{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want v1.PodAffinity
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
		in v1.PodAffinity
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: v1.PodAffinity{},
			},
			want: _default,
		},
		{
			name: "RequiredDuringSchedulingIgnoredDuringExecution - default",
			args: args{
				in: func() v1.PodAffinity {
					subject := v1.PodAffinity{}
					subject.RequiredDuringSchedulingIgnoredDuringExecution = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PreferredDuringSchedulingIgnoredDuringExecution - default",
			args: args{
				in: func() v1.PodAffinity {
					subject := v1.PodAffinity{}
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
		in v1.PodAffinity
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: v1.PodAffinity{},
			},
			want: _default,
		},
		{
			name: "RequiredDuringSchedulingIgnoredDuringExecution - default",
			args: args{
				in: func() v1.PodAffinity {
					subject := v1.PodAffinity{}
					subject.RequiredDuringSchedulingIgnoredDuringExecution = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PreferredDuringSchedulingIgnoredDuringExecution - default",
			args: args{
				in: func() v1.PodAffinity {
					subject := v1.PodAffinity{}
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
