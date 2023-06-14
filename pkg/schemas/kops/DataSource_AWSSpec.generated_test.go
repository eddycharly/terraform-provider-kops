package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceAWSSpec(t *testing.T) {
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
			got := ExpandDataSourceAWSSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceAWSSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceAWSSpecInto(t *testing.T) {
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
			FlattenDataSourceAWSSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceAWSSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceAWSSpec(t *testing.T) {
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
			got := FlattenDataSourceAWSSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceAWSSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
