package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceServiceAccountIssuerDiscoveryConfig(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.ServiceAccountIssuerDiscoveryConfig
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandDataSourceServiceAccountIssuerDiscoveryConfig(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandDataSourceServiceAccountIssuerDiscoveryConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenDataSourceServiceAccountIssuerDiscoveryConfigInto(t *testing.T) {
	_default := map[string]interface{}{
		"discovery_store":          "",
		"enable_aws_oidc_provider": false,
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
			name: "EnableAwsOIDCProvider - default",
			args: args{
				in: func() kops.ServiceAccountIssuerDiscoveryConfig {
					subject := kops.ServiceAccountIssuerDiscoveryConfig{}
					subject.EnableAWSOIDCProvider = false
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenDataSourceServiceAccountIssuerDiscoveryConfigInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceServiceAccountIssuerDiscoveryConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceServiceAccountIssuerDiscoveryConfig(t *testing.T) {
	_default := map[string]interface{}{
		"discovery_store":          "",
		"enable_aws_oidc_provider": false,
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
			name: "EnableAwsOIDCProvider - default",
			args: args{
				in: func() kops.ServiceAccountIssuerDiscoveryConfig {
					subject := kops.ServiceAccountIssuerDiscoveryConfig{}
					subject.EnableAWSOIDCProvider = false
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceServiceAccountIssuerDiscoveryConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceServiceAccountIssuerDiscoveryConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
