package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceServiceAccountExternalPermission(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.ServiceAccountExternalPermission
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandDataSourceServiceAccountExternalPermission(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandDataSourceServiceAccountExternalPermission() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenDataSourceServiceAccountExternalPermissionInto(t *testing.T) {
	type args struct {
		in  kops.ServiceAccountExternalPermission
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
			FlattenDataSourceServiceAccountExternalPermissionInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenDataSourceServiceAccountExternalPermission(t *testing.T) {
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
			want: map[string]interface{}{
				"name":      "",
				"namespace": "",
				"aws":       nil,
			},
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
			want: map[string]interface{}{
				"name":      "",
				"namespace": "",
				"aws":       nil,
			},
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
			want: map[string]interface{}{
				"name":      "",
				"namespace": "",
				"aws":       nil,
			},
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
			want: map[string]interface{}{
				"name":      "",
				"namespace": "",
				"aws":       nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenDataSourceServiceAccountExternalPermission(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenDataSourceServiceAccountExternalPermission() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
