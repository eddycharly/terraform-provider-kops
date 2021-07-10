package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceHTTPProxy(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.HTTPProxy
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandDataSourceHTTPProxy(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandDataSourceHTTPProxy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenDataSourceHTTPProxyInto(t *testing.T) {
	type args struct {
		in  kops.HTTPProxy
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
			FlattenDataSourceHTTPProxyInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenDataSourceHTTPProxy(t *testing.T) {
	type args struct {
		in kops.HTTPProxy
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.HTTPProxy{},
			},
			want: map[string]interface{}{
				"host": "",
				"port": 0,
			},
		},
		{
			name: "Host - default",
			args: args{
				in: func() kops.HTTPProxy {
					subject := kops.HTTPProxy{}
					subject.Host = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
				"host": "",
				"port": 0,
			},
		},
		{
			name: "Port - default",
			args: args{
				in: func() kops.HTTPProxy {
					subject := kops.HTTPProxy{}
					subject.Port = 0
					return subject
				}(),
			},
			want: map[string]interface{}{
				"host": "",
				"port": 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenDataSourceHTTPProxy(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenDataSourceHTTPProxy() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
