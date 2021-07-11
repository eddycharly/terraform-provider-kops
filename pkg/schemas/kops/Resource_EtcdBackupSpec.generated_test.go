package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceEtcdBackupSpec(t *testing.T) {
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
			got := ExpandResourceEtcdBackupSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceEtcdBackupSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceEtcdBackupSpecInto(t *testing.T) {
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
			FlattenResourceEtcdBackupSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceEtcdBackupSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceEtcdBackupSpec(t *testing.T) {
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
			got := FlattenResourceEtcdBackupSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceEtcdBackupSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
