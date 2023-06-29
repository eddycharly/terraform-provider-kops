package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

func TestExpandResourceAssets(t *testing.T) {
	_default := kopsv1alpha2.Assets{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kopsv1alpha2.Assets
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"container_registry": nil,
					"file_repository":    nil,
					"container_proxy":    nil,
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandResourceAssets(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceAssets() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceAssetsInto(t *testing.T) {
	_default := map[string]interface{}{
		"container_registry": nil,
		"file_repository":    nil,
		"container_proxy":    nil,
	}
	type args struct {
		in kopsv1alpha2.Assets
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.Assets{},
			},
			want: _default,
		},
		{
			name: "ContainerRegistry - default",
			args: args{
				in: func() kopsv1alpha2.Assets {
					subject := kopsv1alpha2.Assets{}
					subject.ContainerRegistry = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "FileRepository - default",
			args: args{
				in: func() kopsv1alpha2.Assets {
					subject := kopsv1alpha2.Assets{}
					subject.FileRepository = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ContainerProxy - default",
			args: args{
				in: func() kopsv1alpha2.Assets {
					subject := kopsv1alpha2.Assets{}
					subject.ContainerProxy = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenResourceAssetsInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceAssets() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceAssets(t *testing.T) {
	_default := map[string]interface{}{
		"container_registry": nil,
		"file_repository":    nil,
		"container_proxy":    nil,
	}
	type args struct {
		in kopsv1alpha2.Assets
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.Assets{},
			},
			want: _default,
		},
		{
			name: "ContainerRegistry - default",
			args: args{
				in: func() kopsv1alpha2.Assets {
					subject := kopsv1alpha2.Assets{}
					subject.ContainerRegistry = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "FileRepository - default",
			args: args{
				in: func() kopsv1alpha2.Assets {
					subject := kopsv1alpha2.Assets{}
					subject.FileRepository = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ContainerProxy - default",
			args: args{
				in: func() kopsv1alpha2.Assets {
					subject := kopsv1alpha2.Assets{}
					subject.ContainerProxy = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceAssets(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceAssets() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
