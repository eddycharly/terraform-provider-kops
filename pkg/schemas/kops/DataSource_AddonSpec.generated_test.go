package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceAddonSpec(t *testing.T) {
	_default := kops.AddonSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.AddonSpec
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"manifest": "",
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandDataSourceAddonSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceAddonSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceAddonSpecInto(t *testing.T) {
	_default := map[string]interface{}{
		"manifest": "",
	}
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
			want: _default,
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
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenDataSourceAddonSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceAddonSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceAddonSpec(t *testing.T) {
	_default := map[string]interface{}{
		"manifest": "",
	}
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
			want: _default,
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
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceAddonSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceAddonSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
