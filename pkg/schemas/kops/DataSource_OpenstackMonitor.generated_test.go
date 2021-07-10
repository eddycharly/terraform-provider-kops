package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceOpenstackMonitor(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.OpenstackMonitor
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandDataSourceOpenstackMonitor(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandDataSourceOpenstackMonitor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenDataSourceOpenstackMonitorInto(t *testing.T) {
	type args struct {
		in  kops.OpenstackMonitor
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
			FlattenDataSourceOpenstackMonitorInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenDataSourceOpenstackMonitor(t *testing.T) {
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
			want: map[string]interface{}{
				"delay":       nil,
				"timeout":     nil,
				"max_retries": nil,
			},
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
			want: map[string]interface{}{
				"delay":       nil,
				"timeout":     nil,
				"max_retries": nil,
			},
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
			want: map[string]interface{}{
				"delay":       nil,
				"timeout":     nil,
				"max_retries": nil,
			},
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
			want: map[string]interface{}{
				"delay":       nil,
				"timeout":     nil,
				"max_retries": nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenDataSourceOpenstackMonitor(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenDataSourceOpenstackMonitor() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
