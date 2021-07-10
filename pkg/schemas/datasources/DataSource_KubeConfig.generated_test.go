package schemas

import (
	"reflect"
	"testing"

	"github.com/eddycharly/terraform-provider-kops/pkg/api/datasources"
	"github.com/google/go-cmp/cmp"
)

func TestExpandDataSourceKubeConfig(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want datasources.KubeConfig
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandDataSourceKubeConfig(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandDataSourceKubeConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenDataSourceKubeConfigInto(t *testing.T) {
	type args struct {
		in  datasources.KubeConfig
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
			FlattenDataSourceKubeConfigInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenDataSourceKubeConfig(t *testing.T) {
	type args struct {
		in datasources.KubeConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: datasources.KubeConfig{},
			},
			want: map[string]interface{}{
				"cluster_name":  "",
				"admin":         nil,
				"internal":      false,
				"server":        "",
				"context":       "",
				"namespace":     "",
				"kube_user":     "",
				"kube_password": "",
				"ca_cert":       "",
				"client_cert":   "",
				"client_key":    "",
			},
		},
		{
			name: "ClusterName - default",
			args: args{
				in: func() datasources.KubeConfig {
					subject := datasources.KubeConfig{}
					subject.ClusterName = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
				"cluster_name":  "",
				"admin":         nil,
				"internal":      false,
				"server":        "",
				"context":       "",
				"namespace":     "",
				"kube_user":     "",
				"kube_password": "",
				"ca_cert":       "",
				"client_cert":   "",
				"client_key":    "",
			},
		},
		{
			name: "Admin - default",
			args: args{
				in: func() datasources.KubeConfig {
					subject := datasources.KubeConfig{}
					subject.Admin = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"cluster_name":  "",
				"admin":         nil,
				"internal":      false,
				"server":        "",
				"context":       "",
				"namespace":     "",
				"kube_user":     "",
				"kube_password": "",
				"ca_cert":       "",
				"client_cert":   "",
				"client_key":    "",
			},
		},
		{
			name: "Internal - default",
			args: args{
				in: func() datasources.KubeConfig {
					subject := datasources.KubeConfig{}
					subject.Internal = false
					return subject
				}(),
			},
			want: map[string]interface{}{
				"cluster_name":  "",
				"admin":         nil,
				"internal":      false,
				"server":        "",
				"context":       "",
				"namespace":     "",
				"kube_user":     "",
				"kube_password": "",
				"ca_cert":       "",
				"client_cert":   "",
				"client_key":    "",
			},
		},
		{
			name: "Server - default",
			args: args{
				in: func() datasources.KubeConfig {
					subject := datasources.KubeConfig{}
					subject.Server = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
				"cluster_name":  "",
				"admin":         nil,
				"internal":      false,
				"server":        "",
				"context":       "",
				"namespace":     "",
				"kube_user":     "",
				"kube_password": "",
				"ca_cert":       "",
				"client_cert":   "",
				"client_key":    "",
			},
		},
		{
			name: "Context - default",
			args: args{
				in: func() datasources.KubeConfig {
					subject := datasources.KubeConfig{}
					subject.Context = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
				"cluster_name":  "",
				"admin":         nil,
				"internal":      false,
				"server":        "",
				"context":       "",
				"namespace":     "",
				"kube_user":     "",
				"kube_password": "",
				"ca_cert":       "",
				"client_cert":   "",
				"client_key":    "",
			},
		},
		{
			name: "Namespace - default",
			args: args{
				in: func() datasources.KubeConfig {
					subject := datasources.KubeConfig{}
					subject.Namespace = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
				"cluster_name":  "",
				"admin":         nil,
				"internal":      false,
				"server":        "",
				"context":       "",
				"namespace":     "",
				"kube_user":     "",
				"kube_password": "",
				"ca_cert":       "",
				"client_cert":   "",
				"client_key":    "",
			},
		},
		{
			name: "KubeUser - default",
			args: args{
				in: func() datasources.KubeConfig {
					subject := datasources.KubeConfig{}
					subject.KubeUser = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
				"cluster_name":  "",
				"admin":         nil,
				"internal":      false,
				"server":        "",
				"context":       "",
				"namespace":     "",
				"kube_user":     "",
				"kube_password": "",
				"ca_cert":       "",
				"client_cert":   "",
				"client_key":    "",
			},
		},
		{
			name: "KubePassword - default",
			args: args{
				in: func() datasources.KubeConfig {
					subject := datasources.KubeConfig{}
					subject.KubePassword = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
				"cluster_name":  "",
				"admin":         nil,
				"internal":      false,
				"server":        "",
				"context":       "",
				"namespace":     "",
				"kube_user":     "",
				"kube_password": "",
				"ca_cert":       "",
				"client_cert":   "",
				"client_key":    "",
			},
		},
		{
			name: "CaCert - default",
			args: args{
				in: func() datasources.KubeConfig {
					subject := datasources.KubeConfig{}
					subject.CaCert = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
				"cluster_name":  "",
				"admin":         nil,
				"internal":      false,
				"server":        "",
				"context":       "",
				"namespace":     "",
				"kube_user":     "",
				"kube_password": "",
				"ca_cert":       "",
				"client_cert":   "",
				"client_key":    "",
			},
		},
		{
			name: "ClientCert - default",
			args: args{
				in: func() datasources.KubeConfig {
					subject := datasources.KubeConfig{}
					subject.ClientCert = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
				"cluster_name":  "",
				"admin":         nil,
				"internal":      false,
				"server":        "",
				"context":       "",
				"namespace":     "",
				"kube_user":     "",
				"kube_password": "",
				"ca_cert":       "",
				"client_cert":   "",
				"client_key":    "",
			},
		},
		{
			name: "ClientKey - default",
			args: args{
				in: func() datasources.KubeConfig {
					subject := datasources.KubeConfig{}
					subject.ClientKey = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
				"cluster_name":  "",
				"admin":         nil,
				"internal":      false,
				"server":        "",
				"context":       "",
				"namespace":     "",
				"kube_user":     "",
				"kube_password": "",
				"ca_cert":       "",
				"client_cert":   "",
				"client_key":    "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenDataSourceKubeConfig(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenDataSourceKubeConfig() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
