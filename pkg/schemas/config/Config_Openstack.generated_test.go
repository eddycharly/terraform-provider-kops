package schemas

import (
	"reflect"
	"testing"

	"github.com/eddycharly/terraform-provider-kops/pkg/api/config"
	"github.com/google/go-cmp/cmp"
)

func TestExpandConfigOpenstack(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want config.Openstack
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandConfigOpenstack(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandConfigOpenstack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenConfigOpenstackInto(t *testing.T) {
	type args struct {
		in  config.Openstack
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
			FlattenConfigOpenstackInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenConfigOpenstack(t *testing.T) {
	type args struct {
		in config.Openstack
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: config.Openstack{},
			},
			want: map[string]interface{}{
				"tenant_id":                     "",
				"tenant_name":                   "",
				"project_id":                    "",
				"project_name":                  "",
				"project_domain_id":             "",
				"project_domain_name":           "",
				"domain_id":                     "",
				"domain_name":                   "",
				"username":                      "",
				"password":                      "",
				"auth_url":                      "",
				"region_name":                   "",
				"application_credential_id":     "",
				"application_credential_secret": "",
			},
		},
		{
			name: "TenantId - default",
			args: args{
				in: func() config.Openstack {
					subject := config.Openstack{}
					subject.TenantId = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
				"tenant_id":                     "",
				"tenant_name":                   "",
				"project_id":                    "",
				"project_name":                  "",
				"project_domain_id":             "",
				"project_domain_name":           "",
				"domain_id":                     "",
				"domain_name":                   "",
				"username":                      "",
				"password":                      "",
				"auth_url":                      "",
				"region_name":                   "",
				"application_credential_id":     "",
				"application_credential_secret": "",
			},
		},
		{
			name: "TenantName - default",
			args: args{
				in: func() config.Openstack {
					subject := config.Openstack{}
					subject.TenantName = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
				"tenant_id":                     "",
				"tenant_name":                   "",
				"project_id":                    "",
				"project_name":                  "",
				"project_domain_id":             "",
				"project_domain_name":           "",
				"domain_id":                     "",
				"domain_name":                   "",
				"username":                      "",
				"password":                      "",
				"auth_url":                      "",
				"region_name":                   "",
				"application_credential_id":     "",
				"application_credential_secret": "",
			},
		},
		{
			name: "ProjectId - default",
			args: args{
				in: func() config.Openstack {
					subject := config.Openstack{}
					subject.ProjectId = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
				"tenant_id":                     "",
				"tenant_name":                   "",
				"project_id":                    "",
				"project_name":                  "",
				"project_domain_id":             "",
				"project_domain_name":           "",
				"domain_id":                     "",
				"domain_name":                   "",
				"username":                      "",
				"password":                      "",
				"auth_url":                      "",
				"region_name":                   "",
				"application_credential_id":     "",
				"application_credential_secret": "",
			},
		},
		{
			name: "ProjectName - default",
			args: args{
				in: func() config.Openstack {
					subject := config.Openstack{}
					subject.ProjectName = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
				"tenant_id":                     "",
				"tenant_name":                   "",
				"project_id":                    "",
				"project_name":                  "",
				"project_domain_id":             "",
				"project_domain_name":           "",
				"domain_id":                     "",
				"domain_name":                   "",
				"username":                      "",
				"password":                      "",
				"auth_url":                      "",
				"region_name":                   "",
				"application_credential_id":     "",
				"application_credential_secret": "",
			},
		},
		{
			name: "ProjectDomainId - default",
			args: args{
				in: func() config.Openstack {
					subject := config.Openstack{}
					subject.ProjectDomainId = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
				"tenant_id":                     "",
				"tenant_name":                   "",
				"project_id":                    "",
				"project_name":                  "",
				"project_domain_id":             "",
				"project_domain_name":           "",
				"domain_id":                     "",
				"domain_name":                   "",
				"username":                      "",
				"password":                      "",
				"auth_url":                      "",
				"region_name":                   "",
				"application_credential_id":     "",
				"application_credential_secret": "",
			},
		},
		{
			name: "ProjectDomainName - default",
			args: args{
				in: func() config.Openstack {
					subject := config.Openstack{}
					subject.ProjectDomainName = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
				"tenant_id":                     "",
				"tenant_name":                   "",
				"project_id":                    "",
				"project_name":                  "",
				"project_domain_id":             "",
				"project_domain_name":           "",
				"domain_id":                     "",
				"domain_name":                   "",
				"username":                      "",
				"password":                      "",
				"auth_url":                      "",
				"region_name":                   "",
				"application_credential_id":     "",
				"application_credential_secret": "",
			},
		},
		{
			name: "DomainId - default",
			args: args{
				in: func() config.Openstack {
					subject := config.Openstack{}
					subject.DomainId = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
				"tenant_id":                     "",
				"tenant_name":                   "",
				"project_id":                    "",
				"project_name":                  "",
				"project_domain_id":             "",
				"project_domain_name":           "",
				"domain_id":                     "",
				"domain_name":                   "",
				"username":                      "",
				"password":                      "",
				"auth_url":                      "",
				"region_name":                   "",
				"application_credential_id":     "",
				"application_credential_secret": "",
			},
		},
		{
			name: "DomainName - default",
			args: args{
				in: func() config.Openstack {
					subject := config.Openstack{}
					subject.DomainName = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
				"tenant_id":                     "",
				"tenant_name":                   "",
				"project_id":                    "",
				"project_name":                  "",
				"project_domain_id":             "",
				"project_domain_name":           "",
				"domain_id":                     "",
				"domain_name":                   "",
				"username":                      "",
				"password":                      "",
				"auth_url":                      "",
				"region_name":                   "",
				"application_credential_id":     "",
				"application_credential_secret": "",
			},
		},
		{
			name: "Username - default",
			args: args{
				in: func() config.Openstack {
					subject := config.Openstack{}
					subject.Username = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
				"tenant_id":                     "",
				"tenant_name":                   "",
				"project_id":                    "",
				"project_name":                  "",
				"project_domain_id":             "",
				"project_domain_name":           "",
				"domain_id":                     "",
				"domain_name":                   "",
				"username":                      "",
				"password":                      "",
				"auth_url":                      "",
				"region_name":                   "",
				"application_credential_id":     "",
				"application_credential_secret": "",
			},
		},
		{
			name: "Password - default",
			args: args{
				in: func() config.Openstack {
					subject := config.Openstack{}
					subject.Password = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
				"tenant_id":                     "",
				"tenant_name":                   "",
				"project_id":                    "",
				"project_name":                  "",
				"project_domain_id":             "",
				"project_domain_name":           "",
				"domain_id":                     "",
				"domain_name":                   "",
				"username":                      "",
				"password":                      "",
				"auth_url":                      "",
				"region_name":                   "",
				"application_credential_id":     "",
				"application_credential_secret": "",
			},
		},
		{
			name: "AuthUrl - default",
			args: args{
				in: func() config.Openstack {
					subject := config.Openstack{}
					subject.AuthUrl = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
				"tenant_id":                     "",
				"tenant_name":                   "",
				"project_id":                    "",
				"project_name":                  "",
				"project_domain_id":             "",
				"project_domain_name":           "",
				"domain_id":                     "",
				"domain_name":                   "",
				"username":                      "",
				"password":                      "",
				"auth_url":                      "",
				"region_name":                   "",
				"application_credential_id":     "",
				"application_credential_secret": "",
			},
		},
		{
			name: "RegionName - default",
			args: args{
				in: func() config.Openstack {
					subject := config.Openstack{}
					subject.RegionName = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
				"tenant_id":                     "",
				"tenant_name":                   "",
				"project_id":                    "",
				"project_name":                  "",
				"project_domain_id":             "",
				"project_domain_name":           "",
				"domain_id":                     "",
				"domain_name":                   "",
				"username":                      "",
				"password":                      "",
				"auth_url":                      "",
				"region_name":                   "",
				"application_credential_id":     "",
				"application_credential_secret": "",
			},
		},
		{
			name: "ApplicationCredentialId - default",
			args: args{
				in: func() config.Openstack {
					subject := config.Openstack{}
					subject.ApplicationCredentialId = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
				"tenant_id":                     "",
				"tenant_name":                   "",
				"project_id":                    "",
				"project_name":                  "",
				"project_domain_id":             "",
				"project_domain_name":           "",
				"domain_id":                     "",
				"domain_name":                   "",
				"username":                      "",
				"password":                      "",
				"auth_url":                      "",
				"region_name":                   "",
				"application_credential_id":     "",
				"application_credential_secret": "",
			},
		},
		{
			name: "ApplicationCredentialSecret - default",
			args: args{
				in: func() config.Openstack {
					subject := config.Openstack{}
					subject.ApplicationCredentialSecret = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
				"tenant_id":                     "",
				"tenant_name":                   "",
				"project_id":                    "",
				"project_name":                  "",
				"project_domain_id":             "",
				"project_domain_name":           "",
				"domain_id":                     "",
				"domain_name":                   "",
				"username":                      "",
				"password":                      "",
				"auth_url":                      "",
				"region_name":                   "",
				"application_credential_id":     "",
				"application_credential_secret": "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenConfigOpenstack(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenConfigOpenstack() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
