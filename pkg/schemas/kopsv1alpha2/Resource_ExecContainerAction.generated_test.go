package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

func TestExpandResourceExecContainerAction(t *testing.T) {
	_default := kopsv1alpha2.ExecContainerAction{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kopsv1alpha2.ExecContainerAction
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
			got := ExpandResourceExecContainerAction(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceExecContainerAction() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceExecContainerActionInto(t *testing.T) {
	_default := map[string]interface{}{
		"image":       "",
		"command":     func() []interface{} { return nil }(),
		"environment": func() map[string]interface{} { return nil }(),
	}
	type args struct {
		in kopsv1alpha2.ExecContainerAction
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.ExecContainerAction{},
			},
			want: _default,
		},
		{
			name: "Image - default",
			args: args{
				in: func() kopsv1alpha2.ExecContainerAction {
					subject := kopsv1alpha2.ExecContainerAction{}
					subject.Image = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Command - default",
			args: args{
				in: func() kopsv1alpha2.ExecContainerAction {
					subject := kopsv1alpha2.ExecContainerAction{}
					subject.Command = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Environment - default",
			args: args{
				in: func() kopsv1alpha2.ExecContainerAction {
					subject := kopsv1alpha2.ExecContainerAction{}
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
			FlattenResourceExecContainerActionInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceExecContainerAction() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceExecContainerAction(t *testing.T) {
	_default := map[string]interface{}{
		"image":       "",
		"command":     func() []interface{} { return nil }(),
		"environment": func() map[string]interface{} { return nil }(),
	}
	type args struct {
		in kopsv1alpha2.ExecContainerAction
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.ExecContainerAction{},
			},
			want: _default,
		},
		{
			name: "Image - default",
			args: args{
				in: func() kopsv1alpha2.ExecContainerAction {
					subject := kopsv1alpha2.ExecContainerAction{}
					subject.Image = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Command - default",
			args: args{
				in: func() kopsv1alpha2.ExecContainerAction {
					subject := kopsv1alpha2.ExecContainerAction{}
					subject.Command = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Environment - default",
			args: args{
				in: func() kopsv1alpha2.ExecContainerAction {
					subject := kopsv1alpha2.ExecContainerAction{}
					subject.Environment = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceExecContainerAction(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceExecContainerAction() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
