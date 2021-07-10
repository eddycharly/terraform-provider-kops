package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceUserData(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.UserData
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandResourceUserData(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandResourceUserData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenResourceUserDataInto(t *testing.T) {
	type args struct {
		in  kops.UserData
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
			FlattenResourceUserDataInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenResourceUserData(t *testing.T) {
	type args struct {
		in kops.UserData
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.UserData{},
			},
			want: map[string]interface{}{
				"name":    "",
				"type":    "",
				"content": "",
			},
		},
		{
			name: "Name - default",
			args: args{
				in: func() kops.UserData {
					subject := kops.UserData{}
					subject.Name = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
				"name":    "",
				"type":    "",
				"content": "",
			},
		},
		{
			name: "Type - default",
			args: args{
				in: func() kops.UserData {
					subject := kops.UserData{}
					subject.Type = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
				"name":    "",
				"type":    "",
				"content": "",
			},
		},
		{
			name: "Content - default",
			args: args{
				in: func() kops.UserData {
					subject := kops.UserData{}
					subject.Content = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
				"name":    "",
				"type":    "",
				"content": "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenResourceUserData(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenResourceUserData() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
