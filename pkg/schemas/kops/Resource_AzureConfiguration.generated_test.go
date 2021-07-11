package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceAzureConfiguration(t *testing.T) {
	_default := kops.AzureConfiguration{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.AzureConfiguration
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"subscription_id":     "",
					"tenant_id":           "",
					"resource_group_name": "",
					"route_table_name":    "",
					"admin_user":          "",
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandResourceAzureConfiguration(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceAzureConfiguration() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceAzureConfigurationInto(t *testing.T) {
	_default := map[string]interface{}{
		"subscription_id":     "",
		"tenant_id":           "",
		"resource_group_name": "",
		"route_table_name":    "",
		"admin_user":          "",
	}
	type args struct {
		in kops.AzureConfiguration
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.AzureConfiguration{},
			},
			want: _default,
		},
		{
			name: "SubscriptionID - default",
			args: args{
				in: func() kops.AzureConfiguration {
					subject := kops.AzureConfiguration{}
					subject.SubscriptionID = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "TenantID - default",
			args: args{
				in: func() kops.AzureConfiguration {
					subject := kops.AzureConfiguration{}
					subject.TenantID = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ResourceGroupName - default",
			args: args{
				in: func() kops.AzureConfiguration {
					subject := kops.AzureConfiguration{}
					subject.ResourceGroupName = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "RouteTableName - default",
			args: args{
				in: func() kops.AzureConfiguration {
					subject := kops.AzureConfiguration{}
					subject.RouteTableName = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AdminUser - default",
			args: args{
				in: func() kops.AzureConfiguration {
					subject := kops.AzureConfiguration{}
					subject.AdminUser = ""
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenResourceAzureConfigurationInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceAzureConfiguration() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceAzureConfiguration(t *testing.T) {
	_default := map[string]interface{}{
		"subscription_id":     "",
		"tenant_id":           "",
		"resource_group_name": "",
		"route_table_name":    "",
		"admin_user":          "",
	}
	type args struct {
		in kops.AzureConfiguration
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.AzureConfiguration{},
			},
			want: _default,
		},
		{
			name: "SubscriptionID - default",
			args: args{
				in: func() kops.AzureConfiguration {
					subject := kops.AzureConfiguration{}
					subject.SubscriptionID = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "TenantID - default",
			args: args{
				in: func() kops.AzureConfiguration {
					subject := kops.AzureConfiguration{}
					subject.TenantID = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ResourceGroupName - default",
			args: args{
				in: func() kops.AzureConfiguration {
					subject := kops.AzureConfiguration{}
					subject.ResourceGroupName = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "RouteTableName - default",
			args: args{
				in: func() kops.AzureConfiguration {
					subject := kops.AzureConfiguration{}
					subject.RouteTableName = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AdminUser - default",
			args: args{
				in: func() kops.AzureConfiguration {
					subject := kops.AzureConfiguration{}
					subject.AdminUser = ""
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceAzureConfiguration(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceAzureConfiguration() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
