package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

func TestExpandDataSourceFileAssetSpec(t *testing.T) {
	_default := kopsv1alpha2.FileAssetSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kopsv1alpha2.FileAssetSpec
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"name":      "",
					"path":      "",
					"roles":     func() []interface{} { return nil }(),
					"content":   "",
					"is_base64": false,
					"mode":      "",
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandDataSourceFileAssetSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceFileAssetSpec() mismatch (-want +got):\n%s", diff)
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
		"mode":      "",
	}
	type args struct {
		in kopsv1alpha2.FileAssetSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.FileAssetSpec{},
			},
			want: _default,
		},
		{
			name: "Name - default",
			args: args{
				in: func() kopsv1alpha2.FileAssetSpec {
					subject := kopsv1alpha2.FileAssetSpec{}
					subject.Name = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Path - default",
			args: args{
				in: func() kopsv1alpha2.FileAssetSpec {
					subject := kopsv1alpha2.FileAssetSpec{}
					subject.Path = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Roles - default",
			args: args{
				in: func() kopsv1alpha2.FileAssetSpec {
					subject := kopsv1alpha2.FileAssetSpec{}
					subject.Roles = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Content - default",
			args: args{
				in: func() kopsv1alpha2.FileAssetSpec {
					subject := kopsv1alpha2.FileAssetSpec{}
					subject.Content = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "IsBase64 - default",
			args: args{
				in: func() kopsv1alpha2.FileAssetSpec {
					subject := kopsv1alpha2.FileAssetSpec{}
					subject.IsBase64 = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Mode - default",
			args: args{
				in: func() kopsv1alpha2.FileAssetSpec {
					subject := kopsv1alpha2.FileAssetSpec{}
					subject.Mode = ""
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
		"mode":      "",
	}
	type args struct {
		in kopsv1alpha2.FileAssetSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.FileAssetSpec{},
			},
			want: _default,
		},
		{
			name: "Name - default",
			args: args{
				in: func() kopsv1alpha2.FileAssetSpec {
					subject := kopsv1alpha2.FileAssetSpec{}
					subject.Name = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Path - default",
			args: args{
				in: func() kopsv1alpha2.FileAssetSpec {
					subject := kopsv1alpha2.FileAssetSpec{}
					subject.Path = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Roles - default",
			args: args{
				in: func() kopsv1alpha2.FileAssetSpec {
					subject := kopsv1alpha2.FileAssetSpec{}
					subject.Roles = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Content - default",
			args: args{
				in: func() kopsv1alpha2.FileAssetSpec {
					subject := kopsv1alpha2.FileAssetSpec{}
					subject.Content = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "IsBase64 - default",
			args: args{
				in: func() kopsv1alpha2.FileAssetSpec {
					subject := kopsv1alpha2.FileAssetSpec{}
					subject.IsBase64 = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Mode - default",
			args: args{
				in: func() kopsv1alpha2.FileAssetSpec {
					subject := kopsv1alpha2.FileAssetSpec{}
					subject.Mode = ""
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
