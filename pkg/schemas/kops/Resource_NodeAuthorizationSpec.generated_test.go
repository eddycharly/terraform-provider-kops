package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceNodeAuthorizationSpec(t *testing.T) {
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
			if got := ExpandResourceNodeAuthorizationSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandResourceNodeAuthorizationSpec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenResourceNodeAuthorizationSpecInto(t *testing.T) {
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
			FlattenResourceNodeAuthorizationSpecInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenResourceNodeAuthorizationSpec(t *testing.T) {
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
			if got := FlattenResourceNodeAuthorizationSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenResourceNodeAuthorizationSpec() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
