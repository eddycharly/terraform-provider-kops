package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceRunc(t *testing.T) {
	_default := kops.Runc{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.Runc
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
		in kops.Runc
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.Runc{},
			},
			want: _default,
		},
		{
			name: "Version - default",
			args: args{
				in: func() kops.Runc {
					subject := kops.Runc{}
					subject.Version = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Packages - default",
			args: args{
				in: func() kops.Runc {
					subject := kops.Runc{}
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
		in kops.Runc
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.Runc{},
			},
			want: _default,
		},
		{
			name: "Version - default",
			args: args{
				in: func() kops.Runc {
					subject := kops.Runc{}
					subject.Version = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Packages - default",
			args: args{
				in: func() kops.Runc {
					subject := kops.Runc{}
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
