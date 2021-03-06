package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceOpenstackMonitor(t *testing.T) {
	_default := kops.OpenstackMonitor{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.OpenstackMonitor
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"delay":       nil,
					"timeout":     nil,
					"max_retries": nil,
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandDataSourceOpenstackMonitor(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceOpenstackMonitor() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceOpenstackMonitorInto(t *testing.T) {
	_default := map[string]interface{}{
		"delay":       nil,
		"timeout":     nil,
		"max_retries": nil,
	}
	type args struct {
		in kops.OpenstackMonitor
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.OpenstackMonitor{},
			},
			want: _default,
		},
		{
			name: "Delay - default",
			args: args{
				in: func() kops.OpenstackMonitor {
					subject := kops.OpenstackMonitor{}
					subject.Delay = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Timeout - default",
			args: args{
				in: func() kops.OpenstackMonitor {
					subject := kops.OpenstackMonitor{}
					subject.Timeout = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MaxRetries - default",
			args: args{
				in: func() kops.OpenstackMonitor {
					subject := kops.OpenstackMonitor{}
					subject.MaxRetries = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenDataSourceOpenstackMonitorInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceOpenstackMonitor() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceOpenstackMonitor(t *testing.T) {
	_default := map[string]interface{}{
		"delay":       nil,
		"timeout":     nil,
		"max_retries": nil,
	}
	type args struct {
		in kops.OpenstackMonitor
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.OpenstackMonitor{},
			},
			want: _default,
		},
		{
			name: "Delay - default",
			args: args{
				in: func() kops.OpenstackMonitor {
					subject := kops.OpenstackMonitor{}
					subject.Delay = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Timeout - default",
			args: args{
				in: func() kops.OpenstackMonitor {
					subject := kops.OpenstackMonitor{}
					subject.Timeout = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MaxRetries - default",
			args: args{
				in: func() kops.OpenstackMonitor {
					subject := kops.OpenstackMonitor{}
					subject.MaxRetries = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceOpenstackMonitor(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceOpenstackMonitor() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
