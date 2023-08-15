package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

func TestExpandDataSourceRunc(t *testing.T) {
	_default := kopsv1alpha2.Runc{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kopsv1alpha2.Runc
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"version":  nil,
					"packages": nil,
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandDataSourceRunc(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceRunc() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceRuncInto(t *testing.T) {
	_default := map[string]interface{}{
		"version":  nil,
		"packages": nil,
	}
	type args struct {
		in kopsv1alpha2.Runc
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.Runc{},
			},
			want: _default,
		},
		{
			name: "Version - default",
			args: args{
				in: func() kopsv1alpha2.Runc {
					subject := kopsv1alpha2.Runc{}
					subject.Version = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Packages - default",
			args: args{
				in: func() kopsv1alpha2.Runc {
					subject := kopsv1alpha2.Runc{}
					subject.Packages = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenDataSourceRuncInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceRunc() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceRunc(t *testing.T) {
	_default := map[string]interface{}{
		"version":  nil,
		"packages": nil,
	}
	type args struct {
		in kopsv1alpha2.Runc
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.Runc{},
			},
			want: _default,
		},
		{
			name: "Version - default",
			args: args{
				in: func() kopsv1alpha2.Runc {
					subject := kopsv1alpha2.Runc{}
					subject.Version = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Packages - default",
			args: args{
				in: func() kopsv1alpha2.Runc {
					subject := kopsv1alpha2.Runc{}
					subject.Packages = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceRunc(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceRunc() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
