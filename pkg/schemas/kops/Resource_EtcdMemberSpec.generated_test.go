package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceEtcdMemberSpec(t *testing.T) {
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
			got := ExpandResourceEtcdMemberSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceEtcdMemberSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceEtcdMemberSpecInto(t *testing.T) {
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
			name: "VolumeIOPS - default",
			args: args{
				in: func() kops.EtcdMemberSpec {
					subject := kops.EtcdMemberSpec{}
					subject.VolumeIOPS = nil
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
			name: "KmsKeyID - default",
			args: args{
				in: func() kops.EtcdMemberSpec {
					subject := kops.EtcdMemberSpec{}
					subject.KmsKeyID = nil
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
			FlattenResourceEtcdMemberSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceEtcdMemberSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceEtcdMemberSpec(t *testing.T) {
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
			name: "VolumeIOPS - default",
			args: args{
				in: func() kops.EtcdMemberSpec {
					subject := kops.EtcdMemberSpec{}
					subject.VolumeIOPS = nil
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
			name: "KmsKeyID - default",
			args: args{
				in: func() kops.EtcdMemberSpec {
					subject := kops.EtcdMemberSpec{}
					subject.KmsKeyID = nil
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
			got := FlattenResourceEtcdMemberSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceEtcdMemberSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
