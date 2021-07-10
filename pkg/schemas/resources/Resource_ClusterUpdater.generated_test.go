package schemas

import (
	"reflect"
	"testing"

	"github.com/eddycharly/terraform-provider-kops/pkg/api/resources"
	"github.com/google/go-cmp/cmp"
)

func TestExpandResourceClusterUpdater(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want resources.ClusterUpdater
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandResourceClusterUpdater(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandResourceClusterUpdater() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenResourceClusterUpdaterInto(t *testing.T) {
	_default := map[string]interface{}{
		"revision":     0,
		"cluster_name": "",
		"keepers":      func() map[string]interface{} { return nil }(),
		"apply": func() []map[string]interface{} {
			return []map[string]interface{}{FlattenResourceApplyOptions(resources.ApplyOptions{})}
		}(),
		"rolling_update": func() []map[string]interface{} {
			return []map[string]interface{}{FlattenResourceRollingUpdateOptions(resources.RollingUpdateOptions{})}
		}(),
		"validate": func() []map[string]interface{} {
			return []map[string]interface{}{FlattenResourceValidateOptions(resources.ValidateOptions{})}
		}(),
	}
	type args struct {
		in resources.ClusterUpdater
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: resources.ClusterUpdater{},
			},
			want: _default,
		},
		{
			name: "Revision - default",
			args: args{
				in: func() resources.ClusterUpdater {
					subject := resources.ClusterUpdater{}
					subject.Revision = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ClusterName - default",
			args: args{
				in: func() resources.ClusterUpdater {
					subject := resources.ClusterUpdater{}
					subject.ClusterName = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Keepers - default",
			args: args{
				in: func() resources.ClusterUpdater {
					subject := resources.ClusterUpdater{}
					subject.Keepers = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Apply - default",
			args: args{
				in: func() resources.ClusterUpdater {
					subject := resources.ClusterUpdater{}
					subject.Apply = resources.ApplyOptions{}
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "RollingUpdate - default",
			args: args{
				in: func() resources.ClusterUpdater {
					subject := resources.ClusterUpdater{}
					subject.RollingUpdate = resources.RollingUpdateOptions{}
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Validate - default",
			args: args{
				in: func() resources.ClusterUpdater {
					subject := resources.ClusterUpdater{}
					subject.Validate = resources.ValidateOptions{}
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenResourceClusterUpdaterInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceClusterUpdater() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceClusterUpdater(t *testing.T) {
	_default := map[string]interface{}{
		"revision":     0,
		"cluster_name": "",
		"keepers":      func() map[string]interface{} { return nil }(),
		"apply": func() []map[string]interface{} {
			return []map[string]interface{}{FlattenResourceApplyOptions(resources.ApplyOptions{})}
		}(),
		"rolling_update": func() []map[string]interface{} {
			return []map[string]interface{}{FlattenResourceRollingUpdateOptions(resources.RollingUpdateOptions{})}
		}(),
		"validate": func() []map[string]interface{} {
			return []map[string]interface{}{FlattenResourceValidateOptions(resources.ValidateOptions{})}
		}(),
	}
	type args struct {
		in resources.ClusterUpdater
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: resources.ClusterUpdater{},
			},
			want: _default,
		},
		{
			name: "Revision - default",
			args: args{
				in: func() resources.ClusterUpdater {
					subject := resources.ClusterUpdater{}
					subject.Revision = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ClusterName - default",
			args: args{
				in: func() resources.ClusterUpdater {
					subject := resources.ClusterUpdater{}
					subject.ClusterName = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Keepers - default",
			args: args{
				in: func() resources.ClusterUpdater {
					subject := resources.ClusterUpdater{}
					subject.Keepers = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Apply - default",
			args: args{
				in: func() resources.ClusterUpdater {
					subject := resources.ClusterUpdater{}
					subject.Apply = resources.ApplyOptions{}
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "RollingUpdate - default",
			args: args{
				in: func() resources.ClusterUpdater {
					subject := resources.ClusterUpdater{}
					subject.RollingUpdate = resources.RollingUpdateOptions{}
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Validate - default",
			args: args{
				in: func() resources.ClusterUpdater {
					subject := resources.ClusterUpdater{}
					subject.Validate = resources.ValidateOptions{}
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceClusterUpdater(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceClusterUpdater() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
