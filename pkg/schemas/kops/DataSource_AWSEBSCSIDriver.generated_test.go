package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceAWSEBSCSIDriver(t *testing.T) {
	_default := kops.AWSEBSCSIDriver{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.AWSEBSCSIDriver
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"enabled":             nil,
					"version":             nil,
					"volume_attach_limit": nil,
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandDataSourceAWSEBSCSIDriver(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceAWSEBSCSIDriver() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceAWSEBSCSIDriverInto(t *testing.T) {
	_default := map[string]interface{}{
		"enabled":             nil,
		"version":             nil,
		"volume_attach_limit": nil,
	}
	type args struct {
		in kops.AWSEBSCSIDriver
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.AWSEBSCSIDriver{},
			},
			want: _default,
		},
		{
			name: "Enabled - default",
			args: args{
				in: func() kops.AWSEBSCSIDriver {
					subject := kops.AWSEBSCSIDriver{}
					subject.Enabled = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Version - default",
			args: args{
				in: func() kops.AWSEBSCSIDriver {
					subject := kops.AWSEBSCSIDriver{}
					subject.Version = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "VolumeAttachLimit - default",
			args: args{
				in: func() kops.AWSEBSCSIDriver {
					subject := kops.AWSEBSCSIDriver{}
					subject.VolumeAttachLimit = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenDataSourceAWSEBSCSIDriverInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceAWSEBSCSIDriver() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceAWSEBSCSIDriver(t *testing.T) {
	_default := map[string]interface{}{
		"enabled":             nil,
		"version":             nil,
		"volume_attach_limit": nil,
	}
	type args struct {
		in kops.AWSEBSCSIDriver
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.AWSEBSCSIDriver{},
			},
			want: _default,
		},
		{
			name: "Enabled - default",
			args: args{
				in: func() kops.AWSEBSCSIDriver {
					subject := kops.AWSEBSCSIDriver{}
					subject.Enabled = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Version - default",
			args: args{
				in: func() kops.AWSEBSCSIDriver {
					subject := kops.AWSEBSCSIDriver{}
					subject.Version = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "VolumeAttachLimit - default",
			args: args{
				in: func() kops.AWSEBSCSIDriver {
					subject := kops.AWSEBSCSIDriver{}
					subject.VolumeAttachLimit = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceAWSEBSCSIDriver(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceAWSEBSCSIDriver() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
