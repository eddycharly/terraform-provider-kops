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
	_default := map[string]interface{}{
		"external_network":        nil,
		"dns_servers":             nil,
		"external_subnet":         nil,
		"availability_zone_hints": func() []interface{} { return nil }(),
	}
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
