package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceCNINetworkingSpec(t *testing.T) {
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
			if got := ExpandResourceCNINetworkingSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandResourceCNINetworkingSpec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenResourceCNINetworkingSpecInto(t *testing.T) {
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
			FlattenResourceCNINetworkingSpecInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenResourceCNINetworkingSpec(t *testing.T) {
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
			if got := FlattenResourceCNINetworkingSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenResourceCNINetworkingSpec() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
