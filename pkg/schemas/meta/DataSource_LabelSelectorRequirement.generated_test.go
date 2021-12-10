package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestExpandDataSourceLabelSelectorRequirement(t *testing.T) {
	_default := v1.LabelSelectorRequirement{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want v1.LabelSelectorRequirement
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
			got := ExpandDataSourceLabelSelectorRequirement(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceLabelSelectorRequirement() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceLabelSelectorRequirementInto(t *testing.T) {
	_default := map[string]interface{}{
		"key":      "",
		"operator": "",
		"values":   func() []interface{} { return nil }(),
	}
	type args struct {
		in v1.LabelSelectorRequirement
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: v1.LabelSelectorRequirement{},
			},
			want: _default,
		},
		{
			name: "Key - default",
			args: args{
				in: func() v1.LabelSelectorRequirement {
					subject := v1.LabelSelectorRequirement{}
					subject.Key = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Operator - default",
			args: args{
				in: func() v1.LabelSelectorRequirement {
					subject := v1.LabelSelectorRequirement{}
					subject.Operator = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Values - default",
			args: args{
				in: func() v1.LabelSelectorRequirement {
					subject := v1.LabelSelectorRequirement{}
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
			FlattenDataSourceLabelSelectorRequirementInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceLabelSelectorRequirement() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceLabelSelectorRequirement(t *testing.T) {
	_default := map[string]interface{}{
		"key":      "",
		"operator": "",
		"values":   func() []interface{} { return nil }(),
	}
	type args struct {
		in v1.LabelSelectorRequirement
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: v1.LabelSelectorRequirement{},
			},
			want: _default,
		},
		{
			name: "Key - default",
			args: args{
				in: func() v1.LabelSelectorRequirement {
					subject := v1.LabelSelectorRequirement{}
					subject.Key = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Operator - default",
			args: args{
				in: func() v1.LabelSelectorRequirement {
					subject := v1.LabelSelectorRequirement{}
					subject.Operator = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Values - default",
			args: args{
				in: func() v1.LabelSelectorRequirement {
					subject := v1.LabelSelectorRequirement{}
					subject.Values = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceLabelSelectorRequirement(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceLabelSelectorRequirement() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
