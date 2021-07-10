package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceEgressProxySpec(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.EgressProxySpec
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandResourceEgressProxySpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandResourceEgressProxySpec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenResourceEgressProxySpecInto(t *testing.T) {
	_default := map[string]interface{}{
		"http_proxy": func() []map[string]interface{} {
			return []map[string]interface{}{FlattenResourceHTTPProxy(kops.HTTPProxy{})}
		}(),
		"proxy_excludes": "",
	}
	type args struct {
		in kops.EgressProxySpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.EgressProxySpec{},
			},
			want: _default,
		},
		{
			name: "HTTPProxy - default",
			args: args{
				in: func() kops.EgressProxySpec {
					subject := kops.EgressProxySpec{}
					subject.HTTPProxy = kops.HTTPProxy{}
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ProxyExcludes - default",
			args: args{
				in: func() kops.EgressProxySpec {
					subject := kops.EgressProxySpec{}
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
			FlattenResourceEgressProxySpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceEgressProxySpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceEgressProxySpec(t *testing.T) {
	_default := map[string]interface{}{
		"http_proxy": func() []map[string]interface{} {
			return []map[string]interface{}{FlattenResourceHTTPProxy(kops.HTTPProxy{})}
		}(),
		"proxy_excludes": "",
	}
	type args struct {
		in kops.EgressProxySpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.EgressProxySpec{},
			},
			want: _default,
		},
		{
			name: "HTTPProxy - default",
			args: args{
				in: func() kops.EgressProxySpec {
					subject := kops.EgressProxySpec{}
					subject.HTTPProxy = kops.HTTPProxy{}
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ProxyExcludes - default",
			args: args{
				in: func() kops.EgressProxySpec {
					subject := kops.EgressProxySpec{}
					subject.ProxyExcludes = ""
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceEgressProxySpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceEgressProxySpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
