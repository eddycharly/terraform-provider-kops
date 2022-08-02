package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceCertManagerConfig(t *testing.T) {
	_default := kops.CertManagerConfig{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.CertManagerConfig
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"enabled":        nil,
					"managed":        nil,
					"image":          nil,
					"default_issuer": nil,
					"nameservers":    func() []interface{} { return nil }(),
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandDataSourceCertManagerConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceCertManagerConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceCertManagerConfigInto(t *testing.T) {
	_default := map[string]interface{}{
		"enabled":        nil,
		"managed":        nil,
		"image":          nil,
		"default_issuer": nil,
		"nameservers":    func() []interface{} { return nil }(),
	}
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
		},
		{
			name: "Nameservers - default",
			args: args{
				in: func() kops.CertManagerConfig {
					subject := kops.CertManagerConfig{}
					subject.Nameservers = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenDataSourceCertManagerConfigInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceCertManagerConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceCertManagerConfig(t *testing.T) {
	_default := map[string]interface{}{
		"enabled":        nil,
		"managed":        nil,
		"image":          nil,
		"default_issuer": nil,
		"nameservers":    func() []interface{} { return nil }(),
	}
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
		},
		{
			name: "Nameservers - default",
			args: args{
				in: func() kops.CertManagerConfig {
					subject := kops.CertManagerConfig{}
					subject.Nameservers = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceCertManagerConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceCertManagerConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
