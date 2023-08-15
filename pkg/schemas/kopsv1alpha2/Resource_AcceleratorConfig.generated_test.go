package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

func TestExpandResourceAcceleratorConfig(t *testing.T) {
	_default := kopsv1alpha2.AcceleratorConfig{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kopsv1alpha2.AcceleratorConfig
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"accelerator_count": 0,
					"accelerator_type":  "",
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandResourceAcceleratorConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceAcceleratorConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceAcceleratorConfigInto(t *testing.T) {
	_default := map[string]interface{}{
		"accelerator_count": 0,
		"accelerator_type":  "",
	}
	type args struct {
		in kopsv1alpha2.AcceleratorConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.AcceleratorConfig{},
			},
			want: _default,
		},
		{
			name: "AcceleratorCount - default",
			args: args{
				in: func() kopsv1alpha2.AcceleratorConfig {
					subject := kopsv1alpha2.AcceleratorConfig{}
					subject.AcceleratorCount = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AcceleratorType - default",
			args: args{
				in: func() kopsv1alpha2.AcceleratorConfig {
					subject := kopsv1alpha2.AcceleratorConfig{}
					subject.AcceleratorType = ""
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenResourceAcceleratorConfigInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceAcceleratorConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceAcceleratorConfig(t *testing.T) {
	_default := map[string]interface{}{
		"accelerator_count": 0,
		"accelerator_type":  "",
	}
	type args struct {
		in kopsv1alpha2.AcceleratorConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.AcceleratorConfig{},
			},
			want: _default,
		},
		{
			name: "AcceleratorCount - default",
			args: args{
				in: func() kopsv1alpha2.AcceleratorConfig {
					subject := kopsv1alpha2.AcceleratorConfig{}
					subject.AcceleratorCount = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AcceleratorType - default",
			args: args{
				in: func() kopsv1alpha2.AcceleratorConfig {
					subject := kopsv1alpha2.AcceleratorConfig{}
					subject.AcceleratorType = ""
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceAcceleratorConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceAcceleratorConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
