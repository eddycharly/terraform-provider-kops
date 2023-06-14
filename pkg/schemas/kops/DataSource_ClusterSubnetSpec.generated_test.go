package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceClusterSubnetSpec(t *testing.T) {
	_default := kops.ClusterSubnetSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.ClusterSubnetSpec
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"name":              "",
					"cidr":              "",
					"ipv6_cidr":         "",
					"zone":              "",
					"region":            "",
					"id":                "",
					"egress":            "",
					"type":              "",
					"public_ip":         "",
					"additional_routes": func() []interface{} { return nil }(),
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandDataSourceClusterSubnetSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceClusterSubnetSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceClusterSubnetSpecInto(t *testing.T) {
	_default := map[string]interface{}{
		"name":              "",
		"cidr":              "",
		"ipv6_cidr":         "",
		"zone":              "",
		"region":            "",
		"id":                "",
		"egress":            "",
		"type":              "",
		"public_ip":         "",
		"additional_routes": func() []interface{} { return nil }(),
	}
	type args struct {
		in kops.ClusterSubnetSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.ClusterSubnetSpec{},
			},
			want: _default,
		},
		{
			name: "Name - default",
			args: args{
				in: func() kops.ClusterSubnetSpec {
					subject := kops.ClusterSubnetSpec{}
					subject.Name = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Cidr - default",
			args: args{
				in: func() kops.ClusterSubnetSpec {
					subject := kops.ClusterSubnetSpec{}
					subject.CIDR = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Ipv6Cidr - default",
			args: args{
				in: func() kops.ClusterSubnetSpec {
					subject := kops.ClusterSubnetSpec{}
					subject.IPv6CIDR = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Zone - default",
			args: args{
				in: func() kops.ClusterSubnetSpec {
					subject := kops.ClusterSubnetSpec{}
					subject.Zone = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Region - default",
			args: args{
				in: func() kops.ClusterSubnetSpec {
					subject := kops.ClusterSubnetSpec{}
					subject.Region = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Id - default",
			args: args{
				in: func() kops.ClusterSubnetSpec {
					subject := kops.ClusterSubnetSpec{}
					subject.ID = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Egress - default",
			args: args{
				in: func() kops.ClusterSubnetSpec {
					subject := kops.ClusterSubnetSpec{}
					subject.Egress = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Type - default",
			args: args{
				in: func() kops.ClusterSubnetSpec {
					subject := kops.ClusterSubnetSpec{}
					subject.Type = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PublicIp - default",
			args: args{
				in: func() kops.ClusterSubnetSpec {
					subject := kops.ClusterSubnetSpec{}
					subject.PublicIP = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AdditionalRoutes - default",
			args: args{
				in: func() kops.ClusterSubnetSpec {
					subject := kops.ClusterSubnetSpec{}
					subject.AdditionalRoutes = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenDataSourceClusterSubnetSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceClusterSubnetSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceClusterSubnetSpec(t *testing.T) {
	_default := map[string]interface{}{
		"name":              "",
		"cidr":              "",
		"ipv6_cidr":         "",
		"zone":              "",
		"region":            "",
		"id":                "",
		"egress":            "",
		"type":              "",
		"public_ip":         "",
		"additional_routes": func() []interface{} { return nil }(),
	}
	type args struct {
		in kops.ClusterSubnetSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.ClusterSubnetSpec{},
			},
			want: _default,
		},
		{
			name: "Name - default",
			args: args{
				in: func() kops.ClusterSubnetSpec {
					subject := kops.ClusterSubnetSpec{}
					subject.Name = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Cidr - default",
			args: args{
				in: func() kops.ClusterSubnetSpec {
					subject := kops.ClusterSubnetSpec{}
					subject.CIDR = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Ipv6Cidr - default",
			args: args{
				in: func() kops.ClusterSubnetSpec {
					subject := kops.ClusterSubnetSpec{}
					subject.IPv6CIDR = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Zone - default",
			args: args{
				in: func() kops.ClusterSubnetSpec {
					subject := kops.ClusterSubnetSpec{}
					subject.Zone = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Region - default",
			args: args{
				in: func() kops.ClusterSubnetSpec {
					subject := kops.ClusterSubnetSpec{}
					subject.Region = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Id - default",
			args: args{
				in: func() kops.ClusterSubnetSpec {
					subject := kops.ClusterSubnetSpec{}
					subject.ID = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Egress - default",
			args: args{
				in: func() kops.ClusterSubnetSpec {
					subject := kops.ClusterSubnetSpec{}
					subject.Egress = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Type - default",
			args: args{
				in: func() kops.ClusterSubnetSpec {
					subject := kops.ClusterSubnetSpec{}
					subject.Type = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PublicIp - default",
			args: args{
				in: func() kops.ClusterSubnetSpec {
					subject := kops.ClusterSubnetSpec{}
					subject.PublicIP = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AdditionalRoutes - default",
			args: args{
				in: func() kops.ClusterSubnetSpec {
					subject := kops.ClusterSubnetSpec{}
					subject.AdditionalRoutes = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceClusterSubnetSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceClusterSubnetSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
