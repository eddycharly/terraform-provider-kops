package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceAcceleratorConfig(t *testing.T) {
	_default := kops.AcceleratorConfig{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.AcceleratorConfig
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
			got := ExpandDataSourceAcceleratorConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceAcceleratorConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceAcceleratorConfigInto(t *testing.T) {
	_default := map[string]interface{}{
		"accelerator_count": 0,
		"accelerator_type":  "",
	}
	type args struct {
		in kops.AcceleratorConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.AcceleratorConfig{},
			},
			want: _default,
		},
		{
			name: "AcceleratorCount - default",
			args: args{
				in: func() kops.AcceleratorConfig {
					subject := kops.AcceleratorConfig{}
					subject.AcceleratorCount = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AcceleratorType - default",
			args: args{
				in: func() kops.AcceleratorConfig {
					subject := kops.AcceleratorConfig{}
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
			FlattenDataSourceAcceleratorConfigInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceAcceleratorConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceAcceleratorConfig(t *testing.T) {
	_default := map[string]interface{}{
		"accelerator_count": 0,
		"accelerator_type":  "",
	}
	type args struct {
		in kops.AcceleratorConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.AcceleratorConfig{},
			},
			want: _default,
		},
		{
			name: "AcceleratorCount - default",
			args: args{
				in: func() kops.AcceleratorConfig {
					subject := kops.AcceleratorConfig{}
					subject.AcceleratorCount = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AcceleratorType - default",
			args: args{
				in: func() kops.AcceleratorConfig {
					subject := kops.AcceleratorConfig{}
					subject.AcceleratorType = ""
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceAcceleratorConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceAcceleratorConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
