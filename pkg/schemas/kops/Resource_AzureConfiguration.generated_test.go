package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceAzureConfiguration(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.AzureConfiguration
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandResourceAzureConfiguration(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandResourceAzureConfiguration() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenResourceAzureConfigurationInto(t *testing.T) {
	type args struct {
		in  kops.AzureConfiguration
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
			FlattenResourceAzureConfigurationInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenResourceAzureConfiguration(t *testing.T) {
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
			want: map[string]interface{}{
				"subscription_id":     "",
				"tenant_id":           "",
				"resource_group_name": "",
				"route_table_name":    "",
				"admin_user":          "",
			},
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
			want: map[string]interface{}{
				"subscription_id":     "",
				"tenant_id":           "",
				"resource_group_name": "",
				"route_table_name":    "",
				"admin_user":          "",
			},
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
			want: map[string]interface{}{
				"subscription_id":     "",
				"tenant_id":           "",
				"resource_group_name": "",
				"route_table_name":    "",
				"admin_user":          "",
			},
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
			want: map[string]interface{}{
				"subscription_id":     "",
				"tenant_id":           "",
				"resource_group_name": "",
				"route_table_name":    "",
				"admin_user":          "",
			},
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
			want: map[string]interface{}{
				"subscription_id":     "",
				"tenant_id":           "",
				"resource_group_name": "",
				"route_table_name":    "",
				"admin_user":          "",
			},
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
			want: map[string]interface{}{
				"subscription_id":     "",
				"tenant_id":           "",
				"resource_group_name": "",
				"route_table_name":    "",
				"admin_user":          "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenResourceAzureConfiguration(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenResourceAzureConfiguration() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
