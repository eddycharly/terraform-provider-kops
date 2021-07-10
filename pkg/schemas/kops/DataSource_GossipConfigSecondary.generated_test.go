package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceGossipConfigSecondary(t *testing.T) {
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
			if got := ExpandDataSourceGossipConfigSecondary(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandDataSourceGossipConfigSecondary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenDataSourceGossipConfigSecondaryInto(t *testing.T) {
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
			FlattenDataSourceGossipConfigSecondaryInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenDataSourceGossipConfigSecondary(t *testing.T) {
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
			if got := FlattenDataSourceGossipConfigSecondary(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenDataSourceGossipConfigSecondary() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
