package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	core "k8s.io/api/core/v1"
)

func TestExpandDataSourceNodeAffinity(t *testing.T) {
	_default := core.NodeAffinity{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want core.NodeAffinity
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"required_during_scheduling_ignored_during_execution":  nil,
					"preferred_during_scheduling_ignored_during_execution": func() []interface{} { return nil }(),
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandDataSourceNodeAffinity(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceNodeAffinity() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceNodeAffinityInto(t *testing.T) {
	_default := map[string]interface{}{
		"required_during_scheduling_ignored_during_execution":  nil,
		"preferred_during_scheduling_ignored_during_execution": func() []interface{} { return nil }(),
	}
	type args struct {
		in core.NodeAffinity
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: core.NodeAffinity{},
			},
			want: _default,
		},
		{
			name: "RequiredDuringSchedulingIgnoredDuringExecution - default",
			args: args{
				in: func() core.NodeAffinity {
					subject := core.NodeAffinity{}
					subject.RequiredDuringSchedulingIgnoredDuringExecution = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PreferredDuringSchedulingIgnoredDuringExecution - default",
			args: args{
				in: func() core.NodeAffinity {
					subject := core.NodeAffinity{}
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
			FlattenDataSourceNodeAffinityInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceNodeAffinity() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceNodeAffinity(t *testing.T) {
	_default := map[string]interface{}{
		"required_during_scheduling_ignored_during_execution":  nil,
		"preferred_during_scheduling_ignored_during_execution": func() []interface{} { return nil }(),
	}
	type args struct {
		in core.NodeAffinity
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: core.NodeAffinity{},
			},
			want: _default,
		},
		{
			name: "RequiredDuringSchedulingIgnoredDuringExecution - default",
			args: args{
				in: func() core.NodeAffinity {
					subject := core.NodeAffinity{}
					subject.RequiredDuringSchedulingIgnoredDuringExecution = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PreferredDuringSchedulingIgnoredDuringExecution - default",
			args: args{
				in: func() core.NodeAffinity {
					subject := core.NodeAffinity{}
					subject.PreferredDuringSchedulingIgnoredDuringExecution = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceNodeAffinity(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceNodeAffinity() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
