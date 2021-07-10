package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceLyftVPCNetworkingSpec(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.LyftVPCNetworkingSpec
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandDataSourceLyftVPCNetworkingSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandDataSourceLyftVPCNetworkingSpec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenDataSourceLyftVPCNetworkingSpecInto(t *testing.T) {
	type args struct {
		in  kops.LyftVPCNetworkingSpec
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
			FlattenDataSourceLyftVPCNetworkingSpecInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenDataSourceLyftVPCNetworkingSpec(t *testing.T) {
	type args struct {
		in kops.LyftVPCNetworkingSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.LyftVPCNetworkingSpec{},
			},
			want: map[string]interface{}{
				"subnet_tags": func() map[string]interface{} { return nil }(),
			},
		},
		{
			name: "SubnetTags - default",
			args: args{
				in: func() kops.LyftVPCNetworkingSpec {
					subject := kops.LyftVPCNetworkingSpec{}
					subject.SubnetTags = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"subnet_tags": func() map[string]interface{} { return nil }(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenDataSourceLyftVPCNetworkingSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenDataSourceLyftVPCNetworkingSpec() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
