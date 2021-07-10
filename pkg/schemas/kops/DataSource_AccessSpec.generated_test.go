package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceAccessSpec(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.AccessSpec
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandDataSourceAccessSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandDataSourceAccessSpec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenDataSourceAccessSpecInto(t *testing.T) {
	type args struct {
		in  kops.AccessSpec
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
			FlattenDataSourceAccessSpecInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenDataSourceAccessSpec(t *testing.T) {
	type args struct {
		in kops.AccessSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.AccessSpec{},
			},
			want: map[string]interface{}{
				"dns":           nil,
				"load_balancer": nil,
			},
		},
		{
			name: "Dns - default",
			args: args{
				in: func() kops.AccessSpec {
					subject := kops.AccessSpec{}
					subject.DNS = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"dns":           nil,
				"load_balancer": nil,
			},
		},
		{
			name: "LoadBalancer - default",
			args: args{
				in: func() kops.AccessSpec {
					subject := kops.AccessSpec{}
					subject.LoadBalancer = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"dns":           nil,
				"load_balancer": nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenDataSourceAccessSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenDataSourceAccessSpec() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
