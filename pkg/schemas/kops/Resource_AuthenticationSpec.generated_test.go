package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceAuthenticationSpec(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.AuthenticationSpec
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandResourceAuthenticationSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandResourceAuthenticationSpec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenResourceAuthenticationSpecInto(t *testing.T) {
	type args struct {
		in  kops.AuthenticationSpec
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
			FlattenResourceAuthenticationSpecInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenResourceAuthenticationSpec(t *testing.T) {
	type args struct {
		in kops.AuthenticationSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.AuthenticationSpec{},
			},
			want: map[string]interface{}{
				"kopeio": nil,
				"aws":    nil,
			},
		},
		{
			name: "Kopeio - default",
			args: args{
				in: func() kops.AuthenticationSpec {
					subject := kops.AuthenticationSpec{}
					subject.Kopeio = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"kopeio": nil,
				"aws":    nil,
			},
		},
		{
			name: "Aws - default",
			args: args{
				in: func() kops.AuthenticationSpec {
					subject := kops.AuthenticationSpec{}
					subject.Aws = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"kopeio": nil,
				"aws":    nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenResourceAuthenticationSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenResourceAuthenticationSpec() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
