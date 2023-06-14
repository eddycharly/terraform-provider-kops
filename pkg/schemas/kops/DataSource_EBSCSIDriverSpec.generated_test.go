package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceEBSCSIDriverSpec(t *testing.T) {
	_default := kops.EBSCSIDriverSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.EBSCSIDriverSpec
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"enabled":             nil,
					"managed":             nil,
					"version":             nil,
					"volume_attach_limit": nil,
					"pod_annotations":     func() map[string]interface{} { return nil }(),
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandDataSourceEBSCSIDriverSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceEBSCSIDriverSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceEBSCSIDriverSpecInto(t *testing.T) {
	_default := map[string]interface{}{
		"enabled":             nil,
		"managed":             nil,
		"version":             nil,
		"volume_attach_limit": nil,
		"pod_annotations":     func() map[string]interface{} { return nil }(),
	}
	type args struct {
		in kops.EBSCSIDriverSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.EBSCSIDriverSpec{},
			},
			want: _default,
		},
		{
			name: "Enabled - default",
			args: args{
				in: func() kops.EBSCSIDriverSpec {
					subject := kops.EBSCSIDriverSpec{}
					subject.Enabled = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Managed - default",
			args: args{
				in: func() kops.EBSCSIDriverSpec {
					subject := kops.EBSCSIDriverSpec{}
					subject.Managed = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Version - default",
			args: args{
				in: func() kops.EBSCSIDriverSpec {
					subject := kops.EBSCSIDriverSpec{}
					subject.Version = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "VolumeAttachLimit - default",
			args: args{
				in: func() kops.EBSCSIDriverSpec {
					subject := kops.EBSCSIDriverSpec{}
					subject.VolumeAttachLimit = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PodAnnotations - default",
			args: args{
				in: func() kops.EBSCSIDriverSpec {
					subject := kops.EBSCSIDriverSpec{}
					subject.PodAnnotations = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenDataSourceEBSCSIDriverSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceEBSCSIDriverSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceEBSCSIDriverSpec(t *testing.T) {
	_default := map[string]interface{}{
		"enabled":             nil,
		"managed":             nil,
		"version":             nil,
		"volume_attach_limit": nil,
		"pod_annotations":     func() map[string]interface{} { return nil }(),
	}
	type args struct {
		in kops.EBSCSIDriverSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.EBSCSIDriverSpec{},
			},
			want: _default,
		},
		{
			name: "Enabled - default",
			args: args{
				in: func() kops.EBSCSIDriverSpec {
					subject := kops.EBSCSIDriverSpec{}
					subject.Enabled = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Managed - default",
			args: args{
				in: func() kops.EBSCSIDriverSpec {
					subject := kops.EBSCSIDriverSpec{}
					subject.Managed = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Version - default",
			args: args{
				in: func() kops.EBSCSIDriverSpec {
					subject := kops.EBSCSIDriverSpec{}
					subject.Version = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "VolumeAttachLimit - default",
			args: args{
				in: func() kops.EBSCSIDriverSpec {
					subject := kops.EBSCSIDriverSpec{}
					subject.VolumeAttachLimit = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PodAnnotations - default",
			args: args{
				in: func() kops.EBSCSIDriverSpec {
					subject := kops.EBSCSIDriverSpec{}
					subject.PodAnnotations = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceEBSCSIDriverSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceEBSCSIDriverSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
