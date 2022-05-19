package schemas

import (
	"testing"

	"github.com/eddycharly/terraform-provider-kops/pkg/api/resources"
	"github.com/google/go-cmp/cmp"
)

func TestExpandResourceApplyOptions(t *testing.T) {
	_default := resources.ApplyOptions{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want resources.ApplyOptions
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"skip":                 false,
					"allow_kops_downgrade": false,
					"lifecycle_overrides":  func() map[string]interface{} { return nil }(),
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandResourceApplyOptions(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceApplyOptions() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceApplyOptionsInto(t *testing.T) {
	_default := map[string]interface{}{
		"skip":                 false,
		"allow_kops_downgrade": false,
		"lifecycle_overrides":  func() map[string]interface{} { return nil }(),
	}
	type args struct {
		in resources.ApplyOptions
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: resources.ApplyOptions{},
			},
			want: _default,
		},
		{
			name: "Skip - default",
			args: args{
				in: func() resources.ApplyOptions {
					subject := resources.ApplyOptions{}
					subject.Skip = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AllowKopsDowngrade - default",
			args: args{
				in: func() resources.ApplyOptions {
					subject := resources.ApplyOptions{}
					subject.AllowKopsDowngrade = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LifecycleOverrides - default",
			args: args{
				in: func() resources.ApplyOptions {
					subject := resources.ApplyOptions{}
					subject.LifecycleOverrides = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenResourceApplyOptionsInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceApplyOptions() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceApplyOptions(t *testing.T) {
	_default := map[string]interface{}{
		"skip":                 false,
		"allow_kops_downgrade": false,
		"lifecycle_overrides":  func() map[string]interface{} { return nil }(),
	}
	type args struct {
		in resources.ApplyOptions
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: resources.ApplyOptions{},
			},
			want: _default,
		},
		{
			name: "Skip - default",
			args: args{
				in: func() resources.ApplyOptions {
					subject := resources.ApplyOptions{}
					subject.Skip = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AllowKopsDowngrade - default",
			args: args{
				in: func() resources.ApplyOptions {
					subject := resources.ApplyOptions{}
					subject.AllowKopsDowngrade = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LifecycleOverrides - default",
			args: args{
				in: func() resources.ApplyOptions {
					subject := resources.ApplyOptions{}
					subject.LifecycleOverrides = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceApplyOptions(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceApplyOptions() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
