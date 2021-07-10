package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceOpenstackNetwork(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.OpenstackNetwork
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandDataSourceOpenstackNetwork(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandDataSourceOpenstackNetwork() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenDataSourceOpenstackNetworkInto(t *testing.T) {
	type args struct {
		in  kops.OpenstackNetwork
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
			FlattenDataSourceOpenstackNetworkInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenDataSourceOpenstackNetwork(t *testing.T) {
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
			want: map[string]interface{}{
				"availability_zone_hints": func() []interface{} { return nil }(),
			},
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
			want: map[string]interface{}{
				"availability_zone_hints": func() []interface{} { return nil }(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenDataSourceOpenstackNetwork(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenDataSourceOpenstackNetwork() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
