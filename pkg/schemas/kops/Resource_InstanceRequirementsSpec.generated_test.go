package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceInstanceRequirementsSpec(t *testing.T) {
	_default := kops.InstanceRequirementsSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.InstanceRequirementsSpec
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"cpu":    nil,
					"memory": nil,
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandResourceInstanceRequirementsSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceInstanceRequirementsSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceInstanceRequirementsSpecInto(t *testing.T) {
	_default := map[string]interface{}{
		"cpu":    nil,
		"memory": nil,
	}
	type args struct {
		in kops.InstanceRequirementsSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.InstanceRequirementsSpec{},
			},
			want: _default,
		},
		{
			name: "Cpu - default",
			args: args{
				in: func() kops.InstanceRequirementsSpec {
					subject := kops.InstanceRequirementsSpec{}
					subject.CPU = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Memory - default",
			args: args{
				in: func() kops.InstanceRequirementsSpec {
					subject := kops.InstanceRequirementsSpec{}
					subject.Memory = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenResourceInstanceRequirementsSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceInstanceRequirementsSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceInstanceRequirementsSpec(t *testing.T) {
	_default := map[string]interface{}{
		"cpu":    nil,
		"memory": nil,
	}
	type args struct {
		in kops.InstanceRequirementsSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.InstanceRequirementsSpec{},
			},
			want: _default,
		},
		{
			name: "Cpu - default",
			args: args{
				in: func() kops.InstanceRequirementsSpec {
					subject := kops.InstanceRequirementsSpec{}
					subject.CPU = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Memory - default",
			args: args{
				in: func() kops.InstanceRequirementsSpec {
					subject := kops.InstanceRequirementsSpec{}
					subject.Memory = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceInstanceRequirementsSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceInstanceRequirementsSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
