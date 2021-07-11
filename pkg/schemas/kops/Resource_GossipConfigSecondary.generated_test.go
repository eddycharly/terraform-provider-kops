package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceGossipConfigSecondary(t *testing.T) {
	_default := kops.GossipConfigSecondary{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.GossipConfigSecondary
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"protocol": nil,
					"listen":   nil,
					"secret":   nil,
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandResourceGossipConfigSecondary(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceGossipConfigSecondary() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceGossipConfigSecondaryInto(t *testing.T) {
	_default := map[string]interface{}{
		"protocol": nil,
		"listen":   nil,
		"secret":   nil,
	}
	type args struct {
		in kops.GossipConfigSecondary
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.GossipConfigSecondary{},
			},
			want: _default,
		},
		{
			name: "Protocol - default",
			args: args{
				in: func() kops.GossipConfigSecondary {
					subject := kops.GossipConfigSecondary{}
					subject.Protocol = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Listen - default",
			args: args{
				in: func() kops.GossipConfigSecondary {
					subject := kops.GossipConfigSecondary{}
					subject.Listen = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Secret - default",
			args: args{
				in: func() kops.GossipConfigSecondary {
					subject := kops.GossipConfigSecondary{}
					subject.Secret = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenResourceGossipConfigSecondaryInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceGossipConfigSecondary() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceGossipConfigSecondary(t *testing.T) {
	_default := map[string]interface{}{
		"protocol": nil,
		"listen":   nil,
		"secret":   nil,
	}
	type args struct {
		in kops.GossipConfigSecondary
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.GossipConfigSecondary{},
			},
			want: _default,
		},
		{
			name: "Protocol - default",
			args: args{
				in: func() kops.GossipConfigSecondary {
					subject := kops.GossipConfigSecondary{}
					subject.Protocol = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Listen - default",
			args: args{
				in: func() kops.GossipConfigSecondary {
					subject := kops.GossipConfigSecondary{}
					subject.Listen = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Secret - default",
			args: args{
				in: func() kops.GossipConfigSecondary {
					subject := kops.GossipConfigSecondary{}
					subject.Secret = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceGossipConfigSecondary(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceGossipConfigSecondary() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
