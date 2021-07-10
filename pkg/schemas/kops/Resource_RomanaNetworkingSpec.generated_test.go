package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceRomanaNetworkingSpec(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.RomanaNetworkingSpec
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandResourceRomanaNetworkingSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandResourceRomanaNetworkingSpec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenResourceRomanaNetworkingSpecInto(t *testing.T) {
	_default := map[string]interface{}{
		"daemon_service_ip": "",
		"etcd_service_ip":   "",
	}
	type args struct {
		in kops.RomanaNetworkingSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.RomanaNetworkingSpec{},
			},
			want: _default,
		},
		{
			name: "DaemonServiceIp - default",
			args: args{
				in: func() kops.RomanaNetworkingSpec {
					subject := kops.RomanaNetworkingSpec{}
					subject.DaemonServiceIP = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EtcdServiceIp - default",
			args: args{
				in: func() kops.RomanaNetworkingSpec {
					subject := kops.RomanaNetworkingSpec{}
					subject.EtcdServiceIP = ""
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenResourceRomanaNetworkingSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceRomanaNetworkingSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceRomanaNetworkingSpec(t *testing.T) {
	_default := map[string]interface{}{
		"daemon_service_ip": "",
		"etcd_service_ip":   "",
	}
	type args struct {
		in kops.RomanaNetworkingSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.RomanaNetworkingSpec{},
			},
			want: _default,
		},
		{
			name: "DaemonServiceIp - default",
			args: args{
				in: func() kops.RomanaNetworkingSpec {
					subject := kops.RomanaNetworkingSpec{}
					subject.DaemonServiceIP = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EtcdServiceIp - default",
			args: args{
				in: func() kops.RomanaNetworkingSpec {
					subject := kops.RomanaNetworkingSpec{}
					subject.EtcdServiceIP = ""
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceRomanaNetworkingSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceRomanaNetworkingSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
