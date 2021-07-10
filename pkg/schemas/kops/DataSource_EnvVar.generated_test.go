package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceEnvVar(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.EnvVar
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandDataSourceEnvVar(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandDataSourceEnvVar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenDataSourceEnvVarInto(t *testing.T) {
	type args struct {
		in  kops.EnvVar
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
			FlattenDataSourceEnvVarInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenDataSourceEnvVar(t *testing.T) {
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
			want: map[string]interface{}{
				"name":  "",
				"value": "",
			},
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
			want: map[string]interface{}{
				"name":  "",
				"value": "",
			},
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
			want: map[string]interface{}{
				"name":  "",
				"value": "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenDataSourceEnvVar(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenDataSourceEnvVar() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
