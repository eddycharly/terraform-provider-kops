package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceClusterSubnetSpec(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.ClusterSubnetSpec
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandResourceClusterSubnetSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandResourceClusterSubnetSpec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenResourceClusterSubnetSpecInto(t *testing.T) {
	type args struct {
		in  kops.ClusterSubnetSpec
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
			FlattenResourceClusterSubnetSpecInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenResourceClusterSubnetSpec(t *testing.T) {
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
			want: map[string]interface{}{
				"name":        "",
				"cidr":        "",
				"zone":        "",
				"region":      "",
				"provider_id": "",
				"egress":      "",
				"type":        "",
				"public_ip":   "",
			},
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
			want: map[string]interface{}{
				"name":        "",
				"cidr":        "",
				"zone":        "",
				"region":      "",
				"provider_id": "",
				"egress":      "",
				"type":        "",
				"public_ip":   "",
			},
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
			want: map[string]interface{}{
				"name":        "",
				"cidr":        "",
				"zone":        "",
				"region":      "",
				"provider_id": "",
				"egress":      "",
				"type":        "",
				"public_ip":   "",
			},
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
			want: map[string]interface{}{
				"name":        "",
				"cidr":        "",
				"zone":        "",
				"region":      "",
				"provider_id": "",
				"egress":      "",
				"type":        "",
				"public_ip":   "",
			},
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
			want: map[string]interface{}{
				"name":        "",
				"cidr":        "",
				"zone":        "",
				"region":      "",
				"provider_id": "",
				"egress":      "",
				"type":        "",
				"public_ip":   "",
			},
		},
		{
			name: "ProviderID - default",
			args: args{
				in: func() kops.ClusterSubnetSpec {
					subject := kops.ClusterSubnetSpec{}
					subject.ProviderID = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
				"name":        "",
				"cidr":        "",
				"zone":        "",
				"region":      "",
				"provider_id": "",
				"egress":      "",
				"type":        "",
				"public_ip":   "",
			},
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
			want: map[string]interface{}{
				"name":        "",
				"cidr":        "",
				"zone":        "",
				"region":      "",
				"provider_id": "",
				"egress":      "",
				"type":        "",
				"public_ip":   "",
			},
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
			want: map[string]interface{}{
				"name":        "",
				"cidr":        "",
				"zone":        "",
				"region":      "",
				"provider_id": "",
				"egress":      "",
				"type":        "",
				"public_ip":   "",
			},
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
			want: map[string]interface{}{
				"name":        "",
				"cidr":        "",
				"zone":        "",
				"region":      "",
				"provider_id": "",
				"egress":      "",
				"type":        "",
				"public_ip":   "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenResourceClusterSubnetSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenResourceClusterSubnetSpec() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
