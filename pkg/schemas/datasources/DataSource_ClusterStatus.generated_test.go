package schemas

import (
	"testing"

	"github.com/eddycharly/terraform-provider-kops/pkg/api/datasources"
	"github.com/google/go-cmp/cmp"
)

func TestExpandDataSourceClusterStatus(t *testing.T) {
	_default := datasources.ClusterStatus{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want datasources.ClusterStatus
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"cluster_name":    "",
					"exists":          false,
					"is_valid":        false,
					"needs_update":    false,
					"instance_groups": func() []interface{} { return nil }(),
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandDataSourceClusterStatus(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceClusterStatus() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceClusterStatusInto(t *testing.T) {
	_default := map[string]interface{}{
		"cluster_name":    "",
		"exists":          false,
		"is_valid":        false,
		"needs_update":    false,
		"instance_groups": func() []interface{} { return nil }(),
	}
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenDataSourceClusterStatusInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceClusterStatus() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceClusterStatus(t *testing.T) {
	_default := map[string]interface{}{
		"cluster_name":    "",
		"exists":          false,
		"is_valid":        false,
		"needs_update":    false,
		"instance_groups": func() []interface{} { return nil }(),
	}
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceClusterStatus(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceClusterStatus() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
