package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceEtcdMemberSpec(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.EtcdMemberSpec
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandDataSourceEtcdMemberSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandDataSourceEtcdMemberSpec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenDataSourceEtcdMemberSpecInto(t *testing.T) {
	type args struct {
		in  kops.EtcdMemberSpec
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
			FlattenDataSourceEtcdMemberSpecInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenDataSourceEtcdMemberSpec(t *testing.T) {
	type args struct {
		in kops.EtcdMemberSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.EtcdMemberSpec{},
			},
			want: map[string]interface{}{
				"name":              "",
				"instance_group":    nil,
				"volume_type":       nil,
				"volume_iops":       nil,
				"volume_throughput": nil,
				"volume_size":       nil,
				"kms_key_id":        nil,
				"encrypted_volume":  nil,
			},
		},
		{
			name: "Name - default",
			args: args{
				in: func() kops.EtcdMemberSpec {
					subject := kops.EtcdMemberSpec{}
					subject.Name = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
				"name":              "",
				"instance_group":    nil,
				"volume_type":       nil,
				"volume_iops":       nil,
				"volume_throughput": nil,
				"volume_size":       nil,
				"kms_key_id":        nil,
				"encrypted_volume":  nil,
			},
		},
		{
			name: "InstanceGroup - default",
			args: args{
				in: func() kops.EtcdMemberSpec {
					subject := kops.EtcdMemberSpec{}
					subject.InstanceGroup = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"name":              "",
				"instance_group":    nil,
				"volume_type":       nil,
				"volume_iops":       nil,
				"volume_throughput": nil,
				"volume_size":       nil,
				"kms_key_id":        nil,
				"encrypted_volume":  nil,
			},
		},
		{
			name: "VolumeType - default",
			args: args{
				in: func() kops.EtcdMemberSpec {
					subject := kops.EtcdMemberSpec{}
					subject.VolumeType = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"name":              "",
				"instance_group":    nil,
				"volume_type":       nil,
				"volume_iops":       nil,
				"volume_throughput": nil,
				"volume_size":       nil,
				"kms_key_id":        nil,
				"encrypted_volume":  nil,
			},
		},
		{
			name: "VolumeIops - default",
			args: args{
				in: func() kops.EtcdMemberSpec {
					subject := kops.EtcdMemberSpec{}
					subject.VolumeIops = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"name":              "",
				"instance_group":    nil,
				"volume_type":       nil,
				"volume_iops":       nil,
				"volume_throughput": nil,
				"volume_size":       nil,
				"kms_key_id":        nil,
				"encrypted_volume":  nil,
			},
		},
		{
			name: "VolumeThroughput - default",
			args: args{
				in: func() kops.EtcdMemberSpec {
					subject := kops.EtcdMemberSpec{}
					subject.VolumeThroughput = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"name":              "",
				"instance_group":    nil,
				"volume_type":       nil,
				"volume_iops":       nil,
				"volume_throughput": nil,
				"volume_size":       nil,
				"kms_key_id":        nil,
				"encrypted_volume":  nil,
			},
		},
		{
			name: "VolumeSize - default",
			args: args{
				in: func() kops.EtcdMemberSpec {
					subject := kops.EtcdMemberSpec{}
					subject.VolumeSize = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"name":              "",
				"instance_group":    nil,
				"volume_type":       nil,
				"volume_iops":       nil,
				"volume_throughput": nil,
				"volume_size":       nil,
				"kms_key_id":        nil,
				"encrypted_volume":  nil,
			},
		},
		{
			name: "KmsKeyId - default",
			args: args{
				in: func() kops.EtcdMemberSpec {
					subject := kops.EtcdMemberSpec{}
					subject.KmsKeyId = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"name":              "",
				"instance_group":    nil,
				"volume_type":       nil,
				"volume_iops":       nil,
				"volume_throughput": nil,
				"volume_size":       nil,
				"kms_key_id":        nil,
				"encrypted_volume":  nil,
			},
		},
		{
			name: "EncryptedVolume - default",
			args: args{
				in: func() kops.EtcdMemberSpec {
					subject := kops.EtcdMemberSpec{}
					subject.EncryptedVolume = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"name":              "",
				"instance_group":    nil,
				"volume_type":       nil,
				"volume_iops":       nil,
				"volume_throughput": nil,
				"volume_size":       nil,
				"kms_key_id":        nil,
				"encrypted_volume":  nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenDataSourceEtcdMemberSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenDataSourceEtcdMemberSpec() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
