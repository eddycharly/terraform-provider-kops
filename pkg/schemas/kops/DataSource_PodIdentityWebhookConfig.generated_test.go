package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourcePodIdentityWebhookConfig(t *testing.T) {
	_default := kops.PodIdentityWebhookConfig{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.PodIdentityWebhookConfig
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"enabled":  false,
					"replicas": 0,
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandDataSourcePodIdentityWebhookConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourcePodIdentityWebhookConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourcePodIdentityWebhookConfigInto(t *testing.T) {
	_default := map[string]interface{}{
		"enabled":  false,
		"replicas": 0,
	}
	type args struct {
		in kops.PodIdentityWebhookConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.PodIdentityWebhookConfig{},
			},
			want: _default,
		},
		{
			name: "Enabled - default",
			args: args{
				in: func() kops.PodIdentityWebhookConfig {
					subject := kops.PodIdentityWebhookConfig{}
					subject.Enabled = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Replicas - default",
			args: args{
				in: func() kops.PodIdentityWebhookConfig {
					subject := kops.PodIdentityWebhookConfig{}
					subject.Replicas = 0
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenDataSourcePodIdentityWebhookConfigInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourcePodIdentityWebhookConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourcePodIdentityWebhookConfig(t *testing.T) {
	_default := map[string]interface{}{
		"enabled":  false,
		"replicas": 0,
	}
	type args struct {
		in kops.PodIdentityWebhookConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.PodIdentityWebhookConfig{},
			},
			want: _default,
		},
		{
			name: "Enabled - default",
			args: args{
				in: func() kops.PodIdentityWebhookConfig {
					subject := kops.PodIdentityWebhookConfig{}
					subject.Enabled = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Replicas - default",
			args: args{
				in: func() kops.PodIdentityWebhookConfig {
					subject := kops.PodIdentityWebhookConfig{}
					subject.Replicas = 0
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourcePodIdentityWebhookConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourcePodIdentityWebhookConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
