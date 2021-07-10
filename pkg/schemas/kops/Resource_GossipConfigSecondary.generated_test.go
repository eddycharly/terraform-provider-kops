package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceGossipConfigSecondary(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.GossipConfigSecondary
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandResourceGossipConfigSecondary(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandResourceGossipConfigSecondary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenResourceGossipConfigSecondaryInto(t *testing.T) {
	type args struct {
		in  kops.GossipConfigSecondary
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
			FlattenResourceGossipConfigSecondaryInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenResourceGossipConfigSecondary(t *testing.T) {
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
			want: map[string]interface{}{
				"protocol": nil,
				"listen":   nil,
				"secret":   nil,
			},
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
			want: map[string]interface{}{
				"protocol": nil,
				"listen":   nil,
				"secret":   nil,
			},
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
			want: map[string]interface{}{
				"protocol": nil,
				"listen":   nil,
				"secret":   nil,
			},
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
			want: map[string]interface{}{
				"protocol": nil,
				"listen":   nil,
				"secret":   nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenResourceGossipConfigSecondary(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenResourceGossipConfigSecondary() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
