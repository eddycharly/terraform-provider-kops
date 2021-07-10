package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceDNSSpec(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.DNSSpec
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandResourceDNSSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandResourceDNSSpec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenResourceDNSSpecInto(t *testing.T) {
	type args struct {
		in  kops.DNSSpec
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
			FlattenResourceDNSSpecInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenResourceDNSSpec(t *testing.T) {
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
			want: map[string]interface{}{
				"type": "",
			},
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
			want: map[string]interface{}{
				"type": "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenResourceDNSSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenResourceDNSSpec() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
