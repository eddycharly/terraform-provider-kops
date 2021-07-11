package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceRollingUpdate(t *testing.T) {
	_default := kops.RollingUpdate{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.RollingUpdate
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
			got := ExpandDataSourceRollingUpdate(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceRollingUpdate() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceRollingUpdateInto(t *testing.T) {
	_default := map[string]interface{}{
		"drain_and_terminate": nil,
		"max_unavailable":     nil,
		"max_surge":           nil,
	}
	type args struct {
		in kops.RollingUpdate
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.RollingUpdate{},
			},
			want: _default,
		},
		{
			name: "DrainAndTerminate - default",
			args: args{
				in: func() kops.RollingUpdate {
					subject := kops.RollingUpdate{}
					subject.DrainAndTerminate = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MaxUnavailable - default",
			args: args{
				in: func() kops.RollingUpdate {
					subject := kops.RollingUpdate{}
					subject.MaxUnavailable = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MaxSurge - default",
			args: args{
				in: func() kops.RollingUpdate {
					subject := kops.RollingUpdate{}
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
			FlattenDataSourceRollingUpdateInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceRollingUpdate() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceRollingUpdate(t *testing.T) {
	_default := map[string]interface{}{
		"drain_and_terminate": nil,
		"max_unavailable":     nil,
		"max_surge":           nil,
	}
	type args struct {
		in kops.RollingUpdate
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.RollingUpdate{},
			},
			want: _default,
		},
		{
			name: "DrainAndTerminate - default",
			args: args{
				in: func() kops.RollingUpdate {
					subject := kops.RollingUpdate{}
					subject.DrainAndTerminate = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MaxUnavailable - default",
			args: args{
				in: func() kops.RollingUpdate {
					subject := kops.RollingUpdate{}
					subject.MaxUnavailable = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MaxSurge - default",
			args: args{
				in: func() kops.RollingUpdate {
					subject := kops.RollingUpdate{}
					subject.MaxSurge = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceRollingUpdate(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceRollingUpdate() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
