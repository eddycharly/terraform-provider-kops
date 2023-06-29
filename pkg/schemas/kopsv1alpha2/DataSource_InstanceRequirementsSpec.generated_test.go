package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

func TestExpandDataSourceInstanceRequirementsSpec(t *testing.T) {
	_default := kopsv1alpha2.InstanceRequirementsSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kopsv1alpha2.InstanceRequirementsSpec
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
			got := ExpandDataSourceInstanceRequirementsSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceInstanceRequirementsSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceInstanceRequirementsSpecInto(t *testing.T) {
	_default := map[string]interface{}{
		"cpu":    nil,
		"memory": nil,
	}
	type args struct {
		in kopsv1alpha2.InstanceRequirementsSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.InstanceRequirementsSpec{},
			},
			want: _default,
		},
		{
			name: "Cpu - default",
			args: args{
				in: func() kopsv1alpha2.InstanceRequirementsSpec {
					subject := kopsv1alpha2.InstanceRequirementsSpec{}
					subject.CPU = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Memory - default",
			args: args{
				in: func() kopsv1alpha2.InstanceRequirementsSpec {
					subject := kopsv1alpha2.InstanceRequirementsSpec{}
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
			FlattenDataSourceInstanceRequirementsSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceInstanceRequirementsSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceInstanceRequirementsSpec(t *testing.T) {
	_default := map[string]interface{}{
		"cpu":    nil,
		"memory": nil,
	}
	type args struct {
		in kopsv1alpha2.InstanceRequirementsSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.InstanceRequirementsSpec{},
			},
			want: _default,
		},
		{
			name: "Cpu - default",
			args: args{
				in: func() kopsv1alpha2.InstanceRequirementsSpec {
					subject := kopsv1alpha2.InstanceRequirementsSpec{}
					subject.CPU = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Memory - default",
			args: args{
				in: func() kopsv1alpha2.InstanceRequirementsSpec {
					subject := kopsv1alpha2.InstanceRequirementsSpec{}
					subject.Memory = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceInstanceRequirementsSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceInstanceRequirementsSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
