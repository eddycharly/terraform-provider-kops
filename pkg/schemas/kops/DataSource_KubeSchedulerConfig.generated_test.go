package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceKubeSchedulerConfig(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.KubeSchedulerConfig
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandDataSourceKubeSchedulerConfig(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandDataSourceKubeSchedulerConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenDataSourceKubeSchedulerConfigInto(t *testing.T) {
	type args struct {
		in  kops.KubeSchedulerConfig
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
			FlattenDataSourceKubeSchedulerConfigInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenDataSourceKubeSchedulerConfig(t *testing.T) {
	type args struct {
		in kops.KubeSchedulerConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.KubeSchedulerConfig{},
			},
			want: map[string]interface{}{
				"master":                           "",
				"log_level":                        0,
				"image":                            "",
				"leader_election":                  nil,
				"use_policy_config_map":            nil,
				"feature_gates":                    func() map[string]interface{} { return nil }(),
				"max_persistent_volumes":           nil,
				"qps":                              nil,
				"burst":                            0,
				"authentication_kubeconfig":        "",
				"authorization_kubeconfig":         "",
				"authorization_always_allow_paths": func() []interface{} { return nil }(),
				"enable_profiling":                 nil,
			},
		},
		{
			name: "Master - default",
			args: args{
				in: func() kops.KubeSchedulerConfig {
					subject := kops.KubeSchedulerConfig{}
					subject.Master = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
				"master":                           "",
				"log_level":                        0,
				"image":                            "",
				"leader_election":                  nil,
				"use_policy_config_map":            nil,
				"feature_gates":                    func() map[string]interface{} { return nil }(),
				"max_persistent_volumes":           nil,
				"qps":                              nil,
				"burst":                            0,
				"authentication_kubeconfig":        "",
				"authorization_kubeconfig":         "",
				"authorization_always_allow_paths": func() []interface{} { return nil }(),
				"enable_profiling":                 nil,
			},
		},
		{
			name: "LogLevel - default",
			args: args{
				in: func() kops.KubeSchedulerConfig {
					subject := kops.KubeSchedulerConfig{}
					subject.LogLevel = 0
					return subject
				}(),
			},
			want: map[string]interface{}{
				"master":                           "",
				"log_level":                        0,
				"image":                            "",
				"leader_election":                  nil,
				"use_policy_config_map":            nil,
				"feature_gates":                    func() map[string]interface{} { return nil }(),
				"max_persistent_volumes":           nil,
				"qps":                              nil,
				"burst":                            0,
				"authentication_kubeconfig":        "",
				"authorization_kubeconfig":         "",
				"authorization_always_allow_paths": func() []interface{} { return nil }(),
				"enable_profiling":                 nil,
			},
		},
		{
			name: "Image - default",
			args: args{
				in: func() kops.KubeSchedulerConfig {
					subject := kops.KubeSchedulerConfig{}
					subject.Image = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
				"master":                           "",
				"log_level":                        0,
				"image":                            "",
				"leader_election":                  nil,
				"use_policy_config_map":            nil,
				"feature_gates":                    func() map[string]interface{} { return nil }(),
				"max_persistent_volumes":           nil,
				"qps":                              nil,
				"burst":                            0,
				"authentication_kubeconfig":        "",
				"authorization_kubeconfig":         "",
				"authorization_always_allow_paths": func() []interface{} { return nil }(),
				"enable_profiling":                 nil,
			},
		},
		{
			name: "LeaderElection - default",
			args: args{
				in: func() kops.KubeSchedulerConfig {
					subject := kops.KubeSchedulerConfig{}
					subject.LeaderElection = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"master":                           "",
				"log_level":                        0,
				"image":                            "",
				"leader_election":                  nil,
				"use_policy_config_map":            nil,
				"feature_gates":                    func() map[string]interface{} { return nil }(),
				"max_persistent_volumes":           nil,
				"qps":                              nil,
				"burst":                            0,
				"authentication_kubeconfig":        "",
				"authorization_kubeconfig":         "",
				"authorization_always_allow_paths": func() []interface{} { return nil }(),
				"enable_profiling":                 nil,
			},
		},
		{
			name: "UsePolicyConfigMap - default",
			args: args{
				in: func() kops.KubeSchedulerConfig {
					subject := kops.KubeSchedulerConfig{}
					subject.UsePolicyConfigMap = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"master":                           "",
				"log_level":                        0,
				"image":                            "",
				"leader_election":                  nil,
				"use_policy_config_map":            nil,
				"feature_gates":                    func() map[string]interface{} { return nil }(),
				"max_persistent_volumes":           nil,
				"qps":                              nil,
				"burst":                            0,
				"authentication_kubeconfig":        "",
				"authorization_kubeconfig":         "",
				"authorization_always_allow_paths": func() []interface{} { return nil }(),
				"enable_profiling":                 nil,
			},
		},
		{
			name: "FeatureGates - default",
			args: args{
				in: func() kops.KubeSchedulerConfig {
					subject := kops.KubeSchedulerConfig{}
					subject.FeatureGates = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"master":                           "",
				"log_level":                        0,
				"image":                            "",
				"leader_election":                  nil,
				"use_policy_config_map":            nil,
				"feature_gates":                    func() map[string]interface{} { return nil }(),
				"max_persistent_volumes":           nil,
				"qps":                              nil,
				"burst":                            0,
				"authentication_kubeconfig":        "",
				"authorization_kubeconfig":         "",
				"authorization_always_allow_paths": func() []interface{} { return nil }(),
				"enable_profiling":                 nil,
			},
		},
		{
			name: "MaxPersistentVolumes - default",
			args: args{
				in: func() kops.KubeSchedulerConfig {
					subject := kops.KubeSchedulerConfig{}
					subject.MaxPersistentVolumes = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"master":                           "",
				"log_level":                        0,
				"image":                            "",
				"leader_election":                  nil,
				"use_policy_config_map":            nil,
				"feature_gates":                    func() map[string]interface{} { return nil }(),
				"max_persistent_volumes":           nil,
				"qps":                              nil,
				"burst":                            0,
				"authentication_kubeconfig":        "",
				"authorization_kubeconfig":         "",
				"authorization_always_allow_paths": func() []interface{} { return nil }(),
				"enable_profiling":                 nil,
			},
		},
		{
			name: "Qps - default",
			args: args{
				in: func() kops.KubeSchedulerConfig {
					subject := kops.KubeSchedulerConfig{}
					subject.Qps = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"master":                           "",
				"log_level":                        0,
				"image":                            "",
				"leader_election":                  nil,
				"use_policy_config_map":            nil,
				"feature_gates":                    func() map[string]interface{} { return nil }(),
				"max_persistent_volumes":           nil,
				"qps":                              nil,
				"burst":                            0,
				"authentication_kubeconfig":        "",
				"authorization_kubeconfig":         "",
				"authorization_always_allow_paths": func() []interface{} { return nil }(),
				"enable_profiling":                 nil,
			},
		},
		{
			name: "Burst - default",
			args: args{
				in: func() kops.KubeSchedulerConfig {
					subject := kops.KubeSchedulerConfig{}
					subject.Burst = 0
					return subject
				}(),
			},
			want: map[string]interface{}{
				"master":                           "",
				"log_level":                        0,
				"image":                            "",
				"leader_election":                  nil,
				"use_policy_config_map":            nil,
				"feature_gates":                    func() map[string]interface{} { return nil }(),
				"max_persistent_volumes":           nil,
				"qps":                              nil,
				"burst":                            0,
				"authentication_kubeconfig":        "",
				"authorization_kubeconfig":         "",
				"authorization_always_allow_paths": func() []interface{} { return nil }(),
				"enable_profiling":                 nil,
			},
		},
		{
			name: "AuthenticationKubeconfig - default",
			args: args{
				in: func() kops.KubeSchedulerConfig {
					subject := kops.KubeSchedulerConfig{}
					subject.AuthenticationKubeconfig = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
				"master":                           "",
				"log_level":                        0,
				"image":                            "",
				"leader_election":                  nil,
				"use_policy_config_map":            nil,
				"feature_gates":                    func() map[string]interface{} { return nil }(),
				"max_persistent_volumes":           nil,
				"qps":                              nil,
				"burst":                            0,
				"authentication_kubeconfig":        "",
				"authorization_kubeconfig":         "",
				"authorization_always_allow_paths": func() []interface{} { return nil }(),
				"enable_profiling":                 nil,
			},
		},
		{
			name: "AuthorizationKubeconfig - default",
			args: args{
				in: func() kops.KubeSchedulerConfig {
					subject := kops.KubeSchedulerConfig{}
					subject.AuthorizationKubeconfig = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
				"master":                           "",
				"log_level":                        0,
				"image":                            "",
				"leader_election":                  nil,
				"use_policy_config_map":            nil,
				"feature_gates":                    func() map[string]interface{} { return nil }(),
				"max_persistent_volumes":           nil,
				"qps":                              nil,
				"burst":                            0,
				"authentication_kubeconfig":        "",
				"authorization_kubeconfig":         "",
				"authorization_always_allow_paths": func() []interface{} { return nil }(),
				"enable_profiling":                 nil,
			},
		},
		{
			name: "AuthorizationAlwaysAllowPaths - default",
			args: args{
				in: func() kops.KubeSchedulerConfig {
					subject := kops.KubeSchedulerConfig{}
					subject.AuthorizationAlwaysAllowPaths = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"master":                           "",
				"log_level":                        0,
				"image":                            "",
				"leader_election":                  nil,
				"use_policy_config_map":            nil,
				"feature_gates":                    func() map[string]interface{} { return nil }(),
				"max_persistent_volumes":           nil,
				"qps":                              nil,
				"burst":                            0,
				"authentication_kubeconfig":        "",
				"authorization_kubeconfig":         "",
				"authorization_always_allow_paths": func() []interface{} { return nil }(),
				"enable_profiling":                 nil,
			},
		},
		{
			name: "EnableProfiling - default",
			args: args{
				in: func() kops.KubeSchedulerConfig {
					subject := kops.KubeSchedulerConfig{}
					subject.EnableProfiling = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"master":                           "",
				"log_level":                        0,
				"image":                            "",
				"leader_election":                  nil,
				"use_policy_config_map":            nil,
				"feature_gates":                    func() map[string]interface{} { return nil }(),
				"max_persistent_volumes":           nil,
				"qps":                              nil,
				"burst":                            0,
				"authentication_kubeconfig":        "",
				"authorization_kubeconfig":         "",
				"authorization_always_allow_paths": func() []interface{} { return nil }(),
				"enable_profiling":                 nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenDataSourceKubeSchedulerConfig(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenDataSourceKubeSchedulerConfig() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
