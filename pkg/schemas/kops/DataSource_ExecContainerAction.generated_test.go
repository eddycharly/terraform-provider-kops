package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceExecContainerAction(t *testing.T) {
	_default := kops.ExecContainerAction{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.ExecContainerAction
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"image":       "",
					"command":     func() []interface{} { return nil }(),
					"environment": func() map[string]interface{} { return nil }(),
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandDataSourceExecContainerAction(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceExecContainerAction() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceExecContainerActionInto(t *testing.T) {
	_default := map[string]interface{}{
		"image":       "",
		"command":     func() []interface{} { return nil }(),
		"environment": func() map[string]interface{} { return nil }(),
	}
	type args struct {
		in kops.ExecContainerAction
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.ExecContainerAction{},
			},
			want: _default,
		},
		{
			name: "Image - default",
			args: args{
				in: func() kops.ExecContainerAction {
					subject := kops.ExecContainerAction{}
					subject.Image = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Command - default",
			args: args{
				in: func() kops.ExecContainerAction {
					subject := kops.ExecContainerAction{}
					subject.Command = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Environment - default",
			args: args{
				in: func() kops.ExecContainerAction {
					subject := kops.ExecContainerAction{}
					subject.Environment = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenDataSourceExecContainerActionInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceExecContainerAction() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceExecContainerAction(t *testing.T) {
	_default := map[string]interface{}{
		"image":       "",
		"command":     func() []interface{} { return nil }(),
		"environment": func() map[string]interface{} { return nil }(),
	}
	type args struct {
		in kops.ExecContainerAction
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.ExecContainerAction{},
			},
			want: _default,
		},
		{
			name: "Image - default",
			args: args{
				in: func() kops.ExecContainerAction {
					subject := kops.ExecContainerAction{}
					subject.Image = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Command - default",
			args: args{
				in: func() kops.ExecContainerAction {
					subject := kops.ExecContainerAction{}
					subject.Command = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Environment - default",
			args: args{
				in: func() kops.ExecContainerAction {
					subject := kops.ExecContainerAction{}
					subject.Environment = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceExecContainerAction(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceExecContainerAction() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
