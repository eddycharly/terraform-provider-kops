package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceKopeioAuthenticationSpec(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.KopeioAuthenticationSpec
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandDataSourceKopeioAuthenticationSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandDataSourceKopeioAuthenticationSpec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenDataSourceKopeioAuthenticationSpecInto(t *testing.T) {
	type args struct {
		in  kops.KopeioAuthenticationSpec
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
			FlattenDataSourceKopeioAuthenticationSpecInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenDataSourceKopeioAuthenticationSpec(t *testing.T) {
	type args struct {
		in kops.KopeioAuthenticationSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.KopeioAuthenticationSpec{},
			},
			want: map[string]interface{}{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenDataSourceKopeioAuthenticationSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenDataSourceKopeioAuthenticationSpec() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
