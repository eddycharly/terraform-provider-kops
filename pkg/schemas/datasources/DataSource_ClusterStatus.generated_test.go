package schemas

import (
	"reflect"
	"testing"

	"github.com/eddycharly/terraform-provider-kops/pkg/api/datasources"
	"github.com/google/go-cmp/cmp"
)

func TestExpandDataSourceClusterStatus(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want datasources.ClusterStatus
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandDataSourceClusterStatus(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandDataSourceClusterStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenDataSourceClusterStatusInto(t *testing.T) {
	type args struct {
		in  datasources.ClusterStatus
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
			FlattenDataSourceClusterStatusInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenDataSourceClusterStatus(t *testing.T) {
	type args struct {
		in datasources.ClusterStatus
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: datasources.ClusterStatus{},
			},
			want: map[string]interface{}{
				"cluster_name":    "",
				"exists":          false,
				"is_valid":        false,
				"needs_update":    false,
				"instance_groups": func() []interface{} { return nil }(),
			},
		},
		{
			name: "ClusterName - default",
			args: args{
				in: func() datasources.ClusterStatus {
					subject := datasources.ClusterStatus{}
					subject.ClusterName = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
				"cluster_name":    "",
				"exists":          false,
				"is_valid":        false,
				"needs_update":    false,
				"instance_groups": func() []interface{} { return nil }(),
			},
		},
		{
			name: "Exists - default",
			args: args{
				in: func() datasources.ClusterStatus {
					subject := datasources.ClusterStatus{}
					subject.Exists = false
					return subject
				}(),
			},
			want: map[string]interface{}{
				"cluster_name":    "",
				"exists":          false,
				"is_valid":        false,
				"needs_update":    false,
				"instance_groups": func() []interface{} { return nil }(),
			},
		},
		{
			name: "IsValid - default",
			args: args{
				in: func() datasources.ClusterStatus {
					subject := datasources.ClusterStatus{}
					subject.IsValid = false
					return subject
				}(),
			},
			want: map[string]interface{}{
				"cluster_name":    "",
				"exists":          false,
				"is_valid":        false,
				"needs_update":    false,
				"instance_groups": func() []interface{} { return nil }(),
			},
		},
		{
			name: "NeedsUpdate - default",
			args: args{
				in: func() datasources.ClusterStatus {
					subject := datasources.ClusterStatus{}
					subject.NeedsUpdate = false
					return subject
				}(),
			},
			want: map[string]interface{}{
				"cluster_name":    "",
				"exists":          false,
				"is_valid":        false,
				"needs_update":    false,
				"instance_groups": func() []interface{} { return nil }(),
			},
		},
		{
			name: "InstanceGroups - default",
			args: args{
				in: func() datasources.ClusterStatus {
					subject := datasources.ClusterStatus{}
					subject.InstanceGroups = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"cluster_name":    "",
				"exists":          false,
				"is_valid":        false,
				"needs_update":    false,
				"instance_groups": func() []interface{} { return nil }(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenDataSourceClusterStatus(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenDataSourceClusterStatus() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
