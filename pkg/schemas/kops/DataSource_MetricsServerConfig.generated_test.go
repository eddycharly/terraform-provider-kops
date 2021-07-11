package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceMetricsServerConfig(t *testing.T) {
	_default := kops.MetricsServerConfig{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.MetricsServerConfig
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"enabled":  nil,
					"image":    nil,
					"insecure": nil,
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandDataSourceMetricsServerConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceMetricsServerConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceMetricsServerConfigInto(t *testing.T) {
	_default := map[string]interface{}{
		"enabled":  nil,
		"image":    nil,
		"insecure": nil,
	}
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenDataSourceMetricsServerConfigInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceMetricsServerConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceMetricsServerConfig(t *testing.T) {
	_default := map[string]interface{}{
		"enabled":  nil,
		"image":    nil,
		"insecure": nil,
	}
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceMetricsServerConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceMetricsServerConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
