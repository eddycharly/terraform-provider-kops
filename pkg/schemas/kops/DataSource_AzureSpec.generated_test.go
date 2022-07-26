package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceAzureSpec(t *testing.T) {
	_default := kops.AzureSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.AzureSpec
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
			got := ExpandDataSourceAzureSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceAzureSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceAzureSpecInto(t *testing.T) {
	_default := map[string]interface{}{
		"subscription_id":     "",
		"tenant_id":           "",
		"resource_group_name": "",
		"route_table_name":    "",
		"admin_user":          "",
	}
	type args struct {
		in kops.AzureSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.AzureSpec{},
			},
			want: _default,
		},
		{
			name: "SubscriptionID - default",
			args: args{
				in: func() kops.AzureSpec {
					subject := kops.AzureSpec{}
					subject.SubscriptionID = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "TenantID - default",
			args: args{
				in: func() kops.AzureSpec {
					subject := kops.AzureSpec{}
					subject.TenantID = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ResourceGroupName - default",
			args: args{
				in: func() kops.AzureSpec {
					subject := kops.AzureSpec{}
					subject.ResourceGroupName = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "RouteTableName - default",
			args: args{
				in: func() kops.AzureSpec {
					subject := kops.AzureSpec{}
					subject.RouteTableName = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AdminUser - default",
			args: args{
				in: func() kops.AzureSpec {
					subject := kops.AzureSpec{}
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
			FlattenDataSourceAzureSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceAzureSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceAzureSpec(t *testing.T) {
	_default := map[string]interface{}{
		"subscription_id":     "",
		"tenant_id":           "",
		"resource_group_name": "",
		"route_table_name":    "",
		"admin_user":          "",
	}
	type args struct {
		in kops.AzureSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.AzureSpec{},
			},
			want: _default,
		},
		{
			name: "SubscriptionID - default",
			args: args{
				in: func() kops.AzureSpec {
					subject := kops.AzureSpec{}
					subject.SubscriptionID = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "TenantID - default",
			args: args{
				in: func() kops.AzureSpec {
					subject := kops.AzureSpec{}
					subject.TenantID = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ResourceGroupName - default",
			args: args{
				in: func() kops.AzureSpec {
					subject := kops.AzureSpec{}
					subject.ResourceGroupName = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "RouteTableName - default",
			args: args{
				in: func() kops.AzureSpec {
					subject := kops.AzureSpec{}
					subject.RouteTableName = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AdminUser - default",
			args: args{
				in: func() kops.AzureSpec {
					subject := kops.AzureSpec{}
					subject.AdminUser = ""
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceAzureSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceAzureSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
