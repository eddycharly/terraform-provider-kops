package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceContainerdConfig(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.ContainerdConfig
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandDataSourceContainerdConfig(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandDataSourceContainerdConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenDataSourceContainerdConfigInto(t *testing.T) {
	type args struct {
		in  kops.ContainerdConfig
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
			FlattenDataSourceContainerdConfigInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenDataSourceContainerdConfig(t *testing.T) {
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
			want: map[string]interface{}{
				"address":          nil,
				"config_override":  nil,
				"log_level":        nil,
				"packages":         nil,
				"registry_mirrors": func() map[string]interface{} { return nil }(),
				"root":             nil,
				"skip_install":     false,
				"state":            nil,
				"version":          nil,
			},
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
			want: map[string]interface{}{
				"address":          nil,
				"config_override":  nil,
				"log_level":        nil,
				"packages":         nil,
				"registry_mirrors": func() map[string]interface{} { return nil }(),
				"root":             nil,
				"skip_install":     false,
				"state":            nil,
				"version":          nil,
			},
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
			want: map[string]interface{}{
				"address":          nil,
				"config_override":  nil,
				"log_level":        nil,
				"packages":         nil,
				"registry_mirrors": func() map[string]interface{} { return nil }(),
				"root":             nil,
				"skip_install":     false,
				"state":            nil,
				"version":          nil,
			},
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
			want: map[string]interface{}{
				"address":          nil,
				"config_override":  nil,
				"log_level":        nil,
				"packages":         nil,
				"registry_mirrors": func() map[string]interface{} { return nil }(),
				"root":             nil,
				"skip_install":     false,
				"state":            nil,
				"version":          nil,
			},
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
			want: map[string]interface{}{
				"address":          nil,
				"config_override":  nil,
				"log_level":        nil,
				"packages":         nil,
				"registry_mirrors": func() map[string]interface{} { return nil }(),
				"root":             nil,
				"skip_install":     false,
				"state":            nil,
				"version":          nil,
			},
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
			want: map[string]interface{}{
				"address":          nil,
				"config_override":  nil,
				"log_level":        nil,
				"packages":         nil,
				"registry_mirrors": func() map[string]interface{} { return nil }(),
				"root":             nil,
				"skip_install":     false,
				"state":            nil,
				"version":          nil,
			},
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
			want: map[string]interface{}{
				"address":          nil,
				"config_override":  nil,
				"log_level":        nil,
				"packages":         nil,
				"registry_mirrors": func() map[string]interface{} { return nil }(),
				"root":             nil,
				"skip_install":     false,
				"state":            nil,
				"version":          nil,
			},
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
			want: map[string]interface{}{
				"address":          nil,
				"config_override":  nil,
				"log_level":        nil,
				"packages":         nil,
				"registry_mirrors": func() map[string]interface{} { return nil }(),
				"root":             nil,
				"skip_install":     false,
				"state":            nil,
				"version":          nil,
			},
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
			want: map[string]interface{}{
				"address":          nil,
				"config_override":  nil,
				"log_level":        nil,
				"packages":         nil,
				"registry_mirrors": func() map[string]interface{} { return nil }(),
				"root":             nil,
				"skip_install":     false,
				"state":            nil,
				"version":          nil,
			},
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
			want: map[string]interface{}{
				"address":          nil,
				"config_override":  nil,
				"log_level":        nil,
				"packages":         nil,
				"registry_mirrors": func() map[string]interface{} { return nil }(),
				"root":             nil,
				"skip_install":     false,
				"state":            nil,
				"version":          nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenDataSourceContainerdConfig(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenDataSourceContainerdConfig() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
