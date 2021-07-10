package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceMetricsServerConfig(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.MetricsServerConfig
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandDataSourceMetricsServerConfig(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandDataSourceMetricsServerConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenDataSourceMetricsServerConfigInto(t *testing.T) {
	type args struct {
		in  kops.MetricsServerConfig
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
			FlattenDataSourceMetricsServerConfigInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenDataSourceMetricsServerConfig(t *testing.T) {
	type args struct {
		in kops.MetricsServerConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.MetricsServerConfig{},
			},
			want: map[string]interface{}{
				"enabled":  nil,
				"image":    nil,
				"insecure": nil,
			},
		},
		{
			name: "Enabled - default",
			args: args{
				in: func() kops.MetricsServerConfig {
					subject := kops.MetricsServerConfig{}
					subject.Enabled = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"enabled":  nil,
				"image":    nil,
				"insecure": nil,
			},
		},
		{
			name: "Image - default",
			args: args{
				in: func() kops.MetricsServerConfig {
					subject := kops.MetricsServerConfig{}
					subject.Image = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"enabled":  nil,
				"image":    nil,
				"insecure": nil,
			},
		},
		{
			name: "Insecure - default",
			args: args{
				in: func() kops.MetricsServerConfig {
					subject := kops.MetricsServerConfig{}
					subject.Insecure = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"enabled":  nil,
				"image":    nil,
				"insecure": nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenDataSourceMetricsServerConfig(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenDataSourceMetricsServerConfig() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
