package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceCloudControllerManagerConfig(t *testing.T) {
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
			if got := ExpandDataSourceCloudControllerManagerConfig(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandDataSourceCloudControllerManagerConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenDataSourceCloudControllerManagerConfigInto(t *testing.T) {
	type args struct {
		in  kops.CloudControllerManagerConfig
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
			FlattenDataSourceCloudControllerManagerConfigInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenDataSourceCloudControllerManagerConfig(t *testing.T) {
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
			want: map[string]interface{}{
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
			},
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
			want: map[string]interface{}{
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
			},
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
			want: map[string]interface{}{
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
			},
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
			want: map[string]interface{}{
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
			},
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
			want: map[string]interface{}{
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
			},
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
			want: map[string]interface{}{
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
			},
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
			want: map[string]interface{}{
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
			},
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
			want: map[string]interface{}{
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
			},
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
			want: map[string]interface{}{
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
			},
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
			want: map[string]interface{}{
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
			},
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
			want: map[string]interface{}{
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
			},
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
			want: map[string]interface{}{
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
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenDataSourceCloudControllerManagerConfig(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenDataSourceCloudControllerManagerConfig() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
