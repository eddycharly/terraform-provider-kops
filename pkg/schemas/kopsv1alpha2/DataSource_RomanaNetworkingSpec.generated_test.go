package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

func TestExpandDataSourceRomanaNetworkingSpec(t *testing.T) {
	_default := kopsv1alpha2.RomanaNetworkingSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kopsv1alpha2.RomanaNetworkingSpec
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"daemon_service_ip": "",
					"etcd_service_ip":   "",
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandDataSourceRomanaNetworkingSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceRomanaNetworkingSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceRomanaNetworkingSpecInto(t *testing.T) {
	_default := map[string]interface{}{
		"daemon_service_ip": "",
		"etcd_service_ip":   "",
	}
	type args struct {
		in kopsv1alpha2.RomanaNetworkingSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.RomanaNetworkingSpec{},
			},
			want: _default,
		},
		{
			name: "DaemonServiceIp - default",
			args: args{
				in: func() kopsv1alpha2.RomanaNetworkingSpec {
					subject := kopsv1alpha2.RomanaNetworkingSpec{}
					subject.DaemonServiceIP = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EtcdServiceIp - default",
			args: args{
				in: func() kopsv1alpha2.RomanaNetworkingSpec {
					subject := kopsv1alpha2.RomanaNetworkingSpec{}
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
			FlattenDataSourceRomanaNetworkingSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceRomanaNetworkingSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceRomanaNetworkingSpec(t *testing.T) {
	_default := map[string]interface{}{
		"daemon_service_ip": "",
		"etcd_service_ip":   "",
	}
	type args struct {
		in kopsv1alpha2.RomanaNetworkingSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.RomanaNetworkingSpec{},
			},
			want: _default,
		},
		{
			name: "DaemonServiceIp - default",
			args: args{
				in: func() kopsv1alpha2.RomanaNetworkingSpec {
					subject := kopsv1alpha2.RomanaNetworkingSpec{}
					subject.DaemonServiceIP = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EtcdServiceIp - default",
			args: args{
				in: func() kopsv1alpha2.RomanaNetworkingSpec {
					subject := kopsv1alpha2.RomanaNetworkingSpec{}
					subject.EtcdServiceIP = ""
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceRomanaNetworkingSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceRomanaNetworkingSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
