package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceOpenstackMetadata(t *testing.T) {
	_default := kops.OpenstackMetadata{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.OpenstackMetadata
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
			got := ExpandResourceOpenstackMetadata(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceOpenstackMetadata() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceOpenstackMetadataInto(t *testing.T) {
	_default := map[string]interface{}{
		"config_drive": nil,
	}
	type args struct {
		in kops.OpenstackMetadata
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.OpenstackMetadata{},
			},
			want: _default,
		},
		{
			name: "ConfigDrive - default",
			args: args{
				in: func() kops.OpenstackMetadata {
					subject := kops.OpenstackMetadata{}
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
			FlattenResourceOpenstackMetadataInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceOpenstackMetadata() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceOpenstackMetadata(t *testing.T) {
	_default := map[string]interface{}{
		"config_drive": nil,
	}
	type args struct {
		in kops.OpenstackMetadata
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.OpenstackMetadata{},
			},
			want: _default,
		},
		{
			name: "ConfigDrive - default",
			args: args{
				in: func() kops.OpenstackMetadata {
					subject := kops.OpenstackMetadata{}
					subject.ConfigDrive = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceOpenstackMetadata(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceOpenstackMetadata() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
