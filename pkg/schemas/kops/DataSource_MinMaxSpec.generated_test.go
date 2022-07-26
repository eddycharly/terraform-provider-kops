package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceMinMaxSpec(t *testing.T) {
	_default := kops.MinMaxSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.MinMaxSpec
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
			got := ExpandDataSourceMinMaxSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceMinMaxSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceMinMaxSpecInto(t *testing.T) {
	_default := map[string]interface{}{
		"max": nil,
		"min": nil,
	}
	type args struct {
		in kops.MinMaxSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.MinMaxSpec{},
			},
			want: _default,
		},
		{
			name: "Max - default",
			args: args{
				in: func() kops.MinMaxSpec {
					subject := kops.MinMaxSpec{}
					subject.Max = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Min - default",
			args: args{
				in: func() kops.MinMaxSpec {
					subject := kops.MinMaxSpec{}
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
			FlattenDataSourceMinMaxSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceMinMaxSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceMinMaxSpec(t *testing.T) {
	_default := map[string]interface{}{
		"max": nil,
		"min": nil,
	}
	type args struct {
		in kops.MinMaxSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.MinMaxSpec{},
			},
			want: _default,
		},
		{
			name: "Max - default",
			args: args{
				in: func() kops.MinMaxSpec {
					subject := kops.MinMaxSpec{}
					subject.Max = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Min - default",
			args: args{
				in: func() kops.MinMaxSpec {
					subject := kops.MinMaxSpec{}
					subject.Min = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceMinMaxSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceMinMaxSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
