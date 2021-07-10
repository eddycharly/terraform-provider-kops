package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceSnapshotControllerConfig(t *testing.T) {
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
			if got := ExpandResourceSnapshotControllerConfig(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandResourceSnapshotControllerConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenResourceSnapshotControllerConfigInto(t *testing.T) {
	type args struct {
		in  kops.SnapshotControllerConfig
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
			FlattenResourceSnapshotControllerConfigInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenResourceSnapshotControllerConfig(t *testing.T) {
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
			want: map[string]interface{}{
				"enabled":               nil,
				"install_default_class": false,
			},
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
			want: map[string]interface{}{
				"enabled":               nil,
				"install_default_class": false,
			},
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
			want: map[string]interface{}{
				"enabled":               nil,
				"install_default_class": false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenResourceSnapshotControllerConfig(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenResourceSnapshotControllerConfig() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
