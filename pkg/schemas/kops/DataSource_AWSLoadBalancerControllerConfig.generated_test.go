package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceAWSLoadBalancerControllerConfig(t *testing.T) {
	_default := kops.AWSLoadBalancerControllerConfig{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.AWSLoadBalancerControllerConfig
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"enabled":       nil,
					"version":       nil,
					"enable_waf":    false,
					"enable_wa_fv2": false,
					"enable_shield": false,
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandDataSourceAWSLoadBalancerControllerConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceAWSLoadBalancerControllerConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceAWSLoadBalancerControllerConfigInto(t *testing.T) {
	_default := map[string]interface{}{
		"enabled":       nil,
		"version":       nil,
		"enable_waf":    false,
		"enable_wa_fv2": false,
		"enable_shield": false,
	}
	type args struct {
		in kops.AWSLoadBalancerControllerConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.AWSLoadBalancerControllerConfig{},
			},
			want: _default,
		},
		{
			name: "Enabled - default",
			args: args{
				in: func() kops.AWSLoadBalancerControllerConfig {
					subject := kops.AWSLoadBalancerControllerConfig{}
					subject.Enabled = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Version - default",
			args: args{
				in: func() kops.AWSLoadBalancerControllerConfig {
					subject := kops.AWSLoadBalancerControllerConfig{}
					subject.Version = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableWAF - default",
			args: args{
				in: func() kops.AWSLoadBalancerControllerConfig {
					subject := kops.AWSLoadBalancerControllerConfig{}
					subject.EnableWAF = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableWAFv2 - default",
			args: args{
				in: func() kops.AWSLoadBalancerControllerConfig {
					subject := kops.AWSLoadBalancerControllerConfig{}
					subject.EnableWAFv2 = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableShield - default",
			args: args{
				in: func() kops.AWSLoadBalancerControllerConfig {
					subject := kops.AWSLoadBalancerControllerConfig{}
					subject.EnableShield = false
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenDataSourceAWSLoadBalancerControllerConfigInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceAWSLoadBalancerControllerConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceAWSLoadBalancerControllerConfig(t *testing.T) {
	_default := map[string]interface{}{
		"enabled":       nil,
		"version":       nil,
		"enable_waf":    false,
		"enable_wa_fv2": false,
		"enable_shield": false,
	}
	type args struct {
		in kops.AWSLoadBalancerControllerConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.AWSLoadBalancerControllerConfig{},
			},
			want: _default,
		},
		{
			name: "Enabled - default",
			args: args{
				in: func() kops.AWSLoadBalancerControllerConfig {
					subject := kops.AWSLoadBalancerControllerConfig{}
					subject.Enabled = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Version - default",
			args: args{
				in: func() kops.AWSLoadBalancerControllerConfig {
					subject := kops.AWSLoadBalancerControllerConfig{}
					subject.Version = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableWAF - default",
			args: args{
				in: func() kops.AWSLoadBalancerControllerConfig {
					subject := kops.AWSLoadBalancerControllerConfig{}
					subject.EnableWAF = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableWAFv2 - default",
			args: args{
				in: func() kops.AWSLoadBalancerControllerConfig {
					subject := kops.AWSLoadBalancerControllerConfig{}
					subject.EnableWAFv2 = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableShield - default",
			args: args{
				in: func() kops.AWSLoadBalancerControllerConfig {
					subject := kops.AWSLoadBalancerControllerConfig{}
					subject.EnableShield = false
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceAWSLoadBalancerControllerConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceAWSLoadBalancerControllerConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
