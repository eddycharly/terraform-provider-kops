package schemas

import (
	"reflect"
	"testing"

	"github.com/eddycharly/terraform-provider-kops/pkg/api/resources"
	"github.com/google/go-cmp/cmp"
)

func TestExpandResourceClusterSecrets(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want resources.ClusterSecrets
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandResourceClusterSecrets(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandResourceClusterSecrets() = %v, want %v", got, tt.want)
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
