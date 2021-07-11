package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceSnapshotControllerConfig(t *testing.T) {
	_default := kops.SnapshotControllerConfig{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.SnapshotControllerConfig
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"enabled":               nil,
					"install_default_class": false,
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandResourceSnapshotControllerConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceSnapshotControllerConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceSnapshotControllerConfigInto(t *testing.T) {
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
			FlattenResourceSnapshotControllerConfigInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceSnapshotControllerConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceSnapshotControllerConfig(t *testing.T) {
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
			got := FlattenResourceSnapshotControllerConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceSnapshotControllerConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
