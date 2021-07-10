package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceNetworkingSpec(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.NetworkingSpec
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandDataSourceNetworkingSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandDataSourceNetworkingSpec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenDataSourceNetworkingSpecInto(t *testing.T) {
	type args struct {
		in  kops.NetworkingSpec
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
			FlattenDataSourceNetworkingSpecInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenDataSourceNetworkingSpec(t *testing.T) {
	type args struct {
		in kops.NetworkingSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.NetworkingSpec{},
			},
			want: map[string]interface{}{
				"classic":    nil,
				"kubenet":    nil,
				"external":   nil,
				"cni":        nil,
				"kopeio":     nil,
				"weave":      nil,
				"flannel":    nil,
				"calico":     nil,
				"canal":      nil,
				"kuberouter": nil,
				"romana":     nil,
				"amazon_vpc": nil,
				"cilium":     nil,
				"lyft_vpc":   nil,
				"gce":        nil,
			},
		},
		{
			name: "Classic - default",
			args: args{
				in: func() kops.NetworkingSpec {
					subject := kops.NetworkingSpec{}
					subject.Classic = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"classic":    nil,
				"kubenet":    nil,
				"external":   nil,
				"cni":        nil,
				"kopeio":     nil,
				"weave":      nil,
				"flannel":    nil,
				"calico":     nil,
				"canal":      nil,
				"kuberouter": nil,
				"romana":     nil,
				"amazon_vpc": nil,
				"cilium":     nil,
				"lyft_vpc":   nil,
				"gce":        nil,
			},
		},
		{
			name: "Kubenet - default",
			args: args{
				in: func() kops.NetworkingSpec {
					subject := kops.NetworkingSpec{}
					subject.Kubenet = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"classic":    nil,
				"kubenet":    nil,
				"external":   nil,
				"cni":        nil,
				"kopeio":     nil,
				"weave":      nil,
				"flannel":    nil,
				"calico":     nil,
				"canal":      nil,
				"kuberouter": nil,
				"romana":     nil,
				"amazon_vpc": nil,
				"cilium":     nil,
				"lyft_vpc":   nil,
				"gce":        nil,
			},
		},
		{
			name: "External - default",
			args: args{
				in: func() kops.NetworkingSpec {
					subject := kops.NetworkingSpec{}
					subject.External = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"classic":    nil,
				"kubenet":    nil,
				"external":   nil,
				"cni":        nil,
				"kopeio":     nil,
				"weave":      nil,
				"flannel":    nil,
				"calico":     nil,
				"canal":      nil,
				"kuberouter": nil,
				"romana":     nil,
				"amazon_vpc": nil,
				"cilium":     nil,
				"lyft_vpc":   nil,
				"gce":        nil,
			},
		},
		{
			name: "CNI - default",
			args: args{
				in: func() kops.NetworkingSpec {
					subject := kops.NetworkingSpec{}
					subject.CNI = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"classic":    nil,
				"kubenet":    nil,
				"external":   nil,
				"cni":        nil,
				"kopeio":     nil,
				"weave":      nil,
				"flannel":    nil,
				"calico":     nil,
				"canal":      nil,
				"kuberouter": nil,
				"romana":     nil,
				"amazon_vpc": nil,
				"cilium":     nil,
				"lyft_vpc":   nil,
				"gce":        nil,
			},
		},
		{
			name: "Kopeio - default",
			args: args{
				in: func() kops.NetworkingSpec {
					subject := kops.NetworkingSpec{}
					subject.Kopeio = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"classic":    nil,
				"kubenet":    nil,
				"external":   nil,
				"cni":        nil,
				"kopeio":     nil,
				"weave":      nil,
				"flannel":    nil,
				"calico":     nil,
				"canal":      nil,
				"kuberouter": nil,
				"romana":     nil,
				"amazon_vpc": nil,
				"cilium":     nil,
				"lyft_vpc":   nil,
				"gce":        nil,
			},
		},
		{
			name: "Weave - default",
			args: args{
				in: func() kops.NetworkingSpec {
					subject := kops.NetworkingSpec{}
					subject.Weave = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"classic":    nil,
				"kubenet":    nil,
				"external":   nil,
				"cni":        nil,
				"kopeio":     nil,
				"weave":      nil,
				"flannel":    nil,
				"calico":     nil,
				"canal":      nil,
				"kuberouter": nil,
				"romana":     nil,
				"amazon_vpc": nil,
				"cilium":     nil,
				"lyft_vpc":   nil,
				"gce":        nil,
			},
		},
		{
			name: "Flannel - default",
			args: args{
				in: func() kops.NetworkingSpec {
					subject := kops.NetworkingSpec{}
					subject.Flannel = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"classic":    nil,
				"kubenet":    nil,
				"external":   nil,
				"cni":        nil,
				"kopeio":     nil,
				"weave":      nil,
				"flannel":    nil,
				"calico":     nil,
				"canal":      nil,
				"kuberouter": nil,
				"romana":     nil,
				"amazon_vpc": nil,
				"cilium":     nil,
				"lyft_vpc":   nil,
				"gce":        nil,
			},
		},
		{
			name: "Calico - default",
			args: args{
				in: func() kops.NetworkingSpec {
					subject := kops.NetworkingSpec{}
					subject.Calico = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"classic":    nil,
				"kubenet":    nil,
				"external":   nil,
				"cni":        nil,
				"kopeio":     nil,
				"weave":      nil,
				"flannel":    nil,
				"calico":     nil,
				"canal":      nil,
				"kuberouter": nil,
				"romana":     nil,
				"amazon_vpc": nil,
				"cilium":     nil,
				"lyft_vpc":   nil,
				"gce":        nil,
			},
		},
		{
			name: "Canal - default",
			args: args{
				in: func() kops.NetworkingSpec {
					subject := kops.NetworkingSpec{}
					subject.Canal = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"classic":    nil,
				"kubenet":    nil,
				"external":   nil,
				"cni":        nil,
				"kopeio":     nil,
				"weave":      nil,
				"flannel":    nil,
				"calico":     nil,
				"canal":      nil,
				"kuberouter": nil,
				"romana":     nil,
				"amazon_vpc": nil,
				"cilium":     nil,
				"lyft_vpc":   nil,
				"gce":        nil,
			},
		},
		{
			name: "Kuberouter - default",
			args: args{
				in: func() kops.NetworkingSpec {
					subject := kops.NetworkingSpec{}
					subject.Kuberouter = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"classic":    nil,
				"kubenet":    nil,
				"external":   nil,
				"cni":        nil,
				"kopeio":     nil,
				"weave":      nil,
				"flannel":    nil,
				"calico":     nil,
				"canal":      nil,
				"kuberouter": nil,
				"romana":     nil,
				"amazon_vpc": nil,
				"cilium":     nil,
				"lyft_vpc":   nil,
				"gce":        nil,
			},
		},
		{
			name: "Romana - default",
			args: args{
				in: func() kops.NetworkingSpec {
					subject := kops.NetworkingSpec{}
					subject.Romana = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"classic":    nil,
				"kubenet":    nil,
				"external":   nil,
				"cni":        nil,
				"kopeio":     nil,
				"weave":      nil,
				"flannel":    nil,
				"calico":     nil,
				"canal":      nil,
				"kuberouter": nil,
				"romana":     nil,
				"amazon_vpc": nil,
				"cilium":     nil,
				"lyft_vpc":   nil,
				"gce":        nil,
			},
		},
		{
			name: "AmazonVPC - default",
			args: args{
				in: func() kops.NetworkingSpec {
					subject := kops.NetworkingSpec{}
					subject.AmazonVPC = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"classic":    nil,
				"kubenet":    nil,
				"external":   nil,
				"cni":        nil,
				"kopeio":     nil,
				"weave":      nil,
				"flannel":    nil,
				"calico":     nil,
				"canal":      nil,
				"kuberouter": nil,
				"romana":     nil,
				"amazon_vpc": nil,
				"cilium":     nil,
				"lyft_vpc":   nil,
				"gce":        nil,
			},
		},
		{
			name: "Cilium - default",
			args: args{
				in: func() kops.NetworkingSpec {
					subject := kops.NetworkingSpec{}
					subject.Cilium = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"classic":    nil,
				"kubenet":    nil,
				"external":   nil,
				"cni":        nil,
				"kopeio":     nil,
				"weave":      nil,
				"flannel":    nil,
				"calico":     nil,
				"canal":      nil,
				"kuberouter": nil,
				"romana":     nil,
				"amazon_vpc": nil,
				"cilium":     nil,
				"lyft_vpc":   nil,
				"gce":        nil,
			},
		},
		{
			name: "LyftVPC - default",
			args: args{
				in: func() kops.NetworkingSpec {
					subject := kops.NetworkingSpec{}
					subject.LyftVPC = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"classic":    nil,
				"kubenet":    nil,
				"external":   nil,
				"cni":        nil,
				"kopeio":     nil,
				"weave":      nil,
				"flannel":    nil,
				"calico":     nil,
				"canal":      nil,
				"kuberouter": nil,
				"romana":     nil,
				"amazon_vpc": nil,
				"cilium":     nil,
				"lyft_vpc":   nil,
				"gce":        nil,
			},
		},
		{
			name: "GCE - default",
			args: args{
				in: func() kops.NetworkingSpec {
					subject := kops.NetworkingSpec{}
					subject.GCE = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"classic":    nil,
				"kubenet":    nil,
				"external":   nil,
				"cni":        nil,
				"kopeio":     nil,
				"weave":      nil,
				"flannel":    nil,
				"calico":     nil,
				"canal":      nil,
				"kuberouter": nil,
				"romana":     nil,
				"amazon_vpc": nil,
				"cilium":     nil,
				"lyft_vpc":   nil,
				"gce":        nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenDataSourceNetworkingSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenDataSourceNetworkingSpec() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
