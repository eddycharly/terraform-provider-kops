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
	type args struct {
		in  config.Klog
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
			FlattenConfigKlogInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenConfigKlog(t *testing.T) {
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
			want: map[string]interface{}{
				"verbosity": nil,
			},
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
			want: map[string]interface{}{
				"verbosity": nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenConfigKlog(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenConfigKlog() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
