package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceServiceAccountExternalPermission(t *testing.T) {
	_default := kops.ServiceAccountExternalPermission{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.ServiceAccountExternalPermission
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"name":      "",
					"namespace": "",
					"aws":       nil,
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandResourceServiceAccountExternalPermission(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceServiceAccountExternalPermission() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceServiceAccountExternalPermissionInto(t *testing.T) {
	_default := map[string]interface{}{
		"name":      "",
		"namespace": "",
		"aws":       nil,
	}
	type args struct {
		in kops.ServiceAccountExternalPermission
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.ServiceAccountExternalPermission{},
			},
			want: _default,
		},
		{
			name: "Name - default",
			args: args{
				in: func() kops.ServiceAccountExternalPermission {
					subject := kops.ServiceAccountExternalPermission{}
					subject.Name = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Namespace - default",
			args: args{
				in: func() kops.ServiceAccountExternalPermission {
					subject := kops.ServiceAccountExternalPermission{}
					subject.Namespace = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Aws - default",
			args: args{
				in: func() kops.ServiceAccountExternalPermission {
					subject := kops.ServiceAccountExternalPermission{}
					subject.AWS = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenResourceServiceAccountExternalPermissionInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceServiceAccountExternalPermission() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceServiceAccountExternalPermission(t *testing.T) {
	_default := map[string]interface{}{
		"name":      "",
		"namespace": "",
		"aws":       nil,
	}
	type args struct {
		in kops.ServiceAccountExternalPermission
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.ServiceAccountExternalPermission{},
			},
			want: _default,
		},
		{
			name: "Name - default",
			args: args{
				in: func() kops.ServiceAccountExternalPermission {
					subject := kops.ServiceAccountExternalPermission{}
					subject.Name = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Namespace - default",
			args: args{
				in: func() kops.ServiceAccountExternalPermission {
					subject := kops.ServiceAccountExternalPermission{}
					subject.Namespace = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Aws - default",
			args: args{
				in: func() kops.ServiceAccountExternalPermission {
					subject := kops.ServiceAccountExternalPermission{}
					subject.AWS = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceServiceAccountExternalPermission(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceServiceAccountExternalPermission() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
