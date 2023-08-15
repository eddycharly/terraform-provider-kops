package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

func TestExpandDataSourceOpenstackMetadata(t *testing.T) {
	_default := kopsv1alpha2.OpenstackMetadata{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kopsv1alpha2.OpenstackMetadata
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"config_drive": nil,
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandDataSourceOpenstackMetadata(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceOpenstackMetadata() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceOpenstackMetadataInto(t *testing.T) {
	_default := map[string]interface{}{
		"config_drive": nil,
	}
	type args struct {
		in kopsv1alpha2.OpenstackMetadata
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.OpenstackMetadata{},
			},
			want: _default,
		},
		{
			name: "ConfigDrive - default",
			args: args{
				in: func() kopsv1alpha2.OpenstackMetadata {
					subject := kopsv1alpha2.OpenstackMetadata{}
					subject.ConfigDrive = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenDataSourceOpenstackMetadataInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceOpenstackMetadata() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceOpenstackMetadata(t *testing.T) {
	_default := map[string]interface{}{
		"config_drive": nil,
	}
	type args struct {
		in kopsv1alpha2.OpenstackMetadata
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.OpenstackMetadata{},
			},
			want: _default,
		},
		{
			name: "ConfigDrive - default",
			args: args{
				in: func() kopsv1alpha2.OpenstackMetadata {
					subject := kopsv1alpha2.OpenstackMetadata{}
					subject.ConfigDrive = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceOpenstackMetadata(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceOpenstackMetadata() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
