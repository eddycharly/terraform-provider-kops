package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceRollingUpdate(t *testing.T) {
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
			if got := ExpandDataSourceRollingUpdate(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandDataSourceRollingUpdate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenDataSourceRollingUpdateInto(t *testing.T) {
	type args struct {
		in  kops.RollingUpdate
		out map[string]interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			FlattenDataSourceRollingUpdateInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenDataSourceRollingUpdate(t *testing.T) {
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
			want: map[string]interface{}{
				"drain_and_terminate": nil,
				"max_unavailable":     nil,
				"max_surge":           nil,
			},
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
			want: map[string]interface{}{
				"drain_and_terminate": nil,
				"max_unavailable":     nil,
				"max_surge":           nil,
			},
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
			want: map[string]interface{}{
				"drain_and_terminate": nil,
				"max_unavailable":     nil,
				"max_surge":           nil,
			},
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
			want: map[string]interface{}{
				"drain_and_terminate": nil,
				"max_unavailable":     nil,
				"max_surge":           nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenDataSourceRollingUpdate(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenDataSourceRollingUpdate() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
