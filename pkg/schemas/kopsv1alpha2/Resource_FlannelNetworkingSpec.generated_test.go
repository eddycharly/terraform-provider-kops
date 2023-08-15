package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

func TestExpandResourceFlannelNetworkingSpec(t *testing.T) {
	_default := kopsv1alpha2.FlannelNetworkingSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kopsv1alpha2.FlannelNetworkingSpec
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"backend":                        "",
					"disable_tx_checksum_offloading": false,
					"iptables_resync_seconds":        nil,
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandResourceFlannelNetworkingSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceFlannelNetworkingSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceFlannelNetworkingSpecInto(t *testing.T) {
	_default := map[string]interface{}{
		"backend":                        "",
		"disable_tx_checksum_offloading": false,
		"iptables_resync_seconds":        nil,
	}
	type args struct {
		in kopsv1alpha2.FlannelNetworkingSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.FlannelNetworkingSpec{},
			},
			want: _default,
		},
		{
			name: "Backend - default",
			args: args{
				in: func() kopsv1alpha2.FlannelNetworkingSpec {
					subject := kopsv1alpha2.FlannelNetworkingSpec{}
					subject.Backend = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "DisableTxChecksumOffloading - default",
			args: args{
				in: func() kopsv1alpha2.FlannelNetworkingSpec {
					subject := kopsv1alpha2.FlannelNetworkingSpec{}
					subject.DisableTxChecksumOffloading = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "IptablesResyncSeconds - default",
			args: args{
				in: func() kopsv1alpha2.FlannelNetworkingSpec {
					subject := kopsv1alpha2.FlannelNetworkingSpec{}
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
			FlattenResourceFlannelNetworkingSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceFlannelNetworkingSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceFlannelNetworkingSpec(t *testing.T) {
	_default := map[string]interface{}{
		"backend":                        "",
		"disable_tx_checksum_offloading": false,
		"iptables_resync_seconds":        nil,
	}
	type args struct {
		in kopsv1alpha2.FlannelNetworkingSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.FlannelNetworkingSpec{},
			},
			want: _default,
		},
		{
			name: "Backend - default",
			args: args{
				in: func() kopsv1alpha2.FlannelNetworkingSpec {
					subject := kopsv1alpha2.FlannelNetworkingSpec{}
					subject.Backend = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "DisableTxChecksumOffloading - default",
			args: args{
				in: func() kopsv1alpha2.FlannelNetworkingSpec {
					subject := kopsv1alpha2.FlannelNetworkingSpec{}
					subject.DisableTxChecksumOffloading = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "IptablesResyncSeconds - default",
			args: args{
				in: func() kopsv1alpha2.FlannelNetworkingSpec {
					subject := kopsv1alpha2.FlannelNetworkingSpec{}
					subject.IptablesResyncSeconds = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceFlannelNetworkingSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceFlannelNetworkingSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
