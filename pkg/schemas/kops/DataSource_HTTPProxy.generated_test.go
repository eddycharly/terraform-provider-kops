package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceHTTPProxy(t *testing.T) {
	_default := kops.HTTPProxy{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.HTTPProxy
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"host": "",
					"port": 0,
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandDataSourceHTTPProxy(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceHTTPProxy() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceHTTPProxyInto(t *testing.T) {
	_default := map[string]interface{}{
		"host": "",
		"port": 0,
	}
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
			want: _default,
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
			want: _default,
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
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenDataSourceHTTPProxyInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceHTTPProxy() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceHTTPProxy(t *testing.T) {
	_default := map[string]interface{}{
		"host": "",
		"port": 0,
	}
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
			want: _default,
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
			want: _default,
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
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceHTTPProxy(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceHTTPProxy() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
