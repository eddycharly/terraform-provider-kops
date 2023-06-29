package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

func TestExpandDataSourceEtcdManagerSpec(t *testing.T) {
	_default := kopsv1alpha2.EtcdManagerSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kopsv1alpha2.EtcdManagerSpec
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"image":                   "",
					"env":                     func() []interface{} { return nil }(),
					"backup_interval":         nil,
					"discovery_poll_interval": nil,
					"log_level":               nil,
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandDataSourceEtcdManagerSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceEtcdManagerSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceEtcdManagerSpecInto(t *testing.T) {
	_default := map[string]interface{}{
		"image":                   "",
		"env":                     func() []interface{} { return nil }(),
		"backup_interval":         nil,
		"discovery_poll_interval": nil,
		"log_level":               nil,
	}
	type args struct {
		in kopsv1alpha2.EtcdManagerSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.EtcdManagerSpec{},
			},
			want: _default,
		},
		{
			name: "Image - default",
			args: args{
				in: func() kopsv1alpha2.EtcdManagerSpec {
					subject := kopsv1alpha2.EtcdManagerSpec{}
					subject.Image = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Env - default",
			args: args{
				in: func() kopsv1alpha2.EtcdManagerSpec {
					subject := kopsv1alpha2.EtcdManagerSpec{}
					subject.Env = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "BackupInterval - default",
			args: args{
				in: func() kopsv1alpha2.EtcdManagerSpec {
					subject := kopsv1alpha2.EtcdManagerSpec{}
					subject.BackupInterval = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "DiscoveryPollInterval - default",
			args: args{
				in: func() kopsv1alpha2.EtcdManagerSpec {
					subject := kopsv1alpha2.EtcdManagerSpec{}
					subject.DiscoveryPollInterval = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LogLevel - default",
			args: args{
				in: func() kopsv1alpha2.EtcdManagerSpec {
					subject := kopsv1alpha2.EtcdManagerSpec{}
					subject.LogLevel = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenDataSourceEtcdManagerSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceEtcdManagerSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceEtcdManagerSpec(t *testing.T) {
	_default := map[string]interface{}{
		"image":                   "",
		"env":                     func() []interface{} { return nil }(),
		"backup_interval":         nil,
		"discovery_poll_interval": nil,
		"log_level":               nil,
	}
	type args struct {
		in kopsv1alpha2.EtcdManagerSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.EtcdManagerSpec{},
			},
			want: _default,
		},
		{
			name: "Image - default",
			args: args{
				in: func() kopsv1alpha2.EtcdManagerSpec {
					subject := kopsv1alpha2.EtcdManagerSpec{}
					subject.Image = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Env - default",
			args: args{
				in: func() kopsv1alpha2.EtcdManagerSpec {
					subject := kopsv1alpha2.EtcdManagerSpec{}
					subject.Env = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "BackupInterval - default",
			args: args{
				in: func() kopsv1alpha2.EtcdManagerSpec {
					subject := kopsv1alpha2.EtcdManagerSpec{}
					subject.BackupInterval = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "DiscoveryPollInterval - default",
			args: args{
				in: func() kopsv1alpha2.EtcdManagerSpec {
					subject := kopsv1alpha2.EtcdManagerSpec{}
					subject.DiscoveryPollInterval = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LogLevel - default",
			args: args{
				in: func() kopsv1alpha2.EtcdManagerSpec {
					subject := kopsv1alpha2.EtcdManagerSpec{}
					subject.LogLevel = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceEtcdManagerSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceEtcdManagerSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
