package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceEtcdMemberSpec(t *testing.T) {
	_default := kops.EtcdMemberSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.EtcdMemberSpec
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
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
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandDataSourceEtcdMemberSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceEtcdMemberSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceEtcdMemberSpecInto(t *testing.T) {
	_default := map[string]interface{}{
		"name":              "",
		"instance_group":    nil,
		"volume_type":       nil,
		"volume_iops":       nil,
		"volume_throughput": nil,
		"volume_size":       nil,
		"kms_key_id":        nil,
		"encrypted_volume":  nil,
	}
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenDataSourceEtcdMemberSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceEtcdMemberSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceEtcdMemberSpec(t *testing.T) {
	_default := map[string]interface{}{
		"name":              "",
		"instance_group":    nil,
		"volume_type":       nil,
		"volume_iops":       nil,
		"volume_throughput": nil,
		"volume_size":       nil,
		"kms_key_id":        nil,
		"encrypted_volume":  nil,
	}
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceEtcdMemberSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceEtcdMemberSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
