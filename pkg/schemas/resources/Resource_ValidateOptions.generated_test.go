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
	_default := map[string]interface{}{
		"skip":          false,
		"timeout":       nil,
		"poll_interval": nil,
	}
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenResourceValidateOptionsInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceValidateOptions() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceValidateOptions(t *testing.T) {
	_default := map[string]interface{}{
		"skip":          false,
		"timeout":       nil,
		"poll_interval": nil,
	}
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceValidateOptions(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceValidateOptions() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
