package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourcePackagesConfig(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.PackagesConfig
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandDataSourcePackagesConfig(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandDataSourcePackagesConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenDataSourcePackagesConfigInto(t *testing.T) {
	_default := map[string]interface{}{
		"hash_amd64": nil,
		"hash_arm64": nil,
		"url_amd64":  nil,
		"url_arm64":  nil,
	}
	type args struct {
		in kops.PackagesConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.PackagesConfig{},
			},
			want: _default,
		},
		{
			name: "HashAmd64 - default",
			args: args{
				in: func() kops.PackagesConfig {
					subject := kops.PackagesConfig{}
					subject.HashAmd64 = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "HashArm64 - default",
			args: args{
				in: func() kops.PackagesConfig {
					subject := kops.PackagesConfig{}
					subject.HashArm64 = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "UrlAmd64 - default",
			args: args{
				in: func() kops.PackagesConfig {
					subject := kops.PackagesConfig{}
					subject.UrlAmd64 = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "UrlArm64 - default",
			args: args{
				in: func() kops.PackagesConfig {
					subject := kops.PackagesConfig{}
					subject.UrlArm64 = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenDataSourcePackagesConfigInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourcePackagesConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourcePackagesConfig(t *testing.T) {
	_default := map[string]interface{}{
		"hash_amd64": nil,
		"hash_arm64": nil,
		"url_amd64":  nil,
		"url_arm64":  nil,
	}
	type args struct {
		in kops.PackagesConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.PackagesConfig{},
			},
			want: _default,
		},
		{
			name: "HashAmd64 - default",
			args: args{
				in: func() kops.PackagesConfig {
					subject := kops.PackagesConfig{}
					subject.HashAmd64 = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "HashArm64 - default",
			args: args{
				in: func() kops.PackagesConfig {
					subject := kops.PackagesConfig{}
					subject.HashArm64 = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "UrlAmd64 - default",
			args: args{
				in: func() kops.PackagesConfig {
					subject := kops.PackagesConfig{}
					subject.UrlAmd64 = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "UrlArm64 - default",
			args: args{
				in: func() kops.PackagesConfig {
					subject := kops.PackagesConfig{}
					subject.UrlArm64 = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourcePackagesConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourcePackagesConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
