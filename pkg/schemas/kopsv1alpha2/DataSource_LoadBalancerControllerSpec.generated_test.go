package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

func TestExpandDataSourceLoadBalancerControllerSpec(t *testing.T) {
	_default := kopsv1alpha2.LoadBalancerControllerSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kopsv1alpha2.LoadBalancerControllerSpec
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
			got := ExpandDataSourceLoadBalancerControllerSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceLoadBalancerControllerSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceLoadBalancerControllerSpecInto(t *testing.T) {
	_default := map[string]interface{}{
		"enabled":       nil,
		"version":       nil,
		"enable_waf":    false,
		"enable_wa_fv2": false,
		"enable_shield": false,
	}
	type args struct {
		in kopsv1alpha2.LoadBalancerControllerSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.LoadBalancerControllerSpec{},
			},
			want: _default,
		},
		{
			name: "Enabled - default",
			args: args{
				in: func() kopsv1alpha2.LoadBalancerControllerSpec {
					subject := kopsv1alpha2.LoadBalancerControllerSpec{}
					subject.Enabled = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Version - default",
			args: args{
				in: func() kopsv1alpha2.LoadBalancerControllerSpec {
					subject := kopsv1alpha2.LoadBalancerControllerSpec{}
					subject.Version = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableWAF - default",
			args: args{
				in: func() kopsv1alpha2.LoadBalancerControllerSpec {
					subject := kopsv1alpha2.LoadBalancerControllerSpec{}
					subject.EnableWAF = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableWAFv2 - default",
			args: args{
				in: func() kopsv1alpha2.LoadBalancerControllerSpec {
					subject := kopsv1alpha2.LoadBalancerControllerSpec{}
					subject.EnableWAFv2 = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableShield - default",
			args: args{
				in: func() kopsv1alpha2.LoadBalancerControllerSpec {
					subject := kopsv1alpha2.LoadBalancerControllerSpec{}
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
			FlattenDataSourceLoadBalancerControllerSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceLoadBalancerControllerSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceLoadBalancerControllerSpec(t *testing.T) {
	_default := map[string]interface{}{
		"enabled":       nil,
		"version":       nil,
		"enable_waf":    false,
		"enable_wa_fv2": false,
		"enable_shield": false,
	}
	type args struct {
		in kopsv1alpha2.LoadBalancerControllerSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.LoadBalancerControllerSpec{},
			},
			want: _default,
		},
		{
			name: "Enabled - default",
			args: args{
				in: func() kopsv1alpha2.LoadBalancerControllerSpec {
					subject := kopsv1alpha2.LoadBalancerControllerSpec{}
					subject.Enabled = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Version - default",
			args: args{
				in: func() kopsv1alpha2.LoadBalancerControllerSpec {
					subject := kopsv1alpha2.LoadBalancerControllerSpec{}
					subject.Version = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableWAF - default",
			args: args{
				in: func() kopsv1alpha2.LoadBalancerControllerSpec {
					subject := kopsv1alpha2.LoadBalancerControllerSpec{}
					subject.EnableWAF = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableWAFv2 - default",
			args: args{
				in: func() kopsv1alpha2.LoadBalancerControllerSpec {
					subject := kopsv1alpha2.LoadBalancerControllerSpec{}
					subject.EnableWAFv2 = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableShield - default",
			args: args{
				in: func() kopsv1alpha2.LoadBalancerControllerSpec {
					subject := kopsv1alpha2.LoadBalancerControllerSpec{}
					subject.EnableShield = false
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceLoadBalancerControllerSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceLoadBalancerControllerSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
