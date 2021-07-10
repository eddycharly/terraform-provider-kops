package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceDNSControllerGossipConfig(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.DNSControllerGossipConfig
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandDataSourceDNSControllerGossipConfig(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandDataSourceDNSControllerGossipConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenDataSourceDNSControllerGossipConfigInto(t *testing.T) {
	type args struct {
		in  kops.DNSControllerGossipConfig
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
			FlattenDataSourceDNSControllerGossipConfigInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenDataSourceDNSControllerGossipConfig(t *testing.T) {
	type args struct {
		in kops.DNSControllerGossipConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.DNSControllerGossipConfig{},
			},
			want: map[string]interface{}{
				"protocol":  nil,
				"listen":    nil,
				"secret":    nil,
				"secondary": nil,
				"seed":      nil,
			},
		},
		{
			name: "Protocol - default",
			args: args{
				in: func() kops.DNSControllerGossipConfig {
					subject := kops.DNSControllerGossipConfig{}
					subject.Protocol = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"protocol":  nil,
				"listen":    nil,
				"secret":    nil,
				"secondary": nil,
				"seed":      nil,
			},
		},
		{
			name: "Listen - default",
			args: args{
				in: func() kops.DNSControllerGossipConfig {
					subject := kops.DNSControllerGossipConfig{}
					subject.Listen = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"protocol":  nil,
				"listen":    nil,
				"secret":    nil,
				"secondary": nil,
				"seed":      nil,
			},
		},
		{
			name: "Secret - default",
			args: args{
				in: func() kops.DNSControllerGossipConfig {
					subject := kops.DNSControllerGossipConfig{}
					subject.Secret = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"protocol":  nil,
				"listen":    nil,
				"secret":    nil,
				"secondary": nil,
				"seed":      nil,
			},
		},
		{
			name: "Secondary - default",
			args: args{
				in: func() kops.DNSControllerGossipConfig {
					subject := kops.DNSControllerGossipConfig{}
					subject.Secondary = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"protocol":  nil,
				"listen":    nil,
				"secret":    nil,
				"secondary": nil,
				"seed":      nil,
			},
		},
		{
			name: "Seed - default",
			args: args{
				in: func() kops.DNSControllerGossipConfig {
					subject := kops.DNSControllerGossipConfig{}
					subject.Seed = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"protocol":  nil,
				"listen":    nil,
				"secret":    nil,
				"secondary": nil,
				"seed":      nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenDataSourceDNSControllerGossipConfig(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenDataSourceDNSControllerGossipConfig() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
