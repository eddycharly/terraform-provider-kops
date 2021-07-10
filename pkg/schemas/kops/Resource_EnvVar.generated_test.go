package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceEnvVar(t *testing.T) {
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
			if got := ExpandResourceEnvVar(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandResourceEnvVar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenResourceEnvVarInto(t *testing.T) {
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
			FlattenResourceEnvVarInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenResourceEnvVar(t *testing.T) {
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
			if got := FlattenResourceEnvVar(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenResourceEnvVar() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
