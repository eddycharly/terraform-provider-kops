package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceTopologySpec(t *testing.T) {
	_default := kops.TopologySpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.TopologySpec
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"masters": "",
					"nodes":   "",
					"bastion": nil,
					"dns":     nil,
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandDataSourceTopologySpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceTopologySpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceTopologySpecInto(t *testing.T) {
	_default := map[string]interface{}{
		"masters": "",
		"nodes":   "",
		"bastion": nil,
		"dns":     nil,
	}
	type args struct {
		in kops.TopologySpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.TopologySpec{},
			},
			want: _default,
		},
		{
			name: "Masters - default",
			args: args{
				in: func() kops.TopologySpec {
					subject := kops.TopologySpec{}
					subject.Masters = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Nodes - default",
			args: args{
				in: func() kops.TopologySpec {
					subject := kops.TopologySpec{}
					subject.Nodes = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Bastion - default",
			args: args{
				in: func() kops.TopologySpec {
					subject := kops.TopologySpec{}
					subject.Bastion = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Dns - default",
			args: args{
				in: func() kops.TopologySpec {
					subject := kops.TopologySpec{}
					subject.DNS = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenDataSourceTopologySpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceTopologySpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceTopologySpec(t *testing.T) {
	_default := map[string]interface{}{
		"masters": "",
		"nodes":   "",
		"bastion": nil,
		"dns":     nil,
	}
	type args struct {
		in kops.TopologySpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.TopologySpec{},
			},
			want: _default,
		},
		{
			name: "Masters - default",
			args: args{
				in: func() kops.TopologySpec {
					subject := kops.TopologySpec{}
					subject.Masters = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Nodes - default",
			args: args{
				in: func() kops.TopologySpec {
					subject := kops.TopologySpec{}
					subject.Nodes = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Bastion - default",
			args: args{
				in: func() kops.TopologySpec {
					subject := kops.TopologySpec{}
					subject.Bastion = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Dns - default",
			args: args{
				in: func() kops.TopologySpec {
					subject := kops.TopologySpec{}
					subject.DNS = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceTopologySpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceTopologySpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
