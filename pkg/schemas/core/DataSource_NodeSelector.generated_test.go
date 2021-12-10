package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	v1 "k8s.io/api/core/v1"
)

func TestExpandDataSourceNodeSelector(t *testing.T) {
	_default := v1.NodeSelector{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want v1.NodeSelector
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"node_selector_terms": func() []interface{} { return nil }(),
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandDataSourceNodeSelector(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceNodeSelector() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceNodeSelectorInto(t *testing.T) {
	_default := map[string]interface{}{
		"node_selector_terms": func() []interface{} { return nil }(),
	}
	type args struct {
		in v1.NodeSelector
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: v1.NodeSelector{},
			},
			want: _default,
		},
		{
			name: "NodeSelectorTerms - default",
			args: args{
				in: func() v1.NodeSelector {
					subject := v1.NodeSelector{}
					subject.NodeSelectorTerms = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenDataSourceNodeSelectorInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceNodeSelector() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceNodeSelector(t *testing.T) {
	_default := map[string]interface{}{
		"node_selector_terms": func() []interface{} { return nil }(),
	}
	type args struct {
		in v1.NodeSelector
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: v1.NodeSelector{},
			},
			want: _default,
		},
		{
			name: "NodeSelectorTerms - default",
			args: args{
				in: func() v1.NodeSelector {
					subject := v1.NodeSelector{}
					subject.NodeSelectorTerms = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceNodeSelector(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceNodeSelector() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
