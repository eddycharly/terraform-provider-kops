package schemas

import (
	"testing"

	"github.com/eddycharly/terraform-provider-kops/pkg/api/resources"
	"github.com/google/go-cmp/cmp"
)

func TestExpandResourceClusterSecrets(t *testing.T) {
	_default := resources.ClusterSecrets{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want resources.ClusterSecrets
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"docker_config": "",
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandResourceClusterSecrets(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceClusterSecrets() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceClusterSecretsInto(t *testing.T) {
	_default := map[string]interface{}{
		"docker_config": "",
	}
	type args struct {
		in resources.ClusterSecrets
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: resources.ClusterSecrets{},
			},
			want: _default,
		},
		{
			name: "DockerConfig - default",
			args: args{
				in: func() resources.ClusterSecrets {
					subject := resources.ClusterSecrets{}
					subject.DockerConfig = ""
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenResourceClusterSecretsInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceClusterSecrets() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceClusterSecrets(t *testing.T) {
	_default := map[string]interface{}{
		"docker_config": "",
	}
	type args struct {
		in resources.ClusterSecrets
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: resources.ClusterSecrets{},
			},
			want: _default,
		},
		{
			name: "DockerConfig - default",
			args: args{
				in: func() resources.ClusterSecrets {
					subject := resources.ClusterSecrets{}
					subject.DockerConfig = ""
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceClusterSecrets(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceClusterSecrets() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
