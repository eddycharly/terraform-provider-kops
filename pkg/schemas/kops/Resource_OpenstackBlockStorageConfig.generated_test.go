package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceOpenstackBlockStorageConfig(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.OpenstackBlockStorageConfig
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandResourceOpenstackBlockStorageConfig(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandResourceOpenstackBlockStorageConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenResourceOpenstackBlockStorageConfigInto(t *testing.T) {
	type args struct {
		in  kops.OpenstackBlockStorageConfig
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
			FlattenResourceOpenstackBlockStorageConfigInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenResourceOpenstackBlockStorageConfig(t *testing.T) {
	type args struct {
		in kops.OpenstackBlockStorageConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.OpenstackBlockStorageConfig{},
			},
			want: map[string]interface{}{
				"version":              nil,
				"ignore_az":            nil,
				"override_az":          nil,
				"create_storage_class": nil,
				"csi_plugin_image":     "",
				"csi_topology_support": nil,
			},
		},
		{
			name: "Version - default",
			args: args{
				in: func() kops.OpenstackBlockStorageConfig {
					subject := kops.OpenstackBlockStorageConfig{}
					subject.Version = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"version":              nil,
				"ignore_az":            nil,
				"override_az":          nil,
				"create_storage_class": nil,
				"csi_plugin_image":     "",
				"csi_topology_support": nil,
			},
		},
		{
			name: "IgnoreAZ - default",
			args: args{
				in: func() kops.OpenstackBlockStorageConfig {
					subject := kops.OpenstackBlockStorageConfig{}
					subject.IgnoreAZ = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"version":              nil,
				"ignore_az":            nil,
				"override_az":          nil,
				"create_storage_class": nil,
				"csi_plugin_image":     "",
				"csi_topology_support": nil,
			},
		},
		{
			name: "OverrideAZ - default",
			args: args{
				in: func() kops.OpenstackBlockStorageConfig {
					subject := kops.OpenstackBlockStorageConfig{}
					subject.OverrideAZ = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"version":              nil,
				"ignore_az":            nil,
				"override_az":          nil,
				"create_storage_class": nil,
				"csi_plugin_image":     "",
				"csi_topology_support": nil,
			},
		},
		{
			name: "CreateStorageClass - default",
			args: args{
				in: func() kops.OpenstackBlockStorageConfig {
					subject := kops.OpenstackBlockStorageConfig{}
					subject.CreateStorageClass = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"version":              nil,
				"ignore_az":            nil,
				"override_az":          nil,
				"create_storage_class": nil,
				"csi_plugin_image":     "",
				"csi_topology_support": nil,
			},
		},
		{
			name: "CsiPluginImage - default",
			args: args{
				in: func() kops.OpenstackBlockStorageConfig {
					subject := kops.OpenstackBlockStorageConfig{}
					subject.CSIPluginImage = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
				"version":              nil,
				"ignore_az":            nil,
				"override_az":          nil,
				"create_storage_class": nil,
				"csi_plugin_image":     "",
				"csi_topology_support": nil,
			},
		},
		{
			name: "CsiTopologySupport - default",
			args: args{
				in: func() kops.OpenstackBlockStorageConfig {
					subject := kops.OpenstackBlockStorageConfig{}
					subject.CSITopologySupport = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"version":              nil,
				"ignore_az":            nil,
				"override_az":          nil,
				"create_storage_class": nil,
				"csi_plugin_image":     "",
				"csi_topology_support": nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenResourceOpenstackBlockStorageConfig(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenResourceOpenstackBlockStorageConfig() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
