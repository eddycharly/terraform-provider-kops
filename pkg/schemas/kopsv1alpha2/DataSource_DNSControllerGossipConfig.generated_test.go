package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

func TestExpandDataSourceDNSControllerGossipConfig(t *testing.T) {
	_default := kopsv1alpha2.DNSControllerGossipConfig{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kopsv1alpha2.DNSControllerGossipConfig
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"protocol":  nil,
					"listen":    nil,
					"secret":    nil,
					"secondary": nil,
					"seed":      nil,
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandDataSourceDNSControllerGossipConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceDNSControllerGossipConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceDNSControllerGossipConfigInto(t *testing.T) {
	_default := map[string]interface{}{
		"protocol":  nil,
		"listen":    nil,
		"secret":    nil,
		"secondary": nil,
		"seed":      nil,
	}
	type args struct {
		in kopsv1alpha2.DNSControllerGossipConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.DNSControllerGossipConfig{},
			},
			want: _default,
		},
		{
			name: "Protocol - default",
			args: args{
				in: func() kopsv1alpha2.DNSControllerGossipConfig {
					subject := kopsv1alpha2.DNSControllerGossipConfig{}
					subject.Protocol = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Listen - default",
			args: args{
				in: func() kopsv1alpha2.DNSControllerGossipConfig {
					subject := kopsv1alpha2.DNSControllerGossipConfig{}
					subject.Listen = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Secret - default",
			args: args{
				in: func() kopsv1alpha2.DNSControllerGossipConfig {
					subject := kopsv1alpha2.DNSControllerGossipConfig{}
					subject.Secret = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Secondary - default",
			args: args{
				in: func() kopsv1alpha2.DNSControllerGossipConfig {
					subject := kopsv1alpha2.DNSControllerGossipConfig{}
					subject.Secondary = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Seed - default",
			args: args{
				in: func() kopsv1alpha2.DNSControllerGossipConfig {
					subject := kopsv1alpha2.DNSControllerGossipConfig{}
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
			FlattenDataSourceDNSControllerGossipConfigInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceDNSControllerGossipConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceDNSControllerGossipConfig(t *testing.T) {
	_default := map[string]interface{}{
		"protocol":  nil,
		"listen":    nil,
		"secret":    nil,
		"secondary": nil,
		"seed":      nil,
	}
	type args struct {
		in kopsv1alpha2.DNSControllerGossipConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.DNSControllerGossipConfig{},
			},
			want: _default,
		},
		{
			name: "Protocol - default",
			args: args{
				in: func() kopsv1alpha2.DNSControllerGossipConfig {
					subject := kopsv1alpha2.DNSControllerGossipConfig{}
					subject.Protocol = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Listen - default",
			args: args{
				in: func() kopsv1alpha2.DNSControllerGossipConfig {
					subject := kopsv1alpha2.DNSControllerGossipConfig{}
					subject.Listen = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Secret - default",
			args: args{
				in: func() kopsv1alpha2.DNSControllerGossipConfig {
					subject := kopsv1alpha2.DNSControllerGossipConfig{}
					subject.Secret = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Secondary - default",
			args: args{
				in: func() kopsv1alpha2.DNSControllerGossipConfig {
					subject := kopsv1alpha2.DNSControllerGossipConfig{}
					subject.Secondary = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Seed - default",
			args: args{
				in: func() kopsv1alpha2.DNSControllerGossipConfig {
					subject := kopsv1alpha2.DNSControllerGossipConfig{}
					subject.Seed = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceDNSControllerGossipConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceDNSControllerGossipConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
