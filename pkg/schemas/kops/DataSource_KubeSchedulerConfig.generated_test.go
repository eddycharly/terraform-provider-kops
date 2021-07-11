package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceKubeSchedulerConfig(t *testing.T) {
	_default := kops.KubeSchedulerConfig{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.KubeSchedulerConfig
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
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
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandDataSourceKubeSchedulerConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceKubeSchedulerConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceKubeSchedulerConfigInto(t *testing.T) {
	_default := map[string]interface{}{
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
	}
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenDataSourceKubeSchedulerConfigInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceKubeSchedulerConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceKubeSchedulerConfig(t *testing.T) {
	_default := map[string]interface{}{
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
	}
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceKubeSchedulerConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceKubeSchedulerConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
