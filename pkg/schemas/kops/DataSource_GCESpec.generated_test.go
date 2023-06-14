package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceGCESpec(t *testing.T) {
	_default := kops.GCESpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.GCESpec
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"project":         "",
					"service_account": "",
					"pd_csi_driver":   nil,
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandDataSourceGCESpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceGCESpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceGCESpecInto(t *testing.T) {
	_default := map[string]interface{}{
		"project":         "",
		"service_account": "",
		"pd_csi_driver":   nil,
	}
	type args struct {
		in kops.GCESpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.GCESpec{},
			},
			want: _default,
		},
		{
			name: "Project - default",
			args: args{
				in: func() kops.GCESpec {
					subject := kops.GCESpec{}
					subject.Project = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ServiceAccount - default",
			args: args{
				in: func() kops.GCESpec {
					subject := kops.GCESpec{}
					subject.ServiceAccount = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PDCsiDriver - default",
			args: args{
				in: func() kops.GCESpec {
					subject := kops.GCESpec{}
					subject.PDCSIDriver = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenDataSourceGCESpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceGCESpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceGCESpec(t *testing.T) {
	_default := map[string]interface{}{
		"project":         "",
		"service_account": "",
		"pd_csi_driver":   nil,
	}
	type args struct {
		in kops.GCESpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.GCESpec{},
			},
			want: _default,
		},
		{
			name: "Project - default",
			args: args{
				in: func() kops.GCESpec {
					subject := kops.GCESpec{}
					subject.Project = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ServiceAccount - default",
			args: args{
				in: func() kops.GCESpec {
					subject := kops.GCESpec{}
					subject.ServiceAccount = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PDCsiDriver - default",
			args: args{
				in: func() kops.GCESpec {
					subject := kops.GCESpec{}
					subject.PDCSIDriver = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceGCESpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceGCESpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
