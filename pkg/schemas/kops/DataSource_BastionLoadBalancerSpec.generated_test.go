package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceBastionLoadBalancerSpec(t *testing.T) {
	_default := kops.BastionLoadBalancerSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.BastionLoadBalancerSpec
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"type": "",
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandDataSourceBastionLoadBalancerSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceBastionLoadBalancerSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceBastionLoadBalancerSpecInto(t *testing.T) {
	_default := map[string]interface{}{
		"type": "",
	}
	type args struct {
		in kops.BastionLoadBalancerSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.BastionLoadBalancerSpec{},
			},
			want: _default,
		},
		{
			name: "Type - default",
			args: args{
				in: func() kops.BastionLoadBalancerSpec {
					subject := kops.BastionLoadBalancerSpec{}
					subject.Type = ""
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenDataSourceBastionLoadBalancerSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceBastionLoadBalancerSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceBastionLoadBalancerSpec(t *testing.T) {
	_default := map[string]interface{}{
		"type": "",
	}
	type args struct {
		in kops.BastionLoadBalancerSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.BastionLoadBalancerSpec{},
			},
			want: _default,
		},
		{
			name: "Type - default",
			args: args{
				in: func() kops.BastionLoadBalancerSpec {
					subject := kops.BastionLoadBalancerSpec{}
					subject.Type = ""
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceBastionLoadBalancerSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceBastionLoadBalancerSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
