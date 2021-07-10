package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceCloudControllerManagerConfig(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.CloudControllerManagerConfig
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandResourceCloudControllerManagerConfig(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandResourceCloudControllerManagerConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenResourceCloudControllerManagerConfigInto(t *testing.T) {
	_default := map[string]interface{}{
		"master":                          "",
		"log_level":                       0,
		"image":                           "",
		"cloud_provider":                  "",
		"cluster_name":                    "",
		"cluster_cidr":                    "",
		"allocate_node_cidrs":             nil,
		"configure_cloud_routes":          nil,
		"cidr_allocator_type":             nil,
		"leader_election":                 nil,
		"use_service_account_credentials": nil,
	}
	type args struct {
		in kops.CloudControllerManagerConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.CloudControllerManagerConfig{},
			},
			want: _default,
		},
		{
			name: "Master - default",
			args: args{
				in: func() kops.CloudControllerManagerConfig {
					subject := kops.CloudControllerManagerConfig{}
					subject.Master = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LogLevel - default",
			args: args{
				in: func() kops.CloudControllerManagerConfig {
					subject := kops.CloudControllerManagerConfig{}
					subject.LogLevel = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Image - default",
			args: args{
				in: func() kops.CloudControllerManagerConfig {
					subject := kops.CloudControllerManagerConfig{}
					subject.Image = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CloudProvider - default",
			args: args{
				in: func() kops.CloudControllerManagerConfig {
					subject := kops.CloudControllerManagerConfig{}
					subject.CloudProvider = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ClusterName - default",
			args: args{
				in: func() kops.CloudControllerManagerConfig {
					subject := kops.CloudControllerManagerConfig{}
					subject.ClusterName = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ClusterCidr - default",
			args: args{
				in: func() kops.CloudControllerManagerConfig {
					subject := kops.CloudControllerManagerConfig{}
					subject.ClusterCIDR = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AllocateNodeCidrs - default",
			args: args{
				in: func() kops.CloudControllerManagerConfig {
					subject := kops.CloudControllerManagerConfig{}
					subject.AllocateNodeCIDRs = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ConfigureCloudRoutes - default",
			args: args{
				in: func() kops.CloudControllerManagerConfig {
					subject := kops.CloudControllerManagerConfig{}
					subject.ConfigureCloudRoutes = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CidrAllocatorType - default",
			args: args{
				in: func() kops.CloudControllerManagerConfig {
					subject := kops.CloudControllerManagerConfig{}
					subject.CIDRAllocatorType = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LeaderElection - default",
			args: args{
				in: func() kops.CloudControllerManagerConfig {
					subject := kops.CloudControllerManagerConfig{}
					subject.LeaderElection = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "UseServiceAccountCredentials - default",
			args: args{
				in: func() kops.CloudControllerManagerConfig {
					subject := kops.CloudControllerManagerConfig{}
					subject.UseServiceAccountCredentials = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenResourceCloudControllerManagerConfigInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceCloudControllerManagerConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceCloudControllerManagerConfig(t *testing.T) {
	_default := map[string]interface{}{
		"master":                          "",
		"log_level":                       0,
		"image":                           "",
		"cloud_provider":                  "",
		"cluster_name":                    "",
		"cluster_cidr":                    "",
		"allocate_node_cidrs":             nil,
		"configure_cloud_routes":          nil,
		"cidr_allocator_type":             nil,
		"leader_election":                 nil,
		"use_service_account_credentials": nil,
	}
	type args struct {
		in kops.CloudControllerManagerConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.CloudControllerManagerConfig{},
			},
			want: _default,
		},
		{
			name: "Master - default",
			args: args{
				in: func() kops.CloudControllerManagerConfig {
					subject := kops.CloudControllerManagerConfig{}
					subject.Master = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LogLevel - default",
			args: args{
				in: func() kops.CloudControllerManagerConfig {
					subject := kops.CloudControllerManagerConfig{}
					subject.LogLevel = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Image - default",
			args: args{
				in: func() kops.CloudControllerManagerConfig {
					subject := kops.CloudControllerManagerConfig{}
					subject.Image = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CloudProvider - default",
			args: args{
				in: func() kops.CloudControllerManagerConfig {
					subject := kops.CloudControllerManagerConfig{}
					subject.CloudProvider = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ClusterName - default",
			args: args{
				in: func() kops.CloudControllerManagerConfig {
					subject := kops.CloudControllerManagerConfig{}
					subject.ClusterName = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ClusterCidr - default",
			args: args{
				in: func() kops.CloudControllerManagerConfig {
					subject := kops.CloudControllerManagerConfig{}
					subject.ClusterCIDR = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AllocateNodeCidrs - default",
			args: args{
				in: func() kops.CloudControllerManagerConfig {
					subject := kops.CloudControllerManagerConfig{}
					subject.AllocateNodeCIDRs = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ConfigureCloudRoutes - default",
			args: args{
				in: func() kops.CloudControllerManagerConfig {
					subject := kops.CloudControllerManagerConfig{}
					subject.ConfigureCloudRoutes = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CidrAllocatorType - default",
			args: args{
				in: func() kops.CloudControllerManagerConfig {
					subject := kops.CloudControllerManagerConfig{}
					subject.CIDRAllocatorType = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LeaderElection - default",
			args: args{
				in: func() kops.CloudControllerManagerConfig {
					subject := kops.CloudControllerManagerConfig{}
					subject.LeaderElection = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "UseServiceAccountCredentials - default",
			args: args{
				in: func() kops.CloudControllerManagerConfig {
					subject := kops.CloudControllerManagerConfig{}
					subject.UseServiceAccountCredentials = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceCloudControllerManagerConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceCloudControllerManagerConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
