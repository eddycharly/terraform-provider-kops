package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceDNSControllerGossipConfigSecondary(t *testing.T) {
	_default := kops.DNSControllerGossipConfigSecondary{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.DNSControllerGossipConfigSecondary
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"protocol": nil,
					"listen":   nil,
					"secret":   nil,
					"seed":     nil,
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandDataSourceDNSControllerGossipConfigSecondary(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceDNSControllerGossipConfigSecondary() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceDNSControllerGossipConfigSecondaryInto(t *testing.T) {
	_default := map[string]interface{}{
		"protocol": nil,
		"listen":   nil,
		"secret":   nil,
		"seed":     nil,
	}
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenDataSourceDNSControllerGossipConfigSecondaryInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceDNSControllerGossipConfigSecondary() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceDNSControllerGossipConfigSecondary(t *testing.T) {
	_default := map[string]interface{}{
		"protocol": nil,
		"listen":   nil,
		"secret":   nil,
		"seed":     nil,
	}
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceDNSControllerGossipConfigSecondary(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceDNSControllerGossipConfigSecondary() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
