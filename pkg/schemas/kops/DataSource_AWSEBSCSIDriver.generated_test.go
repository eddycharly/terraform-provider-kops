package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceAWSEBSCSIDriver(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.AWSEBSCSIDriver
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandDataSourceAWSEBSCSIDriver(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandDataSourceAWSEBSCSIDriver() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenDataSourceAWSEBSCSIDriverInto(t *testing.T) {
	type args struct {
		in  kops.AWSEBSCSIDriver
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
			FlattenDataSourceAWSEBSCSIDriverInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenDataSourceAWSEBSCSIDriver(t *testing.T) {
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
			want: map[string]interface{}{
				"enabled":             nil,
				"version":             nil,
				"volume_attach_limit": nil,
			},
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
			want: map[string]interface{}{
				"enabled":             nil,
				"version":             nil,
				"volume_attach_limit": nil,
			},
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
			want: map[string]interface{}{
				"enabled":             nil,
				"version":             nil,
				"volume_attach_limit": nil,
			},
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
			want: map[string]interface{}{
				"enabled":             nil,
				"version":             nil,
				"volume_attach_limit": nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenDataSourceAWSEBSCSIDriver(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenDataSourceAWSEBSCSIDriver() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
