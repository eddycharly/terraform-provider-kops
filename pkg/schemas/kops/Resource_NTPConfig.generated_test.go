package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceNTPConfig(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.NTPConfig
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandResourceNTPConfig(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandResourceNTPConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenResourceNTPConfigInto(t *testing.T) {
	type args struct {
		in  kops.NTPConfig
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
			FlattenResourceNTPConfigInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenResourceNTPConfig(t *testing.T) {
	type args struct {
		in kops.NTPConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.NTPConfig{},
			},
			want: map[string]interface{}{
				"managed": nil,
			},
		},
		{
			name: "Managed - default",
			args: args{
				in: func() kops.NTPConfig {
					subject := kops.NTPConfig{}
					subject.Managed = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"managed": nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenResourceNTPConfig(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenResourceNTPConfig() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
