package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

func TestExpandDataSourceNvidiaGPUConfig(t *testing.T) {
	_default := kopsv1alpha2.NvidiaGPUConfig{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kopsv1alpha2.NvidiaGPUConfig
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"driver_package": "",
					"enabled":        nil,
					"dcgm_exporter":  nil,
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandDataSourceNvidiaGPUConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceNvidiaGPUConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceNvidiaGPUConfigInto(t *testing.T) {
	_default := map[string]interface{}{
		"driver_package": "",
		"enabled":        nil,
		"dcgm_exporter":  nil,
	}
	type args struct {
		in kopsv1alpha2.NvidiaGPUConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.NvidiaGPUConfig{},
			},
			want: _default,
		},
		{
			name: "DriverPackage - default",
			args: args{
				in: func() kopsv1alpha2.NvidiaGPUConfig {
					subject := kopsv1alpha2.NvidiaGPUConfig{}
					subject.DriverPackage = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Enabled - default",
			args: args{
				in: func() kopsv1alpha2.NvidiaGPUConfig {
					subject := kopsv1alpha2.NvidiaGPUConfig{}
					subject.Enabled = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "DCGMExporter - default",
			args: args{
				in: func() kopsv1alpha2.NvidiaGPUConfig {
					subject := kopsv1alpha2.NvidiaGPUConfig{}
					subject.DCGMExporter = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenDataSourceNvidiaGPUConfigInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceNvidiaGPUConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceNvidiaGPUConfig(t *testing.T) {
	_default := map[string]interface{}{
		"driver_package": "",
		"enabled":        nil,
		"dcgm_exporter":  nil,
	}
	type args struct {
		in kopsv1alpha2.NvidiaGPUConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.NvidiaGPUConfig{},
			},
			want: _default,
		},
		{
			name: "DriverPackage - default",
			args: args{
				in: func() kopsv1alpha2.NvidiaGPUConfig {
					subject := kopsv1alpha2.NvidiaGPUConfig{}
					subject.DriverPackage = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Enabled - default",
			args: args{
				in: func() kopsv1alpha2.NvidiaGPUConfig {
					subject := kopsv1alpha2.NvidiaGPUConfig{}
					subject.Enabled = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "DCGMExporter - default",
			args: args{
				in: func() kopsv1alpha2.NvidiaGPUConfig {
					subject := kopsv1alpha2.NvidiaGPUConfig{}
					subject.DCGMExporter = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceNvidiaGPUConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceNvidiaGPUConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
