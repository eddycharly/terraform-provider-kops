package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceOpenstackRouter(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.OpenstackRouter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandResourceOpenstackRouter(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandResourceOpenstackRouter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenResourceOpenstackRouterInto(t *testing.T) {
	type args struct {
		in  kops.OpenstackRouter
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
			FlattenResourceOpenstackRouterInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenResourceOpenstackRouter(t *testing.T) {
	type args struct {
		in kops.OpenstackRouter
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.OpenstackRouter{},
			},
			want: map[string]interface{}{
				"external_network":        nil,
				"dns_servers":             nil,
				"external_subnet":         nil,
				"availability_zone_hints": func() []interface{} { return nil }(),
			},
		},
		{
			name: "ExternalNetwork - default",
			args: args{
				in: func() kops.OpenstackRouter {
					subject := kops.OpenstackRouter{}
					subject.ExternalNetwork = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"external_network":        nil,
				"dns_servers":             nil,
				"external_subnet":         nil,
				"availability_zone_hints": func() []interface{} { return nil }(),
			},
		},
		{
			name: "DnsServers - default",
			args: args{
				in: func() kops.OpenstackRouter {
					subject := kops.OpenstackRouter{}
					subject.DNSServers = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"external_network":        nil,
				"dns_servers":             nil,
				"external_subnet":         nil,
				"availability_zone_hints": func() []interface{} { return nil }(),
			},
		},
		{
			name: "ExternalSubnet - default",
			args: args{
				in: func() kops.OpenstackRouter {
					subject := kops.OpenstackRouter{}
					subject.ExternalSubnet = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"external_network":        nil,
				"dns_servers":             nil,
				"external_subnet":         nil,
				"availability_zone_hints": func() []interface{} { return nil }(),
			},
		},
		{
			name: "AvailabilityZoneHints - default",
			args: args{
				in: func() kops.OpenstackRouter {
					subject := kops.OpenstackRouter{}
					subject.AvailabilityZoneHints = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"external_network":        nil,
				"dns_servers":             nil,
				"external_subnet":         nil,
				"availability_zone_hints": func() []interface{} { return nil }(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenResourceOpenstackRouter(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenResourceOpenstackRouter() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
