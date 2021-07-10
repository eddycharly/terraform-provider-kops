package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceKuberouterNetworkingSpec(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.KuberouterNetworkingSpec
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandDataSourceKuberouterNetworkingSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandDataSourceKuberouterNetworkingSpec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenDataSourceKuberouterNetworkingSpecInto(t *testing.T) {
	type args struct {
		in  kops.KuberouterNetworkingSpec
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
			FlattenDataSourceKuberouterNetworkingSpecInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenDataSourceKuberouterNetworkingSpec(t *testing.T) {
	type args struct {
		in kops.KuberouterNetworkingSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.KuberouterNetworkingSpec{},
			},
			want: map[string]interface{}{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenDataSourceKuberouterNetworkingSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenDataSourceKuberouterNetworkingSpec() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
