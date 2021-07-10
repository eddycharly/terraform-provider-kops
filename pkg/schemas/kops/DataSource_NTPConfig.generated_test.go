package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceNTPConfig(t *testing.T) {
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
			if got := ExpandDataSourceNTPConfig(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandDataSourceNTPConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenDataSourceNTPConfigInto(t *testing.T) {
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
			FlattenDataSourceNTPConfigInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenDataSourceNTPConfig(t *testing.T) {
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
			if got := FlattenDataSourceNTPConfig(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenDataSourceNTPConfig() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
