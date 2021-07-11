package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceVolumeMountSpec(t *testing.T) {
	_default := kops.VolumeMountSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.VolumeMountSpec
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"device":         "",
					"filesystem":     "",
					"format_options": func() []interface{} { return nil }(),
					"mount_options":  func() []interface{} { return nil }(),
					"path":           "",
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandResourceVolumeMountSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceVolumeMountSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceVolumeMountSpecInto(t *testing.T) {
	_default := map[string]interface{}{
		"device":         "",
		"filesystem":     "",
		"format_options": func() []interface{} { return nil }(),
		"mount_options":  func() []interface{} { return nil }(),
		"path":           "",
	}
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenResourceVolumeMountSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceVolumeMountSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceVolumeMountSpec(t *testing.T) {
	_default := map[string]interface{}{
		"device":         "",
		"filesystem":     "",
		"format_options": func() []interface{} { return nil }(),
		"mount_options":  func() []interface{} { return nil }(),
		"path":           "",
	}
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceVolumeMountSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceVolumeMountSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
