package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

func TestExpandResourceHubbleSpec(t *testing.T) {
	_default := kopsv1alpha2.HubbleSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kopsv1alpha2.HubbleSpec
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"enabled": nil,
					"metrics": func() []interface{} { return nil }(),
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandResourceHubbleSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceHubbleSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceHubbleSpecInto(t *testing.T) {
	_default := map[string]interface{}{
		"enabled": nil,
		"metrics": func() []interface{} { return nil }(),
	}
	type args struct {
		in kopsv1alpha2.HubbleSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.HubbleSpec{},
			},
			want: _default,
		},
		{
			name: "Enabled - default",
			args: args{
				in: func() kopsv1alpha2.HubbleSpec {
					subject := kopsv1alpha2.HubbleSpec{}
					subject.Enabled = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Metrics - default",
			args: args{
				in: func() kopsv1alpha2.HubbleSpec {
					subject := kopsv1alpha2.HubbleSpec{}
					subject.Metrics = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenResourceHubbleSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceHubbleSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceHubbleSpec(t *testing.T) {
	_default := map[string]interface{}{
		"enabled": nil,
		"metrics": func() []interface{} { return nil }(),
	}
	type args struct {
		in kopsv1alpha2.HubbleSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.HubbleSpec{},
			},
			want: _default,
		},
		{
			name: "Enabled - default",
			args: args{
				in: func() kopsv1alpha2.HubbleSpec {
					subject := kopsv1alpha2.HubbleSpec{}
					subject.Enabled = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Metrics - default",
			args: args{
				in: func() kopsv1alpha2.HubbleSpec {
					subject := kopsv1alpha2.HubbleSpec{}
					subject.Metrics = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceHubbleSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceHubbleSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
