package schemas

import (
	"reflect"
	"testing"

	"github.com/eddycharly/terraform-provider-kops/pkg/api/resources"
	"github.com/google/go-cmp/cmp"
)

func TestExpandResourceValidateOptions(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want resources.ValidateOptions
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandResourceValidateOptions(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandResourceValidateOptions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenResourceValidateOptionsInto(t *testing.T) {
	type args struct {
		in  resources.ValidateOptions
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
			FlattenResourceValidateOptionsInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenResourceValidateOptions(t *testing.T) {
	type args struct {
		in resources.ValidateOptions
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: resources.ValidateOptions{},
			},
			want: map[string]interface{}{
				"skip":          false,
				"timeout":       nil,
				"poll_interval": nil,
			},
		},
		{
			name: "Skip - default",
			args: args{
				in: func() resources.ValidateOptions {
					subject := resources.ValidateOptions{}
					subject.Skip = false
					return subject
				}(),
			},
			want: map[string]interface{}{
				"skip":          false,
				"timeout":       nil,
				"poll_interval": nil,
			},
		},
		{
			name: "Timeout - default",
			args: args{
				in: func() resources.ValidateOptions {
					subject := resources.ValidateOptions{}
					subject.Timeout = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"skip":          false,
				"timeout":       nil,
				"poll_interval": nil,
			},
		},
		{
			name: "PollInterval - default",
			args: args{
				in: func() resources.ValidateOptions {
					subject := resources.ValidateOptions{}
					subject.PollInterval = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"skip":          false,
				"timeout":       nil,
				"poll_interval": nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenResourceValidateOptions(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenResourceValidateOptions() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
