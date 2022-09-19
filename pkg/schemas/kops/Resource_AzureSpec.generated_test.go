package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceAzureSpec(t *testing.T) {
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
			got := ExpandResourceAzureSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceAzureSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceAzureSpecInto(t *testing.T) {
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
			name: "SubscriptionId - default",
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
			name: "TenantId - default",
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
			FlattenResourceAzureSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceAzureSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceAzureSpec(t *testing.T) {
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
			name: "SubscriptionId - default",
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
			name: "TenantId - default",
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
			got := FlattenResourceAzureSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceAzureSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
