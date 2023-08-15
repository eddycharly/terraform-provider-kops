package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	kopsv1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

func TestExpandResourceUserData(t *testing.T) {
	_default := kopsv1alpha2.UserData{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kopsv1alpha2.UserData
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
			got := ExpandResourceUserData(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceUserData() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceUserDataInto(t *testing.T) {
	_default := map[string]interface{}{
		"name":    "",
		"type":    "",
		"content": "",
	}
	type args struct {
		in kopsv1alpha2.UserData
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.UserData{},
			},
			want: _default,
		},
		{
			name: "Name - default",
			args: args{
				in: func() kopsv1alpha2.UserData {
					subject := kopsv1alpha2.UserData{}
					subject.Name = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Type - default",
			args: args{
				in: func() kopsv1alpha2.UserData {
					subject := kopsv1alpha2.UserData{}
					subject.Type = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Content - default",
			args: args{
				in: func() kopsv1alpha2.UserData {
					subject := kopsv1alpha2.UserData{}
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
			FlattenResourceUserDataInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceUserData() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceUserData(t *testing.T) {
	_default := map[string]interface{}{
		"name":    "",
		"type":    "",
		"content": "",
	}
	type args struct {
		in kopsv1alpha2.UserData
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kopsv1alpha2.UserData{},
			},
			want: _default,
		},
		{
			name: "Name - default",
			args: args{
				in: func() kopsv1alpha2.UserData {
					subject := kopsv1alpha2.UserData{}
					subject.Name = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Type - default",
			args: args{
				in: func() kopsv1alpha2.UserData {
					subject := kopsv1alpha2.UserData{}
					subject.Type = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Content - default",
			args: args{
				in: func() kopsv1alpha2.UserData {
					subject := kopsv1alpha2.UserData{}
					subject.Content = ""
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceUserData(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceUserData() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
