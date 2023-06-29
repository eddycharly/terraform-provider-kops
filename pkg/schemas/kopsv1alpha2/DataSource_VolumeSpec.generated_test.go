package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

func TestExpandDataSourceVolumeSpec(t *testing.T) {
	_default := kopsv1alpha2.VolumeSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kopsv1alpha2.VolumeSpec
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"delete_on_termination": nil,
					"device":                "",
					"encrypted":             nil,
					"iops":                  nil,
					"throughput":            nil,
					"key":                   nil,
					"size":                  0,
					"type":                  "",
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandDataSourceVolumeSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceVolumeSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceVolumeSpecInto(t *testing.T) {
	_default := map[string]interface{}{
		"delete_on_termination": nil,
		"device":                "",
		"encrypted":             nil,
		"iops":                  nil,
		"throughput":            nil,
		"key":                   nil,
		"size":                  0,
		"type":                  "",
	}
	type args struct {
		in kopsv1alpha2.VolumeSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.VolumeSpec{},
			},
			want: _default,
		},
		{
			name: "DeleteOnTermination - default",
			args: args{
				in: func() kopsv1alpha2.VolumeSpec {
					subject := kopsv1alpha2.VolumeSpec{}
					subject.DeleteOnTermination = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Device - default",
			args: args{
				in: func() kopsv1alpha2.VolumeSpec {
					subject := kopsv1alpha2.VolumeSpec{}
					subject.Device = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Encrypted - default",
			args: args{
				in: func() kopsv1alpha2.VolumeSpec {
					subject := kopsv1alpha2.VolumeSpec{}
					subject.Encrypted = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "IOPS - default",
			args: args{
				in: func() kopsv1alpha2.VolumeSpec {
					subject := kopsv1alpha2.VolumeSpec{}
					subject.IOPS = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Throughput - default",
			args: args{
				in: func() kopsv1alpha2.VolumeSpec {
					subject := kopsv1alpha2.VolumeSpec{}
					subject.Throughput = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Key - default",
			args: args{
				in: func() kopsv1alpha2.VolumeSpec {
					subject := kopsv1alpha2.VolumeSpec{}
					subject.Key = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Size - default",
			args: args{
				in: func() kopsv1alpha2.VolumeSpec {
					subject := kopsv1alpha2.VolumeSpec{}
					subject.Size = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Type - default",
			args: args{
				in: func() kopsv1alpha2.VolumeSpec {
					subject := kopsv1alpha2.VolumeSpec{}
					subject.Type = ""
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenDataSourceVolumeSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceVolumeSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceVolumeSpec(t *testing.T) {
	_default := map[string]interface{}{
		"delete_on_termination": nil,
		"device":                "",
		"encrypted":             nil,
		"iops":                  nil,
		"throughput":            nil,
		"key":                   nil,
		"size":                  0,
		"type":                  "",
	}
	type args struct {
		in kopsv1alpha2.VolumeSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.VolumeSpec{},
			},
			want: _default,
		},
		{
			name: "DeleteOnTermination - default",
			args: args{
				in: func() kopsv1alpha2.VolumeSpec {
					subject := kopsv1alpha2.VolumeSpec{}
					subject.DeleteOnTermination = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Device - default",
			args: args{
				in: func() kopsv1alpha2.VolumeSpec {
					subject := kopsv1alpha2.VolumeSpec{}
					subject.Device = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Encrypted - default",
			args: args{
				in: func() kopsv1alpha2.VolumeSpec {
					subject := kopsv1alpha2.VolumeSpec{}
					subject.Encrypted = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "IOPS - default",
			args: args{
				in: func() kopsv1alpha2.VolumeSpec {
					subject := kopsv1alpha2.VolumeSpec{}
					subject.IOPS = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Throughput - default",
			args: args{
				in: func() kopsv1alpha2.VolumeSpec {
					subject := kopsv1alpha2.VolumeSpec{}
					subject.Throughput = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Key - default",
			args: args{
				in: func() kopsv1alpha2.VolumeSpec {
					subject := kopsv1alpha2.VolumeSpec{}
					subject.Key = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Size - default",
			args: args{
				in: func() kopsv1alpha2.VolumeSpec {
					subject := kopsv1alpha2.VolumeSpec{}
					subject.Size = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Type - default",
			args: args{
				in: func() kopsv1alpha2.VolumeSpec {
					subject := kopsv1alpha2.VolumeSpec{}
					subject.Type = ""
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceVolumeSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceVolumeSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
