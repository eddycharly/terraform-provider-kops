package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

func TestExpandResourceCertManagerConfig(t *testing.T) {
	_default := kopsv1alpha2.CertManagerConfig{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kopsv1alpha2.CertManagerConfig
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"enabled":         nil,
					"managed":         nil,
					"image":           nil,
					"default_issuer":  nil,
					"nameservers":     func() []interface{} { return nil }(),
					"hosted_zone_ids": func() []interface{} { return nil }(),
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandResourceCertManagerConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceCertManagerConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceCertManagerConfigInto(t *testing.T) {
	_default := map[string]interface{}{
		"enabled":         nil,
		"managed":         nil,
		"image":           nil,
		"default_issuer":  nil,
		"nameservers":     func() []interface{} { return nil }(),
		"hosted_zone_ids": func() []interface{} { return nil }(),
	}
	type args struct {
		in kopsv1alpha2.CertManagerConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.CertManagerConfig{},
			},
			want: _default,
		},
		{
			name: "Enabled - default",
			args: args{
				in: func() kopsv1alpha2.CertManagerConfig {
					subject := kopsv1alpha2.CertManagerConfig{}
					subject.Enabled = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Managed - default",
			args: args{
				in: func() kopsv1alpha2.CertManagerConfig {
					subject := kopsv1alpha2.CertManagerConfig{}
					subject.Managed = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Image - default",
			args: args{
				in: func() kopsv1alpha2.CertManagerConfig {
					subject := kopsv1alpha2.CertManagerConfig{}
					subject.Image = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "DefaultIssuer - default",
			args: args{
				in: func() kopsv1alpha2.CertManagerConfig {
					subject := kopsv1alpha2.CertManagerConfig{}
					subject.DefaultIssuer = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Nameservers - default",
			args: args{
				in: func() kopsv1alpha2.CertManagerConfig {
					subject := kopsv1alpha2.CertManagerConfig{}
					subject.Nameservers = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "HostedZoneIds - default",
			args: args{
				in: func() kopsv1alpha2.CertManagerConfig {
					subject := kopsv1alpha2.CertManagerConfig{}
					subject.HostedZoneIDs = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenResourceCertManagerConfigInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceCertManagerConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceCertManagerConfig(t *testing.T) {
	_default := map[string]interface{}{
		"enabled":         nil,
		"managed":         nil,
		"image":           nil,
		"default_issuer":  nil,
		"nameservers":     func() []interface{} { return nil }(),
		"hosted_zone_ids": func() []interface{} { return nil }(),
	}
	type args struct {
		in kopsv1alpha2.CertManagerConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.CertManagerConfig{},
			},
			want: _default,
		},
		{
			name: "Enabled - default",
			args: args{
				in: func() kopsv1alpha2.CertManagerConfig {
					subject := kopsv1alpha2.CertManagerConfig{}
					subject.Enabled = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Managed - default",
			args: args{
				in: func() kopsv1alpha2.CertManagerConfig {
					subject := kopsv1alpha2.CertManagerConfig{}
					subject.Managed = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Image - default",
			args: args{
				in: func() kopsv1alpha2.CertManagerConfig {
					subject := kopsv1alpha2.CertManagerConfig{}
					subject.Image = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "DefaultIssuer - default",
			args: args{
				in: func() kopsv1alpha2.CertManagerConfig {
					subject := kopsv1alpha2.CertManagerConfig{}
					subject.DefaultIssuer = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Nameservers - default",
			args: args{
				in: func() kopsv1alpha2.CertManagerConfig {
					subject := kopsv1alpha2.CertManagerConfig{}
					subject.Nameservers = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "HostedZoneIds - default",
			args: args{
				in: func() kopsv1alpha2.CertManagerConfig {
					subject := kopsv1alpha2.CertManagerConfig{}
					subject.HostedZoneIDs = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceCertManagerConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceCertManagerConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
