package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceAWSSpec(t *testing.T) {
	_default := kops.AWSSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.AWSSpec
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"ebs_csi_driver":           nil,
					"node_termination_handler": nil,
					"load_balancer_controller": nil,
					"pod_identity_webhook":     nil,
					"warm_pool":                nil,
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandResourceAWSSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceAWSSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceAWSSpecInto(t *testing.T) {
	_default := map[string]interface{}{
		"ebs_csi_driver":           nil,
		"node_termination_handler": nil,
		"load_balancer_controller": nil,
		"pod_identity_webhook":     nil,
		"warm_pool":                nil,
	}
	type args struct {
		in kops.AWSSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.AWSSpec{},
			},
			want: _default,
		},
		{
			name: "EbsCsiDriver - default",
			args: args{
				in: func() kops.AWSSpec {
					subject := kops.AWSSpec{}
					subject.EBSCSIDriver = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NodeTerminationHandler - default",
			args: args{
				in: func() kops.AWSSpec {
					subject := kops.AWSSpec{}
					subject.NodeTerminationHandler = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LoadBalancerController - default",
			args: args{
				in: func() kops.AWSSpec {
					subject := kops.AWSSpec{}
					subject.LoadBalancerController = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PodIdentityWebhook - default",
			args: args{
				in: func() kops.AWSSpec {
					subject := kops.AWSSpec{}
					subject.PodIdentityWebhook = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "WarmPool - default",
			args: args{
				in: func() kops.AWSSpec {
					subject := kops.AWSSpec{}
					subject.WarmPool = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenResourceAWSSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceAWSSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceAWSSpec(t *testing.T) {
	_default := map[string]interface{}{
		"ebs_csi_driver":           nil,
		"node_termination_handler": nil,
		"load_balancer_controller": nil,
		"pod_identity_webhook":     nil,
		"warm_pool":                nil,
	}
	type args struct {
		in kops.AWSSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.AWSSpec{},
			},
			want: _default,
		},
		{
			name: "EbsCsiDriver - default",
			args: args{
				in: func() kops.AWSSpec {
					subject := kops.AWSSpec{}
					subject.EBSCSIDriver = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NodeTerminationHandler - default",
			args: args{
				in: func() kops.AWSSpec {
					subject := kops.AWSSpec{}
					subject.NodeTerminationHandler = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LoadBalancerController - default",
			args: args{
				in: func() kops.AWSSpec {
					subject := kops.AWSSpec{}
					subject.LoadBalancerController = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PodIdentityWebhook - default",
			args: args{
				in: func() kops.AWSSpec {
					subject := kops.AWSSpec{}
					subject.PodIdentityWebhook = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "WarmPool - default",
			args: args{
				in: func() kops.AWSSpec {
					subject := kops.AWSSpec{}
					subject.WarmPool = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceAWSSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceAWSSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
