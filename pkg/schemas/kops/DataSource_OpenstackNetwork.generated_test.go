package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceOpenstackNetwork(t *testing.T) {
	_default := kops.OpenstackNetwork{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.OpenstackNetwork
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"availability_zone_hints": func() []interface{} { return nil }(),
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandDataSourceOpenstackNetwork(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceOpenstackNetwork() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceOpenstackNetworkInto(t *testing.T) {
	_default := map[string]interface{}{
		"availability_zone_hints": func() []interface{} { return nil }(),
	}
	type args struct {
		in kops.OpenstackNetwork
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.OpenstackNetwork{},
			},
			want: _default,
		},
		{
			name: "AvailabilityZoneHints - default",
			args: args{
				in: func() kops.OpenstackNetwork {
					subject := kops.OpenstackNetwork{}
					subject.AvailabilityZoneHints = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenDataSourceOpenstackNetworkInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceOpenstackNetwork() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceOpenstackNetwork(t *testing.T) {
	_default := map[string]interface{}{
		"availability_zone_hints": func() []interface{} { return nil }(),
	}
	type args struct {
		in kops.OpenstackNetwork
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.OpenstackNetwork{},
			},
			want: _default,
		},
		{
			name: "AvailabilityZoneHints - default",
			args: args{
				in: func() kops.OpenstackNetwork {
					subject := kops.OpenstackNetwork{}
					subject.AvailabilityZoneHints = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceOpenstackNetwork(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceOpenstackNetwork() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
