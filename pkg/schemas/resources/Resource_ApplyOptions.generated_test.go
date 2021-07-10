package schemas

import (
	"reflect"
	"testing"

	"github.com/eddycharly/terraform-provider-kops/pkg/api/resources"
	"github.com/google/go-cmp/cmp"
)

func TestExpandResourceApplyOptions(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want resources.ApplyOptions
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandResourceApplyOptions(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandResourceApplyOptions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenResourceApplyOptionsInto(t *testing.T) {
	type args struct {
		in  resources.ApplyOptions
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
			FlattenResourceApplyOptionsInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenResourceApplyOptions(t *testing.T) {
	type args struct {
		in resources.ApplyOptions
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: resources.ApplyOptions{},
			},
			want: map[string]interface{}{
				"skip":                 false,
				"allow_kops_downgrade": false,
			},
		},
		{
			name: "Skip - default",
			args: args{
				in: func() resources.ApplyOptions {
					subject := resources.ApplyOptions{}
					subject.Skip = false
					return subject
				}(),
			},
			want: map[string]interface{}{
				"skip":                 false,
				"allow_kops_downgrade": false,
			},
		},
		{
			name: "AllowKopsDowngrade - default",
			args: args{
				in: func() resources.ApplyOptions {
					subject := resources.ApplyOptions{}
					subject.AllowKopsDowngrade = false
					return subject
				}(),
			},
			want: map[string]interface{}{
				"skip":                 false,
				"allow_kops_downgrade": false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenResourceApplyOptions(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenResourceApplyOptions() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
