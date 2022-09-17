package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceContainerdConfig(t *testing.T) {
	_default := kops.ContainerdConfig{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.ContainerdConfig
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"address":          nil,
					"config_override":  nil,
					"log_level":        nil,
					"packages":         nil,
					"registry_mirrors": func() []interface{} { return nil }(),
					"root":             nil,
					"skip_install":     false,
					"state":            nil,
					"version":          nil,
					"nvidia_gpu":       nil,
					"runc":             nil,
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandResourceContainerdConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceContainerdConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceContainerdConfigInto(t *testing.T) {
	_default := map[string]interface{}{
		"address":          nil,
		"config_override":  nil,
		"log_level":        nil,
		"packages":         nil,
		"registry_mirrors": func() []interface{} { return nil }(),
		"root":             nil,
		"skip_install":     false,
		"state":            nil,
		"version":          nil,
		"nvidia_gpu":       nil,
		"runc":             nil,
	}
	type args struct {
		in kops.ContainerdConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.ContainerdConfig{},
			},
			want: _default,
		},
		{
			name: "Address - default",
			args: args{
				in: func() kops.ContainerdConfig {
					subject := kops.ContainerdConfig{}
					subject.Address = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ConfigOverride - default",
			args: args{
				in: func() kops.ContainerdConfig {
					subject := kops.ContainerdConfig{}
					subject.ConfigOverride = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LogLevel - default",
			args: args{
				in: func() kops.ContainerdConfig {
					subject := kops.ContainerdConfig{}
					subject.LogLevel = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Packages - default",
			args: args{
				in: func() kops.ContainerdConfig {
					subject := kops.ContainerdConfig{}
					subject.Packages = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "RegistryMirrors - default",
			args: args{
				in: func() kops.ContainerdConfig {
					subject := kops.ContainerdConfig{}
					subject.RegistryMirrors = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Root - default",
			args: args{
				in: func() kops.ContainerdConfig {
					subject := kops.ContainerdConfig{}
					subject.Root = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SkipInstall - default",
			args: args{
				in: func() kops.ContainerdConfig {
					subject := kops.ContainerdConfig{}
					subject.SkipInstall = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "State - default",
			args: args{
				in: func() kops.ContainerdConfig {
					subject := kops.ContainerdConfig{}
					subject.State = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Version - default",
			args: args{
				in: func() kops.ContainerdConfig {
					subject := kops.ContainerdConfig{}
					subject.Version = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NvidiaGPU - default",
			args: args{
				in: func() kops.ContainerdConfig {
					subject := kops.ContainerdConfig{}
					subject.NvidiaGPU = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Runc - default",
			args: args{
				in: func() kops.ContainerdConfig {
					subject := kops.ContainerdConfig{}
					subject.Runc = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenResourceContainerdConfigInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceContainerdConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceContainerdConfig(t *testing.T) {
	_default := map[string]interface{}{
		"address":          nil,
		"config_override":  nil,
		"log_level":        nil,
		"packages":         nil,
		"registry_mirrors": func() []interface{} { return nil }(),
		"root":             nil,
		"skip_install":     false,
		"state":            nil,
		"version":          nil,
		"nvidia_gpu":       nil,
		"runc":             nil,
	}
	type args struct {
		in kops.ContainerdConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.ContainerdConfig{},
			},
			want: _default,
		},
		{
			name: "Address - default",
			args: args{
				in: func() kops.ContainerdConfig {
					subject := kops.ContainerdConfig{}
					subject.Address = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ConfigOverride - default",
			args: args{
				in: func() kops.ContainerdConfig {
					subject := kops.ContainerdConfig{}
					subject.ConfigOverride = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LogLevel - default",
			args: args{
				in: func() kops.ContainerdConfig {
					subject := kops.ContainerdConfig{}
					subject.LogLevel = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Packages - default",
			args: args{
				in: func() kops.ContainerdConfig {
					subject := kops.ContainerdConfig{}
					subject.Packages = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "RegistryMirrors - default",
			args: args{
				in: func() kops.ContainerdConfig {
					subject := kops.ContainerdConfig{}
					subject.RegistryMirrors = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Root - default",
			args: args{
				in: func() kops.ContainerdConfig {
					subject := kops.ContainerdConfig{}
					subject.Root = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SkipInstall - default",
			args: args{
				in: func() kops.ContainerdConfig {
					subject := kops.ContainerdConfig{}
					subject.SkipInstall = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "State - default",
			args: args{
				in: func() kops.ContainerdConfig {
					subject := kops.ContainerdConfig{}
					subject.State = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Version - default",
			args: args{
				in: func() kops.ContainerdConfig {
					subject := kops.ContainerdConfig{}
					subject.Version = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NvidiaGPU - default",
			args: args{
				in: func() kops.ContainerdConfig {
					subject := kops.ContainerdConfig{}
					subject.NvidiaGPU = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Runc - default",
			args: args{
				in: func() kops.ContainerdConfig {
					subject := kops.ContainerdConfig{}
					subject.Runc = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceContainerdConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceContainerdConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
