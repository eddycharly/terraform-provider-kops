package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceEtcdBackupSpec(t *testing.T) {
	_default := kops.EtcdBackupSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.EtcdBackupSpec
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"backup_store": "",
					"image":        "",
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandDataSourceEtcdBackupSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceEtcdBackupSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceEtcdBackupSpecInto(t *testing.T) {
	_default := map[string]interface{}{
		"backup_store": "",
		"image":        "",
	}
	type args struct {
		in kops.EtcdBackupSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.EtcdBackupSpec{},
			},
			want: _default,
		},
		{
			name: "BackupStore - default",
			args: args{
				in: func() kops.EtcdBackupSpec {
					subject := kops.EtcdBackupSpec{}
					subject.BackupStore = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Image - default",
			args: args{
				in: func() kops.EtcdBackupSpec {
					subject := kops.EtcdBackupSpec{}
					subject.Image = ""
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenDataSourceEtcdBackupSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceEtcdBackupSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceEtcdBackupSpec(t *testing.T) {
	_default := map[string]interface{}{
		"backup_store": "",
		"image":        "",
	}
	type args struct {
		in kops.EtcdBackupSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.EtcdBackupSpec{},
			},
			want: _default,
		},
		{
			name: "BackupStore - default",
			args: args{
				in: func() kops.EtcdBackupSpec {
					subject := kops.EtcdBackupSpec{}
					subject.BackupStore = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Image - default",
			args: args{
				in: func() kops.EtcdBackupSpec {
					subject := kops.EtcdBackupSpec{}
					subject.Image = ""
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceEtcdBackupSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceEtcdBackupSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
