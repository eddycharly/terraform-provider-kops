package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceRollingUpdate(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.RollingUpdate
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandResourceRollingUpdate(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandResourceRollingUpdate() = %v, want %v", got, tt.want)
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
			got := FlattenResourceRollingUpdate(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceRollingUpdate() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
