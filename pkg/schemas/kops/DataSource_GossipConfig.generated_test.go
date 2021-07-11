package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceGossipConfig(t *testing.T) {
	_default := kops.GossipConfig{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.GossipConfig
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
		in kops.GossipConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.GossipConfig{},
			},
			want: _default,
		},
		{
			name: "Protocol - default",
			args: args{
				in: func() kops.GossipConfig {
					subject := kops.GossipConfig{}
					subject.Protocol = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Listen - default",
			args: args{
				in: func() kops.GossipConfig {
					subject := kops.GossipConfig{}
					subject.Listen = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Secret - default",
			args: args{
				in: func() kops.GossipConfig {
					subject := kops.GossipConfig{}
					subject.Secret = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Secondary - default",
			args: args{
				in: func() kops.GossipConfig {
					subject := kops.GossipConfig{}
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
		in kops.GossipConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.GossipConfig{},
			},
			want: _default,
		},
		{
			name: "Protocol - default",
			args: args{
				in: func() kops.GossipConfig {
					subject := kops.GossipConfig{}
					subject.Protocol = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Listen - default",
			args: args{
				in: func() kops.GossipConfig {
					subject := kops.GossipConfig{}
					subject.Listen = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Secret - default",
			args: args{
				in: func() kops.GossipConfig {
					subject := kops.GossipConfig{}
					subject.Secret = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Secondary - default",
			args: args{
				in: func() kops.GossipConfig {
					subject := kops.GossipConfig{}
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
