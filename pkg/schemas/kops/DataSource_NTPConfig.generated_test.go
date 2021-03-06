package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceNTPConfig(t *testing.T) {
	_default := kops.NTPConfig{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.NTPConfig
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"managed": nil,
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandDataSourceNTPConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceNTPConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceNTPConfigInto(t *testing.T) {
	_default := map[string]interface{}{
		"managed": nil,
	}
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
			want: _default,
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
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenDataSourceNTPConfigInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceNTPConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceNTPConfig(t *testing.T) {
	_default := map[string]interface{}{
		"managed": nil,
	}
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
			want: _default,
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
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceNTPConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceNTPConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
