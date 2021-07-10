package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceAccessSpec(t *testing.T) {
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
			if got := ExpandResourceAccessSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandResourceAccessSpec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenResourceAccessSpecInto(t *testing.T) {
	_default := map[string]interface{}{
		"dns":           nil,
		"load_balancer": nil,
	}
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
			want: _default,
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
			want: _default,
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
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenResourceAccessSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceAccessSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceAccessSpec(t *testing.T) {
	_default := map[string]interface{}{
		"dns":           nil,
		"load_balancer": nil,
	}
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
			want: _default,
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
			want: _default,
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
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceAccessSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceAccessSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
