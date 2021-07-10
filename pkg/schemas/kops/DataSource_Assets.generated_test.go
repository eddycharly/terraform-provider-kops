package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceAssets(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.Assets
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandDataSourceAssets(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandDataSourceAssets() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenDataSourceAssetsInto(t *testing.T) {
	_default := map[string]interface{}{
		"container_registry": nil,
		"file_repository":    nil,
		"container_proxy":    nil,
	}
	type args struct {
		in kops.Assets
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.Assets{},
			},
			want: _default,
		},
		{
			name: "ContainerRegistry - default",
			args: args{
				in: func() kops.Assets {
					subject := kops.Assets{}
					subject.ContainerRegistry = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "FileRepository - default",
			args: args{
				in: func() kops.Assets {
					subject := kops.Assets{}
					subject.FileRepository = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ContainerProxy - default",
			args: args{
				in: func() kops.Assets {
					subject := kops.Assets{}
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
			FlattenDataSourceAssetsInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceAssets() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceAssets(t *testing.T) {
	_default := map[string]interface{}{
		"container_registry": nil,
		"file_repository":    nil,
		"container_proxy":    nil,
	}
	type args struct {
		in kops.Assets
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.Assets{},
			},
			want: _default,
		},
		{
			name: "ContainerRegistry - default",
			args: args{
				in: func() kops.Assets {
					subject := kops.Assets{}
					subject.ContainerRegistry = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "FileRepository - default",
			args: args{
				in: func() kops.Assets {
					subject := kops.Assets{}
					subject.FileRepository = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ContainerProxy - default",
			args: args{
				in: func() kops.Assets {
					subject := kops.Assets{}
					subject.ContainerProxy = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceAssets(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceAssets() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
