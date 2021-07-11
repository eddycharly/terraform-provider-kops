package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceNetworkingSpec(t *testing.T) {
	_default := kops.NetworkingSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.NetworkingSpec
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
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
	}
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
	}
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
