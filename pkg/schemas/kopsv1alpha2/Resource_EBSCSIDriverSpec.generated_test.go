package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

func TestExpandResourceEBSCSIDriverSpec(t *testing.T) {
	_default := kopsv1alpha2.EBSCSIDriverSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kopsv1alpha2.EBSCSIDriverSpec
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
			got := ExpandResourceEBSCSIDriverSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceEBSCSIDriverSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceEBSCSIDriverSpecInto(t *testing.T) {
	_default := map[string]interface{}{
		"enabled":             nil,
		"managed":             nil,
		"version":             nil,
		"volume_attach_limit": nil,
		"pod_annotations":     func() map[string]interface{} { return nil }(),
	}
	type args struct {
		in kopsv1alpha2.EBSCSIDriverSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.EBSCSIDriverSpec{},
			},
			want: _default,
		},
		{
			name: "Enabled - default",
			args: args{
				in: func() kopsv1alpha2.EBSCSIDriverSpec {
					subject := kopsv1alpha2.EBSCSIDriverSpec{}
					subject.Enabled = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Managed - default",
			args: args{
				in: func() kopsv1alpha2.EBSCSIDriverSpec {
					subject := kopsv1alpha2.EBSCSIDriverSpec{}
					subject.Managed = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Version - default",
			args: args{
				in: func() kopsv1alpha2.EBSCSIDriverSpec {
					subject := kopsv1alpha2.EBSCSIDriverSpec{}
					subject.Version = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "VolumeAttachLimit - default",
			args: args{
				in: func() kopsv1alpha2.EBSCSIDriverSpec {
					subject := kopsv1alpha2.EBSCSIDriverSpec{}
					subject.VolumeAttachLimit = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PodAnnotations - default",
			args: args{
				in: func() kopsv1alpha2.EBSCSIDriverSpec {
					subject := kopsv1alpha2.EBSCSIDriverSpec{}
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
			FlattenResourceEBSCSIDriverSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceEBSCSIDriverSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceEBSCSIDriverSpec(t *testing.T) {
	_default := map[string]interface{}{
		"enabled":             nil,
		"managed":             nil,
		"version":             nil,
		"volume_attach_limit": nil,
		"pod_annotations":     func() map[string]interface{} { return nil }(),
	}
	type args struct {
		in kopsv1alpha2.EBSCSIDriverSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.EBSCSIDriverSpec{},
			},
			want: _default,
		},
		{
			name: "Enabled - default",
			args: args{
				in: func() kopsv1alpha2.EBSCSIDriverSpec {
					subject := kopsv1alpha2.EBSCSIDriverSpec{}
					subject.Enabled = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Managed - default",
			args: args{
				in: func() kopsv1alpha2.EBSCSIDriverSpec {
					subject := kopsv1alpha2.EBSCSIDriverSpec{}
					subject.Managed = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Version - default",
			args: args{
				in: func() kopsv1alpha2.EBSCSIDriverSpec {
					subject := kopsv1alpha2.EBSCSIDriverSpec{}
					subject.Version = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "VolumeAttachLimit - default",
			args: args{
				in: func() kopsv1alpha2.EBSCSIDriverSpec {
					subject := kopsv1alpha2.EBSCSIDriverSpec{}
					subject.VolumeAttachLimit = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PodAnnotations - default",
			args: args{
				in: func() kopsv1alpha2.EBSCSIDriverSpec {
					subject := kopsv1alpha2.EBSCSIDriverSpec{}
					subject.PodAnnotations = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceEBSCSIDriverSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceEBSCSIDriverSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
