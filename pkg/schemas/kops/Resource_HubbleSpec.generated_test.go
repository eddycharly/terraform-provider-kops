package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceHubbleSpec(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.HubbleSpec
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandResourceHubbleSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandResourceHubbleSpec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenResourceHubbleSpecInto(t *testing.T) {
	type args struct {
		in  kops.HubbleSpec
		out map[string]interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			FlattenResourceHubbleSpecInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenResourceHubbleSpec(t *testing.T) {
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
			want: map[string]interface{}{
				"enabled": nil,
				"metrics": func() []interface{} { return nil }(),
			},
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
			want: map[string]interface{}{
				"enabled": nil,
				"metrics": func() []interface{} { return nil }(),
			},
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
			want: map[string]interface{}{
				"enabled": nil,
				"metrics": func() []interface{} { return nil }(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenResourceHubbleSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenResourceHubbleSpec() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
