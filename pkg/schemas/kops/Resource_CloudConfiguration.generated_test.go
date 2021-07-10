package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceCloudConfiguration(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.CloudConfiguration
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandResourceCloudConfiguration(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandResourceCloudConfiguration() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenResourceCloudConfigurationInto(t *testing.T) {
	type args struct {
		in  kops.CloudConfiguration
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
			FlattenResourceCloudConfigurationInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenResourceCloudConfiguration(t *testing.T) {
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
			want: map[string]interface{}{
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
		{
			name: "ManageStorageClasses - default",
			args: args{
				in: func() kops.CloudConfiguration {
					subject := kops.CloudConfiguration{}
					subject.ManageStorageClasses = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
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
		{
			name: "Multizone - default",
			args: args{
				in: func() kops.CloudConfiguration {
					subject := kops.CloudConfiguration{}
					subject.Multizone = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
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
		{
			name: "NodeTags - default",
			args: args{
				in: func() kops.CloudConfiguration {
					subject := kops.CloudConfiguration{}
					subject.NodeTags = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
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
		{
			name: "NodeInstancePrefix - default",
			args: args{
				in: func() kops.CloudConfiguration {
					subject := kops.CloudConfiguration{}
					subject.NodeInstancePrefix = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
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
		{
			name: "GCEServiceAccount - default",
			args: args{
				in: func() kops.CloudConfiguration {
					subject := kops.CloudConfiguration{}
					subject.GCEServiceAccount = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
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
		{
			name: "DisableSecurityGroupIngress - default",
			args: args{
				in: func() kops.CloudConfiguration {
					subject := kops.CloudConfiguration{}
					subject.DisableSecurityGroupIngress = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
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
		{
			name: "ElbSecurityGroup - default",
			args: args{
				in: func() kops.CloudConfiguration {
					subject := kops.CloudConfiguration{}
					subject.ElbSecurityGroup = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
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
		{
			name: "VSphereUsername - default",
			args: args{
				in: func() kops.CloudConfiguration {
					subject := kops.CloudConfiguration{}
					subject.VSphereUsername = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
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
		{
			name: "VSpherePassword - default",
			args: args{
				in: func() kops.CloudConfiguration {
					subject := kops.CloudConfiguration{}
					subject.VSpherePassword = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
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
		{
			name: "VSphereServer - default",
			args: args{
				in: func() kops.CloudConfiguration {
					subject := kops.CloudConfiguration{}
					subject.VSphereServer = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
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
		{
			name: "VSphereDatacenter - default",
			args: args{
				in: func() kops.CloudConfiguration {
					subject := kops.CloudConfiguration{}
					subject.VSphereDatacenter = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
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
		{
			name: "VSphereResourcePool - default",
			args: args{
				in: func() kops.CloudConfiguration {
					subject := kops.CloudConfiguration{}
					subject.VSphereResourcePool = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
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
		{
			name: "VSphereDatastore - default",
			args: args{
				in: func() kops.CloudConfiguration {
					subject := kops.CloudConfiguration{}
					subject.VSphereDatastore = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
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
		{
			name: "VSphereCoreDnsServer - default",
			args: args{
				in: func() kops.CloudConfiguration {
					subject := kops.CloudConfiguration{}
					subject.VSphereCoreDNSServer = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
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
		{
			name: "SpotinstProduct - default",
			args: args{
				in: func() kops.CloudConfiguration {
					subject := kops.CloudConfiguration{}
					subject.SpotinstProduct = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
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
		{
			name: "SpotinstOrientation - default",
			args: args{
				in: func() kops.CloudConfiguration {
					subject := kops.CloudConfiguration{}
					subject.SpotinstOrientation = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
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
		{
			name: "Openstack - default",
			args: args{
				in: func() kops.CloudConfiguration {
					subject := kops.CloudConfiguration{}
					subject.Openstack = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
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
		{
			name: "Azure - default",
			args: args{
				in: func() kops.CloudConfiguration {
					subject := kops.CloudConfiguration{}
					subject.Azure = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
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
		{
			name: "AwsEbsCsiDriver - default",
			args: args{
				in: func() kops.CloudConfiguration {
					subject := kops.CloudConfiguration{}
					subject.AWSEBSCSIDriver = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenResourceCloudConfiguration(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenResourceCloudConfiguration() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
