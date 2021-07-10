package schemas

import (
	"testing"

	"github.com/eddycharly/terraform-provider-kops/pkg/api/datasources"
	"github.com/google/go-cmp/cmp"
)

func TestExpandDataSourceKubeConfig(t *testing.T) {
	_default := datasources.KubeConfig{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want datasources.KubeConfig
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
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
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandDataSourceKubeConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceKubeConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceKubeConfigInto(t *testing.T) {
	_default := map[string]interface{}{
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
	}
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenDataSourceKubeConfigInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceKubeConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceKubeConfig(t *testing.T) {
	_default := map[string]interface{}{
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
	}
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceKubeConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceKubeConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
