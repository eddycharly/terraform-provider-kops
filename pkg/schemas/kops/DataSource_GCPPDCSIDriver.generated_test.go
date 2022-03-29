package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceGCPPDCSIDriver(t *testing.T) {
	_default := kops.GCPPDCSIDriver{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.GCPPDCSIDriver
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
			got := ExpandDataSourceGCPPDCSIDriver(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceGCPPDCSIDriver() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceGCPPDCSIDriverInto(t *testing.T) {
	_default := map[string]interface{}{
		"enabled": nil,
	}
	type args struct {
		in kops.GCPPDCSIDriver
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.GCPPDCSIDriver{},
			},
			want: _default,
		},
		{
			name: "Enabled - default",
			args: args{
				in: func() kops.GCPPDCSIDriver {
					subject := kops.GCPPDCSIDriver{}
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
			FlattenDataSourceGCPPDCSIDriverInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceGCPPDCSIDriver() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceGCPPDCSIDriver(t *testing.T) {
	_default := map[string]interface{}{
		"enabled": nil,
	}
	type args struct {
		in kops.GCPPDCSIDriver
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.GCPPDCSIDriver{},
			},
			want: _default,
		},
		{
			name: "Enabled - default",
			args: args{
				in: func() kops.GCPPDCSIDriver {
					subject := kops.GCPPDCSIDriver{}
					subject.Enabled = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceGCPPDCSIDriver(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceGCPPDCSIDriver() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
