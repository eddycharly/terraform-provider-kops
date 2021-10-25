package schemas

import (
	"testing"

	"github.com/eddycharly/terraform-provider-kops/pkg/api/kube"
	"github.com/google/go-cmp/cmp"
)

func TestExpandDataSourceConfig(t *testing.T) {
	_default := kube.Config{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kube.Config
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"server":        "",
					"context":       "",
					"namespace":     "",
					"kube_user":     "",
					"kube_password": "",
					"ca_certs":      "",
					"client_cert":   "",
					"client_key":    "",
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandDataSourceConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceConfigInto(t *testing.T) {
	_default := map[string]interface{}{
		"server":        "",
		"context":       "",
		"namespace":     "",
		"kube_user":     "",
		"kube_password": "",
		"ca_certs":      "",
		"client_cert":   "",
		"client_key":    "",
	}
	type args struct {
		in kube.Config
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kube.Config{},
			},
			want: _default,
		},
		{
			name: "Server - default",
			args: args{
				in: func() kube.Config {
					subject := kube.Config{}
					subject.Server = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Context - default",
			args: args{
				in: func() kube.Config {
					subject := kube.Config{}
					subject.Context = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Namespace - default",
			args: args{
				in: func() kube.Config {
					subject := kube.Config{}
					subject.Namespace = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "KubeUser - default",
			args: args{
				in: func() kube.Config {
					subject := kube.Config{}
					subject.KubeUser = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "KubePassword - default",
			args: args{
				in: func() kube.Config {
					subject := kube.Config{}
					subject.KubePassword = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CaCerts - default",
			args: args{
				in: func() kube.Config {
					subject := kube.Config{}
					subject.CaCerts = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ClientCert - default",
			args: args{
				in: func() kube.Config {
					subject := kube.Config{}
					subject.ClientCert = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ClientKey - default",
			args: args{
				in: func() kube.Config {
					subject := kube.Config{}
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
			FlattenDataSourceConfigInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceConfig(t *testing.T) {
	_default := map[string]interface{}{
		"server":        "",
		"context":       "",
		"namespace":     "",
		"kube_user":     "",
		"kube_password": "",
		"ca_certs":      "",
		"client_cert":   "",
		"client_key":    "",
	}
	type args struct {
		in kube.Config
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kube.Config{},
			},
			want: _default,
		},
		{
			name: "Server - default",
			args: args{
				in: func() kube.Config {
					subject := kube.Config{}
					subject.Server = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Context - default",
			args: args{
				in: func() kube.Config {
					subject := kube.Config{}
					subject.Context = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Namespace - default",
			args: args{
				in: func() kube.Config {
					subject := kube.Config{}
					subject.Namespace = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "KubeUser - default",
			args: args{
				in: func() kube.Config {
					subject := kube.Config{}
					subject.KubeUser = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "KubePassword - default",
			args: args{
				in: func() kube.Config {
					subject := kube.Config{}
					subject.KubePassword = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CaCerts - default",
			args: args{
				in: func() kube.Config {
					subject := kube.Config{}
					subject.CaCerts = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ClientCert - default",
			args: args{
				in: func() kube.Config {
					subject := kube.Config{}
					subject.ClientCert = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ClientKey - default",
			args: args{
				in: func() kube.Config {
					subject := kube.Config{}
					subject.ClientKey = ""
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
