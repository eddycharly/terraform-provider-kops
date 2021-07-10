package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceNodeAuthorizationSpec(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.NodeAuthorizationSpec
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandDataSourceNodeAuthorizationSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandDataSourceNodeAuthorizationSpec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenDataSourceNodeAuthorizationSpecInto(t *testing.T) {
	type args struct {
		in  kops.NodeAuthorizationSpec
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
			FlattenDataSourceNodeAuthorizationSpecInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenDataSourceNodeAuthorizationSpec(t *testing.T) {
	type args struct {
		in kops.NodeAuthorizationSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.NodeAuthorizationSpec{},
			},
			want: map[string]interface{}{
				"node_authorizer": nil,
			},
		},
		{
			name: "NodeAuthorizer - default",
			args: args{
				in: func() kops.NodeAuthorizationSpec {
					subject := kops.NodeAuthorizationSpec{}
					subject.NodeAuthorizer = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"node_authorizer": nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenDataSourceNodeAuthorizationSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenDataSourceNodeAuthorizationSpec() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
