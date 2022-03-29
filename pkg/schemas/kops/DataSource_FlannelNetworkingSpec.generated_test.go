package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceFlannelNetworkingSpec(t *testing.T) {
	_default := kops.FlannelNetworkingSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.FlannelNetworkingSpec
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"backend":                 "",
					"iptables_resync_seconds": nil,
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandDataSourceFlannelNetworkingSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceFlannelNetworkingSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceFlannelNetworkingSpecInto(t *testing.T) {
	_default := map[string]interface{}{
		"backend":                 "",
		"iptables_resync_seconds": nil,
	}
	type args struct {
		in kops.FlannelNetworkingSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.FlannelNetworkingSpec{},
			},
			want: _default,
		},
		{
			name: "Backend - default",
			args: args{
				in: func() kops.FlannelNetworkingSpec {
					subject := kops.FlannelNetworkingSpec{}
					subject.Backend = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "IptablesResyncSeconds - default",
			args: args{
				in: func() kops.FlannelNetworkingSpec {
					subject := kops.FlannelNetworkingSpec{}
					subject.IptablesResyncSeconds = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenDataSourceFlannelNetworkingSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceFlannelNetworkingSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceFlannelNetworkingSpec(t *testing.T) {
	_default := map[string]interface{}{
		"backend":                 "",
		"iptables_resync_seconds": nil,
	}
	type args struct {
		in kops.FlannelNetworkingSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.FlannelNetworkingSpec{},
			},
			want: _default,
		},
		{
			name: "Backend - default",
			args: args{
				in: func() kops.FlannelNetworkingSpec {
					subject := kops.FlannelNetworkingSpec{}
					subject.Backend = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "IptablesResyncSeconds - default",
			args: args{
				in: func() kops.FlannelNetworkingSpec {
					subject := kops.FlannelNetworkingSpec{}
					subject.IptablesResyncSeconds = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceFlannelNetworkingSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceFlannelNetworkingSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
