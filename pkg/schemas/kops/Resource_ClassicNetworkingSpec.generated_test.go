package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceClassicNetworkingSpec(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.ClassicNetworkingSpec
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandResourceClassicNetworkingSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandResourceClassicNetworkingSpec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenResourceClassicNetworkingSpecInto(t *testing.T) {
	type args struct {
		in  kops.ClassicNetworkingSpec
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
			FlattenResourceClassicNetworkingSpecInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenResourceClassicNetworkingSpec(t *testing.T) {
	type args struct {
		in kops.ClassicNetworkingSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.ClassicNetworkingSpec{},
			},
			want: map[string]interface{}{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenResourceClassicNetworkingSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenResourceClassicNetworkingSpec() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
