package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceFlannelNetworkingSpec(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.FlannelNetworkingSpec
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandResourceFlannelNetworkingSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandResourceFlannelNetworkingSpec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenResourceFlannelNetworkingSpecInto(t *testing.T) {
	type args struct {
		in  kops.FlannelNetworkingSpec
		out map[string]interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			FlattenResourceFlannelNetworkingSpecInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenResourceFlannelNetworkingSpec(t *testing.T) {
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
			want: map[string]interface{}{
				"backend":                        "",
				"disable_tx_checksum_offloading": false,
				"iptables_resync_seconds":        nil,
			},
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
			want: map[string]interface{}{
				"backend":                        "",
				"disable_tx_checksum_offloading": false,
				"iptables_resync_seconds":        nil,
			},
		},
		{
			name: "DisableTxChecksumOffloading - default",
			args: args{
				in: func() kops.FlannelNetworkingSpec {
					subject := kops.FlannelNetworkingSpec{}
					subject.DisableTxChecksumOffloading = false
					return subject
				}(),
			},
			want: map[string]interface{}{
				"backend":                        "",
				"disable_tx_checksum_offloading": false,
				"iptables_resync_seconds":        nil,
			},
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
			want: map[string]interface{}{
				"backend":                        "",
				"disable_tx_checksum_offloading": false,
				"iptables_resync_seconds":        nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenResourceFlannelNetworkingSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenResourceFlannelNetworkingSpec() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
