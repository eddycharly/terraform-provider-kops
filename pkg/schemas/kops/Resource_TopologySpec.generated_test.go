package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceTopologySpec(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.TopologySpec
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandResourceTopologySpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandResourceTopologySpec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenResourceTopologySpecInto(t *testing.T) {
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
			FlattenResourceTopologySpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceTopologySpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceTopologySpec(t *testing.T) {
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
			got := FlattenResourceTopologySpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceTopologySpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
