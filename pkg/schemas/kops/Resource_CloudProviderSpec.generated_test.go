package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceCloudProviderSpec(t *testing.T) {
	_default := kops.CloudProviderSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.CloudProviderSpec
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"aws":       nil,
					"azure":     nil,
					"do":        nil,
					"gce":       nil,
					"hetzner":   nil,
					"openstack": nil,
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandResourceCloudProviderSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceCloudProviderSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceCloudProviderSpecInto(t *testing.T) {
	_default := map[string]interface{}{
		"aws":       nil,
		"azure":     nil,
		"do":        nil,
		"gce":       nil,
		"hetzner":   nil,
		"openstack": nil,
	}
	type args struct {
		in kops.CloudProviderSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.CloudProviderSpec{},
			},
			want: _default,
		},
		{
			name: "Aws - default",
			args: args{
				in: func() kops.CloudProviderSpec {
					subject := kops.CloudProviderSpec{}
					subject.AWS = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Azure - default",
			args: args{
				in: func() kops.CloudProviderSpec {
					subject := kops.CloudProviderSpec{}
					subject.Azure = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "DO - default",
			args: args{
				in: func() kops.CloudProviderSpec {
					subject := kops.CloudProviderSpec{}
					subject.DO = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "GCE - default",
			args: args{
				in: func() kops.CloudProviderSpec {
					subject := kops.CloudProviderSpec{}
					subject.GCE = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Hetzner - default",
			args: args{
				in: func() kops.CloudProviderSpec {
					subject := kops.CloudProviderSpec{}
					subject.Hetzner = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Openstack - default",
			args: args{
				in: func() kops.CloudProviderSpec {
					subject := kops.CloudProviderSpec{}
					subject.Openstack = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenResourceCloudProviderSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceCloudProviderSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceCloudProviderSpec(t *testing.T) {
	_default := map[string]interface{}{
		"aws":       nil,
		"azure":     nil,
		"do":        nil,
		"gce":       nil,
		"hetzner":   nil,
		"openstack": nil,
	}
	type args struct {
		in kops.CloudProviderSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.CloudProviderSpec{},
			},
			want: _default,
		},
		{
			name: "Aws - default",
			args: args{
				in: func() kops.CloudProviderSpec {
					subject := kops.CloudProviderSpec{}
					subject.AWS = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Azure - default",
			args: args{
				in: func() kops.CloudProviderSpec {
					subject := kops.CloudProviderSpec{}
					subject.Azure = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "DO - default",
			args: args{
				in: func() kops.CloudProviderSpec {
					subject := kops.CloudProviderSpec{}
					subject.DO = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "GCE - default",
			args: args{
				in: func() kops.CloudProviderSpec {
					subject := kops.CloudProviderSpec{}
					subject.GCE = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Hetzner - default",
			args: args{
				in: func() kops.CloudProviderSpec {
					subject := kops.CloudProviderSpec{}
					subject.Hetzner = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Openstack - default",
			args: args{
				in: func() kops.CloudProviderSpec {
					subject := kops.CloudProviderSpec{}
					subject.Openstack = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceCloudProviderSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceCloudProviderSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
