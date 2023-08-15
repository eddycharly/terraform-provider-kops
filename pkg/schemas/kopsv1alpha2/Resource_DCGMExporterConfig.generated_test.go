package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

func TestExpandResourceDCGMExporterConfig(t *testing.T) {
	_default := kopsv1alpha2.DCGMExporterConfig{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kopsv1alpha2.DCGMExporterConfig
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"enabled": false,
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandResourceDCGMExporterConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceDCGMExporterConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceDCGMExporterConfigInto(t *testing.T) {
	_default := map[string]interface{}{
		"enabled": false,
	}
	type args struct {
		in kopsv1alpha2.DCGMExporterConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.DCGMExporterConfig{},
			},
			want: _default,
		},
		{
			name: "Enabled - default",
			args: args{
				in: func() kopsv1alpha2.DCGMExporterConfig {
					subject := kopsv1alpha2.DCGMExporterConfig{}
					subject.Enabled = false
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenResourceDCGMExporterConfigInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceDCGMExporterConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceDCGMExporterConfig(t *testing.T) {
	_default := map[string]interface{}{
		"enabled": false,
	}
	type args struct {
		in kopsv1alpha2.DCGMExporterConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.DCGMExporterConfig{},
			},
			want: _default,
		},
		{
			name: "Enabled - default",
			args: args{
				in: func() kopsv1alpha2.DCGMExporterConfig {
					subject := kopsv1alpha2.DCGMExporterConfig{}
					subject.Enabled = false
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceDCGMExporterConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceDCGMExporterConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
