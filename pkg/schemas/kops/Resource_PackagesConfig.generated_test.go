package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourcePackagesConfig(t *testing.T) {
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
			if got := ExpandResourcePackagesConfig(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandResourcePackagesConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenResourcePackagesConfigInto(t *testing.T) {
	type args struct {
		in  kops.PackagesConfig
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
			FlattenResourcePackagesConfigInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenResourcePackagesConfig(t *testing.T) {
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
			want: map[string]interface{}{
				"hash_amd64": nil,
				"hash_arm64": nil,
				"url_amd64":  nil,
				"url_arm64":  nil,
			},
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
			want: map[string]interface{}{
				"hash_amd64": nil,
				"hash_arm64": nil,
				"url_amd64":  nil,
				"url_arm64":  nil,
			},
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
			want: map[string]interface{}{
				"hash_amd64": nil,
				"hash_arm64": nil,
				"url_amd64":  nil,
				"url_arm64":  nil,
			},
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
			want: map[string]interface{}{
				"hash_amd64": nil,
				"hash_arm64": nil,
				"url_amd64":  nil,
				"url_arm64":  nil,
			},
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
			want: map[string]interface{}{
				"hash_amd64": nil,
				"hash_arm64": nil,
				"url_amd64":  nil,
				"url_arm64":  nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenResourcePackagesConfig(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenResourcePackagesConfig() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
