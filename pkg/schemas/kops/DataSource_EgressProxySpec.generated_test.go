package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceEgressProxySpec(t *testing.T) {
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
			if got := ExpandDataSourceEgressProxySpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandDataSourceEgressProxySpec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenDataSourceEgressProxySpecInto(t *testing.T) {
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
			FlattenDataSourceEgressProxySpecInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenDataSourceEgressProxySpec(t *testing.T) {
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
					return []map[string]interface{}{FlattenDataSourceHTTPProxy(kops.HTTPProxy{})}
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
					return []map[string]interface{}{FlattenDataSourceHTTPProxy(kops.HTTPProxy{})}
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
					return []map[string]interface{}{FlattenDataSourceHTTPProxy(kops.HTTPProxy{})}
				}(),
				"proxy_excludes": "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenDataSourceEgressProxySpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenDataSourceEgressProxySpec() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
