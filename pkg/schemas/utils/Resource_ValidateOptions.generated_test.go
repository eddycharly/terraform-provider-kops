package schemas

import (
	"testing"

	"github.com/eddycharly/terraform-provider-kops/pkg/api/utils"
	"github.com/google/go-cmp/cmp"
)

func TestExpandResourceValidateOptions(t *testing.T) {
	_default := utils.ValidateOptions{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want utils.ValidateOptions
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"timeout":       nil,
					"poll_interval": nil,
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandResourceValidateOptions(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceValidateOptions() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceValidateOptionsInto(t *testing.T) {
	_default := map[string]interface{}{
		"timeout":       nil,
		"poll_interval": nil,
	}
	type args struct {
		in utils.ValidateOptions
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: utils.ValidateOptions{},
			},
			want: _default,
		},
		{
			name: "Timeout - default",
			args: args{
				in: func() utils.ValidateOptions {
					subject := utils.ValidateOptions{}
					subject.Timeout = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PollInterval - default",
			args: args{
				in: func() utils.ValidateOptions {
					subject := utils.ValidateOptions{}
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
		"timeout":       nil,
		"poll_interval": nil,
	}
	type args struct {
		in utils.ValidateOptions
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: utils.ValidateOptions{},
			},
			want: _default,
		},
		{
			name: "Timeout - default",
			args: args{
				in: func() utils.ValidateOptions {
					subject := utils.ValidateOptions{}
					subject.Timeout = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PollInterval - default",
			args: args{
				in: func() utils.ValidateOptions {
					subject := utils.ValidateOptions{}
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
