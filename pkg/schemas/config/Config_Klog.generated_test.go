package schemas

import (
	"reflect"
	"testing"

	"github.com/eddycharly/terraform-provider-kops/pkg/api/config"
	"github.com/google/go-cmp/cmp"
)

func TestExpandConfigKlog(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want config.Klog
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandConfigKlog(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandConfigKlog() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenConfigKlogInto(t *testing.T) {
	_default := map[string]interface{}{
		"verbosity": nil,
	}
	type args struct {
		in config.Klog
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: config.Klog{},
			},
			want: _default,
		},
		{
			name: "Verbosity - default",
			args: args{
				in: func() config.Klog {
					subject := config.Klog{}
					subject.Verbosity = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenConfigKlogInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenConfigKlog() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenConfigKlog(t *testing.T) {
	_default := map[string]interface{}{
		"verbosity": nil,
	}
	type args struct {
		in config.Klog
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: config.Klog{},
			},
			want: _default,
		},
		{
			name: "Verbosity - default",
			args: args{
				in: func() config.Klog {
					subject := config.Klog{}
					subject.Verbosity = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenConfigKlog(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenConfigKlog() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
