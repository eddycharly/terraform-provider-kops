package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceAddonSpec(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.AddonSpec
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandDataSourceAddonSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandDataSourceAddonSpec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenDataSourceAddonSpecInto(t *testing.T) {
	type args struct {
		in  kops.AddonSpec
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
			FlattenDataSourceAddonSpecInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenDataSourceAddonSpec(t *testing.T) {
	type args struct {
		in kops.AddonSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.AddonSpec{},
			},
			want: map[string]interface{}{
				"manifest": "",
			},
		},
		{
			name: "Manifest - default",
			args: args{
				in: func() kops.AddonSpec {
					subject := kops.AddonSpec{}
					subject.Manifest = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
				"manifest": "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenDataSourceAddonSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenDataSourceAddonSpec() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
