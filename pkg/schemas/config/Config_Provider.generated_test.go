package schemas

import (
	"testing"

	"github.com/eddycharly/terraform-provider-kops/pkg/api/config"
	"github.com/google/go-cmp/cmp"
)

func TestExpandConfigProvider(t *testing.T) {
	_default := config.Provider{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want config.Provider
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"state_store":   "",
					"aws":           nil,
					"openstack":     nil,
					"klog":          nil,
					"mock":          false,
					"feature_flags": func() []interface{} { return nil }(),
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandConfigProvider(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandConfigProvider() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenConfigProviderInto(t *testing.T) {
	_default := map[string]interface{}{
		"state_store":   "",
		"aws":           nil,
		"openstack":     nil,
		"klog":          nil,
		"mock":          false,
		"feature_flags": func() []interface{} { return nil }(),
	}
	type args struct {
		in config.Provider
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: config.Provider{},
			},
			want: _default,
		},
		{
			name: "StateStore - default",
			args: args{
				in: func() config.Provider {
					subject := config.Provider{}
					subject.StateStore = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Aws - default",
			args: args{
				in: func() config.Provider {
					subject := config.Provider{}
					subject.Aws = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Openstack - default",
			args: args{
				in: func() config.Provider {
					subject := config.Provider{}
					subject.Openstack = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Klog - default",
			args: args{
				in: func() config.Provider {
					subject := config.Provider{}
					subject.Klog = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Mock - default",
			args: args{
				in: func() config.Provider {
					subject := config.Provider{}
					subject.Mock = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "FeatureFlags - default",
			args: args{
				in: func() config.Provider {
					subject := config.Provider{}
					subject.FeatureFlags = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenConfigProviderInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenConfigProvider() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenConfigProvider(t *testing.T) {
	_default := map[string]interface{}{
		"state_store":   "",
		"aws":           nil,
		"openstack":     nil,
		"klog":          nil,
		"mock":          false,
		"feature_flags": func() []interface{} { return nil }(),
	}
	type args struct {
		in config.Provider
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: config.Provider{},
			},
			want: _default,
		},
		{
			name: "StateStore - default",
			args: args{
				in: func() config.Provider {
					subject := config.Provider{}
					subject.StateStore = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Aws - default",
			args: args{
				in: func() config.Provider {
					subject := config.Provider{}
					subject.Aws = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Openstack - default",
			args: args{
				in: func() config.Provider {
					subject := config.Provider{}
					subject.Openstack = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Klog - default",
			args: args{
				in: func() config.Provider {
					subject := config.Provider{}
					subject.Klog = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Mock - default",
			args: args{
				in: func() config.Provider {
					subject := config.Provider{}
					subject.Mock = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "FeatureFlags - default",
			args: args{
				in: func() config.Provider {
					subject := config.Provider{}
					subject.FeatureFlags = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenConfigProvider(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenConfigProvider() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
