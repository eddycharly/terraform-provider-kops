package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceDNSSpec(t *testing.T) {
	_default := kops.DNSSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.DNSSpec
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"type": "",
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandDataSourceDNSSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceDNSSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceDNSSpecInto(t *testing.T) {
	_default := map[string]interface{}{
		"type": "",
	}
	type args struct {
		in kops.DNSSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.DNSSpec{},
			},
			want: _default,
		},
		{
			name: "Type - default",
			args: args{
				in: func() kops.DNSSpec {
					subject := kops.DNSSpec{}
					subject.Type = ""
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenDataSourceDNSSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceDNSSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceDNSSpec(t *testing.T) {
	_default := map[string]interface{}{
		"type": "",
	}
	type args struct {
		in kops.DNSSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.DNSSpec{},
			},
			want: _default,
		},
		{
			name: "Type - default",
			args: args{
				in: func() kops.DNSSpec {
					subject := kops.DNSSpec{}
					subject.Type = ""
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceDNSSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceDNSSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
