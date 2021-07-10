package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceFileAssetSpec(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.FileAssetSpec
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandResourceFileAssetSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandResourceFileAssetSpec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenResourceFileAssetSpecInto(t *testing.T) {
	type args struct {
		in  kops.FileAssetSpec
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
			FlattenResourceFileAssetSpecInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenResourceFileAssetSpec(t *testing.T) {
	type args struct {
		in kops.FileAssetSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.FileAssetSpec{},
			},
			want: map[string]interface{}{
				"name":      "",
				"path":      "",
				"roles":     func() []interface{} { return nil }(),
				"content":   "",
				"is_base64": false,
			},
		},
		{
			name: "Name - default",
			args: args{
				in: func() kops.FileAssetSpec {
					subject := kops.FileAssetSpec{}
					subject.Name = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
				"name":      "",
				"path":      "",
				"roles":     func() []interface{} { return nil }(),
				"content":   "",
				"is_base64": false,
			},
		},
		{
			name: "Path - default",
			args: args{
				in: func() kops.FileAssetSpec {
					subject := kops.FileAssetSpec{}
					subject.Path = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
				"name":      "",
				"path":      "",
				"roles":     func() []interface{} { return nil }(),
				"content":   "",
				"is_base64": false,
			},
		},
		{
			name: "Roles - default",
			args: args{
				in: func() kops.FileAssetSpec {
					subject := kops.FileAssetSpec{}
					subject.Roles = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"name":      "",
				"path":      "",
				"roles":     func() []interface{} { return nil }(),
				"content":   "",
				"is_base64": false,
			},
		},
		{
			name: "Content - default",
			args: args{
				in: func() kops.FileAssetSpec {
					subject := kops.FileAssetSpec{}
					subject.Content = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
				"name":      "",
				"path":      "",
				"roles":     func() []interface{} { return nil }(),
				"content":   "",
				"is_base64": false,
			},
		},
		{
			name: "IsBase64 - default",
			args: args{
				in: func() kops.FileAssetSpec {
					subject := kops.FileAssetSpec{}
					subject.IsBase64 = false
					return subject
				}(),
			},
			want: map[string]interface{}{
				"name":      "",
				"path":      "",
				"roles":     func() []interface{} { return nil }(),
				"content":   "",
				"is_base64": false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenResourceFileAssetSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenResourceFileAssetSpec() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
