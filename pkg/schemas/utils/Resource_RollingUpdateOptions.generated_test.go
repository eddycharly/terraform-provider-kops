package schemas

import (
	"testing"

	"github.com/eddycharly/terraform-provider-kops/pkg/api/utils"
	"github.com/google/go-cmp/cmp"
)

func TestExpandResourceRollingUpdateOptions(t *testing.T) {
	_default := utils.RollingUpdateOptions{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want utils.RollingUpdateOptions
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"master_interval":     nil,
					"node_interval":       nil,
					"bastion_interval":    nil,
					"fail_on_drain_error": false,
					"fail_on_validate":    false,
					"instance_groups":     func() []interface{} { return nil }(),
					"post_drain_delay":    nil,
					"validation_timeout":  nil,
					"validate_count":      nil,
					"cloud_only":          false,
					"force":               false,
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandResourceRollingUpdateOptions(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceRollingUpdateOptions() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceRollingUpdateOptionsInto(t *testing.T) {
	_default := map[string]interface{}{
		"master_interval":     nil,
		"node_interval":       nil,
		"bastion_interval":    nil,
		"fail_on_drain_error": false,
		"fail_on_validate":    false,
		"instance_groups":     func() []interface{} { return nil }(),
		"post_drain_delay":    nil,
		"validation_timeout":  nil,
		"validate_count":      nil,
		"cloud_only":          false,
		"force":               false,
	}
	type args struct {
		in utils.RollingUpdateOptions
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: utils.RollingUpdateOptions{},
			},
			want: _default,
		},
		{
			name: "MasterInterval - default",
			args: args{
				in: func() utils.RollingUpdateOptions {
					subject := utils.RollingUpdateOptions{}
					subject.MasterInterval = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NodeInterval - default",
			args: args{
				in: func() utils.RollingUpdateOptions {
					subject := utils.RollingUpdateOptions{}
					subject.NodeInterval = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "BastionInterval - default",
			args: args{
				in: func() utils.RollingUpdateOptions {
					subject := utils.RollingUpdateOptions{}
					subject.BastionInterval = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "FailOnDrainError - default",
			args: args{
				in: func() utils.RollingUpdateOptions {
					subject := utils.RollingUpdateOptions{}
					subject.FailOnDrainError = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "FailOnValidate - default",
			args: args{
				in: func() utils.RollingUpdateOptions {
					subject := utils.RollingUpdateOptions{}
					subject.FailOnValidate = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "InstanceGroups - default",
			args: args{
				in: func() utils.RollingUpdateOptions {
					subject := utils.RollingUpdateOptions{}
					subject.InstanceGroups = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PostDrainDelay - default",
			args: args{
				in: func() utils.RollingUpdateOptions {
					subject := utils.RollingUpdateOptions{}
					subject.PostDrainDelay = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ValidationTimeout - default",
			args: args{
				in: func() utils.RollingUpdateOptions {
					subject := utils.RollingUpdateOptions{}
					subject.ValidationTimeout = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ValidateCount - default",
			args: args{
				in: func() utils.RollingUpdateOptions {
					subject := utils.RollingUpdateOptions{}
					subject.ValidateCount = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CloudOnly - default",
			args: args{
				in: func() utils.RollingUpdateOptions {
					subject := utils.RollingUpdateOptions{}
					subject.CloudOnly = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Force - default",
			args: args{
				in: func() utils.RollingUpdateOptions {
					subject := utils.RollingUpdateOptions{}
					subject.Force = false
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenResourceRollingUpdateOptionsInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceRollingUpdateOptions() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceRollingUpdateOptions(t *testing.T) {
	_default := map[string]interface{}{
		"master_interval":     nil,
		"node_interval":       nil,
		"bastion_interval":    nil,
		"fail_on_drain_error": false,
		"fail_on_validate":    false,
		"instance_groups":     func() []interface{} { return nil }(),
		"post_drain_delay":    nil,
		"validation_timeout":  nil,
		"validate_count":      nil,
		"cloud_only":          false,
		"force":               false,
	}
	type args struct {
		in utils.RollingUpdateOptions
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: utils.RollingUpdateOptions{},
			},
			want: _default,
		},
		{
			name: "MasterInterval - default",
			args: args{
				in: func() utils.RollingUpdateOptions {
					subject := utils.RollingUpdateOptions{}
					subject.MasterInterval = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NodeInterval - default",
			args: args{
				in: func() utils.RollingUpdateOptions {
					subject := utils.RollingUpdateOptions{}
					subject.NodeInterval = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "BastionInterval - default",
			args: args{
				in: func() utils.RollingUpdateOptions {
					subject := utils.RollingUpdateOptions{}
					subject.BastionInterval = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "FailOnDrainError - default",
			args: args{
				in: func() utils.RollingUpdateOptions {
					subject := utils.RollingUpdateOptions{}
					subject.FailOnDrainError = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "FailOnValidate - default",
			args: args{
				in: func() utils.RollingUpdateOptions {
					subject := utils.RollingUpdateOptions{}
					subject.FailOnValidate = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "InstanceGroups - default",
			args: args{
				in: func() utils.RollingUpdateOptions {
					subject := utils.RollingUpdateOptions{}
					subject.InstanceGroups = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PostDrainDelay - default",
			args: args{
				in: func() utils.RollingUpdateOptions {
					subject := utils.RollingUpdateOptions{}
					subject.PostDrainDelay = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ValidationTimeout - default",
			args: args{
				in: func() utils.RollingUpdateOptions {
					subject := utils.RollingUpdateOptions{}
					subject.ValidationTimeout = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ValidateCount - default",
			args: args{
				in: func() utils.RollingUpdateOptions {
					subject := utils.RollingUpdateOptions{}
					subject.ValidateCount = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CloudOnly - default",
			args: args{
				in: func() utils.RollingUpdateOptions {
					subject := utils.RollingUpdateOptions{}
					subject.CloudOnly = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Force - default",
			args: args{
				in: func() utils.RollingUpdateOptions {
					subject := utils.RollingUpdateOptions{}
					subject.Force = false
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceRollingUpdateOptions(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceRollingUpdateOptions() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
