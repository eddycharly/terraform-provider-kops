package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceDNSAccessSpec(t *testing.T) {
	_default := kops.DNSAccessSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.DNSAccessSpec
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandResourceDNSAccessSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceDNSAccessSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceDNSAccessSpecInto(t *testing.T) {
	_default := map[string]interface{}{}
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
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenResourceDNSAccessSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceDNSAccessSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceDNSAccessSpec(t *testing.T) {
	_default := map[string]interface{}{}
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
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceDNSAccessSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceDNSAccessSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
