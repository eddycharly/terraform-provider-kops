package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceVolumeSpec(t *testing.T) {
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
			if got := ExpandResourceVolumeSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandResourceVolumeSpec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenResourceVolumeSpecInto(t *testing.T) {
	type args struct {
		in  kops.VolumeSpec
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
			FlattenResourceVolumeSpecInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenResourceVolumeSpec(t *testing.T) {
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
			want: map[string]interface{}{
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
		{
			name: "DeleteOnTermination - default",
			args: args{
				in: func() kops.VolumeSpec {
					subject := kops.VolumeSpec{}
					subject.DeleteOnTermination = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
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
		{
			name: "Device - default",
			args: args{
				in: func() kops.VolumeSpec {
					subject := kops.VolumeSpec{}
					subject.Device = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
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
		{
			name: "Encrypted - default",
			args: args{
				in: func() kops.VolumeSpec {
					subject := kops.VolumeSpec{}
					subject.Encrypted = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
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
		{
			name: "Iops - default",
			args: args{
				in: func() kops.VolumeSpec {
					subject := kops.VolumeSpec{}
					subject.Iops = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
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
		{
			name: "Throughput - default",
			args: args{
				in: func() kops.VolumeSpec {
					subject := kops.VolumeSpec{}
					subject.Throughput = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
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
		{
			name: "Key - default",
			args: args{
				in: func() kops.VolumeSpec {
					subject := kops.VolumeSpec{}
					subject.Key = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
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
		{
			name: "Size - default",
			args: args{
				in: func() kops.VolumeSpec {
					subject := kops.VolumeSpec{}
					subject.Size = 0
					return subject
				}(),
			},
			want: map[string]interface{}{
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
		{
			name: "Type - default",
			args: args{
				in: func() kops.VolumeSpec {
					subject := kops.VolumeSpec{}
					subject.Type = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenResourceVolumeSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenResourceVolumeSpec() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
