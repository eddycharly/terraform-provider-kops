package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

func TestExpandResourceMinMaxSpec(t *testing.T) {
	_default := kopsv1alpha2.MinMaxSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kopsv1alpha2.MinMaxSpec
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"max": nil,
					"min": nil,
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandResourceMinMaxSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceMinMaxSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceMinMaxSpecInto(t *testing.T) {
	_default := map[string]interface{}{
		"max": nil,
		"min": nil,
	}
	type args struct {
		in kopsv1alpha2.MinMaxSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.MinMaxSpec{},
			},
			want: _default,
		},
		{
			name: "Max - default",
			args: args{
				in: func() kopsv1alpha2.MinMaxSpec {
					subject := kopsv1alpha2.MinMaxSpec{}
					subject.Max = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Min - default",
			args: args{
				in: func() kopsv1alpha2.MinMaxSpec {
					subject := kopsv1alpha2.MinMaxSpec{}
					subject.Min = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenResourceMinMaxSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceMinMaxSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceMinMaxSpec(t *testing.T) {
	_default := map[string]interface{}{
		"max": nil,
		"min": nil,
	}
	type args struct {
		in kopsv1alpha2.MinMaxSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.MinMaxSpec{},
			},
			want: _default,
		},
		{
			name: "Max - default",
			args: args{
				in: func() kopsv1alpha2.MinMaxSpec {
					subject := kopsv1alpha2.MinMaxSpec{}
					subject.Max = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Min - default",
			args: args{
				in: func() kopsv1alpha2.MinMaxSpec {
					subject := kopsv1alpha2.MinMaxSpec{}
					subject.Min = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceMinMaxSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceMinMaxSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
