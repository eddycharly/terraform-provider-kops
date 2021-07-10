package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceRomanaNetworkingSpec(t *testing.T) {
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
			if got := ExpandDataSourceRomanaNetworkingSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandDataSourceRomanaNetworkingSpec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenDataSourceRomanaNetworkingSpecInto(t *testing.T) {
	type args struct {
		in  kops.RomanaNetworkingSpec
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
			FlattenDataSourceRomanaNetworkingSpecInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenDataSourceRomanaNetworkingSpec(t *testing.T) {
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
			want: map[string]interface{}{
				"daemon_service_ip": "",
				"etcd_service_ip":   "",
			},
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
			want: map[string]interface{}{
				"daemon_service_ip": "",
				"etcd_service_ip":   "",
			},
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
			want: map[string]interface{}{
				"daemon_service_ip": "",
				"etcd_service_ip":   "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenDataSourceRomanaNetworkingSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenDataSourceRomanaNetworkingSpec() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
