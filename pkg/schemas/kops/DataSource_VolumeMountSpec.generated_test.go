package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceVolumeMountSpec(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.VolumeMountSpec
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandDataSourceVolumeMountSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandDataSourceVolumeMountSpec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenDataSourceVolumeMountSpecInto(t *testing.T) {
	type args struct {
		in  kops.VolumeMountSpec
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
			FlattenDataSourceVolumeMountSpecInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenDataSourceVolumeMountSpec(t *testing.T) {
	type args struct {
		in kops.VolumeMountSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.VolumeMountSpec{},
			},
			want: map[string]interface{}{
				"device":         "",
				"filesystem":     "",
				"format_options": func() []interface{} { return nil }(),
				"mount_options":  func() []interface{} { return nil }(),
				"path":           "",
			},
		},
		{
			name: "Device - default",
			args: args{
				in: func() kops.VolumeMountSpec {
					subject := kops.VolumeMountSpec{}
					subject.Device = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
				"device":         "",
				"filesystem":     "",
				"format_options": func() []interface{} { return nil }(),
				"mount_options":  func() []interface{} { return nil }(),
				"path":           "",
			},
		},
		{
			name: "Filesystem - default",
			args: args{
				in: func() kops.VolumeMountSpec {
					subject := kops.VolumeMountSpec{}
					subject.Filesystem = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
				"device":         "",
				"filesystem":     "",
				"format_options": func() []interface{} { return nil }(),
				"mount_options":  func() []interface{} { return nil }(),
				"path":           "",
			},
		},
		{
			name: "FormatOptions - default",
			args: args{
				in: func() kops.VolumeMountSpec {
					subject := kops.VolumeMountSpec{}
					subject.FormatOptions = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"device":         "",
				"filesystem":     "",
				"format_options": func() []interface{} { return nil }(),
				"mount_options":  func() []interface{} { return nil }(),
				"path":           "",
			},
		},
		{
			name: "MountOptions - default",
			args: args{
				in: func() kops.VolumeMountSpec {
					subject := kops.VolumeMountSpec{}
					subject.MountOptions = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"device":         "",
				"filesystem":     "",
				"format_options": func() []interface{} { return nil }(),
				"mount_options":  func() []interface{} { return nil }(),
				"path":           "",
			},
		},
		{
			name: "Path - default",
			args: args{
				in: func() kops.VolumeMountSpec {
					subject := kops.VolumeMountSpec{}
					subject.Path = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
				"device":         "",
				"filesystem":     "",
				"format_options": func() []interface{} { return nil }(),
				"mount_options":  func() []interface{} { return nil }(),
				"path":           "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenDataSourceVolumeMountSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenDataSourceVolumeMountSpec() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
