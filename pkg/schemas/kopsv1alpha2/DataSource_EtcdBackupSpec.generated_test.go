package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

func TestExpandDataSourceEtcdBackupSpec(t *testing.T) {
	_default := kopsv1alpha2.EtcdBackupSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kopsv1alpha2.EtcdBackupSpec
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
		in kopsv1alpha2.EtcdBackupSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.EtcdBackupSpec{},
			},
			want: _default,
		},
		{
			name: "BackupStore - default",
			args: args{
				in: func() kopsv1alpha2.EtcdBackupSpec {
					subject := kopsv1alpha2.EtcdBackupSpec{}
					subject.BackupStore = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Image - default",
			args: args{
				in: func() kopsv1alpha2.EtcdBackupSpec {
					subject := kopsv1alpha2.EtcdBackupSpec{}
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
		in kopsv1alpha2.EtcdBackupSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.EtcdBackupSpec{},
			},
			want: _default,
		},
		{
			name: "BackupStore - default",
			args: args{
				in: func() kopsv1alpha2.EtcdBackupSpec {
					subject := kopsv1alpha2.EtcdBackupSpec{}
					subject.BackupStore = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Image - default",
			args: args{
				in: func() kopsv1alpha2.EtcdBackupSpec {
					subject := kopsv1alpha2.EtcdBackupSpec{}
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
