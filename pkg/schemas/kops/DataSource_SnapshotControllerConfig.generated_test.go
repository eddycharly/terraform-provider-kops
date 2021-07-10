package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceSnapshotControllerConfig(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.SnapshotControllerConfig
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandDataSourceSnapshotControllerConfig(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandDataSourceSnapshotControllerConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenDataSourceSnapshotControllerConfigInto(t *testing.T) {
	_default := map[string]interface{}{
		"enabled":               nil,
		"install_default_class": false,
	}
	type args struct {
		in kops.SnapshotControllerConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.SnapshotControllerConfig{},
			},
			want: _default,
		},
		{
			name: "Enabled - default",
			args: args{
				in: func() kops.SnapshotControllerConfig {
					subject := kops.SnapshotControllerConfig{}
					subject.Enabled = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "InstallDefaultClass - default",
			args: args{
				in: func() kops.SnapshotControllerConfig {
					subject := kops.SnapshotControllerConfig{}
					subject.InstallDefaultClass = false
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenDataSourceSnapshotControllerConfigInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceSnapshotControllerConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceSnapshotControllerConfig(t *testing.T) {
	_default := map[string]interface{}{
		"enabled":               nil,
		"install_default_class": false,
	}
	type args struct {
		in kops.SnapshotControllerConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.SnapshotControllerConfig{},
			},
			want: _default,
		},
		{
			name: "Enabled - default",
			args: args{
				in: func() kops.SnapshotControllerConfig {
					subject := kops.SnapshotControllerConfig{}
					subject.Enabled = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "InstallDefaultClass - default",
			args: args{
				in: func() kops.SnapshotControllerConfig {
					subject := kops.SnapshotControllerConfig{}
					subject.InstallDefaultClass = false
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceSnapshotControllerConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceSnapshotControllerConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
