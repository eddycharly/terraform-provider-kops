package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceAssets(t *testing.T) {
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
			if got := ExpandResourceAssets(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandResourceAssets() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenResourceAssetsInto(t *testing.T) {
	type args struct {
		in  kops.Assets
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
			FlattenResourceAssetsInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenResourceAssets(t *testing.T) {
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
			want: map[string]interface{}{
				"container_registry": nil,
				"file_repository":    nil,
				"container_proxy":    nil,
			},
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
			want: map[string]interface{}{
				"container_registry": nil,
				"file_repository":    nil,
				"container_proxy":    nil,
			},
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
			want: map[string]interface{}{
				"container_registry": nil,
				"file_repository":    nil,
				"container_proxy":    nil,
			},
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
			want: map[string]interface{}{
				"container_registry": nil,
				"file_repository":    nil,
				"container_proxy":    nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenResourceAssets(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenResourceAssets() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
