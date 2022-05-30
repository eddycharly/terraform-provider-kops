package schemas

import (
	"testing"

	"github.com/eddycharly/terraform-provider-kops/pkg/api/resources"
	"github.com/google/go-cmp/cmp"
)

func TestExpandResourceClusterUpdater(t *testing.T) {
	_default := resources.ClusterUpdater{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want resources.ClusterUpdater
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"revision":         0,
					"provider_version": "",
					"cluster_name":     "",
					"keepers":          func() map[string]interface{} { return nil }(),
					"apply":            func() []interface{} { return []interface{}{FlattenResourceApplyOptions(resources.ApplyOptions{})} }(),
					"rolling_update": func() []interface{} {
						return []interface{}{FlattenResourceRollingUpdateOptions(resources.RollingUpdateOptions{})}
					}(),
					"validate": func() []interface{} {
						return []interface{}{FlattenResourceValidateOptions(resources.ValidateOptions{})}
					}(),
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandResourceClusterUpdater(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceClusterUpdater() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceClusterUpdaterInto(t *testing.T) {
	_default := map[string]interface{}{
		"revision":         0,
		"provider_version": "",
		"cluster_name":     "",
		"keepers":          func() map[string]interface{} { return nil }(),
		"apply":            func() []interface{} { return []interface{}{FlattenResourceApplyOptions(resources.ApplyOptions{})} }(),
		"rolling_update": func() []interface{} {
			return []interface{}{FlattenResourceRollingUpdateOptions(resources.RollingUpdateOptions{})}
		}(),
		"validate": func() []interface{} {
			return []interface{}{FlattenResourceValidateOptions(resources.ValidateOptions{})}
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
			name: "ProviderVersion - default",
			args: args{
				in: func() resources.ClusterUpdater {
					subject := resources.ClusterUpdater{}
					subject.ProviderVersion = ""
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
		"revision":         0,
		"provider_version": "",
		"cluster_name":     "",
		"keepers":          func() map[string]interface{} { return nil }(),
		"apply":            func() []interface{} { return []interface{}{FlattenResourceApplyOptions(resources.ApplyOptions{})} }(),
		"rolling_update": func() []interface{} {
			return []interface{}{FlattenResourceRollingUpdateOptions(resources.RollingUpdateOptions{})}
		}(),
		"validate": func() []interface{} {
			return []interface{}{FlattenResourceValidateOptions(resources.ValidateOptions{})}
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
			name: "ProviderVersion - default",
			args: args{
				in: func() resources.ClusterUpdater {
					subject := resources.ClusterUpdater{}
					subject.ProviderVersion = ""
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
