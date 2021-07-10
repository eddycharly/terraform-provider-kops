package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceWarmPoolSpec(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.WarmPoolSpec
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandDataSourceWarmPoolSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandDataSourceWarmPoolSpec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenDataSourceWarmPoolSpecInto(t *testing.T) {
	type args struct {
		in  kops.WarmPoolSpec
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
			FlattenDataSourceWarmPoolSpecInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenDataSourceWarmPoolSpec(t *testing.T) {
	type args struct {
		in kops.WarmPoolSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.WarmPoolSpec{},
			},
			want: map[string]interface{}{
				"min_size":              0,
				"max_size":              nil,
				"enable_lifecycle_hook": false,
			},
		},
		{
			name: "MinSize - default",
			args: args{
				in: func() kops.WarmPoolSpec {
					subject := kops.WarmPoolSpec{}
					subject.MinSize = 0
					return subject
				}(),
			},
			want: map[string]interface{}{
				"min_size":              0,
				"max_size":              nil,
				"enable_lifecycle_hook": false,
			},
		},
		{
			name: "MaxSize - default",
			args: args{
				in: func() kops.WarmPoolSpec {
					subject := kops.WarmPoolSpec{}
					subject.MaxSize = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"min_size":              0,
				"max_size":              nil,
				"enable_lifecycle_hook": false,
			},
		},
		{
			name: "EnableLifecycleHook - default",
			args: args{
				in: func() kops.WarmPoolSpec {
					subject := kops.WarmPoolSpec{}
					subject.EnableLifecycleHook = false
					return subject
				}(),
			},
			want: map[string]interface{}{
				"min_size":              0,
				"max_size":              nil,
				"enable_lifecycle_hook": false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenDataSourceWarmPoolSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenDataSourceWarmPoolSpec() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
