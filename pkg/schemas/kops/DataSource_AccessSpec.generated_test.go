package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceAccessSpec(t *testing.T) {
	_default := kops.AccessSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.AccessSpec
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"dns":           nil,
					"load_balancer": nil,
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandDataSourceAccessSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceAccessSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceAccessSpecInto(t *testing.T) {
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
			FlattenDataSourceAccessSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceAccessSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceAccessSpec(t *testing.T) {
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
			got := FlattenDataSourceAccessSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceAccessSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
