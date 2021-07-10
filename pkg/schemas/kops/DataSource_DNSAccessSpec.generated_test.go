package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceDNSAccessSpec(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.DNSAccessSpec
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandDataSourceDNSAccessSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandDataSourceDNSAccessSpec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenDataSourceDNSAccessSpecInto(t *testing.T) {
	type args struct {
		in  kops.DNSAccessSpec
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
			FlattenDataSourceDNSAccessSpecInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenDataSourceDNSAccessSpec(t *testing.T) {
	type args struct {
		in kops.DNSAccessSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.DNSAccessSpec{},
			},
			want: map[string]interface{}{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenDataSourceDNSAccessSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenDataSourceDNSAccessSpec() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
