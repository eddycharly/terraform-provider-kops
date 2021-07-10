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
	type args struct {
		in  kops.EgressProxySpec
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
			FlattenResourceEgressProxySpecInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenResourceEgressProxySpec(t *testing.T) {
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
			want: map[string]interface{}{
				"http_proxy": func() []map[string]interface{} {
					return []map[string]interface{}{FlattenResourceHTTPProxy(kops.HTTPProxy{})}
				}(),
				"proxy_excludes": "",
			},
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
			want: map[string]interface{}{
				"http_proxy": func() []map[string]interface{} {
					return []map[string]interface{}{FlattenResourceHTTPProxy(kops.HTTPProxy{})}
				}(),
				"proxy_excludes": "",
			},
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
			want: map[string]interface{}{
				"http_proxy": func() []map[string]interface{} {
					return []map[string]interface{}{FlattenResourceHTTPProxy(kops.HTTPProxy{})}
				}(),
				"proxy_excludes": "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenResourceEgressProxySpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenResourceEgressProxySpec() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
