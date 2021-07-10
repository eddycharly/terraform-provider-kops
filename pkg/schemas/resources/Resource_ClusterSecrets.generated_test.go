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
	type args struct {
		in  resources.ClusterSecrets
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
			FlattenResourceClusterSecretsInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenResourceClusterSecrets(t *testing.T) {
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
			want: map[string]interface{}{
				"docker_config": "",
			},
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
			want: map[string]interface{}{
				"docker_config": "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenResourceClusterSecrets(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenResourceClusterSecrets() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
