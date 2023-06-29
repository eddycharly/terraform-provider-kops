package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

func TestExpandResourceNetworkingSpec(t *testing.T) {
	_default := kopsv1alpha2.NetworkingSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kopsv1alpha2.NetworkingSpec
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"network_id":               "",
					"network_cidr":             "",
					"additional_network_cidrs": func() []interface{} { return nil }(),
					"subnet":                   func() []interface{} { return nil }(),
					"tag_subnets":              nil,
					"topology":                 nil,
					"egress_proxy":             nil,
					"non_masquerade_cidr":      "",
					"pod_cidr":                 "",
					"service_cluster_ip_range": "",
					"isolate_control_plane":    nil,
					"classic":                  nil,
					"kubenet":                  nil,
					"external":                 nil,
					"cni":                      nil,
					"kopeio":                   nil,
					"weave":                    nil,
					"flannel":                  nil,
					"calico":                   nil,
					"canal":                    nil,
					"kube_router":              nil,
					"romana":                   nil,
					"amazon_vpc":               nil,
					"cilium":                   nil,
					"lyft_vpc":                 nil,
					"gce":                      nil,
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandResourceNetworkingSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceNetworkingSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceNetworkingSpecInto(t *testing.T) {
	_default := map[string]interface{}{
		"network_id":               "",
		"network_cidr":             "",
		"additional_network_cidrs": func() []interface{} { return nil }(),
		"subnet":                   func() []interface{} { return nil }(),
		"tag_subnets":              nil,
		"topology":                 nil,
		"egress_proxy":             nil,
		"non_masquerade_cidr":      "",
		"pod_cidr":                 "",
		"service_cluster_ip_range": "",
		"isolate_control_plane":    nil,
		"classic":                  nil,
		"kubenet":                  nil,
		"external":                 nil,
		"cni":                      nil,
		"kopeio":                   nil,
		"weave":                    nil,
		"flannel":                  nil,
		"calico":                   nil,
		"canal":                    nil,
		"kube_router":              nil,
		"romana":                   nil,
		"amazon_vpc":               nil,
		"cilium":                   nil,
		"lyft_vpc":                 nil,
		"gce":                      nil,
	}
	type args struct {
		in kopsv1alpha2.NetworkingSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.NetworkingSpec{},
			},
			want: _default,
		},
		{
			name: "NetworkId - default",
			args: args{
				in: func() kopsv1alpha2.NetworkingSpec {
					subject := kopsv1alpha2.NetworkingSpec{}
					subject.NetworkID = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NetworkCidr - default",
			args: args{
				in: func() kopsv1alpha2.NetworkingSpec {
					subject := kopsv1alpha2.NetworkingSpec{}
					subject.NetworkCIDR = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AdditionalNetworkCidrs - default",
			args: args{
				in: func() kopsv1alpha2.NetworkingSpec {
					subject := kopsv1alpha2.NetworkingSpec{}
					subject.AdditionalNetworkCIDRs = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Subnet - default",
			args: args{
				in: func() kopsv1alpha2.NetworkingSpec {
					subject := kopsv1alpha2.NetworkingSpec{}
					subject.Subnets = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "TagSubnets - default",
			args: args{
				in: func() kopsv1alpha2.NetworkingSpec {
					subject := kopsv1alpha2.NetworkingSpec{}
					subject.TagSubnets = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Topology - default",
			args: args{
				in: func() kopsv1alpha2.NetworkingSpec {
					subject := kopsv1alpha2.NetworkingSpec{}
					subject.Topology = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EgressProxy - default",
			args: args{
				in: func() kopsv1alpha2.NetworkingSpec {
					subject := kopsv1alpha2.NetworkingSpec{}
					subject.EgressProxy = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NonMasqueradeCidr - default",
			args: args{
				in: func() kopsv1alpha2.NetworkingSpec {
					subject := kopsv1alpha2.NetworkingSpec{}
					subject.NonMasqueradeCIDR = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PodCidr - default",
			args: args{
				in: func() kopsv1alpha2.NetworkingSpec {
					subject := kopsv1alpha2.NetworkingSpec{}
					subject.PodCIDR = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ServiceClusterIpRange - default",
			args: args{
				in: func() kopsv1alpha2.NetworkingSpec {
					subject := kopsv1alpha2.NetworkingSpec{}
					subject.ServiceClusterIPRange = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "IsolateControlPlane - default",
			args: args{
				in: func() kopsv1alpha2.NetworkingSpec {
					subject := kopsv1alpha2.NetworkingSpec{}
					subject.IsolateControlPlane = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Classic - default",
			args: args{
				in: func() kopsv1alpha2.NetworkingSpec {
					subject := kopsv1alpha2.NetworkingSpec{}
					subject.Classic = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Kubenet - default",
			args: args{
				in: func() kopsv1alpha2.NetworkingSpec {
					subject := kopsv1alpha2.NetworkingSpec{}
					subject.Kubenet = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "External - default",
			args: args{
				in: func() kopsv1alpha2.NetworkingSpec {
					subject := kopsv1alpha2.NetworkingSpec{}
					subject.External = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CNI - default",
			args: args{
				in: func() kopsv1alpha2.NetworkingSpec {
					subject := kopsv1alpha2.NetworkingSpec{}
					subject.CNI = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Kopeio - default",
			args: args{
				in: func() kopsv1alpha2.NetworkingSpec {
					subject := kopsv1alpha2.NetworkingSpec{}
					subject.Kopeio = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Weave - default",
			args: args{
				in: func() kopsv1alpha2.NetworkingSpec {
					subject := kopsv1alpha2.NetworkingSpec{}
					subject.Weave = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Flannel - default",
			args: args{
				in: func() kopsv1alpha2.NetworkingSpec {
					subject := kopsv1alpha2.NetworkingSpec{}
					subject.Flannel = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Calico - default",
			args: args{
				in: func() kopsv1alpha2.NetworkingSpec {
					subject := kopsv1alpha2.NetworkingSpec{}
					subject.Calico = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Canal - default",
			args: args{
				in: func() kopsv1alpha2.NetworkingSpec {
					subject := kopsv1alpha2.NetworkingSpec{}
					subject.Canal = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "KubeRouter - default",
			args: args{
				in: func() kopsv1alpha2.NetworkingSpec {
					subject := kopsv1alpha2.NetworkingSpec{}
					subject.KubeRouter = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Romana - default",
			args: args{
				in: func() kopsv1alpha2.NetworkingSpec {
					subject := kopsv1alpha2.NetworkingSpec{}
					subject.Romana = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AmazonVPC - default",
			args: args{
				in: func() kopsv1alpha2.NetworkingSpec {
					subject := kopsv1alpha2.NetworkingSpec{}
					subject.AmazonVPC = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Cilium - default",
			args: args{
				in: func() kopsv1alpha2.NetworkingSpec {
					subject := kopsv1alpha2.NetworkingSpec{}
					subject.Cilium = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LyftVPC - default",
			args: args{
				in: func() kopsv1alpha2.NetworkingSpec {
					subject := kopsv1alpha2.NetworkingSpec{}
					subject.LyftVPC = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "GCE - default",
			args: args{
				in: func() kopsv1alpha2.NetworkingSpec {
					subject := kopsv1alpha2.NetworkingSpec{}
					subject.GCE = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenResourceNetworkingSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceNetworkingSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceNetworkingSpec(t *testing.T) {
	_default := map[string]interface{}{
		"network_id":               "",
		"network_cidr":             "",
		"additional_network_cidrs": func() []interface{} { return nil }(),
		"subnet":                   func() []interface{} { return nil }(),
		"tag_subnets":              nil,
		"topology":                 nil,
		"egress_proxy":             nil,
		"non_masquerade_cidr":      "",
		"pod_cidr":                 "",
		"service_cluster_ip_range": "",
		"isolate_control_plane":    nil,
		"classic":                  nil,
		"kubenet":                  nil,
		"external":                 nil,
		"cni":                      nil,
		"kopeio":                   nil,
		"weave":                    nil,
		"flannel":                  nil,
		"calico":                   nil,
		"canal":                    nil,
		"kube_router":              nil,
		"romana":                   nil,
		"amazon_vpc":               nil,
		"cilium":                   nil,
		"lyft_vpc":                 nil,
		"gce":                      nil,
	}
	type args struct {
		in kopsv1alpha2.NetworkingSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.NetworkingSpec{},
			},
			want: _default,
		},
		{
			name: "NetworkId - default",
			args: args{
				in: func() kopsv1alpha2.NetworkingSpec {
					subject := kopsv1alpha2.NetworkingSpec{}
					subject.NetworkID = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NetworkCidr - default",
			args: args{
				in: func() kopsv1alpha2.NetworkingSpec {
					subject := kopsv1alpha2.NetworkingSpec{}
					subject.NetworkCIDR = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AdditionalNetworkCidrs - default",
			args: args{
				in: func() kopsv1alpha2.NetworkingSpec {
					subject := kopsv1alpha2.NetworkingSpec{}
					subject.AdditionalNetworkCIDRs = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Subnet - default",
			args: args{
				in: func() kopsv1alpha2.NetworkingSpec {
					subject := kopsv1alpha2.NetworkingSpec{}
					subject.Subnets = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "TagSubnets - default",
			args: args{
				in: func() kopsv1alpha2.NetworkingSpec {
					subject := kopsv1alpha2.NetworkingSpec{}
					subject.TagSubnets = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Topology - default",
			args: args{
				in: func() kopsv1alpha2.NetworkingSpec {
					subject := kopsv1alpha2.NetworkingSpec{}
					subject.Topology = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EgressProxy - default",
			args: args{
				in: func() kopsv1alpha2.NetworkingSpec {
					subject := kopsv1alpha2.NetworkingSpec{}
					subject.EgressProxy = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NonMasqueradeCidr - default",
			args: args{
				in: func() kopsv1alpha2.NetworkingSpec {
					subject := kopsv1alpha2.NetworkingSpec{}
					subject.NonMasqueradeCIDR = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PodCidr - default",
			args: args{
				in: func() kopsv1alpha2.NetworkingSpec {
					subject := kopsv1alpha2.NetworkingSpec{}
					subject.PodCIDR = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ServiceClusterIpRange - default",
			args: args{
				in: func() kopsv1alpha2.NetworkingSpec {
					subject := kopsv1alpha2.NetworkingSpec{}
					subject.ServiceClusterIPRange = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "IsolateControlPlane - default",
			args: args{
				in: func() kopsv1alpha2.NetworkingSpec {
					subject := kopsv1alpha2.NetworkingSpec{}
					subject.IsolateControlPlane = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Classic - default",
			args: args{
				in: func() kopsv1alpha2.NetworkingSpec {
					subject := kopsv1alpha2.NetworkingSpec{}
					subject.Classic = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Kubenet - default",
			args: args{
				in: func() kopsv1alpha2.NetworkingSpec {
					subject := kopsv1alpha2.NetworkingSpec{}
					subject.Kubenet = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "External - default",
			args: args{
				in: func() kopsv1alpha2.NetworkingSpec {
					subject := kopsv1alpha2.NetworkingSpec{}
					subject.External = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CNI - default",
			args: args{
				in: func() kopsv1alpha2.NetworkingSpec {
					subject := kopsv1alpha2.NetworkingSpec{}
					subject.CNI = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Kopeio - default",
			args: args{
				in: func() kopsv1alpha2.NetworkingSpec {
					subject := kopsv1alpha2.NetworkingSpec{}
					subject.Kopeio = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Weave - default",
			args: args{
				in: func() kopsv1alpha2.NetworkingSpec {
					subject := kopsv1alpha2.NetworkingSpec{}
					subject.Weave = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Flannel - default",
			args: args{
				in: func() kopsv1alpha2.NetworkingSpec {
					subject := kopsv1alpha2.NetworkingSpec{}
					subject.Flannel = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Calico - default",
			args: args{
				in: func() kopsv1alpha2.NetworkingSpec {
					subject := kopsv1alpha2.NetworkingSpec{}
					subject.Calico = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Canal - default",
			args: args{
				in: func() kopsv1alpha2.NetworkingSpec {
					subject := kopsv1alpha2.NetworkingSpec{}
					subject.Canal = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "KubeRouter - default",
			args: args{
				in: func() kopsv1alpha2.NetworkingSpec {
					subject := kopsv1alpha2.NetworkingSpec{}
					subject.KubeRouter = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Romana - default",
			args: args{
				in: func() kopsv1alpha2.NetworkingSpec {
					subject := kopsv1alpha2.NetworkingSpec{}
					subject.Romana = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AmazonVPC - default",
			args: args{
				in: func() kopsv1alpha2.NetworkingSpec {
					subject := kopsv1alpha2.NetworkingSpec{}
					subject.AmazonVPC = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Cilium - default",
			args: args{
				in: func() kopsv1alpha2.NetworkingSpec {
					subject := kopsv1alpha2.NetworkingSpec{}
					subject.Cilium = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LyftVPC - default",
			args: args{
				in: func() kopsv1alpha2.NetworkingSpec {
					subject := kopsv1alpha2.NetworkingSpec{}
					subject.LyftVPC = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "GCE - default",
			args: args{
				in: func() kopsv1alpha2.NetworkingSpec {
					subject := kopsv1alpha2.NetworkingSpec{}
					subject.GCE = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceNetworkingSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceNetworkingSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
