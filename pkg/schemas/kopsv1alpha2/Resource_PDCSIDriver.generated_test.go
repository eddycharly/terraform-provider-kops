package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

func TestExpandResourcePDCSIDriver(t *testing.T) {
	_default := kopsv1alpha2.PDCSIDriver{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kopsv1alpha2.PDCSIDriver
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"enabled": nil,
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandResourcePDCSIDriver(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourcePDCSIDriver() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourcePDCSIDriverInto(t *testing.T) {
	_default := map[string]interface{}{
		"enabled": nil,
	}
	type args struct {
		in kopsv1alpha2.PDCSIDriver
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.PDCSIDriver{},
			},
			want: _default,
		},
		{
			name: "Enabled - default",
			args: args{
				in: func() kopsv1alpha2.PDCSIDriver {
					subject := kopsv1alpha2.PDCSIDriver{}
					subject.Enabled = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenResourcePDCSIDriverInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourcePDCSIDriver() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourcePDCSIDriver(t *testing.T) {
	_default := map[string]interface{}{
		"enabled": nil,
	}
	type args struct {
		in kopsv1alpha2.PDCSIDriver
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.PDCSIDriver{},
			},
			want: _default,
		},
		{
			name: "Enabled - default",
			args: args{
				in: func() kopsv1alpha2.PDCSIDriver {
					subject := kopsv1alpha2.PDCSIDriver{}
					subject.Enabled = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourcePDCSIDriver(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourcePDCSIDriver() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
