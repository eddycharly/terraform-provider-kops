package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceHubbleSpec(t *testing.T) {
	_default := kops.HubbleSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.HubbleSpec
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
			got := ExpandDataSourceHubbleSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceHubbleSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceHubbleSpecInto(t *testing.T) {
	_default := map[string]interface{}{
		"enabled": nil,
		"metrics": func() []interface{} { return nil }(),
	}
	type args struct {
		in kops.HubbleSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.HubbleSpec{},
			},
			want: _default,
		},
		{
			name: "Enabled - default",
			args: args{
				in: func() kops.HubbleSpec {
					subject := kops.HubbleSpec{}
					subject.Enabled = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Metrics - default",
			args: args{
				in: func() kops.HubbleSpec {
					subject := kops.HubbleSpec{}
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
			FlattenDataSourceHubbleSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceHubbleSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceHubbleSpec(t *testing.T) {
	_default := map[string]interface{}{
		"enabled": nil,
		"metrics": func() []interface{} { return nil }(),
	}
	type args struct {
		in kops.HubbleSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.HubbleSpec{},
			},
			want: _default,
		},
		{
			name: "Enabled - default",
			args: args{
				in: func() kops.HubbleSpec {
					subject := kops.HubbleSpec{}
					subject.Enabled = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Metrics - default",
			args: args{
				in: func() kops.HubbleSpec {
					subject := kops.HubbleSpec{}
					subject.Metrics = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceHubbleSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceHubbleSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
