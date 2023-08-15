package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

func TestExpandDataSourceKarpenterConfig(t *testing.T) {
	_default := kopsv1alpha2.KarpenterConfig{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kopsv1alpha2.KarpenterConfig
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
			got := ExpandDataSourceKarpenterConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceKarpenterConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceKarpenterConfigInto(t *testing.T) {
	_default := map[string]interface{}{
		"enabled": false,
	}
	type args struct {
		in kopsv1alpha2.KarpenterConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.KarpenterConfig{},
			},
			want: _default,
		},
		{
			name: "Enabled - default",
			args: args{
				in: func() kopsv1alpha2.KarpenterConfig {
					subject := kopsv1alpha2.KarpenterConfig{}
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
			FlattenDataSourceKarpenterConfigInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceKarpenterConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceKarpenterConfig(t *testing.T) {
	_default := map[string]interface{}{
		"enabled": false,
	}
	type args struct {
		in kopsv1alpha2.KarpenterConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.KarpenterConfig{},
			},
			want: _default,
		},
		{
			name: "Enabled - default",
			args: args{
				in: func() kopsv1alpha2.KarpenterConfig {
					subject := kopsv1alpha2.KarpenterConfig{}
					subject.Enabled = false
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceKarpenterConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceKarpenterConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
