package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceNvidiaGPUConfig(t *testing.T) {
	_default := kops.NvidiaGPUConfig{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.NvidiaGPUConfig
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"driver_package": "",
					"enabled":        nil,
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandResourceNvidiaGPUConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceNvidiaGPUConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceNvidiaGPUConfigInto(t *testing.T) {
	_default := map[string]interface{}{
		"driver_package": "",
		"enabled":        nil,
	}
	type args struct {
		in kops.NvidiaGPUConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.NvidiaGPUConfig{},
			},
			want: _default,
		},
		{
			name: "DriverPackage - default",
			args: args{
				in: func() kops.NvidiaGPUConfig {
					subject := kops.NvidiaGPUConfig{}
					subject.DriverPackage = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Enabled - default",
			args: args{
				in: func() kops.NvidiaGPUConfig {
					subject := kops.NvidiaGPUConfig{}
					subject.Enabled = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenResourceNvidiaGPUConfigInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceNvidiaGPUConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceNvidiaGPUConfig(t *testing.T) {
	_default := map[string]interface{}{
		"driver_package": "",
		"enabled":        nil,
	}
	type args struct {
		in kops.NvidiaGPUConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.NvidiaGPUConfig{},
			},
			want: _default,
		},
		{
			name: "DriverPackage - default",
			args: args{
				in: func() kops.NvidiaGPUConfig {
					subject := kops.NvidiaGPUConfig{}
					subject.DriverPackage = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Enabled - default",
			args: args{
				in: func() kops.NvidiaGPUConfig {
					subject := kops.NvidiaGPUConfig{}
					subject.Enabled = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceNvidiaGPUConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceNvidiaGPUConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
