package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

func TestExpandDataSourceCloudControllerManagerConfig(t *testing.T) {
	_default := kopsv1alpha2.CloudControllerManagerConfig{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kopsv1alpha2.CloudControllerManagerConfig
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"master":                          "",
					"log_level":                       0,
					"image":                           "",
					"cloud_provider":                  "",
					"cluster_name":                    "",
					"allow_untagged_cloud":            nil,
					"cluster_cidr":                    "",
					"allocate_node_cidrs":             nil,
					"configure_cloud_routes":          nil,
					"controllers":                     func() []interface{} { return nil }(),
					"cidr_allocator_type":             nil,
					"leader_election":                 nil,
					"use_service_account_credentials": nil,
					"enable_leader_migration":         nil,
					"cpu_request":                     nil,
					"node_status_update_frequency":    nil,
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandDataSourceCloudControllerManagerConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceCloudControllerManagerConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceCloudControllerManagerConfigInto(t *testing.T) {
	_default := map[string]interface{}{
		"master":                          "",
		"log_level":                       0,
		"image":                           "",
		"cloud_provider":                  "",
		"cluster_name":                    "",
		"allow_untagged_cloud":            nil,
		"cluster_cidr":                    "",
		"allocate_node_cidrs":             nil,
		"configure_cloud_routes":          nil,
		"controllers":                     func() []interface{} { return nil }(),
		"cidr_allocator_type":             nil,
		"leader_election":                 nil,
		"use_service_account_credentials": nil,
		"enable_leader_migration":         nil,
		"cpu_request":                     nil,
		"node_status_update_frequency":    nil,
	}
	type args struct {
		in kopsv1alpha2.CloudControllerManagerConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.CloudControllerManagerConfig{},
			},
			want: _default,
		},
		{
			name: "Master - default",
			args: args{
				in: func() kopsv1alpha2.CloudControllerManagerConfig {
					subject := kopsv1alpha2.CloudControllerManagerConfig{}
					subject.Master = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LogLevel - default",
			args: args{
				in: func() kopsv1alpha2.CloudControllerManagerConfig {
					subject := kopsv1alpha2.CloudControllerManagerConfig{}
					subject.LogLevel = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Image - default",
			args: args{
				in: func() kopsv1alpha2.CloudControllerManagerConfig {
					subject := kopsv1alpha2.CloudControllerManagerConfig{}
					subject.Image = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CloudProvider - default",
			args: args{
				in: func() kopsv1alpha2.CloudControllerManagerConfig {
					subject := kopsv1alpha2.CloudControllerManagerConfig{}
					subject.CloudProvider = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ClusterName - default",
			args: args{
				in: func() kopsv1alpha2.CloudControllerManagerConfig {
					subject := kopsv1alpha2.CloudControllerManagerConfig{}
					subject.ClusterName = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AllowUntaggedCloud - default",
			args: args{
				in: func() kopsv1alpha2.CloudControllerManagerConfig {
					subject := kopsv1alpha2.CloudControllerManagerConfig{}
					subject.AllowUntaggedCloud = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ClusterCidr - default",
			args: args{
				in: func() kopsv1alpha2.CloudControllerManagerConfig {
					subject := kopsv1alpha2.CloudControllerManagerConfig{}
					subject.ClusterCIDR = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AllocateNodeCidrs - default",
			args: args{
				in: func() kopsv1alpha2.CloudControllerManagerConfig {
					subject := kopsv1alpha2.CloudControllerManagerConfig{}
					subject.AllocateNodeCIDRs = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ConfigureCloudRoutes - default",
			args: args{
				in: func() kopsv1alpha2.CloudControllerManagerConfig {
					subject := kopsv1alpha2.CloudControllerManagerConfig{}
					subject.ConfigureCloudRoutes = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Controllers - default",
			args: args{
				in: func() kopsv1alpha2.CloudControllerManagerConfig {
					subject := kopsv1alpha2.CloudControllerManagerConfig{}
					subject.Controllers = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CidrAllocatorType - default",
			args: args{
				in: func() kopsv1alpha2.CloudControllerManagerConfig {
					subject := kopsv1alpha2.CloudControllerManagerConfig{}
					subject.CIDRAllocatorType = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LeaderElection - default",
			args: args{
				in: func() kopsv1alpha2.CloudControllerManagerConfig {
					subject := kopsv1alpha2.CloudControllerManagerConfig{}
					subject.LeaderElection = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "UseServiceAccountCredentials - default",
			args: args{
				in: func() kopsv1alpha2.CloudControllerManagerConfig {
					subject := kopsv1alpha2.CloudControllerManagerConfig{}
					subject.UseServiceAccountCredentials = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableLeaderMigration - default",
			args: args{
				in: func() kopsv1alpha2.CloudControllerManagerConfig {
					subject := kopsv1alpha2.CloudControllerManagerConfig{}
					subject.EnableLeaderMigration = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CpuRequest - default",
			args: args{
				in: func() kopsv1alpha2.CloudControllerManagerConfig {
					subject := kopsv1alpha2.CloudControllerManagerConfig{}
					subject.CPURequest = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NodeStatusUpdateFrequency - default",
			args: args{
				in: func() kopsv1alpha2.CloudControllerManagerConfig {
					subject := kopsv1alpha2.CloudControllerManagerConfig{}
					subject.NodeStatusUpdateFrequency = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenDataSourceCloudControllerManagerConfigInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceCloudControllerManagerConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceCloudControllerManagerConfig(t *testing.T) {
	_default := map[string]interface{}{
		"master":                          "",
		"log_level":                       0,
		"image":                           "",
		"cloud_provider":                  "",
		"cluster_name":                    "",
		"allow_untagged_cloud":            nil,
		"cluster_cidr":                    "",
		"allocate_node_cidrs":             nil,
		"configure_cloud_routes":          nil,
		"controllers":                     func() []interface{} { return nil }(),
		"cidr_allocator_type":             nil,
		"leader_election":                 nil,
		"use_service_account_credentials": nil,
		"enable_leader_migration":         nil,
		"cpu_request":                     nil,
		"node_status_update_frequency":    nil,
	}
	type args struct {
		in kopsv1alpha2.CloudControllerManagerConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.CloudControllerManagerConfig{},
			},
			want: _default,
		},
		{
			name: "Master - default",
			args: args{
				in: func() kopsv1alpha2.CloudControllerManagerConfig {
					subject := kopsv1alpha2.CloudControllerManagerConfig{}
					subject.Master = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LogLevel - default",
			args: args{
				in: func() kopsv1alpha2.CloudControllerManagerConfig {
					subject := kopsv1alpha2.CloudControllerManagerConfig{}
					subject.LogLevel = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Image - default",
			args: args{
				in: func() kopsv1alpha2.CloudControllerManagerConfig {
					subject := kopsv1alpha2.CloudControllerManagerConfig{}
					subject.Image = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CloudProvider - default",
			args: args{
				in: func() kopsv1alpha2.CloudControllerManagerConfig {
					subject := kopsv1alpha2.CloudControllerManagerConfig{}
					subject.CloudProvider = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ClusterName - default",
			args: args{
				in: func() kopsv1alpha2.CloudControllerManagerConfig {
					subject := kopsv1alpha2.CloudControllerManagerConfig{}
					subject.ClusterName = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AllowUntaggedCloud - default",
			args: args{
				in: func() kopsv1alpha2.CloudControllerManagerConfig {
					subject := kopsv1alpha2.CloudControllerManagerConfig{}
					subject.AllowUntaggedCloud = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ClusterCidr - default",
			args: args{
				in: func() kopsv1alpha2.CloudControllerManagerConfig {
					subject := kopsv1alpha2.CloudControllerManagerConfig{}
					subject.ClusterCIDR = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AllocateNodeCidrs - default",
			args: args{
				in: func() kopsv1alpha2.CloudControllerManagerConfig {
					subject := kopsv1alpha2.CloudControllerManagerConfig{}
					subject.AllocateNodeCIDRs = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ConfigureCloudRoutes - default",
			args: args{
				in: func() kopsv1alpha2.CloudControllerManagerConfig {
					subject := kopsv1alpha2.CloudControllerManagerConfig{}
					subject.ConfigureCloudRoutes = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Controllers - default",
			args: args{
				in: func() kopsv1alpha2.CloudControllerManagerConfig {
					subject := kopsv1alpha2.CloudControllerManagerConfig{}
					subject.Controllers = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CidrAllocatorType - default",
			args: args{
				in: func() kopsv1alpha2.CloudControllerManagerConfig {
					subject := kopsv1alpha2.CloudControllerManagerConfig{}
					subject.CIDRAllocatorType = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LeaderElection - default",
			args: args{
				in: func() kopsv1alpha2.CloudControllerManagerConfig {
					subject := kopsv1alpha2.CloudControllerManagerConfig{}
					subject.LeaderElection = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "UseServiceAccountCredentials - default",
			args: args{
				in: func() kopsv1alpha2.CloudControllerManagerConfig {
					subject := kopsv1alpha2.CloudControllerManagerConfig{}
					subject.UseServiceAccountCredentials = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableLeaderMigration - default",
			args: args{
				in: func() kopsv1alpha2.CloudControllerManagerConfig {
					subject := kopsv1alpha2.CloudControllerManagerConfig{}
					subject.EnableLeaderMigration = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CpuRequest - default",
			args: args{
				in: func() kopsv1alpha2.CloudControllerManagerConfig {
					subject := kopsv1alpha2.CloudControllerManagerConfig{}
					subject.CPURequest = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NodeStatusUpdateFrequency - default",
			args: args{
				in: func() kopsv1alpha2.CloudControllerManagerConfig {
					subject := kopsv1alpha2.CloudControllerManagerConfig{}
					subject.NodeStatusUpdateFrequency = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceCloudControllerManagerConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceCloudControllerManagerConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
