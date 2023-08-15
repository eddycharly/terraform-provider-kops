package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

func TestExpandDataSourceOpenstackNetwork(t *testing.T) {
	_default := kopsv1alpha2.OpenstackNetwork{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kopsv1alpha2.OpenstackNetwork
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"availability_zone_hints": func() []interface{} { return nil }(),
					"ipv6_support_disabled":   nil,
					"public_network_names":    func() []interface{} { return nil }(),
					"internal_network_names":  func() []interface{} { return nil }(),
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandDataSourceOpenstackNetwork(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceOpenstackNetwork() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceOpenstackNetworkInto(t *testing.T) {
	_default := map[string]interface{}{
		"availability_zone_hints": func() []interface{} { return nil }(),
		"ipv6_support_disabled":   nil,
		"public_network_names":    func() []interface{} { return nil }(),
		"internal_network_names":  func() []interface{} { return nil }(),
	}
	type args struct {
		in kopsv1alpha2.OpenstackNetwork
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.OpenstackNetwork{},
			},
			want: _default,
		},
		{
			name: "AvailabilityZoneHints - default",
			args: args{
				in: func() kopsv1alpha2.OpenstackNetwork {
					subject := kopsv1alpha2.OpenstackNetwork{}
					subject.AvailabilityZoneHints = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Ipv6SupportDisabled - default",
			args: args{
				in: func() kopsv1alpha2.OpenstackNetwork {
					subject := kopsv1alpha2.OpenstackNetwork{}
					subject.IPv6SupportDisabled = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PublicNetworkNames - default",
			args: args{
				in: func() kopsv1alpha2.OpenstackNetwork {
					subject := kopsv1alpha2.OpenstackNetwork{}
					subject.PublicNetworkNames = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "InternalNetworkNames - default",
			args: args{
				in: func() kopsv1alpha2.OpenstackNetwork {
					subject := kopsv1alpha2.OpenstackNetwork{}
					subject.InternalNetworkNames = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenDataSourceOpenstackNetworkInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceOpenstackNetwork() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceOpenstackNetwork(t *testing.T) {
	_default := map[string]interface{}{
		"availability_zone_hints": func() []interface{} { return nil }(),
		"ipv6_support_disabled":   nil,
		"public_network_names":    func() []interface{} { return nil }(),
		"internal_network_names":  func() []interface{} { return nil }(),
	}
	type args struct {
		in kopsv1alpha2.OpenstackNetwork
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.OpenstackNetwork{},
			},
			want: _default,
		},
		{
			name: "AvailabilityZoneHints - default",
			args: args{
				in: func() kopsv1alpha2.OpenstackNetwork {
					subject := kopsv1alpha2.OpenstackNetwork{}
					subject.AvailabilityZoneHints = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Ipv6SupportDisabled - default",
			args: args{
				in: func() kopsv1alpha2.OpenstackNetwork {
					subject := kopsv1alpha2.OpenstackNetwork{}
					subject.IPv6SupportDisabled = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PublicNetworkNames - default",
			args: args{
				in: func() kopsv1alpha2.OpenstackNetwork {
					subject := kopsv1alpha2.OpenstackNetwork{}
					subject.PublicNetworkNames = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "InternalNetworkNames - default",
			args: args{
				in: func() kopsv1alpha2.OpenstackNetwork {
					subject := kopsv1alpha2.OpenstackNetwork{}
					subject.InternalNetworkNames = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceOpenstackNetwork(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceOpenstackNetwork() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
