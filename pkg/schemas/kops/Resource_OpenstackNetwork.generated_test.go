package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceOpenstackNetwork(t *testing.T) {
	_default := kops.OpenstackNetwork{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.OpenstackNetwork
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
			got := ExpandResourceOpenstackNetwork(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceOpenstackNetwork() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceOpenstackNetworkInto(t *testing.T) {
	_default := map[string]interface{}{
		"availability_zone_hints": func() []interface{} { return nil }(),
		"ipv6_support_disabled":   nil,
		"public_network_names":    func() []interface{} { return nil }(),
		"internal_network_names":  func() []interface{} { return nil }(),
	}
	type args struct {
		in kops.OpenstackNetwork
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.OpenstackNetwork{},
			},
			want: _default,
		},
		{
			name: "AvailabilityZoneHints - default",
			args: args{
				in: func() kops.OpenstackNetwork {
					subject := kops.OpenstackNetwork{}
					subject.AvailabilityZoneHints = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Ipv6SupportDisabled - default",
			args: args{
				in: func() kops.OpenstackNetwork {
					subject := kops.OpenstackNetwork{}
					subject.IPv6SupportDisabled = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PublicNetworkNames - default",
			args: args{
				in: func() kops.OpenstackNetwork {
					subject := kops.OpenstackNetwork{}
					subject.PublicNetworkNames = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "InternalNetworkNames - default",
			args: args{
				in: func() kops.OpenstackNetwork {
					subject := kops.OpenstackNetwork{}
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
			FlattenResourceOpenstackNetworkInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceOpenstackNetwork() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceOpenstackNetwork(t *testing.T) {
	_default := map[string]interface{}{
		"availability_zone_hints": func() []interface{} { return nil }(),
		"ipv6_support_disabled":   nil,
		"public_network_names":    func() []interface{} { return nil }(),
		"internal_network_names":  func() []interface{} { return nil }(),
	}
	type args struct {
		in kops.OpenstackNetwork
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.OpenstackNetwork{},
			},
			want: _default,
		},
		{
			name: "AvailabilityZoneHints - default",
			args: args{
				in: func() kops.OpenstackNetwork {
					subject := kops.OpenstackNetwork{}
					subject.AvailabilityZoneHints = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Ipv6SupportDisabled - default",
			args: args{
				in: func() kops.OpenstackNetwork {
					subject := kops.OpenstackNetwork{}
					subject.IPv6SupportDisabled = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PublicNetworkNames - default",
			args: args{
				in: func() kops.OpenstackNetwork {
					subject := kops.OpenstackNetwork{}
					subject.PublicNetworkNames = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "InternalNetworkNames - default",
			args: args{
				in: func() kops.OpenstackNetwork {
					subject := kops.OpenstackNetwork{}
					subject.InternalNetworkNames = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceOpenstackNetwork(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceOpenstackNetwork() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
