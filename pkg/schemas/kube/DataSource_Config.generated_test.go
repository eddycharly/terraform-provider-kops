package schemas

import (
	"reflect"
	"testing"

	"github.com/eddycharly/terraform-provider-kops/pkg/api/kube"
	"github.com/google/go-cmp/cmp"
)

func TestExpandDataSourceConfig(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kube.Config
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandDataSourceConfig(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandDataSourceConfig() = %v, want %v", got, tt.want)
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
		"ca_cert":       "",
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
			name: "CaCert - default",
			args: args{
				in: func() kube.Config {
					subject := kube.Config{}
					subject.CaCert = ""
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
		"ca_cert":       "",
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
			name: "CaCert - default",
			args: args{
				in: func() kube.Config {
					subject := kube.Config{}
					subject.CaCert = ""
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
