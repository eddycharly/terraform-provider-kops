package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops/v1alpha2"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

func TestExpandDataSourceEgressProxySpec(t *testing.T) {
	_default := kopsv1alpha2.EgressProxySpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kopsv1alpha2.EgressProxySpec
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"http_proxy":     func() []interface{} { return []interface{}{FlattenDataSourceHTTPProxy(kopsv1alpha2.HTTPProxy{})} }(),
					"proxy_excludes": "",
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandDataSourceEgressProxySpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceEgressProxySpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceEgressProxySpecInto(t *testing.T) {
	_default := map[string]interface{}{
		"http_proxy":     func() []interface{} { return []interface{}{FlattenDataSourceHTTPProxy(kopsv1alpha2.HTTPProxy{})} }(),
		"proxy_excludes": "",
	}
	type args struct {
		in kopsv1alpha2.EgressProxySpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.EgressProxySpec{},
			},
			want: _default,
		},
		{
			name: "HTTPProxy - default",
			args: args{
				in: func() kopsv1alpha2.EgressProxySpec {
					subject := kopsv1alpha2.EgressProxySpec{}
					subject.HTTPProxy = v1alpha2.HTTPProxy{}
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ProxyExcludes - default",
			args: args{
				in: func() kopsv1alpha2.EgressProxySpec {
					subject := kopsv1alpha2.EgressProxySpec{}
					subject.ProxyExcludes = ""
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenDataSourceEgressProxySpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceEgressProxySpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceEgressProxySpec(t *testing.T) {
	_default := map[string]interface{}{
		"http_proxy":     func() []interface{} { return []interface{}{FlattenDataSourceHTTPProxy(kopsv1alpha2.HTTPProxy{})} }(),
		"proxy_excludes": "",
	}
	type args struct {
		in kopsv1alpha2.EgressProxySpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.EgressProxySpec{},
			},
			want: _default,
		},
		{
			name: "HTTPProxy - default",
			args: args{
				in: func() kopsv1alpha2.EgressProxySpec {
					subject := kopsv1alpha2.EgressProxySpec{}
					subject.HTTPProxy = v1alpha2.HTTPProxy{}
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ProxyExcludes - default",
			args: args{
				in: func() kopsv1alpha2.EgressProxySpec {
					subject := kopsv1alpha2.EgressProxySpec{}
					subject.ProxyExcludes = ""
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceEgressProxySpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceEgressProxySpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
