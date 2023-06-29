package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

func TestExpandResourceContainerdConfig(t *testing.T) {
	_default := kopsv1alpha2.ContainerdConfig{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kopsv1alpha2.ContainerdConfig
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
		in kopsv1alpha2.ContainerdConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.ContainerdConfig{},
			},
			want: _default,
		},
		{
			name: "Address - default",
			args: args{
				in: func() kopsv1alpha2.ContainerdConfig {
					subject := kopsv1alpha2.ContainerdConfig{}
					subject.Address = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ConfigOverride - default",
			args: args{
				in: func() kopsv1alpha2.ContainerdConfig {
					subject := kopsv1alpha2.ContainerdConfig{}
					subject.ConfigOverride = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LogLevel - default",
			args: args{
				in: func() kopsv1alpha2.ContainerdConfig {
					subject := kopsv1alpha2.ContainerdConfig{}
					subject.LogLevel = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Packages - default",
			args: args{
				in: func() kopsv1alpha2.ContainerdConfig {
					subject := kopsv1alpha2.ContainerdConfig{}
					subject.Packages = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "RegistryMirrors - default",
			args: args{
				in: func() kopsv1alpha2.ContainerdConfig {
					subject := kopsv1alpha2.ContainerdConfig{}
					subject.RegistryMirrors = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Root - default",
			args: args{
				in: func() kopsv1alpha2.ContainerdConfig {
					subject := kopsv1alpha2.ContainerdConfig{}
					subject.Root = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SkipInstall - default",
			args: args{
				in: func() kopsv1alpha2.ContainerdConfig {
					subject := kopsv1alpha2.ContainerdConfig{}
					subject.SkipInstall = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "State - default",
			args: args{
				in: func() kopsv1alpha2.ContainerdConfig {
					subject := kopsv1alpha2.ContainerdConfig{}
					subject.State = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Version - default",
			args: args{
				in: func() kopsv1alpha2.ContainerdConfig {
					subject := kopsv1alpha2.ContainerdConfig{}
					subject.Version = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NvidiaGPU - default",
			args: args{
				in: func() kopsv1alpha2.ContainerdConfig {
					subject := kopsv1alpha2.ContainerdConfig{}
					subject.NvidiaGPU = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Runc - default",
			args: args{
				in: func() kopsv1alpha2.ContainerdConfig {
					subject := kopsv1alpha2.ContainerdConfig{}
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
		in kopsv1alpha2.ContainerdConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.ContainerdConfig{},
			},
			want: _default,
		},
		{
			name: "Address - default",
			args: args{
				in: func() kopsv1alpha2.ContainerdConfig {
					subject := kopsv1alpha2.ContainerdConfig{}
					subject.Address = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ConfigOverride - default",
			args: args{
				in: func() kopsv1alpha2.ContainerdConfig {
					subject := kopsv1alpha2.ContainerdConfig{}
					subject.ConfigOverride = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LogLevel - default",
			args: args{
				in: func() kopsv1alpha2.ContainerdConfig {
					subject := kopsv1alpha2.ContainerdConfig{}
					subject.LogLevel = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Packages - default",
			args: args{
				in: func() kopsv1alpha2.ContainerdConfig {
					subject := kopsv1alpha2.ContainerdConfig{}
					subject.Packages = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "RegistryMirrors - default",
			args: args{
				in: func() kopsv1alpha2.ContainerdConfig {
					subject := kopsv1alpha2.ContainerdConfig{}
					subject.RegistryMirrors = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Root - default",
			args: args{
				in: func() kopsv1alpha2.ContainerdConfig {
					subject := kopsv1alpha2.ContainerdConfig{}
					subject.Root = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SkipInstall - default",
			args: args{
				in: func() kopsv1alpha2.ContainerdConfig {
					subject := kopsv1alpha2.ContainerdConfig{}
					subject.SkipInstall = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "State - default",
			args: args{
				in: func() kopsv1alpha2.ContainerdConfig {
					subject := kopsv1alpha2.ContainerdConfig{}
					subject.State = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Version - default",
			args: args{
				in: func() kopsv1alpha2.ContainerdConfig {
					subject := kopsv1alpha2.ContainerdConfig{}
					subject.Version = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NvidiaGPU - default",
			args: args{
				in: func() kopsv1alpha2.ContainerdConfig {
					subject := kopsv1alpha2.ContainerdConfig{}
					subject.NvidiaGPU = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Runc - default",
			args: args{
				in: func() kopsv1alpha2.ContainerdConfig {
					subject := kopsv1alpha2.ContainerdConfig{}
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
