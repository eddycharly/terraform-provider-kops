package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceMetricsServerConfig(t *testing.T) {
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
			if got := ExpandResourceMetricsServerConfig(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandResourceMetricsServerConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenResourceMetricsServerConfigInto(t *testing.T) {
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
			FlattenResourceMetricsServerConfigInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceMetricsServerConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceMetricsServerConfig(t *testing.T) {
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
			got := FlattenResourceMetricsServerConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceMetricsServerConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
