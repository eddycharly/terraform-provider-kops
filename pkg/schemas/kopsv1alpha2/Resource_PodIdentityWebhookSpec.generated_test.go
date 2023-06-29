package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

func TestExpandResourcePodIdentityWebhookSpec(t *testing.T) {
	_default := kopsv1alpha2.PodIdentityWebhookSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kopsv1alpha2.PodIdentityWebhookSpec
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
			got := ExpandResourcePodIdentityWebhookSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourcePodIdentityWebhookSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourcePodIdentityWebhookSpecInto(t *testing.T) {
	_default := map[string]interface{}{
		"enabled":  false,
		"replicas": 0,
	}
	type args struct {
		in kopsv1alpha2.PodIdentityWebhookSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.PodIdentityWebhookSpec{},
			},
			want: _default,
		},
		{
			name: "Enabled - default",
			args: args{
				in: func() kopsv1alpha2.PodIdentityWebhookSpec {
					subject := kopsv1alpha2.PodIdentityWebhookSpec{}
					subject.Enabled = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Replicas - default",
			args: args{
				in: func() kopsv1alpha2.PodIdentityWebhookSpec {
					subject := kopsv1alpha2.PodIdentityWebhookSpec{}
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
			FlattenResourcePodIdentityWebhookSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourcePodIdentityWebhookSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourcePodIdentityWebhookSpec(t *testing.T) {
	_default := map[string]interface{}{
		"enabled":  false,
		"replicas": 0,
	}
	type args struct {
		in kopsv1alpha2.PodIdentityWebhookSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.PodIdentityWebhookSpec{},
			},
			want: _default,
		},
		{
			name: "Enabled - default",
			args: args{
				in: func() kopsv1alpha2.PodIdentityWebhookSpec {
					subject := kopsv1alpha2.PodIdentityWebhookSpec{}
					subject.Enabled = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Replicas - default",
			args: args{
				in: func() kopsv1alpha2.PodIdentityWebhookSpec {
					subject := kopsv1alpha2.PodIdentityWebhookSpec{}
					subject.Replicas = 0
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourcePodIdentityWebhookSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourcePodIdentityWebhookSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
