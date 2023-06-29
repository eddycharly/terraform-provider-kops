package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

func TestExpandDataSourceAddonSpec(t *testing.T) {
	_default := kopsv1alpha2.AddonSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kopsv1alpha2.AddonSpec
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
		in kopsv1alpha2.AddonSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.AddonSpec{},
			},
			want: _default,
		},
		{
			name: "Manifest - default",
			args: args{
				in: func() kopsv1alpha2.AddonSpec {
					subject := kopsv1alpha2.AddonSpec{}
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
		in kopsv1alpha2.AddonSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.AddonSpec{},
			},
			want: _default,
		},
		{
			name: "Manifest - default",
			args: args{
				in: func() kopsv1alpha2.AddonSpec {
					subject := kopsv1alpha2.AddonSpec{}
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
