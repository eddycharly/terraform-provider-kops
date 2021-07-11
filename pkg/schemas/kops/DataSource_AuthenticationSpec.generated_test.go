package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceAuthenticationSpec(t *testing.T) {
	_default := kops.AuthenticationSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.AuthenticationSpec
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"kopeio": nil,
					"aws":    nil,
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandDataSourceAuthenticationSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceAuthenticationSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceAuthenticationSpecInto(t *testing.T) {
	_default := map[string]interface{}{
		"kopeio": nil,
		"aws":    nil,
	}
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
			want: _default,
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
			want: _default,
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
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenDataSourceAuthenticationSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceAuthenticationSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceAuthenticationSpec(t *testing.T) {
	_default := map[string]interface{}{
		"kopeio": nil,
		"aws":    nil,
	}
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
			want: _default,
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
			want: _default,
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
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceAuthenticationSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceAuthenticationSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
