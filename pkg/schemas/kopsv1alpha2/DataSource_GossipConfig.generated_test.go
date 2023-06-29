package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

func TestExpandDataSourceGossipConfig(t *testing.T) {
	_default := kopsv1alpha2.GossipConfig{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kopsv1alpha2.GossipConfig
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"protocol":  nil,
					"listen":    nil,
					"secret":    nil,
					"secondary": nil,
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandDataSourceGossipConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceGossipConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceGossipConfigInto(t *testing.T) {
	_default := map[string]interface{}{
		"protocol":  nil,
		"listen":    nil,
		"secret":    nil,
		"secondary": nil,
	}
	type args struct {
		in kopsv1alpha2.GossipConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.GossipConfig{},
			},
			want: _default,
		},
		{
			name: "Protocol - default",
			args: args{
				in: func() kopsv1alpha2.GossipConfig {
					subject := kopsv1alpha2.GossipConfig{}
					subject.Protocol = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Listen - default",
			args: args{
				in: func() kopsv1alpha2.GossipConfig {
					subject := kopsv1alpha2.GossipConfig{}
					subject.Listen = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Secret - default",
			args: args{
				in: func() kopsv1alpha2.GossipConfig {
					subject := kopsv1alpha2.GossipConfig{}
					subject.Secret = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Secondary - default",
			args: args{
				in: func() kopsv1alpha2.GossipConfig {
					subject := kopsv1alpha2.GossipConfig{}
					subject.Secondary = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenDataSourceGossipConfigInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceGossipConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceGossipConfig(t *testing.T) {
	_default := map[string]interface{}{
		"protocol":  nil,
		"listen":    nil,
		"secret":    nil,
		"secondary": nil,
	}
	type args struct {
		in kopsv1alpha2.GossipConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.GossipConfig{},
			},
			want: _default,
		},
		{
			name: "Protocol - default",
			args: args{
				in: func() kopsv1alpha2.GossipConfig {
					subject := kopsv1alpha2.GossipConfig{}
					subject.Protocol = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Listen - default",
			args: args{
				in: func() kopsv1alpha2.GossipConfig {
					subject := kopsv1alpha2.GossipConfig{}
					subject.Listen = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Secret - default",
			args: args{
				in: func() kopsv1alpha2.GossipConfig {
					subject := kopsv1alpha2.GossipConfig{}
					subject.Secret = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Secondary - default",
			args: args{
				in: func() kopsv1alpha2.GossipConfig {
					subject := kopsv1alpha2.GossipConfig{}
					subject.Secondary = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceGossipConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceGossipConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
