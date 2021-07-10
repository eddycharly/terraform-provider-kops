package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceCNINetworkingSpec(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.CNINetworkingSpec
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandDataSourceCNINetworkingSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandDataSourceCNINetworkingSpec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenDataSourceCNINetworkingSpecInto(t *testing.T) {
	type args struct {
		in  kops.CNINetworkingSpec
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
			FlattenDataSourceCNINetworkingSpecInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenDataSourceCNINetworkingSpec(t *testing.T) {
	type args struct {
		in kops.CNINetworkingSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.CNINetworkingSpec{},
			},
			want: map[string]interface{}{
				"uses_secondary_ip": false,
			},
		},
		{
			name: "UsesSecondaryIp - default",
			args: args{
				in: func() kops.CNINetworkingSpec {
					subject := kops.CNINetworkingSpec{}
					subject.UsesSecondaryIP = false
					return subject
				}(),
			},
			want: map[string]interface{}{
				"uses_secondary_ip": false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenDataSourceCNINetworkingSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenDataSourceCNINetworkingSpec() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
