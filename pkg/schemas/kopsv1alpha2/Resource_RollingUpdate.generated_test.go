package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

func TestExpandResourceRollingUpdate(t *testing.T) {
	_default := kopsv1alpha2.RollingUpdate{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kopsv1alpha2.RollingUpdate
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"drain_and_terminate": nil,
					"max_unavailable":     nil,
					"max_surge":           nil,
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandResourceRollingUpdate(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceRollingUpdate() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceRollingUpdateInto(t *testing.T) {
	_default := map[string]interface{}{
		"drain_and_terminate": nil,
		"max_unavailable":     nil,
		"max_surge":           nil,
	}
	type args struct {
		in kopsv1alpha2.RollingUpdate
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.RollingUpdate{},
			},
			want: _default,
		},
		{
			name: "DrainAndTerminate - default",
			args: args{
				in: func() kopsv1alpha2.RollingUpdate {
					subject := kopsv1alpha2.RollingUpdate{}
					subject.DrainAndTerminate = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MaxUnavailable - default",
			args: args{
				in: func() kopsv1alpha2.RollingUpdate {
					subject := kopsv1alpha2.RollingUpdate{}
					subject.MaxUnavailable = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MaxSurge - default",
			args: args{
				in: func() kopsv1alpha2.RollingUpdate {
					subject := kopsv1alpha2.RollingUpdate{}
					subject.MaxSurge = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenResourceRollingUpdateInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceRollingUpdate() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceRollingUpdate(t *testing.T) {
	_default := map[string]interface{}{
		"drain_and_terminate": nil,
		"max_unavailable":     nil,
		"max_surge":           nil,
	}
	type args struct {
		in kopsv1alpha2.RollingUpdate
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.RollingUpdate{},
			},
			want: _default,
		},
		{
			name: "DrainAndTerminate - default",
			args: args{
				in: func() kopsv1alpha2.RollingUpdate {
					subject := kopsv1alpha2.RollingUpdate{}
					subject.DrainAndTerminate = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MaxUnavailable - default",
			args: args{
				in: func() kopsv1alpha2.RollingUpdate {
					subject := kopsv1alpha2.RollingUpdate{}
					subject.MaxUnavailable = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MaxSurge - default",
			args: args{
				in: func() kopsv1alpha2.RollingUpdate {
					subject := kopsv1alpha2.RollingUpdate{}
					subject.MaxSurge = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceRollingUpdate(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceRollingUpdate() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
