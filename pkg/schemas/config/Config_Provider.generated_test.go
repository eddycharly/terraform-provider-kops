package schemas

import (
	"reflect"
	"testing"

	"github.com/eddycharly/terraform-provider-kops/pkg/api/config"
	"github.com/google/go-cmp/cmp"
)

func TestExpandConfigProvider(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want config.Provider
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandConfigProvider(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandConfigProvider() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenConfigProviderInto(t *testing.T) {
	type args struct {
		in  config.Provider
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
			FlattenConfigProviderInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenConfigProvider(t *testing.T) {
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
			want: map[string]interface{}{
				"state_store": "",
				"aws":         nil,
				"openstack":   nil,
				"klog":        nil,
			},
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
			want: map[string]interface{}{
				"state_store": "",
				"aws":         nil,
				"openstack":   nil,
				"klog":        nil,
			},
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
			want: map[string]interface{}{
				"state_store": "",
				"aws":         nil,
				"openstack":   nil,
				"klog":        nil,
			},
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
			want: map[string]interface{}{
				"state_store": "",
				"aws":         nil,
				"openstack":   nil,
				"klog":        nil,
			},
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
			want: map[string]interface{}{
				"state_store": "",
				"aws":         nil,
				"openstack":   nil,
				"klog":        nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenConfigProvider(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenConfigProvider() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
