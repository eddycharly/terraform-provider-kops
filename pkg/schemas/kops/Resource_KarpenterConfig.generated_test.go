package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceKarpenterConfig(t *testing.T) {
	_default := kops.KarpenterConfig{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.KarpenterConfig
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
			got := ExpandResourceKarpenterConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceKarpenterConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceKarpenterConfigInto(t *testing.T) {
	_default := map[string]interface{}{
		"enabled": false,
	}
	type args struct {
		in kops.KarpenterConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.KarpenterConfig{},
			},
			want: _default,
		},
		{
			name: "Enabled - default",
			args: args{
				in: func() kops.KarpenterConfig {
					subject := kops.KarpenterConfig{}
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
			FlattenResourceKarpenterConfigInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceKarpenterConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceKarpenterConfig(t *testing.T) {
	_default := map[string]interface{}{
		"enabled": false,
	}
	type args struct {
		in kops.KarpenterConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.KarpenterConfig{},
			},
			want: _default,
		},
		{
			name: "Enabled - default",
			args: args{
				in: func() kops.KarpenterConfig {
					subject := kops.KarpenterConfig{}
					subject.Enabled = false
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceKarpenterConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceKarpenterConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
