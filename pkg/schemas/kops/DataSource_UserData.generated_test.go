package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceUserData(t *testing.T) {
	_default := kops.UserData{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.UserData
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"name":    "",
					"type":    "",
					"content": "",
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandDataSourceUserData(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceUserData() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceUserDataInto(t *testing.T) {
	_default := map[string]interface{}{
		"name":    "",
		"type":    "",
		"content": "",
	}
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenDataSourceUserDataInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceUserData() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceUserData(t *testing.T) {
	_default := map[string]interface{}{
		"name":    "",
		"type":    "",
		"content": "",
	}
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceUserData(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceUserData() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
