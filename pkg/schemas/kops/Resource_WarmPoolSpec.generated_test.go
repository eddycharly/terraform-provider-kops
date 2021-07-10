package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceWarmPoolSpec(t *testing.T) {
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
			if got := ExpandResourceWarmPoolSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandResourceWarmPoolSpec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenResourceWarmPoolSpecInto(t *testing.T) {
	_default := map[string]interface{}{
		"min_size":              0,
		"max_size":              nil,
		"enable_lifecycle_hook": false,
	}
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenResourceWarmPoolSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceWarmPoolSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceWarmPoolSpec(t *testing.T) {
	_default := map[string]interface{}{
		"min_size":              0,
		"max_size":              nil,
		"enable_lifecycle_hook": false,
	}
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceWarmPoolSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceWarmPoolSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
