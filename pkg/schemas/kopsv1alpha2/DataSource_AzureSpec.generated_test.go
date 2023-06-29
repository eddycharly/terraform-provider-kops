package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

func TestExpandDataSourceAzureSpec(t *testing.T) {
	_default := kopsv1alpha2.AzureSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kopsv1alpha2.AzureSpec
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
		in kopsv1alpha2.AzureSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.AzureSpec{},
			},
			want: _default,
		},
		{
			name: "SubscriptionId - default",
			args: args{
				in: func() kopsv1alpha2.AzureSpec {
					subject := kopsv1alpha2.AzureSpec{}
					subject.SubscriptionID = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "TenantId - default",
			args: args{
				in: func() kopsv1alpha2.AzureSpec {
					subject := kopsv1alpha2.AzureSpec{}
					subject.TenantID = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ResourceGroupName - default",
			args: args{
				in: func() kopsv1alpha2.AzureSpec {
					subject := kopsv1alpha2.AzureSpec{}
					subject.ResourceGroupName = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "RouteTableName - default",
			args: args{
				in: func() kopsv1alpha2.AzureSpec {
					subject := kopsv1alpha2.AzureSpec{}
					subject.RouteTableName = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AdminUser - default",
			args: args{
				in: func() kopsv1alpha2.AzureSpec {
					subject := kopsv1alpha2.AzureSpec{}
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
		in kopsv1alpha2.AzureSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.AzureSpec{},
			},
			want: _default,
		},
		{
			name: "SubscriptionId - default",
			args: args{
				in: func() kopsv1alpha2.AzureSpec {
					subject := kopsv1alpha2.AzureSpec{}
					subject.SubscriptionID = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "TenantId - default",
			args: args{
				in: func() kopsv1alpha2.AzureSpec {
					subject := kopsv1alpha2.AzureSpec{}
					subject.TenantID = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ResourceGroupName - default",
			args: args{
				in: func() kopsv1alpha2.AzureSpec {
					subject := kopsv1alpha2.AzureSpec{}
					subject.ResourceGroupName = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "RouteTableName - default",
			args: args{
				in: func() kopsv1alpha2.AzureSpec {
					subject := kopsv1alpha2.AzureSpec{}
					subject.RouteTableName = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AdminUser - default",
			args: args{
				in: func() kopsv1alpha2.AzureSpec {
					subject := kopsv1alpha2.AzureSpec{}
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
