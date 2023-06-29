package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

func TestExpandResourceServiceAccountExternalPermission(t *testing.T) {
	_default := kopsv1alpha2.ServiceAccountExternalPermission{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kopsv1alpha2.ServiceAccountExternalPermission
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
		in kopsv1alpha2.ServiceAccountExternalPermission
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.ServiceAccountExternalPermission{},
			},
			want: _default,
		},
		{
			name: "Name - default",
			args: args{
				in: func() kopsv1alpha2.ServiceAccountExternalPermission {
					subject := kopsv1alpha2.ServiceAccountExternalPermission{}
					subject.Name = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Namespace - default",
			args: args{
				in: func() kopsv1alpha2.ServiceAccountExternalPermission {
					subject := kopsv1alpha2.ServiceAccountExternalPermission{}
					subject.Namespace = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Aws - default",
			args: args{
				in: func() kopsv1alpha2.ServiceAccountExternalPermission {
					subject := kopsv1alpha2.ServiceAccountExternalPermission{}
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
		in kopsv1alpha2.ServiceAccountExternalPermission
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.ServiceAccountExternalPermission{},
			},
			want: _default,
		},
		{
			name: "Name - default",
			args: args{
				in: func() kopsv1alpha2.ServiceAccountExternalPermission {
					subject := kopsv1alpha2.ServiceAccountExternalPermission{}
					subject.Name = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Namespace - default",
			args: args{
				in: func() kopsv1alpha2.ServiceAccountExternalPermission {
					subject := kopsv1alpha2.ServiceAccountExternalPermission{}
					subject.Namespace = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Aws - default",
			args: args{
				in: func() kopsv1alpha2.ServiceAccountExternalPermission {
					subject := kopsv1alpha2.ServiceAccountExternalPermission{}
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
