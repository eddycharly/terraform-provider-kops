package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

func TestExpandResourceMetricsServerConfig(t *testing.T) {
	_default := kopsv1alpha2.MetricsServerConfig{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kopsv1alpha2.MetricsServerConfig
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
			got := ExpandResourceMetricsServerConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceMetricsServerConfig() mismatch (-want +got):\n%s", diff)
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
		in kopsv1alpha2.MetricsServerConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.MetricsServerConfig{},
			},
			want: _default,
		},
		{
			name: "Enabled - default",
			args: args{
				in: func() kopsv1alpha2.MetricsServerConfig {
					subject := kopsv1alpha2.MetricsServerConfig{}
					subject.Enabled = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Image - default",
			args: args{
				in: func() kopsv1alpha2.MetricsServerConfig {
					subject := kopsv1alpha2.MetricsServerConfig{}
					subject.Image = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Insecure - default",
			args: args{
				in: func() kopsv1alpha2.MetricsServerConfig {
					subject := kopsv1alpha2.MetricsServerConfig{}
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
		in kopsv1alpha2.MetricsServerConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.MetricsServerConfig{},
			},
			want: _default,
		},
		{
			name: "Enabled - default",
			args: args{
				in: func() kopsv1alpha2.MetricsServerConfig {
					subject := kopsv1alpha2.MetricsServerConfig{}
					subject.Enabled = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Image - default",
			args: args{
				in: func() kopsv1alpha2.MetricsServerConfig {
					subject := kopsv1alpha2.MetricsServerConfig{}
					subject.Image = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Insecure - default",
			args: args{
				in: func() kopsv1alpha2.MetricsServerConfig {
					subject := kopsv1alpha2.MetricsServerConfig{}
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
