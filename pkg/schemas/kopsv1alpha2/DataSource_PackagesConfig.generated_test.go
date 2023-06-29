package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

func TestExpandDataSourcePackagesConfig(t *testing.T) {
	_default := kopsv1alpha2.PackagesConfig{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kopsv1alpha2.PackagesConfig
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"hash_amd64": nil,
					"hash_arm64": nil,
					"url_amd64":  nil,
					"url_arm64":  nil,
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandDataSourcePackagesConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourcePackagesConfig() mismatch (-want +got):\n%s", diff)
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
		in kopsv1alpha2.PackagesConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.PackagesConfig{},
			},
			want: _default,
		},
		{
			name: "HashAmd64 - default",
			args: args{
				in: func() kopsv1alpha2.PackagesConfig {
					subject := kopsv1alpha2.PackagesConfig{}
					subject.HashAmd64 = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "HashArm64 - default",
			args: args{
				in: func() kopsv1alpha2.PackagesConfig {
					subject := kopsv1alpha2.PackagesConfig{}
					subject.HashArm64 = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "UrlAmd64 - default",
			args: args{
				in: func() kopsv1alpha2.PackagesConfig {
					subject := kopsv1alpha2.PackagesConfig{}
					subject.UrlAmd64 = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "UrlArm64 - default",
			args: args{
				in: func() kopsv1alpha2.PackagesConfig {
					subject := kopsv1alpha2.PackagesConfig{}
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
		in kopsv1alpha2.PackagesConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.PackagesConfig{},
			},
			want: _default,
		},
		{
			name: "HashAmd64 - default",
			args: args{
				in: func() kopsv1alpha2.PackagesConfig {
					subject := kopsv1alpha2.PackagesConfig{}
					subject.HashAmd64 = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "HashArm64 - default",
			args: args{
				in: func() kopsv1alpha2.PackagesConfig {
					subject := kopsv1alpha2.PackagesConfig{}
					subject.HashArm64 = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "UrlAmd64 - default",
			args: args{
				in: func() kopsv1alpha2.PackagesConfig {
					subject := kopsv1alpha2.PackagesConfig{}
					subject.UrlAmd64 = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "UrlArm64 - default",
			args: args{
				in: func() kopsv1alpha2.PackagesConfig {
					subject := kopsv1alpha2.PackagesConfig{}
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
