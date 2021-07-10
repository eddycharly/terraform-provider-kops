package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceCloudConfiguration(t *testing.T) {
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
					"gce_service_account":            "",
					"disable_security_group_ingress": nil,
					"elb_security_group":             nil,
					"v_sphere_username":              nil,
					"v_sphere_password":              nil,
					"v_sphere_server":                nil,
					"v_sphere_datacenter":            nil,
					"v_sphere_resource_pool":         nil,
					"v_sphere_datastore":             nil,
					"v_sphere_core_dns_server":       nil,
					"spotinst_product":               nil,
					"spotinst_orientation":           nil,
					"openstack":                      nil,
					"azure":                          nil,
					"aws_ebs_csi_driver":             nil,
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandDataSourceCloudConfiguration(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceCloudConfiguration() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceCloudConfigurationInto(t *testing.T) {
	_default := map[string]interface{}{
		"manage_storage_classes":         nil,
		"multizone":                      nil,
		"node_tags":                      nil,
		"node_instance_prefix":           nil,
		"gce_service_account":            "",
		"disable_security_group_ingress": nil,
		"elb_security_group":             nil,
		"v_sphere_username":              nil,
		"v_sphere_password":              nil,
		"v_sphere_server":                nil,
		"v_sphere_datacenter":            nil,
		"v_sphere_resource_pool":         nil,
		"v_sphere_datastore":             nil,
		"v_sphere_core_dns_server":       nil,
		"spotinst_product":               nil,
		"spotinst_orientation":           nil,
		"openstack":                      nil,
		"azure":                          nil,
		"aws_ebs_csi_driver":             nil,
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
			name: "VSphereUsername - default",
			args: args{
				in: func() kops.CloudConfiguration {
					subject := kops.CloudConfiguration{}
					subject.VSphereUsername = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "VSpherePassword - default",
			args: args{
				in: func() kops.CloudConfiguration {
					subject := kops.CloudConfiguration{}
					subject.VSpherePassword = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "VSphereServer - default",
			args: args{
				in: func() kops.CloudConfiguration {
					subject := kops.CloudConfiguration{}
					subject.VSphereServer = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "VSphereDatacenter - default",
			args: args{
				in: func() kops.CloudConfiguration {
					subject := kops.CloudConfiguration{}
					subject.VSphereDatacenter = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "VSphereResourcePool - default",
			args: args{
				in: func() kops.CloudConfiguration {
					subject := kops.CloudConfiguration{}
					subject.VSphereResourcePool = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "VSphereDatastore - default",
			args: args{
				in: func() kops.CloudConfiguration {
					subject := kops.CloudConfiguration{}
					subject.VSphereDatastore = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "VSphereCoreDnsServer - default",
			args: args{
				in: func() kops.CloudConfiguration {
					subject := kops.CloudConfiguration{}
					subject.VSphereCoreDNSServer = nil
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
			name: "Openstack - default",
			args: args{
				in: func() kops.CloudConfiguration {
					subject := kops.CloudConfiguration{}
					subject.Openstack = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Azure - default",
			args: args{
				in: func() kops.CloudConfiguration {
					subject := kops.CloudConfiguration{}
					subject.Azure = nil
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenDataSourceCloudConfigurationInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceCloudConfiguration() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceCloudConfiguration(t *testing.T) {
	_default := map[string]interface{}{
		"manage_storage_classes":         nil,
		"multizone":                      nil,
		"node_tags":                      nil,
		"node_instance_prefix":           nil,
		"gce_service_account":            "",
		"disable_security_group_ingress": nil,
		"elb_security_group":             nil,
		"v_sphere_username":              nil,
		"v_sphere_password":              nil,
		"v_sphere_server":                nil,
		"v_sphere_datacenter":            nil,
		"v_sphere_resource_pool":         nil,
		"v_sphere_datastore":             nil,
		"v_sphere_core_dns_server":       nil,
		"spotinst_product":               nil,
		"spotinst_orientation":           nil,
		"openstack":                      nil,
		"azure":                          nil,
		"aws_ebs_csi_driver":             nil,
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
			name: "VSphereUsername - default",
			args: args{
				in: func() kops.CloudConfiguration {
					subject := kops.CloudConfiguration{}
					subject.VSphereUsername = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "VSpherePassword - default",
			args: args{
				in: func() kops.CloudConfiguration {
					subject := kops.CloudConfiguration{}
					subject.VSpherePassword = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "VSphereServer - default",
			args: args{
				in: func() kops.CloudConfiguration {
					subject := kops.CloudConfiguration{}
					subject.VSphereServer = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "VSphereDatacenter - default",
			args: args{
				in: func() kops.CloudConfiguration {
					subject := kops.CloudConfiguration{}
					subject.VSphereDatacenter = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "VSphereResourcePool - default",
			args: args{
				in: func() kops.CloudConfiguration {
					subject := kops.CloudConfiguration{}
					subject.VSphereResourcePool = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "VSphereDatastore - default",
			args: args{
				in: func() kops.CloudConfiguration {
					subject := kops.CloudConfiguration{}
					subject.VSphereDatastore = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "VSphereCoreDnsServer - default",
			args: args{
				in: func() kops.CloudConfiguration {
					subject := kops.CloudConfiguration{}
					subject.VSphereCoreDNSServer = nil
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
			name: "Openstack - default",
			args: args{
				in: func() kops.CloudConfiguration {
					subject := kops.CloudConfiguration{}
					subject.Openstack = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Azure - default",
			args: args{
				in: func() kops.CloudConfiguration {
					subject := kops.CloudConfiguration{}
					subject.Azure = nil
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceCloudConfiguration(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceCloudConfiguration() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
