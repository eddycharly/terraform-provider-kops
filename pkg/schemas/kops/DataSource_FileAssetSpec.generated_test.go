package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceFileAssetSpec(t *testing.T) {
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
			if got := ExpandDataSourceFileAssetSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandDataSourceFileAssetSpec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenDataSourceFileAssetSpecInto(t *testing.T) {
	_default := map[string]interface{}{
		"name":      "",
		"path":      "",
		"roles":     func() []interface{} { return nil }(),
		"content":   "",
		"is_base64": false,
	}
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenDataSourceFileAssetSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceFileAssetSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceFileAssetSpec(t *testing.T) {
	_default := map[string]interface{}{
		"name":      "",
		"path":      "",
		"roles":     func() []interface{} { return nil }(),
		"content":   "",
		"is_base64": false,
	}
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceFileAssetSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceFileAssetSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
