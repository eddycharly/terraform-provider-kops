package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

func TestExpandResourceTopologySpec(t *testing.T) {
	_default := kopsv1alpha2.TopologySpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kopsv1alpha2.TopologySpec
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"control_plane": "",
					"nodes":         "",
					"bastion":       nil,
					"dns":           "",
					"legacy_dns":    nil,
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandResourceTopologySpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceTopologySpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceTopologySpecInto(t *testing.T) {
	_default := map[string]interface{}{
		"control_plane": "",
		"nodes":         "",
		"bastion":       nil,
		"dns":           "",
		"legacy_dns":    nil,
	}
	type args struct {
		in kopsv1alpha2.TopologySpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.TopologySpec{},
			},
			want: _default,
		},
		{
			name: "ControlPlane - default",
			args: args{
				in: func() kopsv1alpha2.TopologySpec {
					subject := kopsv1alpha2.TopologySpec{}
					subject.ControlPlane = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Nodes - default",
			args: args{
				in: func() kopsv1alpha2.TopologySpec {
					subject := kopsv1alpha2.TopologySpec{}
					subject.Nodes = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Bastion - default",
			args: args{
				in: func() kopsv1alpha2.TopologySpec {
					subject := kopsv1alpha2.TopologySpec{}
					subject.Bastion = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Dns - default",
			args: args{
				in: func() kopsv1alpha2.TopologySpec {
					subject := kopsv1alpha2.TopologySpec{}
					subject.DNS = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LegacyDns - default",
			args: args{
				in: func() kopsv1alpha2.TopologySpec {
					subject := kopsv1alpha2.TopologySpec{}
					subject.LegacyDNS = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenResourceTopologySpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceTopologySpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceTopologySpec(t *testing.T) {
	_default := map[string]interface{}{
		"control_plane": "",
		"nodes":         "",
		"bastion":       nil,
		"dns":           "",
		"legacy_dns":    nil,
	}
	type args struct {
		in kopsv1alpha2.TopologySpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.TopologySpec{},
			},
			want: _default,
		},
		{
			name: "ControlPlane - default",
			args: args{
				in: func() kopsv1alpha2.TopologySpec {
					subject := kopsv1alpha2.TopologySpec{}
					subject.ControlPlane = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Nodes - default",
			args: args{
				in: func() kopsv1alpha2.TopologySpec {
					subject := kopsv1alpha2.TopologySpec{}
					subject.Nodes = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Bastion - default",
			args: args{
				in: func() kopsv1alpha2.TopologySpec {
					subject := kopsv1alpha2.TopologySpec{}
					subject.Bastion = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Dns - default",
			args: args{
				in: func() kopsv1alpha2.TopologySpec {
					subject := kopsv1alpha2.TopologySpec{}
					subject.DNS = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LegacyDns - default",
			args: args{
				in: func() kopsv1alpha2.TopologySpec {
					subject := kopsv1alpha2.TopologySpec{}
					subject.LegacyDNS = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceTopologySpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceTopologySpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
