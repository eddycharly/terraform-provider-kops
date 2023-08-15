package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

func TestExpandResourceOpenstackBlockStorageConfig(t *testing.T) {
	_default := kopsv1alpha2.OpenstackBlockStorageConfig{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kopsv1alpha2.OpenstackBlockStorageConfig
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"version":                     nil,
					"ignore_az":                   nil,
					"override_az":                 nil,
					"ignore_volume_micro_version": nil,
					"create_storage_class":        nil,
					"csi_plugin_image":            "",
					"csi_topology_support":        nil,
					"cluster_name":                "",
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandResourceOpenstackBlockStorageConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceOpenstackBlockStorageConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceOpenstackBlockStorageConfigInto(t *testing.T) {
	_default := map[string]interface{}{
		"version":                     nil,
		"ignore_az":                   nil,
		"override_az":                 nil,
		"ignore_volume_micro_version": nil,
		"create_storage_class":        nil,
		"csi_plugin_image":            "",
		"csi_topology_support":        nil,
		"cluster_name":                "",
	}
	type args struct {
		in kopsv1alpha2.OpenstackBlockStorageConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.OpenstackBlockStorageConfig{},
			},
			want: _default,
		},
		{
			name: "Version - default",
			args: args{
				in: func() kopsv1alpha2.OpenstackBlockStorageConfig {
					subject := kopsv1alpha2.OpenstackBlockStorageConfig{}
					subject.Version = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "IgnoreAZ - default",
			args: args{
				in: func() kopsv1alpha2.OpenstackBlockStorageConfig {
					subject := kopsv1alpha2.OpenstackBlockStorageConfig{}
					subject.IgnoreAZ = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "OverrideAZ - default",
			args: args{
				in: func() kopsv1alpha2.OpenstackBlockStorageConfig {
					subject := kopsv1alpha2.OpenstackBlockStorageConfig{}
					subject.OverrideAZ = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "IgnoreVolumeMicroVersion - default",
			args: args{
				in: func() kopsv1alpha2.OpenstackBlockStorageConfig {
					subject := kopsv1alpha2.OpenstackBlockStorageConfig{}
					subject.IgnoreVolumeMicroVersion = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CreateStorageClass - default",
			args: args{
				in: func() kopsv1alpha2.OpenstackBlockStorageConfig {
					subject := kopsv1alpha2.OpenstackBlockStorageConfig{}
					subject.CreateStorageClass = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CsiPluginImage - default",
			args: args{
				in: func() kopsv1alpha2.OpenstackBlockStorageConfig {
					subject := kopsv1alpha2.OpenstackBlockStorageConfig{}
					subject.CSIPluginImage = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CsiTopologySupport - default",
			args: args{
				in: func() kopsv1alpha2.OpenstackBlockStorageConfig {
					subject := kopsv1alpha2.OpenstackBlockStorageConfig{}
					subject.CSITopologySupport = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ClusterName - default",
			args: args{
				in: func() kopsv1alpha2.OpenstackBlockStorageConfig {
					subject := kopsv1alpha2.OpenstackBlockStorageConfig{}
					subject.ClusterName = ""
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenResourceOpenstackBlockStorageConfigInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceOpenstackBlockStorageConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceOpenstackBlockStorageConfig(t *testing.T) {
	_default := map[string]interface{}{
		"version":                     nil,
		"ignore_az":                   nil,
		"override_az":                 nil,
		"ignore_volume_micro_version": nil,
		"create_storage_class":        nil,
		"csi_plugin_image":            "",
		"csi_topology_support":        nil,
		"cluster_name":                "",
	}
	type args struct {
		in kopsv1alpha2.OpenstackBlockStorageConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.OpenstackBlockStorageConfig{},
			},
			want: _default,
		},
		{
			name: "Version - default",
			args: args{
				in: func() kopsv1alpha2.OpenstackBlockStorageConfig {
					subject := kopsv1alpha2.OpenstackBlockStorageConfig{}
					subject.Version = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "IgnoreAZ - default",
			args: args{
				in: func() kopsv1alpha2.OpenstackBlockStorageConfig {
					subject := kopsv1alpha2.OpenstackBlockStorageConfig{}
					subject.IgnoreAZ = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "OverrideAZ - default",
			args: args{
				in: func() kopsv1alpha2.OpenstackBlockStorageConfig {
					subject := kopsv1alpha2.OpenstackBlockStorageConfig{}
					subject.OverrideAZ = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "IgnoreVolumeMicroVersion - default",
			args: args{
				in: func() kopsv1alpha2.OpenstackBlockStorageConfig {
					subject := kopsv1alpha2.OpenstackBlockStorageConfig{}
					subject.IgnoreVolumeMicroVersion = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CreateStorageClass - default",
			args: args{
				in: func() kopsv1alpha2.OpenstackBlockStorageConfig {
					subject := kopsv1alpha2.OpenstackBlockStorageConfig{}
					subject.CreateStorageClass = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CsiPluginImage - default",
			args: args{
				in: func() kopsv1alpha2.OpenstackBlockStorageConfig {
					subject := kopsv1alpha2.OpenstackBlockStorageConfig{}
					subject.CSIPluginImage = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CsiTopologySupport - default",
			args: args{
				in: func() kopsv1alpha2.OpenstackBlockStorageConfig {
					subject := kopsv1alpha2.OpenstackBlockStorageConfig{}
					subject.CSITopologySupport = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ClusterName - default",
			args: args{
				in: func() kopsv1alpha2.OpenstackBlockStorageConfig {
					subject := kopsv1alpha2.OpenstackBlockStorageConfig{}
					subject.ClusterName = ""
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceOpenstackBlockStorageConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceOpenstackBlockStorageConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
