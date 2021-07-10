package schemas

import (
	"reflect"
	"testing"

	"github.com/eddycharly/terraform-provider-kops/pkg/api/resources"
	"github.com/google/go-cmp/cmp"
)

func TestExpandResourceRollingUpdateOptions(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want resources.RollingUpdateOptions
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandResourceRollingUpdateOptions(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandResourceRollingUpdateOptions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenResourceRollingUpdateOptionsInto(t *testing.T) {
	type args struct {
		in  resources.RollingUpdateOptions
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
			FlattenResourceRollingUpdateOptionsInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenResourceRollingUpdateOptions(t *testing.T) {
	type args struct {
		in resources.RollingUpdateOptions
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: resources.RollingUpdateOptions{},
			},
			want: map[string]interface{}{
				"skip":                false,
				"master_interval":     nil,
				"node_interval":       nil,
				"bastion_interval":    nil,
				"fail_on_drain_error": false,
				"fail_on_validate":    false,
				"post_drain_delay":    nil,
				"validation_timeout":  nil,
				"validate_count":      nil,
				"cloud_only":          false,
				"force":               false,
			},
		},
		{
			name: "Skip - default",
			args: args{
				in: func() resources.RollingUpdateOptions {
					subject := resources.RollingUpdateOptions{}
					subject.Skip = false
					return subject
				}(),
			},
			want: map[string]interface{}{
				"skip":                false,
				"master_interval":     nil,
				"node_interval":       nil,
				"bastion_interval":    nil,
				"fail_on_drain_error": false,
				"fail_on_validate":    false,
				"post_drain_delay":    nil,
				"validation_timeout":  nil,
				"validate_count":      nil,
				"cloud_only":          false,
				"force":               false,
			},
		},
		{
			name: "MasterInterval - default",
			args: args{
				in: func() resources.RollingUpdateOptions {
					subject := resources.RollingUpdateOptions{}
					subject.MasterInterval = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"skip":                false,
				"master_interval":     nil,
				"node_interval":       nil,
				"bastion_interval":    nil,
				"fail_on_drain_error": false,
				"fail_on_validate":    false,
				"post_drain_delay":    nil,
				"validation_timeout":  nil,
				"validate_count":      nil,
				"cloud_only":          false,
				"force":               false,
			},
		},
		{
			name: "NodeInterval - default",
			args: args{
				in: func() resources.RollingUpdateOptions {
					subject := resources.RollingUpdateOptions{}
					subject.NodeInterval = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"skip":                false,
				"master_interval":     nil,
				"node_interval":       nil,
				"bastion_interval":    nil,
				"fail_on_drain_error": false,
				"fail_on_validate":    false,
				"post_drain_delay":    nil,
				"validation_timeout":  nil,
				"validate_count":      nil,
				"cloud_only":          false,
				"force":               false,
			},
		},
		{
			name: "BastionInterval - default",
			args: args{
				in: func() resources.RollingUpdateOptions {
					subject := resources.RollingUpdateOptions{}
					subject.BastionInterval = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"skip":                false,
				"master_interval":     nil,
				"node_interval":       nil,
				"bastion_interval":    nil,
				"fail_on_drain_error": false,
				"fail_on_validate":    false,
				"post_drain_delay":    nil,
				"validation_timeout":  nil,
				"validate_count":      nil,
				"cloud_only":          false,
				"force":               false,
			},
		},
		{
			name: "FailOnDrainError - default",
			args: args{
				in: func() resources.RollingUpdateOptions {
					subject := resources.RollingUpdateOptions{}
					subject.FailOnDrainError = false
					return subject
				}(),
			},
			want: map[string]interface{}{
				"skip":                false,
				"master_interval":     nil,
				"node_interval":       nil,
				"bastion_interval":    nil,
				"fail_on_drain_error": false,
				"fail_on_validate":    false,
				"post_drain_delay":    nil,
				"validation_timeout":  nil,
				"validate_count":      nil,
				"cloud_only":          false,
				"force":               false,
			},
		},
		{
			name: "FailOnValidate - default",
			args: args{
				in: func() resources.RollingUpdateOptions {
					subject := resources.RollingUpdateOptions{}
					subject.FailOnValidate = false
					return subject
				}(),
			},
			want: map[string]interface{}{
				"skip":                false,
				"master_interval":     nil,
				"node_interval":       nil,
				"bastion_interval":    nil,
				"fail_on_drain_error": false,
				"fail_on_validate":    false,
				"post_drain_delay":    nil,
				"validation_timeout":  nil,
				"validate_count":      nil,
				"cloud_only":          false,
				"force":               false,
			},
		},
		{
			name: "PostDrainDelay - default",
			args: args{
				in: func() resources.RollingUpdateOptions {
					subject := resources.RollingUpdateOptions{}
					subject.PostDrainDelay = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"skip":                false,
				"master_interval":     nil,
				"node_interval":       nil,
				"bastion_interval":    nil,
				"fail_on_drain_error": false,
				"fail_on_validate":    false,
				"post_drain_delay":    nil,
				"validation_timeout":  nil,
				"validate_count":      nil,
				"cloud_only":          false,
				"force":               false,
			},
		},
		{
			name: "ValidationTimeout - default",
			args: args{
				in: func() resources.RollingUpdateOptions {
					subject := resources.RollingUpdateOptions{}
					subject.ValidationTimeout = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"skip":                false,
				"master_interval":     nil,
				"node_interval":       nil,
				"bastion_interval":    nil,
				"fail_on_drain_error": false,
				"fail_on_validate":    false,
				"post_drain_delay":    nil,
				"validation_timeout":  nil,
				"validate_count":      nil,
				"cloud_only":          false,
				"force":               false,
			},
		},
		{
			name: "ValidateCount - default",
			args: args{
				in: func() resources.RollingUpdateOptions {
					subject := resources.RollingUpdateOptions{}
					subject.ValidateCount = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"skip":                false,
				"master_interval":     nil,
				"node_interval":       nil,
				"bastion_interval":    nil,
				"fail_on_drain_error": false,
				"fail_on_validate":    false,
				"post_drain_delay":    nil,
				"validation_timeout":  nil,
				"validate_count":      nil,
				"cloud_only":          false,
				"force":               false,
			},
		},
		{
			name: "CloudOnly - default",
			args: args{
				in: func() resources.RollingUpdateOptions {
					subject := resources.RollingUpdateOptions{}
					subject.CloudOnly = false
					return subject
				}(),
			},
			want: map[string]interface{}{
				"skip":                false,
				"master_interval":     nil,
				"node_interval":       nil,
				"bastion_interval":    nil,
				"fail_on_drain_error": false,
				"fail_on_validate":    false,
				"post_drain_delay":    nil,
				"validation_timeout":  nil,
				"validate_count":      nil,
				"cloud_only":          false,
				"force":               false,
			},
		},
		{
			name: "Force - default",
			args: args{
				in: func() resources.RollingUpdateOptions {
					subject := resources.RollingUpdateOptions{}
					subject.Force = false
					return subject
				}(),
			},
			want: map[string]interface{}{
				"skip":                false,
				"master_interval":     nil,
				"node_interval":       nil,
				"bastion_interval":    nil,
				"fail_on_drain_error": false,
				"fail_on_validate":    false,
				"post_drain_delay":    nil,
				"validation_timeout":  nil,
				"validate_count":      nil,
				"cloud_only":          false,
				"force":               false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenResourceRollingUpdateOptions(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenResourceRollingUpdateOptions() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
