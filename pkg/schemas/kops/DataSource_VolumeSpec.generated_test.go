package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceVolumeSpec(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.VolumeSpec
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandDataSourceVolumeSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandDataSourceVolumeSpec() = %v, want %v", got, tt.want)
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
		in kops.VolumeSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.VolumeSpec{},
			},
			want: _default,
		},
		{
			name: "DeleteOnTermination - default",
			args: args{
				in: func() kops.VolumeSpec {
					subject := kops.VolumeSpec{}
					subject.DeleteOnTermination = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Device - default",
			args: args{
				in: func() kops.VolumeSpec {
					subject := kops.VolumeSpec{}
					subject.Device = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Encrypted - default",
			args: args{
				in: func() kops.VolumeSpec {
					subject := kops.VolumeSpec{}
					subject.Encrypted = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Iops - default",
			args: args{
				in: func() kops.VolumeSpec {
					subject := kops.VolumeSpec{}
					subject.Iops = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Throughput - default",
			args: args{
				in: func() kops.VolumeSpec {
					subject := kops.VolumeSpec{}
					subject.Throughput = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Key - default",
			args: args{
				in: func() kops.VolumeSpec {
					subject := kops.VolumeSpec{}
					subject.Key = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Size - default",
			args: args{
				in: func() kops.VolumeSpec {
					subject := kops.VolumeSpec{}
					subject.Size = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Type - default",
			args: args{
				in: func() kops.VolumeSpec {
					subject := kops.VolumeSpec{}
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
		in kops.VolumeSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.VolumeSpec{},
			},
			want: _default,
		},
		{
			name: "DeleteOnTermination - default",
			args: args{
				in: func() kops.VolumeSpec {
					subject := kops.VolumeSpec{}
					subject.DeleteOnTermination = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Device - default",
			args: args{
				in: func() kops.VolumeSpec {
					subject := kops.VolumeSpec{}
					subject.Device = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Encrypted - default",
			args: args{
				in: func() kops.VolumeSpec {
					subject := kops.VolumeSpec{}
					subject.Encrypted = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Iops - default",
			args: args{
				in: func() kops.VolumeSpec {
					subject := kops.VolumeSpec{}
					subject.Iops = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Throughput - default",
			args: args{
				in: func() kops.VolumeSpec {
					subject := kops.VolumeSpec{}
					subject.Throughput = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Key - default",
			args: args{
				in: func() kops.VolumeSpec {
					subject := kops.VolumeSpec{}
					subject.Key = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Size - default",
			args: args{
				in: func() kops.VolumeSpec {
					subject := kops.VolumeSpec{}
					subject.Size = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Type - default",
			args: args{
				in: func() kops.VolumeSpec {
					subject := kops.VolumeSpec{}
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
