package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceExecContainerAction(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.ExecContainerAction
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandResourceExecContainerAction(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandResourceExecContainerAction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenResourceExecContainerActionInto(t *testing.T) {
	type args struct {
		in  kops.ExecContainerAction
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
			FlattenResourceExecContainerActionInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenResourceExecContainerAction(t *testing.T) {
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
			want: map[string]interface{}{
				"image":       "",
				"command":     func() []interface{} { return nil }(),
				"environment": func() map[string]interface{} { return nil }(),
			},
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
			want: map[string]interface{}{
				"image":       "",
				"command":     func() []interface{} { return nil }(),
				"environment": func() map[string]interface{} { return nil }(),
			},
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
			want: map[string]interface{}{
				"image":       "",
				"command":     func() []interface{} { return nil }(),
				"environment": func() map[string]interface{} { return nil }(),
			},
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
			want: map[string]interface{}{
				"image":       "",
				"command":     func() []interface{} { return nil }(),
				"environment": func() map[string]interface{} { return nil }(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenResourceExecContainerAction(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenResourceExecContainerAction() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
