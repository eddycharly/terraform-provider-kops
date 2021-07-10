package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceCertManagerConfig(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.CertManagerConfig
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandResourceCertManagerConfig(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandResourceCertManagerConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenResourceCertManagerConfigInto(t *testing.T) {
	type args struct {
		in  kops.CertManagerConfig
		out map[string]interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			FlattenResourceCertManagerConfigInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenResourceCertManagerConfig(t *testing.T) {
	type args struct {
		in kops.CertManagerConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.CertManagerConfig{},
			},
			want: map[string]interface{}{
				"enabled":        nil,
				"managed":        nil,
				"image":          nil,
				"default_issuer": nil,
			},
		},
		{
			name: "Enabled - default",
			args: args{
				in: func() kops.CertManagerConfig {
					subject := kops.CertManagerConfig{}
					subject.Enabled = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"enabled":        nil,
				"managed":        nil,
				"image":          nil,
				"default_issuer": nil,
			},
		},
		{
			name: "Managed - default",
			args: args{
				in: func() kops.CertManagerConfig {
					subject := kops.CertManagerConfig{}
					subject.Managed = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"enabled":        nil,
				"managed":        nil,
				"image":          nil,
				"default_issuer": nil,
			},
		},
		{
			name: "Image - default",
			args: args{
				in: func() kops.CertManagerConfig {
					subject := kops.CertManagerConfig{}
					subject.Image = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"enabled":        nil,
				"managed":        nil,
				"image":          nil,
				"default_issuer": nil,
			},
		},
		{
			name: "DefaultIssuer - default",
			args: args{
				in: func() kops.CertManagerConfig {
					subject := kops.CertManagerConfig{}
					subject.DefaultIssuer = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"enabled":        nil,
				"managed":        nil,
				"image":          nil,
				"default_issuer": nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenResourceCertManagerConfig(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenResourceCertManagerConfig() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
