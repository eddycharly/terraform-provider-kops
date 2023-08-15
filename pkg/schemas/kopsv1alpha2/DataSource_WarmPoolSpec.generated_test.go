package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

func TestExpandDataSourceWarmPoolSpec(t *testing.T) {
	_default := kopsv1alpha2.WarmPoolSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kopsv1alpha2.WarmPoolSpec
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"min_size":              0,
					"max_size":              nil,
					"enable_lifecycle_hook": false,
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandDataSourceWarmPoolSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceWarmPoolSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceWarmPoolSpecInto(t *testing.T) {
	_default := map[string]interface{}{
		"min_size":              0,
		"max_size":              nil,
		"enable_lifecycle_hook": false,
	}
	type args struct {
		in kopsv1alpha2.WarmPoolSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.WarmPoolSpec{},
			},
			want: _default,
		},
		{
			name: "MinSize - default",
			args: args{
				in: func() kopsv1alpha2.WarmPoolSpec {
					subject := kopsv1alpha2.WarmPoolSpec{}
					subject.MinSize = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MaxSize - default",
			args: args{
				in: func() kopsv1alpha2.WarmPoolSpec {
					subject := kopsv1alpha2.WarmPoolSpec{}
					subject.MaxSize = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableLifecycleHook - default",
			args: args{
				in: func() kopsv1alpha2.WarmPoolSpec {
					subject := kopsv1alpha2.WarmPoolSpec{}
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
			FlattenDataSourceWarmPoolSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceWarmPoolSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceWarmPoolSpec(t *testing.T) {
	_default := map[string]interface{}{
		"min_size":              0,
		"max_size":              nil,
		"enable_lifecycle_hook": false,
	}
	type args struct {
		in kopsv1alpha2.WarmPoolSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.WarmPoolSpec{},
			},
			want: _default,
		},
		{
			name: "MinSize - default",
			args: args{
				in: func() kopsv1alpha2.WarmPoolSpec {
					subject := kopsv1alpha2.WarmPoolSpec{}
					subject.MinSize = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MaxSize - default",
			args: args{
				in: func() kopsv1alpha2.WarmPoolSpec {
					subject := kopsv1alpha2.WarmPoolSpec{}
					subject.MaxSize = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableLifecycleHook - default",
			args: args{
				in: func() kopsv1alpha2.WarmPoolSpec {
					subject := kopsv1alpha2.WarmPoolSpec{}
					subject.EnableLifecycleHook = false
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceWarmPoolSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceWarmPoolSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
