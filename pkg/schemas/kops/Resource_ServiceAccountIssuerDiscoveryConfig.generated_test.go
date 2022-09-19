package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceServiceAccountIssuerDiscoveryConfig(t *testing.T) {
	_default := kops.ServiceAccountIssuerDiscoveryConfig{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.ServiceAccountIssuerDiscoveryConfig
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"discovery_store":          "",
					"enable_aws_oidc_provider": false,
					"additional_audiences":     func() []interface{} { return nil }(),
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandResourceServiceAccountIssuerDiscoveryConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceServiceAccountIssuerDiscoveryConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceServiceAccountIssuerDiscoveryConfigInto(t *testing.T) {
	_default := map[string]interface{}{
		"discovery_store":          "",
		"enable_aws_oidc_provider": false,
		"additional_audiences":     func() []interface{} { return nil }(),
	}
	type args struct {
		in kops.ServiceAccountIssuerDiscoveryConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.ServiceAccountIssuerDiscoveryConfig{},
			},
			want: _default,
		},
		{
			name: "DiscoveryStore - default",
			args: args{
				in: func() kops.ServiceAccountIssuerDiscoveryConfig {
					subject := kops.ServiceAccountIssuerDiscoveryConfig{}
					subject.DiscoveryStore = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableAwsOidcProvider - default",
			args: args{
				in: func() kops.ServiceAccountIssuerDiscoveryConfig {
					subject := kops.ServiceAccountIssuerDiscoveryConfig{}
					subject.EnableAWSOIDCProvider = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AdditionalAudiences - default",
			args: args{
				in: func() kops.ServiceAccountIssuerDiscoveryConfig {
					subject := kops.ServiceAccountIssuerDiscoveryConfig{}
					subject.AdditionalAudiences = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenResourceServiceAccountIssuerDiscoveryConfigInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceServiceAccountIssuerDiscoveryConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceServiceAccountIssuerDiscoveryConfig(t *testing.T) {
	_default := map[string]interface{}{
		"discovery_store":          "",
		"enable_aws_oidc_provider": false,
		"additional_audiences":     func() []interface{} { return nil }(),
	}
	type args struct {
		in kops.ServiceAccountIssuerDiscoveryConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.ServiceAccountIssuerDiscoveryConfig{},
			},
			want: _default,
		},
		{
			name: "DiscoveryStore - default",
			args: args{
				in: func() kops.ServiceAccountIssuerDiscoveryConfig {
					subject := kops.ServiceAccountIssuerDiscoveryConfig{}
					subject.DiscoveryStore = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableAwsOidcProvider - default",
			args: args{
				in: func() kops.ServiceAccountIssuerDiscoveryConfig {
					subject := kops.ServiceAccountIssuerDiscoveryConfig{}
					subject.EnableAWSOIDCProvider = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AdditionalAudiences - default",
			args: args{
				in: func() kops.ServiceAccountIssuerDiscoveryConfig {
					subject := kops.ServiceAccountIssuerDiscoveryConfig{}
					subject.AdditionalAudiences = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceServiceAccountIssuerDiscoveryConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceServiceAccountIssuerDiscoveryConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
