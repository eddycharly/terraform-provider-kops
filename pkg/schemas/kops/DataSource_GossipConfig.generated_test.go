package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceGossipConfig(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.GossipConfig
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandDataSourceGossipConfig(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandDataSourceGossipConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenDataSourceGossipConfigInto(t *testing.T) {
	type args struct {
		in  kops.GossipConfig
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
			FlattenDataSourceGossipConfigInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenDataSourceGossipConfig(t *testing.T) {
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
			want: map[string]interface{}{
				"protocol":  nil,
				"listen":    nil,
				"secret":    nil,
				"secondary": nil,
			},
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
			want: map[string]interface{}{
				"protocol":  nil,
				"listen":    nil,
				"secret":    nil,
				"secondary": nil,
			},
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
			want: map[string]interface{}{
				"protocol":  nil,
				"listen":    nil,
				"secret":    nil,
				"secondary": nil,
			},
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
			want: map[string]interface{}{
				"protocol":  nil,
				"listen":    nil,
				"secret":    nil,
				"secondary": nil,
			},
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
			want: map[string]interface{}{
				"protocol":  nil,
				"listen":    nil,
				"secret":    nil,
				"secondary": nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenDataSourceGossipConfig(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenDataSourceGossipConfig() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
