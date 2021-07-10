package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceEnvVar(t *testing.T) {
	_default := kops.EnvVar{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.EnvVar
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"name":  "",
					"value": "",
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandDataSourceEnvVar(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceEnvVar() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceEnvVarInto(t *testing.T) {
	_default := map[string]interface{}{
		"name":  "",
		"value": "",
	}
	type args struct {
		in kops.EnvVar
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.EnvVar{},
			},
			want: _default,
		},
		{
			name: "Name - default",
			args: args{
				in: func() kops.EnvVar {
					subject := kops.EnvVar{}
					subject.Name = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Value - default",
			args: args{
				in: func() kops.EnvVar {
					subject := kops.EnvVar{}
					subject.Value = ""
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenDataSourceEnvVarInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceEnvVar() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceEnvVar(t *testing.T) {
	_default := map[string]interface{}{
		"name":  "",
		"value": "",
	}
	type args struct {
		in kops.EnvVar
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.EnvVar{},
			},
			want: _default,
		},
		{
			name: "Name - default",
			args: args{
				in: func() kops.EnvVar {
					subject := kops.EnvVar{}
					subject.Name = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Value - default",
			args: args{
				in: func() kops.EnvVar {
					subject := kops.EnvVar{}
					subject.Value = ""
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceEnvVar(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceEnvVar() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
