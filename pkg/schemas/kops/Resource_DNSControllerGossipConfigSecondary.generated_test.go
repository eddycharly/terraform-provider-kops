package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceDNSControllerGossipConfigSecondary(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.DNSControllerGossipConfigSecondary
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandResourceDNSControllerGossipConfigSecondary(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandResourceDNSControllerGossipConfigSecondary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenResourceDNSControllerGossipConfigSecondaryInto(t *testing.T) {
	type args struct {
		in  kops.DNSControllerGossipConfigSecondary
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
			FlattenResourceDNSControllerGossipConfigSecondaryInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenResourceDNSControllerGossipConfigSecondary(t *testing.T) {
	type args struct {
		in kops.DNSControllerGossipConfigSecondary
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.DNSControllerGossipConfigSecondary{},
			},
			want: map[string]interface{}{
				"protocol": nil,
				"listen":   nil,
				"secret":   nil,
				"seed":     nil,
			},
		},
		{
			name: "Protocol - default",
			args: args{
				in: func() kops.DNSControllerGossipConfigSecondary {
					subject := kops.DNSControllerGossipConfigSecondary{}
					subject.Protocol = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"protocol": nil,
				"listen":   nil,
				"secret":   nil,
				"seed":     nil,
			},
		},
		{
			name: "Listen - default",
			args: args{
				in: func() kops.DNSControllerGossipConfigSecondary {
					subject := kops.DNSControllerGossipConfigSecondary{}
					subject.Listen = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"protocol": nil,
				"listen":   nil,
				"secret":   nil,
				"seed":     nil,
			},
		},
		{
			name: "Secret - default",
			args: args{
				in: func() kops.DNSControllerGossipConfigSecondary {
					subject := kops.DNSControllerGossipConfigSecondary{}
					subject.Secret = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"protocol": nil,
				"listen":   nil,
				"secret":   nil,
				"seed":     nil,
			},
		},
		{
			name: "Seed - default",
			args: args{
				in: func() kops.DNSControllerGossipConfigSecondary {
					subject := kops.DNSControllerGossipConfigSecondary{}
					subject.Seed = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"protocol": nil,
				"listen":   nil,
				"secret":   nil,
				"seed":     nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenResourceDNSControllerGossipConfigSecondary(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenResourceDNSControllerGossipConfigSecondary() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
