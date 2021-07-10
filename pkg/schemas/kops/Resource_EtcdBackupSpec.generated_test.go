package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceEtcdBackupSpec(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.EtcdBackupSpec
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandResourceEtcdBackupSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandResourceEtcdBackupSpec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenResourceEtcdBackupSpecInto(t *testing.T) {
	type args struct {
		in  kops.EtcdBackupSpec
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
			FlattenResourceEtcdBackupSpecInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenResourceEtcdBackupSpec(t *testing.T) {
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
			want: map[string]interface{}{
				"backup_store": "",
				"image":        "",
			},
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
			want: map[string]interface{}{
				"backup_store": "",
				"image":        "",
			},
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
			want: map[string]interface{}{
				"backup_store": "",
				"image":        "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenResourceEtcdBackupSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenResourceEtcdBackupSpec() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
