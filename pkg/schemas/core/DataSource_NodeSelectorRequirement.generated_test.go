package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	v1 "k8s.io/api/core/v1"
)

func TestExpandDataSourceNodeSelectorRequirement(t *testing.T) {
	_default := v1.NodeSelectorRequirement{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want v1.NodeSelectorRequirement
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"key":      "",
					"operator": "",
					"values":   func() []interface{} { return nil }(),
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandDataSourceNodeSelectorRequirement(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceNodeSelectorRequirement() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceNodeSelectorRequirementInto(t *testing.T) {
	_default := map[string]interface{}{
		"key":      "",
		"operator": "",
		"values":   func() []interface{} { return nil }(),
	}
	type args struct {
		in v1.NodeSelectorRequirement
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: v1.NodeSelectorRequirement{},
			},
			want: _default,
		},
		{
			name: "Key - default",
			args: args{
				in: func() v1.NodeSelectorRequirement {
					subject := v1.NodeSelectorRequirement{}
					subject.Key = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Operator - default",
			args: args{
				in: func() v1.NodeSelectorRequirement {
					subject := v1.NodeSelectorRequirement{}
					subject.Operator = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Values - default",
			args: args{
				in: func() v1.NodeSelectorRequirement {
					subject := v1.NodeSelectorRequirement{}
					subject.Values = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenDataSourceNodeSelectorRequirementInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceNodeSelectorRequirement() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceNodeSelectorRequirement(t *testing.T) {
	_default := map[string]interface{}{
		"key":      "",
		"operator": "",
		"values":   func() []interface{} { return nil }(),
	}
	type args struct {
		in v1.NodeSelectorRequirement
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: v1.NodeSelectorRequirement{},
			},
			want: _default,
		},
		{
			name: "Key - default",
			args: args{
				in: func() v1.NodeSelectorRequirement {
					subject := v1.NodeSelectorRequirement{}
					subject.Key = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Operator - default",
			args: args{
				in: func() v1.NodeSelectorRequirement {
					subject := v1.NodeSelectorRequirement{}
					subject.Operator = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Values - default",
			args: args{
				in: func() v1.NodeSelectorRequirement {
					subject := v1.NodeSelectorRequirement{}
					subject.Values = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceNodeSelectorRequirement(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceNodeSelectorRequirement() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
