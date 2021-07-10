package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceKopeioNetworkingSpec(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.KopeioNetworkingSpec
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandResourceKopeioNetworkingSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandResourceKopeioNetworkingSpec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenResourceKopeioNetworkingSpecInto(t *testing.T) {
	type args struct {
		in  kops.KopeioNetworkingSpec
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
			FlattenResourceKopeioNetworkingSpecInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenResourceKopeioNetworkingSpec(t *testing.T) {
	type args struct {
		in kops.KopeioNetworkingSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.KopeioNetworkingSpec{},
			},
			want: map[string]interface{}{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenResourceKopeioNetworkingSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenResourceKopeioNetworkingSpec() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
