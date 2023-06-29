package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

func TestExpandDataSourceHTTPProxy(t *testing.T) {
	_default := kopsv1alpha2.HTTPProxy{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kopsv1alpha2.HTTPProxy
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
		in kopsv1alpha2.HTTPProxy
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.HTTPProxy{},
			},
			want: _default,
		},
		{
			name: "Host - default",
			args: args{
				in: func() kopsv1alpha2.HTTPProxy {
					subject := kopsv1alpha2.HTTPProxy{}
					subject.Host = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Port - default",
			args: args{
				in: func() kopsv1alpha2.HTTPProxy {
					subject := kopsv1alpha2.HTTPProxy{}
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
		in kopsv1alpha2.HTTPProxy
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.HTTPProxy{},
			},
			want: _default,
		},
		{
			name: "Host - default",
			args: args{
				in: func() kopsv1alpha2.HTTPProxy {
					subject := kopsv1alpha2.HTTPProxy{}
					subject.Host = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Port - default",
			args: args{
				in: func() kopsv1alpha2.HTTPProxy {
					subject := kopsv1alpha2.HTTPProxy{}
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
