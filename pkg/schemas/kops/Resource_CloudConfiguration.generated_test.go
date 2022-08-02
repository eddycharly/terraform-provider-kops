package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceCloudConfiguration(t *testing.T) {
	_default := kops.CloudConfiguration{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.CloudConfiguration
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"manage_storage_classes":         nil,
					"multizone":                      nil,
					"node_tags":                      nil,
					"node_instance_prefix":           nil,
					"node_ip_families":               func() []interface{} { return nil }(),
					"gce_service_account":            "",
					"disable_security_group_ingress": nil,
					"elb_security_group":             nil,
					"spotinst_product":               nil,
					"spotinst_orientation":           nil,
					"aws_ebs_csi_driver":             nil,
					"gcp_pd_csi_driver":              nil,
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandResourceCloudConfiguration(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceCloudConfiguration() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceCloudConfigurationInto(t *testing.T) {
	_default := map[string]interface{}{
		"manage_storage_classes":         nil,
		"multizone":                      nil,
		"node_tags":                      nil,
		"node_instance_prefix":           nil,
		"node_ip_families":               func() []interface{} { return nil }(),
		"gce_service_account":            "",
		"disable_security_group_ingress": nil,
		"elb_security_group":             nil,
		"spotinst_product":               nil,
		"spotinst_orientation":           nil,
		"aws_ebs_csi_driver":             nil,
		"gcp_pd_csi_driver":              nil,
	}
	type args struct {
		in kops.CloudConfiguration
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.CloudConfiguration{},
			},
			want: _default,
		},
		{
			name: "ManageStorageClasses - default",
			args: args{
				in: func() kops.CloudConfiguration {
					subject := kops.CloudConfiguration{}
					subject.ManageStorageClasses = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Multizone - default",
			args: args{
				in: func() kops.CloudConfiguration {
					subject := kops.CloudConfiguration{}
					subject.Multizone = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NodeTags - default",
			args: args{
				in: func() kops.CloudConfiguration {
					subject := kops.CloudConfiguration{}
					subject.NodeTags = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NodeInstancePrefix - default",
			args: args{
				in: func() kops.CloudConfiguration {
					subject := kops.CloudConfiguration{}
					subject.NodeInstancePrefix = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NodeIpFamilies - default",
			args: args{
				in: func() kops.CloudConfiguration {
					subject := kops.CloudConfiguration{}
					subject.NodeIPFamilies = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "GCEServiceAccount - default",
			args: args{
				in: func() kops.CloudConfiguration {
					subject := kops.CloudConfiguration{}
					subject.GCEServiceAccount = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "DisableSecurityGroupIngress - default",
			args: args{
				in: func() kops.CloudConfiguration {
					subject := kops.CloudConfiguration{}
					subject.DisableSecurityGroupIngress = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ElbSecurityGroup - default",
			args: args{
				in: func() kops.CloudConfiguration {
					subject := kops.CloudConfiguration{}
					subject.ElbSecurityGroup = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SpotinstProduct - default",
			args: args{
				in: func() kops.CloudConfiguration {
					subject := kops.CloudConfiguration{}
					subject.SpotinstProduct = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SpotinstOrientation - default",
			args: args{
				in: func() kops.CloudConfiguration {
					subject := kops.CloudConfiguration{}
					subject.SpotinstOrientation = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AwsEbsCsiDriver - default",
			args: args{
				in: func() kops.CloudConfiguration {
					subject := kops.CloudConfiguration{}
					subject.AWSEBSCSIDriver = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "GcpPDCsiDriver - default",
			args: args{
				in: func() kops.CloudConfiguration {
					subject := kops.CloudConfiguration{}
					subject.GCPPDCSIDriver = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenResourceCloudConfigurationInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceCloudConfiguration() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceCloudConfiguration(t *testing.T) {
	_default := map[string]interface{}{
		"manage_storage_classes":         nil,
		"multizone":                      nil,
		"node_tags":                      nil,
		"node_instance_prefix":           nil,
		"node_ip_families":               func() []interface{} { return nil }(),
		"gce_service_account":            "",
		"disable_security_group_ingress": nil,
		"elb_security_group":             nil,
		"spotinst_product":               nil,
		"spotinst_orientation":           nil,
		"aws_ebs_csi_driver":             nil,
		"gcp_pd_csi_driver":              nil,
	}
	type args struct {
		in kops.CloudConfiguration
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.CloudConfiguration{},
			},
			want: _default,
		},
		{
			name: "ManageStorageClasses - default",
			args: args{
				in: func() kops.CloudConfiguration {
					subject := kops.CloudConfiguration{}
					subject.ManageStorageClasses = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Multizone - default",
			args: args{
				in: func() kops.CloudConfiguration {
					subject := kops.CloudConfiguration{}
					subject.Multizone = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NodeTags - default",
			args: args{
				in: func() kops.CloudConfiguration {
					subject := kops.CloudConfiguration{}
					subject.NodeTags = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NodeInstancePrefix - default",
			args: args{
				in: func() kops.CloudConfiguration {
					subject := kops.CloudConfiguration{}
					subject.NodeInstancePrefix = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NodeIpFamilies - default",
			args: args{
				in: func() kops.CloudConfiguration {
					subject := kops.CloudConfiguration{}
					subject.NodeIPFamilies = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "GCEServiceAccount - default",
			args: args{
				in: func() kops.CloudConfiguration {
					subject := kops.CloudConfiguration{}
					subject.GCEServiceAccount = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "DisableSecurityGroupIngress - default",
			args: args{
				in: func() kops.CloudConfiguration {
					subject := kops.CloudConfiguration{}
					subject.DisableSecurityGroupIngress = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ElbSecurityGroup - default",
			args: args{
				in: func() kops.CloudConfiguration {
					subject := kops.CloudConfiguration{}
					subject.ElbSecurityGroup = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SpotinstProduct - default",
			args: args{
				in: func() kops.CloudConfiguration {
					subject := kops.CloudConfiguration{}
					subject.SpotinstProduct = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SpotinstOrientation - default",
			args: args{
				in: func() kops.CloudConfiguration {
					subject := kops.CloudConfiguration{}
					subject.SpotinstOrientation = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AwsEbsCsiDriver - default",
			args: args{
				in: func() kops.CloudConfiguration {
					subject := kops.CloudConfiguration{}
					subject.AWSEBSCSIDriver = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "GcpPDCsiDriver - default",
			args: args{
				in: func() kops.CloudConfiguration {
					subject := kops.CloudConfiguration{}
					subject.GCPPDCSIDriver = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceCloudConfiguration(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceCloudConfiguration() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
