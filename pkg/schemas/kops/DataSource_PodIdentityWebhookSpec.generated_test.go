package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourcePodIdentityWebhookSpec(t *testing.T) {
	_default := kops.PodIdentityWebhookSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.PodIdentityWebhookSpec
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
			got := ExpandDataSourcePodIdentityWebhookSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourcePodIdentityWebhookSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourcePodIdentityWebhookSpecInto(t *testing.T) {
	_default := map[string]interface{}{
		"enabled":  false,
		"replicas": 0,
	}
	type args struct {
		in kops.PodIdentityWebhookSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.PodIdentityWebhookSpec{},
			},
			want: _default,
		},
		{
			name: "Enabled - default",
			args: args{
				in: func() kops.PodIdentityWebhookSpec {
					subject := kops.PodIdentityWebhookSpec{}
					subject.Enabled = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Replicas - default",
			args: args{
				in: func() kops.PodIdentityWebhookSpec {
					subject := kops.PodIdentityWebhookSpec{}
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
			FlattenDataSourcePodIdentityWebhookSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourcePodIdentityWebhookSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourcePodIdentityWebhookSpec(t *testing.T) {
	_default := map[string]interface{}{
		"enabled":  false,
		"replicas": 0,
	}
	type args struct {
		in kops.PodIdentityWebhookSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.PodIdentityWebhookSpec{},
			},
			want: _default,
		},
		{
			name: "Enabled - default",
			args: args{
				in: func() kops.PodIdentityWebhookSpec {
					subject := kops.PodIdentityWebhookSpec{}
					subject.Enabled = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Replicas - default",
			args: args{
				in: func() kops.PodIdentityWebhookSpec {
					subject := kops.PodIdentityWebhookSpec{}
					subject.Replicas = 0
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourcePodIdentityWebhookSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourcePodIdentityWebhookSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
