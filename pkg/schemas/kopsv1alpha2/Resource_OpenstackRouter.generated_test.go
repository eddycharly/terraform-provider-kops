package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

func TestExpandResourceOpenstackRouter(t *testing.T) {
	_default := kopsv1alpha2.OpenstackRouter{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kopsv1alpha2.OpenstackRouter
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"external_network":        nil,
					"dns_servers":             nil,
					"external_subnet":         nil,
					"availability_zone_hints": func() []interface{} { return nil }(),
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandResourceOpenstackRouter(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceOpenstackRouter() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceOpenstackRouterInto(t *testing.T) {
	_default := map[string]interface{}{
		"external_network":        nil,
		"dns_servers":             nil,
		"external_subnet":         nil,
		"availability_zone_hints": func() []interface{} { return nil }(),
	}
	type args struct {
		in kopsv1alpha2.OpenstackRouter
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.OpenstackRouter{},
			},
			want: _default,
		},
		{
			name: "ExternalNetwork - default",
			args: args{
				in: func() kopsv1alpha2.OpenstackRouter {
					subject := kopsv1alpha2.OpenstackRouter{}
					subject.ExternalNetwork = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "DnsServers - default",
			args: args{
				in: func() kopsv1alpha2.OpenstackRouter {
					subject := kopsv1alpha2.OpenstackRouter{}
					subject.DNSServers = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ExternalSubnet - default",
			args: args{
				in: func() kopsv1alpha2.OpenstackRouter {
					subject := kopsv1alpha2.OpenstackRouter{}
					subject.ExternalSubnet = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AvailabilityZoneHints - default",
			args: args{
				in: func() kopsv1alpha2.OpenstackRouter {
					subject := kopsv1alpha2.OpenstackRouter{}
					subject.AvailabilityZoneHints = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenResourceOpenstackRouterInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceOpenstackRouter() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceOpenstackRouter(t *testing.T) {
	_default := map[string]interface{}{
		"external_network":        nil,
		"dns_servers":             nil,
		"external_subnet":         nil,
		"availability_zone_hints": func() []interface{} { return nil }(),
	}
	type args struct {
		in kopsv1alpha2.OpenstackRouter
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.OpenstackRouter{},
			},
			want: _default,
		},
		{
			name: "ExternalNetwork - default",
			args: args{
				in: func() kopsv1alpha2.OpenstackRouter {
					subject := kopsv1alpha2.OpenstackRouter{}
					subject.ExternalNetwork = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "DnsServers - default",
			args: args{
				in: func() kopsv1alpha2.OpenstackRouter {
					subject := kopsv1alpha2.OpenstackRouter{}
					subject.DNSServers = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ExternalSubnet - default",
			args: args{
				in: func() kopsv1alpha2.OpenstackRouter {
					subject := kopsv1alpha2.OpenstackRouter{}
					subject.ExternalSubnet = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AvailabilityZoneHints - default",
			args: args{
				in: func() kopsv1alpha2.OpenstackRouter {
					subject := kopsv1alpha2.OpenstackRouter{}
					subject.AvailabilityZoneHints = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceOpenstackRouter(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceOpenstackRouter() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
