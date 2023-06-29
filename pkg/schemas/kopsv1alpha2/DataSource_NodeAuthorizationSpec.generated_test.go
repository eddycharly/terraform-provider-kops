package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

func TestExpandDataSourceNodeAuthorizationSpec(t *testing.T) {
	_default := kopsv1alpha2.NodeAuthorizationSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kopsv1alpha2.NodeAuthorizationSpec
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"node_authorizer": nil,
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandDataSourceNodeAuthorizationSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceNodeAuthorizationSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceNodeAuthorizationSpecInto(t *testing.T) {
	_default := map[string]interface{}{
		"node_authorizer": nil,
	}
	type args struct {
		in kopsv1alpha2.NodeAuthorizationSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.NodeAuthorizationSpec{},
			},
			want: _default,
		},
		{
			name: "NodeAuthorizer - default",
			args: args{
				in: func() kopsv1alpha2.NodeAuthorizationSpec {
					subject := kopsv1alpha2.NodeAuthorizationSpec{}
					subject.NodeAuthorizer = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenDataSourceNodeAuthorizationSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceNodeAuthorizationSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceNodeAuthorizationSpec(t *testing.T) {
	_default := map[string]interface{}{
		"node_authorizer": nil,
	}
	type args struct {
		in kopsv1alpha2.NodeAuthorizationSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.NodeAuthorizationSpec{},
			},
			want: _default,
		},
		{
			name: "NodeAuthorizer - default",
			args: args{
				in: func() kopsv1alpha2.NodeAuthorizationSpec {
					subject := kopsv1alpha2.NodeAuthorizationSpec{}
					subject.NodeAuthorizer = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceNodeAuthorizationSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceNodeAuthorizationSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
