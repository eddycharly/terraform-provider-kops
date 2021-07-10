package schemas

import (
	"testing"

	"github.com/eddycharly/terraform-provider-kops/pkg/api/config"
	"github.com/google/go-cmp/cmp"
)

func TestExpandConfigOpenstack(t *testing.T) {
	_default := config.Openstack{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want config.Openstack
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
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
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandConfigOpenstack(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandConfigOpenstack() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenConfigOpenstackInto(t *testing.T) {
	_default := map[string]interface{}{
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
	}
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenConfigOpenstackInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenConfigOpenstack() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenConfigOpenstack(t *testing.T) {
	_default := map[string]interface{}{
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
	}
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenConfigOpenstack(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenConfigOpenstack() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
