package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestExpandResourceLabelSelectorRequirement(t *testing.T) {
	_default := meta.LabelSelectorRequirement{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want meta.LabelSelectorRequirement
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
			got := ExpandResourceLabelSelectorRequirement(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceLabelSelectorRequirement() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceLabelSelectorRequirementInto(t *testing.T) {
	_default := map[string]interface{}{
		"key":      "",
		"operator": "",
		"values":   func() []interface{} { return nil }(),
	}
	type args struct {
		in meta.LabelSelectorRequirement
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: meta.LabelSelectorRequirement{},
			},
			want: _default,
		},
		{
			name: "Key - default",
			args: args{
				in: func() meta.LabelSelectorRequirement {
					subject := meta.LabelSelectorRequirement{}
					subject.Key = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Operator - default",
			args: args{
				in: func() meta.LabelSelectorRequirement {
					subject := meta.LabelSelectorRequirement{}
					subject.Operator = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Values - default",
			args: args{
				in: func() meta.LabelSelectorRequirement {
					subject := meta.LabelSelectorRequirement{}
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
			FlattenResourceLabelSelectorRequirementInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceLabelSelectorRequirement() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceLabelSelectorRequirement(t *testing.T) {
	_default := map[string]interface{}{
		"key":      "",
		"operator": "",
		"values":   func() []interface{} { return nil }(),
	}
	type args struct {
		in meta.LabelSelectorRequirement
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: meta.LabelSelectorRequirement{},
			},
			want: _default,
		},
		{
			name: "Key - default",
			args: args{
				in: func() meta.LabelSelectorRequirement {
					subject := meta.LabelSelectorRequirement{}
					subject.Key = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Operator - default",
			args: args{
				in: func() meta.LabelSelectorRequirement {
					subject := meta.LabelSelectorRequirement{}
					subject.Operator = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Values - default",
			args: args{
				in: func() meta.LabelSelectorRequirement {
					subject := meta.LabelSelectorRequirement{}
					subject.Values = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceLabelSelectorRequirement(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceLabelSelectorRequirement() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
