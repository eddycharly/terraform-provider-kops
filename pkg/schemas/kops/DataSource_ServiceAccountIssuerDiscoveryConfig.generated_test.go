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
	type args struct {
		in  kops.ServiceAccountIssuerDiscoveryConfig
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
			FlattenDataSourceServiceAccountIssuerDiscoveryConfigInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenDataSourceServiceAccountIssuerDiscoveryConfig(t *testing.T) {
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
			want: map[string]interface{}{
				"discovery_store":          "",
				"enable_aws_oidc_provider": false,
			},
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
			want: map[string]interface{}{
				"discovery_store":          "",
				"enable_aws_oidc_provider": false,
			},
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
			want: map[string]interface{}{
				"discovery_store":          "",
				"enable_aws_oidc_provider": false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenDataSourceServiceAccountIssuerDiscoveryConfig(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenDataSourceServiceAccountIssuerDiscoveryConfig() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
